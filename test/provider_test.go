package test

import (
    "testing"

    "github.com/hashicorp/terraform/terraform"
    "github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
)

func TestProvider(t *testing.T) {
    if err := instaclustr.Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = instaclustr.Provider()
}
