
BIN_NAME=terraform-provider-instaclustr


# for VERSION, don't add prefix "v", e.g., use "1.9.8" instead of "v1.9.8" as it could break circleCI stuff


VERSION=1.24.0

INSTALL_FOLDER=$(HOME)/.terraform.d/plugins/terraform.instaclustr.com/instaclustr/instaclustr/$(VERSION)/darwin_amd64

.PHONY: install clean all build test testacc testtarget release_version
release_version:
	@echo v$(VERSION)

all: build

clean:
	-rm -rf bin/$(BIN_NAME)_v$(VERSION)
	-rm -rf $(INSTALL_FOLDER)

build:
	go build $(FLAGS) -o bin/$(BIN_NAME)_v$(VERSION) main.go

test:
	cd instaclustr && go test -v -timeout 120m -count=1 -coverprofile coverage.out -json ./... > report.json
	@cd instaclustr && cat report.json | sed -n '/Output/p' | jq '.Output' # Prettify the report.json file to print it to stdout

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
ifndef IC_AWS_ACCESS_KEY
	@echo "IC_AWS_ACCESS_KEY must be set for acceptance tests (Kafka Connect custom connector bucket)"
	@exit 1
endif
ifndef IC_AWS_SECRET_KEY
	@echo "IC_AWS_SECRET_KEY must be set for acceptance test (Kafka Connect custom connector bucket)"
	@exit 1
endif
ifndef IC_S3_BUCKET_NAME
	@echo "IC_S3_BUCKET_NAME must be set for acceptance test (Kafka Connect custom connector bucket)"
	@exit 1
endif
ifndef IC_AZURE_STORAGE_ACCOUNT_NAME
	@echo "IC_AZURE_STORAGE_ACCOUNT_NAME must be set for acceptance test (Kafka Connect custom connector bucket)"
	@exit 1
endif
ifndef IC_AZURE_STORAGE_ACCOUNT_KEY
	@echo "IC_AZURE_STORAGE_ACCOUNT_KEY must be set for acceptance test (Kafka Connect custom connector bucket)"
	@exit 1
endif
ifndef IC_AZURE_STORAGE_CONTAINER_NAME
	@echo "IC_AZURE_STORAGE_CONTAINER_NAME must be set for acceptance test (Kafka Connect custom connector bucket)"
	@exit 1
endif
	cd acc_test && TF_ACC=1 go test -v -timeout 200m -count=1 -parallel=6



install:
	@if [ ! -d $(INSTALL_FOLDER) ]; then \
		echo "$(INSTALL_FOLDER) doesn't exist, creating..."; \
		mkdir -p $(INSTALL_FOLDER); \
	fi
	cp ./bin/$(BIN_NAME)_v$(VERSION) $(INSTALL_FOLDER)
