package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
)

func TestAccPCICluster(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/pci_valid.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	updatedConfig := strings.Replace(oriConfig, "testcluster", "newcluster", 1)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckPCIResourceDeleted("valid", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckPCIResourceValid("valid"),
					testCheckPCIResourceCreated("valid", hostname, username, apiKey),
				),
			},
			{
				Config:      updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("valid"),
					testCheckResourceCreated("valid", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestAccPCIClusterResize(t *testing.T) {
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	validConfig, _ := ioutil.ReadFile("data/valid_with_resizable_pci_cluster.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	originalConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	validResizeConfig := strings.Replace(originalConfig, "resizeable-small(r5-l)-v2", "resizeable-small(r5-xl)-v2", 1)
	invalidResizeClassConfig := strings.Replace(originalConfig, "resizeable-small(r5-l)-v2", "resizeable-large(r5-xl)-v2", 1)
	invalidResizeConfig := strings.Replace(originalConfig, "resizeable-small(r5-l)-v2", "t3.medium", 1)


	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckPCIResourceDeleted("resizable_pci_cluster", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: originalConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckPCIResourceValid("resizable_pci_cluster"),
					testCheckPCIResourceCreated("resizable_pci_cluster", hostname, username, apiKey),
					checkClusterRunning("resizable_pci_cluster", hostname, username, apiKey),
					testCheckContactIPCorrect("resizable_pci_cluster", hostname, username, apiKey, 3),
				),
			},
			{
				Config: validResizeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("instaclustr_cluster.resizable_pci_cluster", "cluster_name", "tf-resizable-test"),
					resource.TestCheckResourceAttr("instaclustr_cluster.resizable_pci_cluster", "node_size", "resizeable-small(r5-xl)-v2"),
					testCheckClusterResize("resizable_pci_cluster", hostname, username, apiKey, "resizeable-small(r5-xl)-v2"),
				),
			},
			{
				Config:      invalidResizeClassConfig,
				ExpectError: regexp.MustCompile("Cannot resize nodes"),
			},
			{
				Config:      invalidResizeConfig,
				ExpectError: regexp.MustCompile("Cannot resize nodes"),
			},
		},
	})
}

func TestAccPCIClusterInvalid(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/pci_invalid.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, apiKey, hostname),
				ExpectError: regexp.MustCompile("Error creating cluster"),
			},
		},
	})
}

func testCheckPCIResourceValid(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster."+resourceName]
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

func testCheckPCIResourceCreated(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster."+resourceName]
		id := resourceState.Primary.Attributes["cluster_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		cluster, err := client.ReadCluster(id)
		if err != nil {
			return fmt.Errorf("Failed to read cluster %s: %s", id, err)
		}
		if cluster.ID != id {
			return fmt.Errorf("Cluster expected %s but got %s", id, cluster.ID)
		}
		return nil
	}
}

func testCheckPCIResourceDeleted(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster."+resourceName]
		id := resourceState.Primary.Attributes["cluster_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		err := client.DeleteCluster(id)
		if err == nil {
			return fmt.Errorf("Cluster %s still exists.", id)
		}
		return nil
	}
}
