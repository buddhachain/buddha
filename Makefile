

DOCKERIMAGENAME = buddhachain/buddha-server
PKGNAME = github.com/buddhachain/buddha
PREFIX = $(shell pwd)
export DOCKERIMAGENAME
export PREFIX

buddha:
	@echo "Building buddha...."
	go build -o sampleconfig/buddha -ldflags "-linkmode external -extldflags '-static' -s -w" -mod=vendor $(PKGNAME)/cmd/apiserver

event:
	@echo "Building buddha...."
	cd eventserver && go build -ldflags "-linkmode external -extldflags '-static' -s -w" -mod=vendor -o eventserver

image:
	@echo "Building image...."
	-@docker images |grep "$(DOCKERIMAGENAME)"|grep "latest"|awk '{print $$3}'|xargs docker rmi -f
	docker build -t $DOCKERIMAGENAME .

all: buddha image

.PHONY: \
	all \
	buddha \
	event \
	image