package github

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"githubhook/util"

	githttp "github.com/google/go-github/v32/github"
	"github.com/sirupsen/logrus"
)

type GitClient struct {
	Logger *logrus.Logger
	base   string

	// cacheDir is the local git cache dir for fast cloning
	cacheDir string

	// git is the path of git binary
	git string

	token string
}

// NewClient creates a new fully operational GitHub client.
func NewGitClient(token string) (*GitClient, error) {
	g, err := exec.LookPath("git")
	if err != nil {
		return nil, err
	}

	t, err := ioutil.TempDir("", "git")
	if err != nil {
		return nil, err
	}

	base := "https://" + token + "@github.com"

	return &GitClient{
		token:    token,
		base:     base,
		cacheDir: t,
		git:      g,
		Logger:   logrus.New(),
	}, nil
}

func (c *GitClient) Clean() error {
	return os.RemoveAll(c.cacheDir)
}

// Clone clones a repository.
// This function may take a long time if it is the first time cloning the repo.
// In that case, it must do a full git mirror clone. For large repos, this can
// take a while. Once that is done, it will do a git fetch instead of a clone,
// which will usually take at most a few seconds
func (c *GitClient) Clone(repo string) (string, error) {
	ac := NewAPIClient(c.token)

	reps, err := ac.ListRepos()
	if err != nil {
		return "", err
	}

	var cRepo *githttp.Repository
	for _, r := range reps {
		if r.Name != nil && *r.Name == repo {
			cRepo = r
			break
		}
	}

	if cRepo == nil {
		return "", fmt.Errorf("repository with the name: %s, not found for the authenticated user", repo)
	}

	c.Logger.Infof("login user: %s, for repo: %s", *cRepo.GetOwner().Login, repo)

	remote := c.base + "/" + path.Join(*cRepo.GetOwner().Login, repo) + ".git"
	c.Logger.Infof("remote url for clone: %s", remote)

	cachePath := filepath.Join(c.cacheDir, repo) + ".git"
	c.Logger.Infof("cachePath directory: %s", cachePath)

	if _, err = os.Stat(cachePath); err != nil {
		if os.IsNotExist(err) {
			// Cache miss, clone it now.
			c.Logger.Infof("cache directory not found for the repo. Cloning for the first time: %s", repo)
			if err = os.Mkdir(filepath.Dir(cachePath), os.ModePerm); err != nil && !os.IsExist(err) {
				return "", err
			}
			if b, rErr := util.RetryCmd(c.Logger, "", c.git, "clone", "--mirror", remote, cachePath); rErr != nil {
				return "", fmt.Errorf("git cachePath clone error: %v. output: %s", rErr, string(b))
			}
		} else {
			return "", err
		}
	} else {
		// Cache hit. Do a git fetch to keep updated.
		c.Logger.Infof("cache hit. Fetching %s.", repo)
		if b, rErr := util.RetryCmd(c.Logger, cachePath, c.git, "fetch"); rErr != nil {
			return "", fmt.Errorf("git fetch error: %v. output: %s", rErr, string(b))
		}
	}

	newLoc, err := ioutil.TempDir("", "git")
	if err != nil {
		return newLoc, err
	}

	c.Logger.Infof("Cloning in the new Location: %s", newLoc)
	if b, err := exec.Command(c.git, "clone", cachePath, newLoc).CombinedOutput(); err != nil {
		return "", fmt.Errorf("git repo clone error: %v. output: %s", err, string(b))
	}
	return c.cacheDir, nil
}
