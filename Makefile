BIN_NAME="terraform-provider-instaclustr"
VERSION=0.0.1

.PHONY: install clean all build test testacc

all: build

clean:
	rm $(BIN_NAME)_v$(VERSION)
	rm -fr vendor

build:
	go build -o $(BIN_NAME)_v$(VERSION) main.go

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
	cd test && TF_ACC=1 go test -v -timeout 120m -count=1

install:
	@if [ ! -d "$(HOME)/.terraform.d/plugins/" ]; then \
		echo "$(HOME)/.terraform.d/plugins/ doesn't exist, creating..."; \
		mkdir -p $(HOME)/.terraform.d/plugins/; \
	fi
	cp $(BIN_NAME)_v$(VERSION) $(HOME)/.terraform.d/plugins/
