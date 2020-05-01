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


func TestAccCluster(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey)
	updatedConfig := strings.Replace(oriConfig, "testcluster", "newcluster", 1)
	hostname := getoptionalenv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { AccTestEnvVarsCheck(t) },
		CheckDestroy: testCheckResourceDeleted("valid", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("valid"),
					testCheckResourceCreated("valid", hostname, username, apiKey),
				),
			},
			{
				Config:      updatedConfig,
				ExpectError: regexp.MustCompile("The cluster doesn't support update"),
			},
		},
	})
}

func TestAccClusterResize(t *testing.T) {
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	validConfig, _ := ioutil.ReadFile("data/valid_with_resizable_cluster.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey)
	validResizeConfig := strings.Replace(oriConfig, "resizeable-small(r5-l)", "resizeable-small(r5-xl)", 1)
	validResizeConfig = strings.Replace(validResizeConfig, "tf-resizable-test", "tf-resizable-partial-test", 1)
	invalidResizeClassConfig := strings.Replace(oriConfig, "resizeable-small(r5-l)", "resizeable-large(r5-xl)", 1)
	invalidResizeConfig := strings.Replace(oriConfig, "resizeable-small(r5-l)", "t3.medium", 1)

	hostname := getoptionalenv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { AccTestEnvVarsCheck(t) },
		CheckDestroy: testCheckResourceDeleted("resizable_cluster", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("resizable_cluster"),
					testCheckResourceCreated("resizable_cluster", hostname, username, apiKey),
				),
			},
			{
				Config: validResizeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("instaclustr_cluster.resizable_cluster", "cluster_name", "tf-resizable-test"),
					resource.TestCheckResourceAttr("instaclustr_cluster.resizable_cluster", "node_size", "resizeable-small(r5-xl)"),
					testCheckClusterResize(hostname, username, apiKey, "resizeable-small(r5-xl)"),
				),
				ExpectNonEmptyPlan: true,
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

func TestAccClusterInvalid(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/invalid.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { AccTestEnvVarsCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, apiKey),
				ExpectError: regexp.MustCompile("Error creating cluster"),
			},
		},
	})
}

func TestAccClusterInvalidBundleOptionCombo(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/invalid_with_wrong_bundle_option.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { AccTestEnvVarsCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, apiKey),
				ExpectError: regexp.MustCompile("Error creating cluster"),
			},
		},
	})
}

func TestAccClusterCustomVPC(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid_with_custom_vpc.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	providerAccountName := os.Getenv("IC_PROV_ACC_NAME")
	providerVpcId := os.Getenv("IC_PROV_VPC_ID")
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, providerAccountName, providerVpcId)
	hostname := getoptionalenv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { AccTestEnvVarsCheck(t) },
		CheckDestroy: testCheckResourceDeleted("vpc_cluster", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("vpc_cluster"),
					testCheckResourceCreated("vpc_cluster", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestAccClusterCustomVPCInvalid(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/invalid_with_custom_vpc.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	providerAccountName := os.Getenv("IC_PROV_ACC_NAME")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { AccTestEnvVarsCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, apiKey, providerAccountName),
				ExpectError: regexp.MustCompile("Error creating cluster"),
			},
		},
	})
}

func testCheckResourceValid(resourceName string) resource.TestCheckFunc {
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

func testCheckResourceCreated(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
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

func testCheckResourceDeleted(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster."+resourceName]
		id := resourceState.Primary.Attributes["cluster_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		err := client.DeleteCluster(id)
		if err == nil {
			return fmt.Errorf("Cluster %s still exists", id)
		}
		return nil
	}
}

func testCheckClusterResize(hostname, username, apiKey, expectedNodeSize string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster.resizable_cluster"]
		id := resourceState.Primary.Attributes["cluster_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)

		cluster, err := client.ReadCluster(id)
		if err != nil {
			return fmt.Errorf("Failed to read cluster %s: %s", id, err)
		}
		targetNodeSize := cluster.DataCentres[0].ResizeTargetNodeSize
		if targetNodeSize != expectedNodeSize {
			return fmt.Errorf("Expected cluster to be resized to %s", expectedNodeSize)
		}
		return nil
	}
}
