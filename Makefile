.PHONY: docker

EXECUTABLE ?= drone-sentry
IMAGE ?= urbint/drone-sentry

docker:
  go get
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(EXECUTABLE)
	docker build --rm -t $(IMAGE) .
