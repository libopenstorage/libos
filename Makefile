HAS_SDKTEST := $(shell command -v sdk-test 2> /dev/null)
BRANCH	:= $(shell git rev-parse --abbrev-ref HEAD)

ifeq ($(TRAVIS_BRANCH), master)
MOCKSDKSERVERTAG := latest
else

ifeq ($(BRANCH), master)
MOCKSDKSERVERTAG := latest
else
MOCKSDKSERVERTAG := $(shell go run tools/sdkver/sdkver.go)
endif

endif

REGISTRY = openstorage
IMAGE_MOCKSDKSERVER := $(REGISTRY)/mock-sdk-server:$(MOCKSDKSERVERTAG)

ifndef TAGS
TAGS := daemon
endif

ifndef PKGS
PKGS := $(shell go list ./... 2>&1 | grep -v 'vendor' | grep -v 'sanity')
endif

ifeq ($(BUILD_TYPE),debug)
BUILDFLAGS := -gcflags "-N -l"
endif

ifdef HAVE_BTRFS
TAGS+=btrfs_noversion have_btrfs
endif

ifdef HAVE_CHAINFS
TAGS+=have_chainfs
endif

ifndef PROTOC
PROTOC = protoc
endif

ifndef PROTOS_PATH
PROTOS_PATH = $(GOPATH)/src
endif

ifndef PROTOSRC_PATH
PROTOSRC_PATH = $(PROTOS_PATH)/github.com/libopenstorage/openstorage
endif

OSDSANITY:=cmd/osd-sanity/osd-sanity

.PHONY: \
	all \
	deps \
	update-deps \
	test-deps \
	update-test-deps \
	vendor-update \
	vendor-without-update \
	vendor \
	build \
	install \
	proto \
	lint \
	vet \
	packr \
	errcheck \
	pretest \
	test \
	docs \
	docker-build-osd-dev \
	docker-build \
	docker-test \
	docker-build-osd-internal \
	docker-build-osd \
	launch \
	launch-local-btrfs \
	install-flexvolume-plugin \
	$(OSDSANITY)-install \
	$(OSDSANITY)-clean \
	clean \
	generate \
	generate-mockfiles \
	e2e \
	verify \
	sdk-check-version


all: build $(OSDSANITY)

# TOOLS build rules
#
$(GOPATH)/bin/golint:
	@echo "Installing missing $@ ..."
	go get -u github.com/golang/lint/golint

$(GOPATH)/bin/errcheck:
	@echo "Installing missing $@ ..."
	go get -u github.com/kisielk/errcheck

$(GOPATH)/bin/protoc-gen-go:
	@echo "Installing missing $@ ..."
	go get -u github.com/golang/protobuf/protoc-gen-go

$(GOPATH)/bin/protoc-gen-grpc-gateway:
	@echo "Installing missing $@ ..."
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

$(GOPATH)/bin/protoc-gen-swagger:
	@echo "Installing missing $@ ..."
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

$(GOPATH)/bin/govendor:
	@echo "Installing missing $@ ..."
	go get -u github.com/kardianos/govendor

$(GOPATH)/bin/packr:
	@echo "Installing missing $@ ..."
	go get -u github.com/gobuffalo/packr/...

$(GOPATH)/bin/cover:
	@echo "Installing missing $@ ..."
	go get -u golang.org/x/tools/cmd/cover

$(GOPATH)/bin/gotestcover:
	@echo "Installing missing $@ ..."
	go get -u github.com/pierrre/gotestcover

# DEPS build rules
#

deps:
	go get -d -v $(PKGS)

update-deps:
	go get -d -v -u -f $(PKGS)

test-deps:
	go get -d -v -t $(PKGS)

update-test-deps:
	go get -tags "$(TAGS)" -d -v -t -u -f $(PKGS)

vendor-update:
	GOOS=linux GOARCH=amd64 go get -tags "daemon btrfs_noversion have_btrfs have_chainfs" -d -v -t -u -f $(PKGS)

vendor-without-update: $(GOPATH)/bin/govendor
	rm -rf vendor
	govendor init
	GOOS=linux GOARCH=amd64 govendor add +external
	GOOS=linux GOARCH=amd64 govendor update +vendor
	GOOS=linux GOARCH=amd64 govendor add +external
	GOOS=linux GOARCH=amd64 govendor update +vendor

vendor: vendor-update vendor-without-update

build: packr
	go build -tags "$(TAGS)" $(BUILDFLAGS) $(PKGS)

install: packr $(OSDSANITY)-install
	go install -gcflags="all=-N -l" -tags "$(TAGS)" $(PKGS)
	go install github.com/libopenstorage/openstorage/cmd/osd-token-generator

$(OSDSANITY):
	@$(MAKE) -C cmd/osd-sanity

$(OSDSANITY)-install:
	@$(MAKE) -C cmd/osd-sanity install

$(OSDSANITY)-clean:
	@$(MAKE) -C cmd/osd-sanity clean

docker-build-proto:
	docker build -t quay.io/openstorage/osd-proto -f Dockerfile.proto .

docker-proto: $(GOPATH)/bin/protoc-gen-go
	docker run \
		--privileged --rm \
		-v $(shell pwd):/go/src/github.com/libopenstorage/openstorage \
		-e "GOPATH=/go" \
		-e "DOCKER_PROTO=yes" \
		-e "PATH=/bin:/usr/bin:/usr/local/bin:/go/bin" \
		quay.io/openstorage/osd-proto \
			make proto mockgen

proto: $(GOPATH)/bin/protoc-gen-go $(GOPATH)/bin/protoc-gen-grpc-gateway $(GOPATH)/bin/protoc-gen-swagger
ifndef DOCKER_PROTO
	$(error Do not run directly. Run 'make docker-proto' instead.)
endif

	@echo ">>> Generating protobuf definitions from api/api.proto"
	$(PROTOC) -I $(PROTOSRC_PATH) \
		-I /usr/local/include \
		-I $(PROTOS_PATH)/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		$(PROTOSRC_PATH)/api/api.proto
	$(PROTOC) -I $(PROTOSRC_PATH) \
		-I /usr/local/include \
		-I $(PROTOS_PATH)/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		$(PROTOSRC_PATH)/api/api.proto
	$(PROTOC) -I $(PROTOSRC_PATH) \
		-I /usr/local/include \
		-I $(PROTOS_PATH)/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:$(PROTOSRC_PATH)/api/server/sdk \
		$(PROTOSRC_PATH)/api/api.proto
	@echo ">>> Upgrading swagger 2.0 to openapi 3.0"
	mv api/server/sdk/api/api.swagger.json api/server/sdk/api/20api.swagger.json
	swagger2openapi api/server/sdk/api/20api.swagger.json -o api/server/sdk/api/api.swagger.json
	rm -f api/server/sdk/api/20api.swagger.json
	@echo ">>> Generating grpc protobuf definitions from pkg/flexvolume/flexvolume.proto"
	$(PROTOC) -I/usr/local/include -I$(PROTOSRC_PATH) -I$(PROTOS_PATH)/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. $(PROTOSRC_PATH)/pkg/flexvolume/flexvolume.proto
	$(PROTOC) -I/usr/local/include -I$(PROTOSRC_PATH) -I$(PROTOS_PATH)/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. $(PROTOSRC_PATH)/pkg/flexvolume/flexvolume.proto
	@echo ">>> Generating protobuf definitions from pkg/jsonpb/testing/testing.proto"
	$(PROTOC) -I $(PROTOSRC_PATH) $(PROTOSRC_PATH)/pkg/jsonpb/testing/testing.proto --go_out=plugins=grpc:.
	@echo ">>> Updating SDK versions"
	go run tools/sdkver/sdkver.go --swagger api/server/sdk/api/api.swagger.json

lint: $(GOPATH)/bin/golint
	golint $(PKGS)

vet:
	go vet $(PKGS)

errcheck: $(GOPATH)/bin/errcheck
	errcheck -tags "$(TAGS)" $(PKGS)

pretest: lint vet errcheck

install-sdk-test:
ifndef HAS_SDKTEST
	@echo "Installing sdk-test"
	@-go get -d -u github.com/libopenstorage/sdk-test 1>/dev/null 2>&1
	@(cd $(GOPATH)/src/github.com/libopenstorage/sdk-test/cmd/sdk-test && make install )
endif

test-sdk: install-sdk-test launch-sdk
	timeout 30 sh -c 'until curl --silent -X GET -d {} http://localhost:9110/v1/clusters/inspectcurrent | grep STATUS_OK; do sleep 1; done'
	sdk-test -ginkgo.noColor -ginkgo.noisySkippings=false -sdk.endpoint=localhost:9100 -sdk.cpg=$(GOPATH)/src/github.com/libopenstorage/sdk-test/cmd/sdk-test/cb.yaml

test: packr
	go test -tags "$(TAGS)" $(TESTFLAGS) $(PKGS)

docs:
	go generate ./cmd/osd/main.go

packr: $(GOPATH)/bin/packr
	packr clean
	packr

generate-mockfiles:
	go generate $(PKGS)

generate: docs generate-mockfiles

sdk: docker-proto docker-build-mock-sdk-server

docker-build-mock-sdk-server: packr
	rm -rf _tmp
	mkdir -p _tmp
	CGO_ENABLED=0 GOOS=linux go build \
				-a -ldflags '-extldflags "-static"' \
				-tags "$(TAGS)" \
				-o ./_tmp/osd \
				./cmd/osd
	docker build -t $(IMAGE_MOCKSDKSERVER) -f Dockerfile.sdk .
	rm -rf _tmp

docker-build-osd-dev-base:
	docker build -t quay.io/openstorage/osd-dev-base -f Dockerfile.osd-dev-base .

push-mock-sdk-server: docker-build-mock-sdk-server
	docker push $(IMAGE_MOCKSDKSERVER)

docker-build-osd-dev:
	# This image is local only and will not be pushed
	docker build -t openstorage/osd-dev -f Dockerfile.osd-dev .

docker-build: docker-build-osd-dev
	docker run \
		--privileged \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-e AWS_ACCESS_KEY_ID \
		-e AWS_SECRET_ACCESS_KEY \
		-e "TAGS=$(TAGS)" \
		-e "PKGS=$(PKGS)" \
		-e "BUILDFLAGS=$(BUILDFLAGS)" \
		openstorage/osd-dev \
			make build

docker-test: docker-build-osd-dev
	docker run \
		--privileged \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v /mnt:/mnt \
		-e AWS_REGION \
		-e AWS_ZONE \
		-e AWS_INSTANCE_NAME \
		-e AWS_ACCESS_KEY_ID \
		-e AWS_SECRET_ACCESS_KEY \
		-e GOOGLE_APPLICATION_CREDENTIALS \
		-e GCE_INSTANCE_NAME \
		-e GCE_INSTANCE_ZONE \
		-e GCE_INSTANCE_PROJECT \
		-e AZURE_INSTANCE_NAME \
		-e AZURE_SUBSCRIPTION_ID \
		-e AZURE_RESOURCE_GROUP_NAME \
		-e AZURE_ENVIRONMENT \
		-e AZURE_TENANT_ID \
		-e AZURE_CLIENT_ID \
		-e AZURE_CLIENT_SECRET \
		-e "TAGS=$(TAGS)" \
		-e "PKGS=$(PKGS)" \
		-e "BUILDFLAGS=$(BUILDFLAGS)" \
		-e "TESTFLAGS=$(TESTFLAGS)" \
		-e "GO111MODULE=auto" \
		openstorage/osd-dev \
			make test

docker-build-osd-internal:
	rm -rf _tmp
	mkdir -p _tmp
	go build -a -tags "$(TAGS)" -o _tmp/osd cmd/osd/main.go
	docker build -t quay.io/openstorage/osd -f Dockerfile.osd .

docker-build-osd: docker-build-osd-dev
	docker run \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-e "TAGS=$(TAGS)" \
		-e "PKGS=$(PKGS)" \
		-e "BUILDFLAGS=$(BUILDFLAGS)" \
		openstorage/osd-dev \
			make docker-build-osd-internal

launch-sdk-quick:
	@-docker stop sdk > /dev/null 2>&1
	docker run --rm --name sdk \
		-d -p 9110:9110 -p 9100:9100 \
		$(IMAGE_MOCKSDKSERVER)

launch-sdk: sdk launch-sdk-quick

launch: docker-build-osd
	docker run \
		--privileged \
		-d \
		-v $(shell pwd)/etc:/etc \
		-v /run/docker/plugins:/run/docker/plugins \
		-v /var/lib/osd/:/var/lib/osd/ \
		-p 9005:9005 \
		-p 9100:9100 \
		-p 9110:9110 \
		quay.io/openstorage/osd -d -f /etc/config/config.yaml

# must set HAVE_BTRFS
launch-local-btrfs: install
	sudo bash -x etc/btrfs/init.sh
	sudo $(shell which osd) -d -f etc/config/config_btrfs.yaml

install-flexvolume:
	go install -a -tags "$(TAGS)" github.com/libopenstorage/openstorage/pkg/flexvolume github.com/libopenstorage/openstorage/pkg/flexvolume/cmd/flexvolume

install-flexvolume-plugin: install-flexvolume
	sudo rm -rf /usr/libexec/kubernetes/kubelet/volume/exec-plugins/openstorage~openstorage
	sudo mkdir -p /usr/libexec/kubernetes/kubelet/volume/exec-plugins/openstorage~openstorage
	sudo chmod 777 /usr/libexec/kubernetes/kubelet/volume/exec-plugins/openstorage~openstorage
	cp $(GOPATH)/bin/flexvolume /usr/libexec/kubernetes/kubelet/volume/exec-plugins/openstorage~openstorage/openstorage

clean: $(OSDSANITY)-clean
	go clean -i $(PKGS)
	packr clean

# Generate test-coverage HTML report
# - note: the 'go test -coverprofile...' does append results, so we're merging individual pkgs in for-loop
coverage: packr $(GOPATH)/bin/cover $(GOPATH)/bin/gotestcover
	gotestcover -coverprofile=coverage.out $(PKGS)
	go tool cover -html=coverage.out -o coverage.html
	@echo "INFO: Summary of coverage"
	go tool cover -func=coverage.out
	@cp coverage.out coverage.html /mnt/ && \
	echo "INFO: libopenstorage coverage saved at /mnt/coverage.{html,out}"

docker-coverage: docker-build-osd-dev
	docker run \
		--privileged \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v /mnt:/mnt \
		-e AWS_ACCESS_KEY_ID \
		-e AWS_SECRET_ACCESS_KEY \
		-e "TAGS=$(TAGS)" \
		-e "PKGS=$(PKGS)" \
		-e "BUILDFLAGS=$(BUILDFLAGS)" \
		-e "TESTFLAGS=$(TESTFLAGS)" \
		openstorage/osd-dev \
			make coverage

docker-images: docker-build-proto docker-build-osd-dev-base
push-docker-images: docker-images
	docker push quay.io/openstorage/osd-dev-base
	docker push quay.io/openstorage/osd-proto

# This needs to be adjusted for each release branch according
# to the SDK Version.
# For master (until released), major should be 0 and patch should be 0.
# For release branches, major and minor should be frozen.
#
# If you think you need to adjust this, you are doing something wrong.
sdk-check-version:
	go run tools/sdkver/sdkver.go --check-major=0 --check-minor=101

mockgen:
	go get github.com/golang/mock/gomock
	go get github.com/golang/mock/mockgen
	mockgen -destination=api/mock/mock_storagepool.go -package=mock github.com/libopenstorage/openstorage/api OpenStoragePoolServer,OpenStoragePoolClient
	mockgen -destination=api/mock/mock_cluster.go -package=mock github.com/libopenstorage/openstorage/api OpenStorageClusterServer,OpenStorageClusterClient
	mockgen -destination=api/mock/mock_node.go -package=mock github.com/libopenstorage/openstorage/api OpenStorageNodeServer,OpenStorageNodeClient
	mockgen -destination=api/mock/mock_diags.go -package=mock github.com/libopenstorage/openstorage/api OpenStorageDiagsServer,OpenStorageDiagsClient
	mockgen -destination=api/mock/mock_volume.go -package=mock github.com/libopenstorage/openstorage/api OpenStorageVolumeServer,OpenStorageVolumeClient
	mockgen -destination=cluster/mock/cluster.mock.go -package=mock github.com/libopenstorage/openstorage/cluster Cluster
	mockgen -destination=api/mock/mock_fstrim.go -package=mock github.com/libopenstorage/openstorage/api OpenStorageFilesystemTrimServer,OpenStorageFilesystemTrimClient
	mockgen -destination=api/mock/mock_fscheck.go -package=mock github.com/libopenstorage/openstorage/api OpenStorageFilesystemCheckServer,OpenStorageFilesystemCheckClient
	mockgen -destination=api/server/mock/mock_schedops_k8s.go -package=mock github.com/portworx/sched-ops/k8s/core Ops

e2e: docker-build-osd
	cd test && ./run.bash

verify: vet sdk-check-version docker-test e2e
