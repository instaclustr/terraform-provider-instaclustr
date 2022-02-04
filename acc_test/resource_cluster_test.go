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

func AccClusterResourceTestSteps(t *testing.T, testAccProviders map[string]terraform.ResourceProvider, validConfig []byte) {
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)

	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	updatedConfig := strings.Replace(oriConfig, "testcluster", "newcluster", 1)
	newToOldVersionConfig := strings.Replace(updatedConfig, `version = "3.11.8"`, `version = "apache-cassandra-3.11.8.ic4"`, 1)

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
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("valid"),
					testCheckResourceCreated("valid", hostname, username, apiKey),
				),
			},
			{
				Config: newToOldVersionConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("valid"),
				),
				PlanOnly: true,
			},
			{
				Config:  updatedConfig,
				Destroy: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceDeleted("valid", hostname, username, apiKey),
				),
			},
			{
				Config: newToOldVersionConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("valid"),
					testCheckResourceCreated("valid", hostname, username, apiKey),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("valid"),
				),
				PlanOnly: true,
			},
		},
	})
}

func AccGCPClusterResourceTestSteps(t *testing.T, testAccProviders map[string]terraform.ResourceProvider, validConfig []byte) {
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)

	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	updatedConfig := strings.Replace(oriConfig, "testcluster", "newcluster", 1)

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
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("gcp_valid"),
					testCheckResourceCreated("gcp_valid", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestAccClusterSingleDC(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}

	validSingleDCClusterConfig, _ := ioutil.ReadFile("data/valid.tf")
	AccClusterResourceTestSteps(t, testAccProviders, validSingleDCClusterConfig)
}

func TestAccClusterMultiDC(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}

	validMultiDCClusterConfig, _ := ioutil.ReadFile("data/valid_multi_DC_provisioning.tf")
	AccClusterResourceTestSteps(t, testAccProviders, validMultiDCClusterConfig)
}

func TestKafkaConnectClusterCreateInstaclustrAWS(t *testing.T) {
	if v := os.Getenv("IC_TEST_KAFKA_CONNECT"); v == "" {
		t.Skip("Skipping TestKafkaConnectClusterCreateInstaclustrAWS because IC_TEST_KAFKA_CONNECT is not set")
	}
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validKCConfig, _ := ioutil.ReadFile("data/valid_kafka_connect_instaclustr_aws.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	kafkaClusterId := os.Getenv("IC_TARGET_KAFKA_CLUSTER_ID")
	awsAccessKey := os.Getenv("IC_AWS_ACCESS_KEY")
	awsSecretKey := os.Getenv("IC_AWS_SECRET_KEY")
	S3BucketName := os.Getenv("IC_S3_BUCKET_NAME")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriKCConfig := fmt.Sprintf(string(validKCConfig), username, apiKey, hostname, kafkaClusterId, awsAccessKey, awsSecretKey, S3BucketName)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted("validKC", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriKCConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("validKC"),
					testCheckResourceCreated("validKC", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestKafkaConnectClusterCreateNonInstaclustrAZURE(t *testing.T) {
	if v := os.Getenv("IC_TEST_KAFKA_CONNECT"); v == "" {
		t.Skip("Skipping TestKafkaConnectClusterCreateNonInstaclustrAZURE because IC_TEST_KAFKA_CONNECT is not set")
	}
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validKCConfig, _ := ioutil.ReadFile("data/valid_kafka_connect_non_instaclustr_azure.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	azureStorageAccountName := os.Getenv("IC_AZURE_STORAGE_ACCOUNT_NAME")
	azureStorageAccountKey := os.Getenv("IC_AZURE_STORAGE_ACCOUNT_KEY")
	azureStorageContainerName := os.Getenv("IC_AZURE_STORAGE_CONTAINER_NAME")
	sslEnabledProtocols := os.Getenv("IC_SSL_ENABLED_PROTOCOLS")
	sslTruststorePassword := os.Getenv("IC_SSL_TRUSTSTORE_PASSWORD")
	sslProtocol := os.Getenv("IC_SSL_PROTOCOL")
	securityProtocol := os.Getenv("IC_SECURITY_PROTOCOL")
	saslMechanism := os.Getenv("IC_SASL_MECHANISM")
	saslJaasConfig := os.Getenv("IC_SASL_JAAS_CONFIG")
	bootstrapServers := os.Getenv("IC_BOOTSTRAP_SERVER")
	truststore := os.Getenv("IC_TRUSTSTORE")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriKCConfig := fmt.Sprintf(string(validKCConfig), username, apiKey, hostname, azureStorageAccountName,
		azureStorageAccountKey, azureStorageContainerName, sslEnabledProtocols, sslTruststorePassword,
		sslProtocol, securityProtocol, saslMechanism, saslJaasConfig, bootstrapServers, truststore)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted("validKC", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriKCConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("validKC"),
					testCheckResourceCreated("validKC", hostname, username, apiKey),
				),
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
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	resourceName := "resizable_cluster"
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	validResizeConfig := strings.Replace(oriConfig, "resizeable-small(r5-l)-v2", "resizeable-small(r5-xl)-v2", 1)
	invalidResizeClassConfig := strings.Replace(oriConfig, "resizeable-small(r5-l)-v2", "resizeable-large(r5-xl)-v2", 1)
	invalidResizeConfig := strings.Replace(oriConfig, "resizeable-small(r5-l)-v2", "t3.medium", 1)

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted(resourceName, hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid(resourceName),
					testCheckResourceCreated(resourceName, hostname, username, apiKey),
					checkClusterRunning(resourceName, hostname, username, apiKey),
					resource.TestCheckResourceAttr("instaclustr_cluster.resizable_cluster", "data_centre_custom_name", "AWS_VPC_US_EAST_1_name"),
					testCheckContactIPCorrect(resourceName, hostname, username, apiKey, 3, 3),
				),
			},
			{
				Config: validResizeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("instaclustr_cluster.resizable_cluster", "cluster_name", "tf-resizable-test"),
					resource.TestCheckResourceAttr("instaclustr_cluster.resizable_cluster", "node_size", "resizeable-small(r5-xl)-v2"),
					testCheckClusterResize("resizable_cluster", hostname, username, apiKey, "resizeable-small(r5-xl)-v2"),
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

func TestAccClusterInvalid(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/invalid.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, apiKey, hostname),
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
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, apiKey, hostname),
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
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	providerAccountName := os.Getenv("IC_PROV_ACC_NAME")
	providerVpcId := os.Getenv("IC_PROV_VPC_ID")
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname, providerAccountName, providerVpcId)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
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
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	providerAccountName := os.Getenv("IC_PROV_ACC_NAME")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, hostname, apiKey, providerAccountName),
				ExpectError: regexp.MustCompile("Error creating cluster"),
			},
		},
	})
}

func TestAccElasticsearchClusterResize(t *testing.T) {
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	validConfig, _ := ioutil.ReadFile("data/valid_with_resizable_elasticsearch_cluster.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	resourceName := "resizable_cluster"
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	validResizeConfig := strings.Replace(oriConfig, `kibana_node_size = "t3.small-v2",`, `kibana_node_size = "m5xl-400-v2",`, 1)
	invalidResizeConfig := strings.Replace(oriConfig, `kibana_node_size = "t3.small-v2",`, `kibana_node_size = "t3.small",`, 1)

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted(resourceName, hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid(resourceName),
					testCheckResourceCreated(resourceName, hostname, username, apiKey),
					checkClusterRunning(resourceName, hostname, username, apiKey),
					testCheckContactIPCorrect(resourceName, hostname, username, apiKey, 3, 3),
				),
			},
			{
				PreConfig: func() {
					fmt.Println("Sleep for 1 minutes to wait for Elasticsearch cluster to be ready for resize.")
					time.Sleep(1 * time.Minute)
				},
				Config:      invalidResizeConfig,
				ExpectError: regexp.MustCompile("t3.small"),
			},
			{
				Config: validResizeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("instaclustr_cluster.resizable_cluster", "cluster_name", "tf-resizable-test"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOpenSearchClusterResize(t *testing.T) {
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	validConfig, _ := ioutil.ReadFile("data/valid_with_resizable_opensearch_cluster.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	resourceName := "resizable_cluster"
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	validResizeConfig := strings.Replace(oriConfig, `opensearch_dashboards_node_size = "t3.small-v2",`, `opensearch_dashboards_node_size = "m5xl-400-v2",`, 1)
	invalidResizeConfig := strings.Replace(oriConfig, `opensearch_dashboards_node_size = "t3.small-v2",`, `opensearch_dashboards_node_size = "t3.small",`, 1)

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted(resourceName, hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid(resourceName),
					testCheckResourceCreated(resourceName, hostname, username, apiKey),
					checkClusterRunning(resourceName, hostname, username, apiKey),
					testCheckContactIPCorrect(resourceName, hostname, username, apiKey, 3, 3),
				),
			},
			{
				PreConfig: func() {
					fmt.Println("Sleep for 1 minutes to wait for OpenSearch cluster to be ready for resize.")
					time.Sleep(1 * time.Minute)
				},
				Config:      invalidResizeConfig,
				ExpectError: regexp.MustCompile("t3.small"),
			},
			{
				Config: validResizeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("instaclustr_cluster.resizable_cluster", "cluster_name", "tf-resizable-test"),
				),
				ExpectNonEmptyPlan: true,
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

func testCheckClusterResize(resourceName, hostname, username, apiKey, expectedNodeSize string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster."+resourceName]
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

func TestValidPostgresqlClusterCreate(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid_postgresql_cluster_create.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted("validPostgresql", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("validPostgresql"),
					testCheckResourceCreated("validPostgresql", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestValidElasticsearchClusterCreate(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid_elasticsearch_cluster_create.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted("validElasticsearch", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("validElasticsearch"),
					testCheckResourceCreated("validElasticsearch", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestValidOpenSearchClusterCreate(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid_opensearch_cluster_create.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted("validOpenSearch", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("validOpenSearch"),
					testCheckResourceCreated("validOpenSearch", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestInvalidElasticsearchClusterCreate(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	invalidConfig, _ := ioutil.ReadFile("data/invalid_elasticsearch_cluster_create.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(invalidConfig), username, apiKey, hostname),
				ExpectError: regexp.MustCompile("When 'dedicated_master_nodes' is not true , data_node_size can be either null or equal to master_node_size."),
			},
		},
	})
}

func TestInvalidOpenSearchClusterCreate(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	invalidConfig, _ := ioutil.ReadFile("data/invalid_opensearch_cluster_create.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(invalidConfig), username, apiKey, hostname),
				ExpectError: regexp.MustCompile("When 'dedicated_master_nodes' is not true , data_node_size can be either null or equal to master_node_size."),
			},
		},
	})
}

func TestValidApacheZookeeperClusterCreate(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid_apache_zookeeper.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted("validApacheZookeeper", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("validApacheZookeeper"),
					testCheckResourceCreated("validApacheZookeeper", hostname, username, apiKey),
				),
			},
		},
	})
}

func TestAccClusterCredentials(t *testing.T) {
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	validConfig, _ := ioutil.ReadFile("data/valid_with_password_and_client_encryption.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted("valid_with_password_and_client_encryption", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckClusterCredentials(hostname, username, apiKey),
				),
			},
		},
	})
}

func TestCheckSingleDCRefreshToMultiDC(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}

	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)

	validSingleDCConfig, _ := ioutil.ReadFile("data/valid_single_dc_cluster.tf")
	singleDCConfig := fmt.Sprintf(string(validSingleDCConfig), username, apiKey, hostname)

	validMultiDCConfig, _ := ioutil.ReadFile("data/valid_multi_dc_cluster.tf")
	multiDCConfig := fmt.Sprintf(string(validMultiDCConfig), username, apiKey, hostname)

	attributesConflictWithDataCentres := []string{"data_centre",
		"node_size", "rack_allocation", "network", "cluster_provider", "bundle"}

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted("dc_test_cluster", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config:             singleDCConfig,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("instaclustr_cluster.dc_test_cluster", "data_centre_custom_name", "AWS_VPC_US_EAST_1_name"),
					addDCtoCluster("dc_test_cluster", hostname, username, apiKey, "data/valid_add_dc.json"),
				),
			},
			{
				Config: multiDCConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceStateAttributesDeleted("dc_test_cluster", attributesConflictWithDataCentres...),
				),
			},
		},
	})
}

func testCheckClusterCredentials(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["data.instaclustr_cluster_credentials.cluster_credentials"]

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		clusterId := resourceState.Primary.Attributes["cluster_id"]

		clusterCredentials, err := client.ReadCluster(clusterId)
		if err != nil {
			return fmt.Errorf("Failed to read Cluster Credentials from %s: %s", clusterId, err)
		}

		if clusterCredentials.InstaclustrUserPassword != resourceState.Primary.Attributes["cluster_password"] {
			return fmt.Errorf("Password of the cluster and resource are different")
		}

		if clusterCredentials.ClusterCertificateDownload != resourceState.Primary.Attributes["certificate_download"] {
			return fmt.Errorf("Certificate download link of the cluster and resource are different")
		}

		if clusterCredentials.ClusterCertificateDownload == "disabled" {
			return fmt.Errorf("Client encryption is disabled")
		}

		return nil
	}
}

func testCheckResourceStateAttributesDeleted(resourceName string, attributeNames ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster."+resourceName]
		attributes := resourceState.Primary.Attributes
		for _, givenAttribute := range attributeNames {
			exist := false
			for _, attribute := range attributes {
				if attribute == givenAttribute {
					exist = true
					break
				}
			}
			if exist {
				return fmt.Errorf("Attribute %s exists", givenAttribute)
			}
		}
		return nil
	}
}
