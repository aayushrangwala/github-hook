# If the "make" command is run without any argument, the first default goal target will run.
# A default goal is the first target without a (.) at the begining of its name.
# The default goal can be override by specifying ".DEFAULT_GOAL=<target>".
# The convention is to keep "all: <all targets, space separated>" at the begining of the Makefile

# export <ENV_VAR> command in make exports that env var with its value

# CGO_ENABLED is an env var used at the time of compiling and building the programs.
# It needs to be enabled (1) for cross compiling and disabled (0) for native builds
export CGO_ENABLED=0

# GO111MODULE is the env var used by the mod tool (go.mod file) is useful for enabling the module behaviour
export GO111MODULE=on

# GOOS is used to build static linked binary for a go program by setting the os
export GOOS=linux

# GOOS is used to build static linked binary for a go program by setting the Architecture
export GOARCH=amd64

# Declaring the binary name
GO_APP_BINARY ?= www

# Declaration for project name
PROJECT ?= github-hook

# Declaration for the docker image name
IMAGE ?= www

# Declaring and calculating the version (tag) for the docker image in the format: <data>-<commit>
VERSION ?= $(shell date +v%Y%m%d)-$(shell git describe --tags --always)

DOCKER_HUB_USER ?= aayushrangwala

# A Verb with some commands under it is called as target in Makefile.
# Target is used to run as an argument along with "make" command. It basically runs the commands defined under it

# all target runs all the targets specified.
all: lint test coverage run clean

# fmt target to format the go code
fmt:
	go fmt ./...

# vet to run the vet linter on the go code
vet:
	go vet ./...

# lint target is used to run the golangci-lint binary to check for the linting errors
lint:
	golangci-lint run --skip-dirs='(vendor)' -vc ./.golangci.yaml ./...

# yaml-lint will run the linter for all the yaml files in the root directory or sub-directory
yaml-lint:
	yamllint -c .yamllint.conf ./

# test is the target to run the tests for all the directories and sub directories
test:
	go test -v ./... -coverprofile coverage.out

# coverage taret will run a script which will check the test coverage of the project, if it is greater than 85% or not
coverage:
	scripts/gocoverage.sh

# dep target will sync the dependencies in the project
dep:
	go mod vendor

# run is the target used to compile and build the program (main.go) by calling the 'build' target and run
run: build
	./$(GO_APP_BINARY)

# build target is used to only to compile and build the program (main.go) with running fmt and vet targets also
build: clean fmt vet dep
	go build -o $(GO_APP_BINARY)

# clean is the target which will clean the object files in the temporary source directory and the binary
# ,which are created at the time of build
clean:
	go clean
	rm -f $(GO_APP_BINARY) coverage.out

# build target will call the docker build command to build the docker image by giving the dockerfile and current context
docker-build:
	docker build -t docker.io/$(PROJECT)/$(IMAGE):$(VERSION) \
		-t docker.io/$(PROJECT)/$(IMAGE):latest -f Dockerfile .

# push target will push the docker image
docker-push: docker-login
	docker push docker.io/$(PROJECT)/$(IMAGE):$(VERSION)
	docker push docker.io/$(PROJECT)/$(IMAGE):latest

# run target will run the docker image
docker-run: docker-build
	docker run -p 8888:8888 docker.io/$(PROJECT)/$(IMAGE):$(VERSION) --network="host"

docker-login:
	cat creds | docker login -u=$(DOCKER_HUB_USER) --password-stdin

# This target will build and push the code in a commit push which can also be used to create a release
ci-release: docker-build docker-push

# .PHONY is a special built in target which is used to specify the target names explicitely
# so that it is not conflicted with the file names and also it improves performance
.PHONY: clean lint test coverage build run docker-build docker-push docker-run docker-login ci-release

