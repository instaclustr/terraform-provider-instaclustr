package test

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
	"os"
	"testing"
	"time"
)

func AccTestEnvVarsCheck(t *testing.T) {
	if v := os.Getenv("IC_USERNAME"); v == "" {
		t.Fatal("IC_USERNAME for provisioning API must be set for acceptance tests")
	}
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY for provisioning API must be set for acceptance tests")
	}
	if v := os.Getenv("KMS_ARN"); v == "" {
		t.Fatal("KMS_ARN for AccEBS encryption must be set for acceptance tests")
	}
	if v := os.Getenv("IC_PROV_ACC_NAME"); v == "" {
		t.Fatal("IC_PROV_ACC_NAME for provisioning API must be set for acceptance tests")
	}
	if v := os.Getenv("IC_PROV_VPC_ID"); v == "" {
		t.Fatal("IC_PROV_VPC_ID for provisioning API must be set for acceptance tests")
	}
	if x := os.Getenv("IC_TEST_KAFKA_CONNECT"); x != "" {
		env_vars := []string{"IC_TARGET_KAFKA_CLUSTER_ID", "IC_AWS_ACCESS_KEY", "IC_AWS_SECRET_KEY", "IC_S3_BUCKET_NAME",
			"IC_AZURE_STORAGE_ACCOUNT_NAME", "IC_AZURE_STORAGE_ACCOUNT_KEY", "IC_AZURE_STORAGE_CONTAINER_NAME",
			"IC_SSL_ENABLED_PROTOCOLS", "IC_SSL_TRUSTSTORE_PASSWORD", "IC_SSL_PROTOCOL", "IC_SECURITY_PROTOCOL",
			"IC_SASL_MECHANISM", "IC_SASL_JAAS_CONFIG", "IC_BOOTSTRAP_SERVERS", "IC_TRUSTSTORE"}
		for _, s := range env_vars {
			if v := os.Getenv(s); v == "" {
				fatalMessage := fmt.Sprintf("When IC_TEST_KAFKA_CONNECT is set, %s must be set for acceptance tests", s)
				t.Fatal(fatalMessage, s)
			}
		}
	}
}

func checkAccVariablesSet(t *testing.T, envVars []string) {
	for i := 0; i < len(envVars); i++ {
		switch envVars[i] {
		case "IC_USERNAME":
			if v := os.Getenv("IC_USERNAME"); v == "" {
				t.Fatal("IC_USERNAME for provisioning API must be set for acceptance tests")
			}
		case "IC_API_KEY":
			if v := os.Getenv("IC_API_KEY"); v == "" {
				t.Fatal("IC_API_KEY for provisioning API must be set for acceptance tests")
			}
		case "KMS_ARN":
			if v := os.Getenv("KMS_ARN"); v == "" {
				t.Fatal("KMS_ARN for AccEBS encryption must be set for acceptance tests")
			}
		case "IC_PROV_ACC_NAME":
			if v := os.Getenv("IC_PROV_ACC_NAME"); v == "" {
				t.Fatal("IC_PROV_ACC_NAME for provisioning API must be set for acceptance tests")
			}
		case "IC_PROV_VPC_ID":
			if v := os.Getenv("IC_PROV_VPC_ID"); v == "" {
				t.Fatal("IC_PROV_VPC_ID for provisioning API must be set for acceptance tests")
			}
		case "IC_TEST_KAFKA_CONNECT":
			if x := os.Getenv("IC_TEST_KAFKA_CONNECT"); x != "" {
				env_vars := []string{"IC_TARGET_KAFKA_CLUSTER_ID", "IC_AWS_ACCESS_KEY", "IC_AWS_SECRET_KEY", "IC_S3_BUCKET_NAME",
					"IC_AZURE_STORAGE_ACCOUNT_NAME", "IC_AZURE_STORAGE_ACCOUNT_KEY", "IC_AZURE_STORAGE_CONTAINER_NAME",
					"IC_SSL_ENABLED_PROTOCOLS", "IC_SSL_TRUSTSTORE_PASSWORD", "IC_SSL_PROTOCOL", "IC_SECURITY_PROTOCOL",
					"IC_SASL_MECHANISM", "IC_SASL_JAAS_CONFIG", "IC_BOOTSTRAP_SERVERS", "IC_TRUSTSTORE"}
				for _, s := range env_vars {
					if v := os.Getenv(s); v == "" {
						fatalMessage := fmt.Sprintf("When IC_TEST_KAFKA_CONNECT is set, %s must be set for acceptance tests", s)
						t.Fatal(fatalMessage, s)
					}
				}
			}
		}
	}
}

func getOptionalEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func checkClusterRunning(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster." + resourceName]
		id := resourceState.Primary.Attributes["cluster_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)

		const ClusterReadInterval = 5
		const WaitForClusterTimeout = 30 * 60
		var latestStatus string
		timePassed := 0
		fmt.Print("\033[s")
		for {
			cluster, err := client.ReadCluster(id)
			if err != nil {
				fmt.Printf("\n")
				return fmt.Errorf("[Error] Error retrieving cluster info: %s", err)
			}
			latestStatus = cluster.ClusterStatus
			if cluster.ClusterStatus == "RUNNING" {
				break
			}
			if timePassed > WaitForClusterTimeout {
				fmt.Printf("\n")
				return fmt.Errorf("[Error] Timed out waiting for cluster to have the status 'RUNNING'. Current cluster status is '%s'", latestStatus)
			}
			timePassed += ClusterReadInterval
			fmt.Printf("\033[u\033[K%ds has elapsed while waiting for the cluster to reach RUNNING.\n", timePassed)
			time.Sleep(ClusterReadInterval * time.Second)
		}
		fmt.Printf("\n")
		return nil
	}
}
