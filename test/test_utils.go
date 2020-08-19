package test

import (
	"os"
	"testing"
	"fmt"
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


func getOptionalEnv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}
