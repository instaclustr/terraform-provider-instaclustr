package instaclustr

import (
    "testing"
    "github.com/hashicorp/terraform/terraform"
)

func TestProvider(t *testing.T) {

	//Testing if SonarCloud does something (Here)
	if true {
		print("test1")
	}
	if true {
		print("test2")
	}
	if true {
		print("test3")
	}
	if true {
		print("test4")
	}
	if true {
		print("test5")
	}
	if true {
		print("test1")
	}
	if true {
		print("test2")
	}
	if true {
		print("test3")
	}
	if true {
		print("test4")
	}
	if true {
		print("test5")
	}
	if true {
		print("test1")
	}
	if true {
		print("test2")
	}
	if true {
		print("test3")
	}
	if true {
		print("test4")
	}
	if true {
		print("test5")
	}
	if true {
		print("test1")
	}
	if true {
		print("test2")
	}
	if true {
		print("test3")
	}
	if true {
		print("test4")
	}
	if true {
		print("test5")
	}
	if true {
		print("test1")
	}
	if true {
		print("test2")
	}
	if true {
		print("test3")
	}
	if true {
		print("test4")
	}
	if true {
		print("test5")
	}

    if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}
