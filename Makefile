BIN_NAME="terraform-provider-instaclustr"
VERSION=v1.3.0

.PHONY: install clean all build test testacc

all: build

clean:
	rm $(BIN_NAME)_$(VERSION)
	rm -fr vendor

build:
	go build -o bin/$(BIN_NAME)_$(VERSION) main.go

test:
	cd test && go test -v -timeout 120m -count=1

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
	cd test && TF_ACC=1 go test -v -timeout 120m -count=1

install:
	@if [ ! -d "$(HOME)/.terraform.d/plugins/" ]; then \
		echo "$(HOME)/.terraform.d/plugins/ doesn't exist, creating..."; \
		mkdir -p $(HOME)/.terraform.d/plugins/; \
	fi
	cp ./bin/$(BIN_NAME)_$(VERSION) $(HOME)/.terraform.d/plugins/
