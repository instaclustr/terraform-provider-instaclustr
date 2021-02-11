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
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	updatedConfig := strings.Replace(oriConfig, "testcluster", "newcluster", 1)
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
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { AccTestEnvVarsCheck(t) },
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
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { AccTestEnvVarsCheck(t) },
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
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	validResizeConfig := strings.Replace(oriConfig, "resizeable-small(r5-l)", "resizeable-small(r5-xl)", 1)
	validResizeConfig = strings.Replace(validResizeConfig, "tf-resizable-test", "tf-resizable-partial-test", 1)
	invalidResizeClassConfig := strings.Replace(oriConfig, "resizeable-small(r5-l)", "resizeable-large(r5-xl)", 1)
	invalidResizeConfig := strings.Replace(oriConfig, "resizeable-small(r5-l)", "t3.medium", 1)

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
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { AccTestEnvVarsCheck(t) },
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
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { AccTestEnvVarsCheck(t) },
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
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	providerAccountName := os.Getenv("IC_PROV_ACC_NAME")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { AccTestEnvVarsCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(validConfig), username, hostname, apiKey, providerAccountName),
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

func TestValidRedisClusterCreate(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	validConfig, _ := ioutil.ReadFile("data/valid_redis_cluster_create.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck: func() {
			checkAccVariablesSet(t, []string{
				"IC_USERNAME",
				"IC_API_KEY",
			})
		},
		CheckDestroy: testCheckResourceDeleted("validRedis", hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid("validRedis"),
					testCheckResourceCreated("validRedis", hostname, username, apiKey),
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

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { AccTestEnvVarsCheck(t) },
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
