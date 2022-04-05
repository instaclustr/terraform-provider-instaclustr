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

func TestRedisResource(t *testing.T) {
	testAccProvider := instaclustr.Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": testAccProvider,
	}
	resourceName := "validRedis"
	validConfig, _ := ioutil.ReadFile("data/valid_redis_cluster_create.tf")
	invalidRedisConfig, _ := ioutil.ReadFile("data/invalid_redis_cluster_create.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)

	validResizeConfig := strings.Replace(oriConfig, `node_size = "t3.small-20-r"`, `node_size = "t3.medium-80-r"`, 1)
	invalidResizeConfig := strings.Replace(oriConfig, `node_size = "t3.small-20-r"`, `node_size = "t3.small"`, 1)

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testCheckResourceDeleted(resourceName, hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config:      fmt.Sprintf(string(invalidRedisConfig), username, apiKey, hostname),
				ExpectError: regexp.MustCompile("'rack_allocation' is not supported in REDIS"),
			},
			{
				Config: oriConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValid(resourceName),
					testCheckResourceCreated(resourceName, hostname, username, apiKey),
					checkClusterRunning(resourceName, hostname, username, apiKey),
				),
			},
			{
				PreConfig: func() {
					fmt.Println("Sleep for 5 minutes to wait for Redis cluster to be ready for resize.")
					time.Sleep(5 * time.Minute)
				},
				Config:      invalidResizeConfig,
				ExpectError: regexp.MustCompile("t3.small"),
			},
			{
				Config: validResizeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("instaclustr_cluster."+resourceName, "cluster_name", "tf-redis-test"),
					resource.TestCheckResourceAttr("instaclustr_cluster."+resourceName, "node_size", "t3.medium-80-r"),
					testCheckClusterResize(resourceName, hostname, username, apiKey, "t3.medium-80-r"),
				),
			},
		},
	})
}
