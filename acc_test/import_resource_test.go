package test

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestAccCluster_importBasic(t *testing.T) {

	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
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
				Config:            oriConfig,
				ResourceName:      "instaclustr_cluster.valid",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestAccKafkaCluster_importBasic(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid_kafka.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
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
				Config:            oriConfig,
				ResourceName:      "instaclustr_cluster.valid",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccEncryptionKey_importBasic(t *testing.T) {
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
			{
				Config:            oriConfig,
				ResourceName:      "instaclustr_encryption_key.valid",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccFirewallRule_importBasic(t *testing.T) {
	testProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	tfFile, _ := ioutil.ReadFile("data/valid_with_firewall_rule.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	config := fmt.Sprintf(string(tfFile), username, apiKey, hostname)

	resource.Test(t, resource.TestCase{
		Providers:    testProviders,
		CheckDestroy: checkFirewallRuleDeleted(hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					checkFirewallRuleState,
					checkFirewallRuleCreated(hostname, username, apiKey),
				),
			},
			{
				Config:            config,
				ResourceName:      "instaclustr_firewall_rule.valid_with_firewall_rule",
				ImportState:       true,
				ImportStateIdFunc: testAccFirewallRuleImportStateIdFunc("instaclustr_firewall_rule.valid_with_firewall_rule"),
				ImportStateVerify: true,
			},
		},
	})
}

func testAccFirewallRuleImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		return fmt.Sprintf("%s&%s", rs.Primary.Attributes["cluster_id"], rs.Primary.Attributes["rule_cidr"]), nil
	}
}

func TestKafkaUserResource_importBasic(t *testing.T) {
	testProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}

	configBytes1, _ := ioutil.ReadFile("data/kafka_user_create_cluster.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)

	zookeeperNodeSize := "zk-developer-t3.small-20"

	createClusterConfig := fmt.Sprintf(string(configBytes1), username, apiKey, hostname, zookeeperNodeSize)
	validResizeConfig := strings.Replace(createClusterConfig, `t3.medium-80-gp2`, `r5.xlarge-800-gp2`, 1)
	invalidResizeConfig := strings.Replace(createClusterConfig, `t3.medium-80-gp2`, `t3.small-20-gp2`, 1)
	resourceName := "kafka_cluster"

	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: createClusterConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValidKafka("instaclustr_cluster.kafka_cluster"),
					checkClusterRunning("kafka_cluster", hostname, username, apiKey),
					testCheckContactIPCorrect("kafka_cluster", hostname, username, apiKey, 3, 3),
				),
			},
			{
				Config:            createClusterConfig,
				ResourceName:      "instaclustr_cluster.kafka_cluster",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// These are are cluster level attributes.
					"bundle",
					"cluster_provider",
					"pci_compliant_cluster",
					"rack_allocation",
					// wait_for_state is a creation option to ensure IP contact points are ready for other parts of the infrastructure. It cannot be imported.
					"wait_for_state",
				},
			},
			{
				PreConfig: func() {
					time.Sleep(6 * time.Minute)
				},
				Config:      invalidResizeConfig,
				ExpectError: regexp.MustCompile("t3.small-20-gp2"),
			},
			{
				Config: validResizeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("instaclustr_cluster."+resourceName, "cluster_name", "example_kafka_tf_test"),
					testCheckClusterResize(resourceName, hostname, username, apiKey, "r5.xlarge-800-gp2"),
				),
			},
		},
	})
}

func TestAccVpcPeering_importBasic(t *testing.T) {
	testProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	tfFile, _ := ioutil.ReadFile("data/valid_with_vpc_peering.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	config := fmt.Sprintf(string(tfFile), username, apiKey, hostname)
	resource.Test(t, resource.TestCase{
		Providers:    testProviders,
		CheckDestroy: checkVpcPeeringDeleted(hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					checkVpcPeeringState,
					checkVpcPeeringCreated(hostname, username, apiKey),
				),
			},
			{
				Config:            config,
				ResourceName:      "instaclustr_vpc_peering.valid_with_vpc_peering",
				ImportState:       true,
				ImportStateIdFunc: testAccVpcPeeringImportStateIdFunc("instaclustr_vpc_peering.valid_with_vpc_peering"),
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVpcPeeringImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		return fmt.Sprintf("%s&%s", rs.Primary.Attributes["cluster_id"], rs.Primary.Attributes["vpc_peering_id"]), nil
	}
}
