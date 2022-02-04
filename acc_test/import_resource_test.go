package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
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
	resource.ParallelTest(t, resource.TestCase{
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

func AccMultiDcCluster_importBasicTestSteps(t *testing.T, testAccProviders map[string]terraform.ResourceProvider, validConfig []byte) {
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)

	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	resource.ParallelTest(t, resource.TestCase{
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

func TestAccMultiDcCluster_importBasic(t *testing.T) {

	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}

	validConfig, _ := ioutil.ReadFile("data/valid_multi_DC_provisioning.tf")
	AccMultiDcCluster_importBasicTestSteps(t, testAccProviders, validConfig)

	validConfig, _ = ioutil.ReadFile("data/valid_multi_DC_provisioning_2_DC_6_nodes.tf")
	AccMultiDcCluster_importBasicTestSteps(t, testAccProviders, validConfig)

	validConfig, _ = ioutil.ReadFile("data/valid_multi_DC_provisioning_with_different_providers.tf")
	AccMultiDcCluster_importBasicTestSteps(t, testAccProviders, validConfig)
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

	kafkaNodeSize := "KFK-PRD-r6g.large-250"
	kafkaVersion := "2.7.1"

	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname, kafkaNodeSize, kafkaVersion)
	resource.ParallelTest(t, resource.TestCase{
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
	// Not running this test parallelly since we only have 1 test encryption key
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

	resource.ParallelTest(t, resource.TestCase{
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

	kafkaNodeSize := "KFK-DEV-t4g.medium-80"

	kafkaVersion := "2.7.1"

	zookeeperNodeSize := "KDZ-DEV-t4g.small-30"

	createClusterConfig := fmt.Sprintf(string(configBytes1), username, apiKey, hostname, kafkaNodeSize, kafkaVersion, zookeeperNodeSize)
	validResizeConfig := strings.Replace(createClusterConfig, `KFK-DEV-t4g.medium-80`, `KFK-PRD-r6g.xlarge-800`, 1)
	invalidResizeConfig := strings.Replace(createClusterConfig, `KFK-DEV-t4g.medium-80`, `KFK-DEV-t4g.small-30`, 1)
	resourceName := "kafka_cluster"

	resource.ParallelTest(t, resource.TestCase{
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
					fmt.Println("Sleep for 6 minutes to wait for Kafka cluster to be ready for resize.")
					time.Sleep(6 * time.Minute)
				},
				Config:      invalidResizeConfig,
				ExpectError: regexp.MustCompile("There are no suitable replacement modes for cluster data centre"),
			},
			{
				Config: validResizeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("instaclustr_cluster."+resourceName, "cluster_name", "example_kafka_tf_test"),
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
	resource.ParallelTest(t, resource.TestCase{
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

func TestGCPAccCluster_importBasic(t *testing.T) {

	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted("gcp_valid", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("gcp_valid"),
					testCheckResourceCreated("gcp_valid", hostname, username, apiKey),
				),
			},
			{
				Config:            oriConfig,
				ResourceName:      "instaclustr_cluster.gcp_valid",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestGCPAccVpcPeering_importBasic(t *testing.T) {
	testProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	tfFile, _ := ioutil.ReadFile("data/valid_with_vpc_peering.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	config := fmt.Sprintf(string(tfFile), username, apiKey, hostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testProviders,
		CheckDestroy: checkGCPVpcPeeringDeleted(hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					checkGCPVpcPeeringState,
					checkGCPVpcPeeringCreated(hostname, username, apiKey),
				),
			},
			{
				Config:       config,
				ResourceName: "instaclustr_vpc_peering_gcp.gcp_example",

				ImportState:       true,
				ImportStateIdFunc: testAccVpcPeeringImportStateIdFunc("instaclustr_vpc_peering_gcp.gcp_example"),
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
