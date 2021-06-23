package test

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"testing"
)

func TestKafkaUserResource(t *testing.T) {
	testProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}

	configBytes1, _ := ioutil.ReadFile("data/kafka_user_create_cluster.tf")
	configBytes2, _ := ioutil.ReadFile("data/kafka_user_create_user.tf")
	configBytes3, _ := ioutil.ReadFile("data/kafka_user_user_list.tf")
	configBytes4, _ := ioutil.ReadFile("data/invalid_kafka_user_create.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)

	kafkaUsername1 := "charlie1"
	kafkaUsername2 := "charlie2"
	kafkaUsername3 := "charlie3"
	kafkaUsername4 := "charlie4"
	oldPassword := "charlie123!"
	newPassword := "charlie123standard!"
	zookeeperNodeSize := "zk-developer-t3.small-20"

	createClusterConfig := fmt.Sprintf(string(configBytes1), username, apiKey, hostname, zookeeperNodeSize)
	createKafkaUserConfig := fmt.Sprintf(string(configBytes2), username, apiKey, hostname, zookeeperNodeSize,
		kafkaUsername1, oldPassword,
		kafkaUsername2, oldPassword,
		kafkaUsername3, oldPassword)
	createKafkaUserListConfig := fmt.Sprintf(string(configBytes3), username, apiKey, hostname, zookeeperNodeSize,
		kafkaUsername1, oldPassword,
		kafkaUsername2, oldPassword,
		kafkaUsername3, oldPassword)
	updateKafkaUserConfig := fmt.Sprintf(string(configBytes3), username, apiKey, hostname, zookeeperNodeSize,
		kafkaUsername1, newPassword,
		kafkaUsername2, newPassword,
		kafkaUsername3, newPassword)
	invalidKafkaUserCreateConfig := fmt.Sprintf(string(configBytes4), username, apiKey, hostname, zookeeperNodeSize,
		kafkaUsername4, oldPassword)

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
				Config: createKafkaUserConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValidKafka("instaclustr_kafka_user.kafka_user_charlie"),
					checkClusterRunning("kafka_cluster", hostname, username, apiKey),
					checkKafkaUserCreated(hostname, username, apiKey),
				),
			},
			{
				Config: createKafkaUserListConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValidKafka("data.instaclustr_kafka_user_list.kafka_user_list"),
					checkKafkaUserListCreated(hostname, username, apiKey),
				),
			},
			// Currently there is no easy way to check that the password for a Kafka user has been changed.
			// So, we just have to trust that successful API query changed the kafka user password.
			{
				Config: updateKafkaUserConfig,
				Check:  checkKafkaUserUpdated(newPassword),
			},
			// Can't rely on the resource destruction because we need the destruction to happen in order and checked,
			// i.e., we need to destroy the kafka user resources first.
			{
				Config: createClusterConfig,
				Check:  resource.ComposeTestCheckFunc(checkKafkaUserDeleted(kafkaUsername1, hostname, username, apiKey),
					checkKafkaUserDeleted(kafkaUsername2, hostname, username, apiKey),
					checkKafkaUserDeleted(kafkaUsername3, hostname, username, apiKey)),
			},
			{
				Config: invalidKafkaUserCreateConfig,
				ExpectError: regexp.MustCompile("invalid value for the 'sasl-scram-mechanism' option"),
			},
		},
	})
}

func testCheckResourceValidKafka(resourceName string) resource.TestCheckFunc {
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

func checkKafkaUserCreated(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		userResources := [3]string{
			"instaclustr_kafka_user.kafka_user_charlie",
			"instaclustr_kafka_user.kafka_user_charlie_scram-sha-512",
			"instaclustr_kafka_user.kafka_user_charlie_empty_options",
		}

		OUTER:
		for _, resourceName := range userResources {
			resourceState := s.Modules[0].Resources[resourceName]

			client := new(instaclustr.APIClient)
			client.InitClient(hostname, username, apiKey)
			kafka_username := resourceState.Primary.Attributes["username"]
			clusterId := resourceState.Primary.Attributes["cluster_id"]

			usernameList, err := client.ReadKafkaUserList(clusterId)
			if err != nil {
				return fmt.Errorf("Failed to read Kafka user list from %s: %s", clusterId, err)
			}
			for _, str := range usernameList {
				if kafka_username == str {
					continue OUTER
				}
			}
			return fmt.Errorf("User %s is not found within the username list of %s", username, clusterId)
		}
		return nil
	}
}

func checkKafkaUserUpdated(newPassword string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		userResources := [3]string{
			"instaclustr_kafka_user.kafka_user_charlie",
			"instaclustr_kafka_user.kafka_user_charlie_scram-sha-512",
			"instaclustr_kafka_user.kafka_user_charlie_empty_options",
		}

		for _, resourceName := range userResources {
			resourceState := s.Modules[0].Resources[resourceName]
			if resourceState == nil {
				return fmt.Errorf("%s resource not found in state.", resourceName)
			}

			instanceState := resourceState.Primary
			if instanceState == nil {
				return fmt.Errorf("resource has no primary instance")
			}

			if instanceState.Attributes["password"] != newPassword {
				return fmt.Errorf("The new password in the terraform state is not as expected after update: %s != %s", instanceState.Attributes["password"], newPassword)
			}
		}
		return nil
	}
}

func checkKafkaUserDeleted(kafka_username, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// the resource for the kafka user has been deleted, therefore we need to get the cluster id from the cluster resource
		resourceState := s.Modules[0].Resources["instaclustr_cluster.kafka_cluster"]
		clusterId := resourceState.Primary.Attributes["cluster_id"]

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)

		usernameList, err := client.ReadKafkaUserList(clusterId)
		if err != nil {
			return fmt.Errorf("Failed to read Kafka user list from %s: %s", clusterId, err)
		}
		for _, str := range usernameList {
			if kafka_username == str {
				return fmt.Errorf("Kafka user %s still exists in %s", kafka_username, clusterId)
			}
		}
		return nil
	}
}

func checkKafkaUserListCreated(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["data.instaclustr_kafka_user_list.kafka_user_list"]

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		clusterId := resourceState.Primary.Attributes["cluster_id"]

		usernameList, err := client.ReadKafkaUserList(clusterId)
		if err != nil {
			return fmt.Errorf("Failed to read Kafka user list from %s: %s", clusterId, err)
		}

		resourceListLen, _ := strconv.Atoi(resourceState.Primary.Attributes["username_list.#"])
		if resourceListLen != len(usernameList) {
			return fmt.Errorf("List of Kafka users of the Kafka cluster and resource are different (Length %d != %d). ", resourceListLen, len(usernameList))
		}

		for index, kafka_username := range usernameList {
			resourceUser := resourceState.Primary.Attributes[fmt.Sprintf("username_list.%d", index)]
			if resourceUser != kafka_username {
				return fmt.Errorf("List of Kafka users of the Kafka cluster and resource are different (Index %d: %s != %s). ", index, resourceUser, kafka_username)
			}
		}

		return nil
	}
}
