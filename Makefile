

DOCKERIMAGENAME = buddhachain/buddha-server

PREFIX = $(shell pwd)
export DOCKERIMAGENAME
export PREFIX

buddha:
	@echo "Building buddha...."
	go build -ldflags "-linkmode external -extldflags '-static' -s -w" -o sampleconfig/buddha

image:
	@echo "Building image...."
	-@docker images |grep "$(DOCKERIMAGENAME)"|grep "latest"|awk '{print $$3}'|xargs docker rmi -f
	docker build -t $DOCKERIMAGENAME .

all: buddha image

.PHONY: \
	all \
	buddha \
	image