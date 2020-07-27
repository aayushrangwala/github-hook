package util

import (
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

// retryCmd will retry the command a few times with backoff. Use this for any
// commands that will be talking to GitHub, such as clones or fetches.
func RetryCmd(l *logrus.Logger, dir, cmdStr string, arg ...string) ([]byte, error) {
	var b []byte
	var err error
	sleepyTime := time.Second
	for i := 0; i < 3; i++ {
		cmd := exec.Command(cmdStr, arg...)
		cmd.Dir = dir
		b, err = cmd.CombinedOutput()
		if err != nil {
			l.Warningf("Running %s %v returned error %v with output %s.", cmdStr, arg, err, string(b))
			time.Sleep(sleepyTime)
			sleepyTime *= 2
			continue
		}
		break
	}
	return b, err
}
