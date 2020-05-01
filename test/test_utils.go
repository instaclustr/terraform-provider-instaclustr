package test

import (
	"os"
	"testing"
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
}


func getoptionalenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}