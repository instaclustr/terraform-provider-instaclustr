
BIN_NAME="terraform-provider-instaclustr"


# for VERSION, don't add prefix "v", e.g., use "1.9.8" instead of "v1.9.8" as it could break circleCI stuff
VERSION=1.9.9
INSTALL_FOLDER=$(HOME)/.terraform.d/plugins/terraform.instaclustr.com/instaclustr/instaclustr/$(VERSION)/darwin_amd64


.PHONY: install clean all build test testacc testtarget release_version
release_version:
	@echo v$(VERSION)

all: build

clean:
	rm $(BIN_NAME)_v$(VERSION)
	rm -fr vendor

build:
	go build $(FLAGS) -o bin/$(BIN_NAME)_v$(VERSION) main.go

test:
	cd instaclustr && go test -v -timeout 120m -count=1 -coverprofile coverage.out ./...

testacc:
ifndef IC_USERNAME
	@echo "IC_USERNAME for provisioning API must be set for acceptance tests"
	@exit 1
endif
ifndef IC_API_KEY
	@echo "IC_API_KEY for provisioning API must be set for acceptance tests"
	@exit 1
endif
ifndef KMS_ARN
	@echo "KMS_ARN for provisioning API must be set for acceptance tests"
	@exit 1
endif
ifndef IC_PROV_ACC_NAME
	@echo "IC_PROV_ACC_NAME for provisioning API must be set for acceptance tests"
	@exit 1
endif
ifndef IC_PROV_VPC_ID
	@echo "IC_PROV_VPC_ID for provisioning API must be set for acceptance tests"
	@exit 1
endif
	cd acc_test && TF_ACC=1 go test -v -timeout 120m -count=1


install:
	@if [ ! -d $(INSTALL_FOLDER) ]; then \
		echo "$(INSTALL_FOLDER) doesn't exist, creating..."; \
		mkdir -p $(INSTALL_FOLDER); \
	fi
	cp ./bin/$(BIN_NAME)_v$(VERSION) $(INSTALL_FOLDER)
