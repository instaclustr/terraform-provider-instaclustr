package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
)

func TestAccEBSKey(t *testing.T) {
	testAccEBSKeyProvider := instaclustr.Provider()
	testAccEBSKeyProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccEBSKeyProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid_encryption_key.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	providerAccountName := os.Getenv("IC_PROV_ACC_NAME")
	kmsArn := os.Getenv("KMS_ARN")
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname, kmsArn, providerAccountName)
	resource.Test(t, resource.TestCase{
		Providers:    testAccEBSKeyProviders,
		CheckDestroy: testCheckAccEBSResourceDeleted("valid", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckAccEBSResourceValid("valid"),
					testCheckAccEBSResourceCreated("valid", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestAccEBSKeyInvalid(t *testing.T) {
	testAccEBSKeyProvider := instaclustr.Provider()
	testAccEBSKeyProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccEBSKeyProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/invalid_encryption_key.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	kmsArn := os.Getenv("KMS_ARN")
	resource.Test(t, resource.TestCase{
		Providers: testAccEBSKeyProviders,
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, apiKey, hostname, kmsArn),
				ExpectError: regexp.MustCompile("Error adding encryption key"),
			},
		},
	})
}

func testCheckAccEBSResourceValid(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_encryption_key."+resourceName]
		if resourceState == nil {
			return fmt.Errorf("%s: resource not found in state", resourceName)
		}

		instanceState := resourceState.Primary
		if instanceState == nil {
			return fmt.Errorf("resource has no primary instance")
		}
		return nil
	}
}

func testCheckAccEBSResourceCreated(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_encryption_key."+resourceName]
		id := resourceState.Primary.Attributes["key_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		resource, err := client.ReadEncryptionKey(id)
		if err != nil {
			return fmt.Errorf("Failed to read encryption key %s: %s", id, err)
		}
		if resource.ID != id {
			return fmt.Errorf("Encryption key expected %s but got %s", id, resource.ID)
		}
		return nil
	}
}

func testCheckAccEBSResourceDeleted(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_encryption_key."+resourceName]
		id := resourceState.Primary.Attributes["key_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		err := client.DeleteEncryptionKey(id)
		if err == nil {
			return fmt.Errorf("Encryption key %s still exists", id)
		}
		return nil
	}
}
