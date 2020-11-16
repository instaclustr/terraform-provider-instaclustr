package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
	"strconv"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
)

func TestKafkaRestProxyUserResource(t *testing.T) {
	testProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}

	configBytes1, _ := ioutil.ReadFile("data/kafka_rest_proxy_user_create_cluster.tf")
	configBytes2, _ := ioutil.ReadFile("data/kafka_rest_proxy_user_update_user.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	zookeeperNodeSize := "zk-developer-t3.small-20"

	kafkaRestProxyUserName := "ickafkarest"
	restProxyNewPassword := "RestProxyTest123test!"

	createClusterConfig := fmt.Sprintf(string(configBytes1), username, apiKey, hostname, zookeeperNodeSize)
	updateKafkaRestProxyUserConfig := fmt.Sprintf(string(configBytes2), username, apiKey, hostname, zookeeperNodeSize, kafkaRestProxyUserName, restProxyNewPassword)

	resource.Test(t, resource.TestCase{
		Providers:    testProviders,
		Steps: []resource.TestStep{
			{
				Config: createClusterConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValidKafkaRestProxy("instaclustr_cluster.kafka_rest_proxy_cluster"),
					checkKafkaRestProxyClusterRunning(hostname, username, apiKey),
				),
			},
			{
				Config: updateKafkaRestProxyUserConfig,
				Check: checkKafkaRestProxyUserUpdated(restProxyNewPassword),
			},
		},
	})
}

func testCheckResourceValidKafkaRestProxy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources[resourceName]
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

func checkKafkaRestProxyClusterRunning(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster.kafka_rest_proxy_cluster"]
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

func checkKafkaRestProxyUserUpdated(newPassword string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_kafka_rest_proxy_user.kafka_rest_proxy_user_update"]
		if resourceState == nil {
			return fmt.Errorf("instaclustr_kafka_rest_proxy_user.kafka_rest_proxy_user_update resource not found in state")
		}

		instanceState := resourceState.Primary
		if instanceState == nil {
			return fmt.Errorf("resource has no primary instance")
		}

		if instanceState.Attributes["password"] != newPassword {
			return fmt.Errorf("The new password in the terraform state is not as expected after update: %s != %s", instanceState.Attributes["password"], newPassword)
		}
		return nil
	}
}
