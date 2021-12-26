IMAGE_FLAG := $(shell git rev-parse --abbrev-ref HEAD) #get image flag from current branch name.
IMAGE_TAG  ?= registry.hamdocker.ir/artmis/legolas:${IMAGE_FLAG}

compile:
	go build --ldflags "-X main.CommitRefName=$(COMMIT_REF_SLUG) -X main.CommitSHA=$(COMMIT_SHORT_SHA) -X main.BuildDate=$(CURRENT_DATETIME) -linkmode external -extldflags '-static'" -v -o ./legolas .

code-gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

build-image:
	docker build -t ${IMAGE_TAG} .

build-image-dev:
	docker build -t legolas:latest .
