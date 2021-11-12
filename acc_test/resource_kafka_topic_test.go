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

func TestKafkaTopicResource(t *testing.T) {
	testProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}

	configBytes1, _ := ioutil.ReadFile("data/kafka_topic_create_cluster.tf")
	configBytes2, _ := ioutil.ReadFile("data/kafka_topic_create_topic.tf")
	configBytes3, _ := ioutil.ReadFile("data/kafka_topic_topic_list.tf")
	configBytes4, _ := ioutil.ReadFile("data/invalid_kafka_topic_create_duplicate.tf")
	configBytes5, _ := ioutil.ReadFile("data/kafka_topic_config_update.tf")

	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)

	topic1 := "test1"
	topic2 := "test2"
	kafkaNodeSize := "KFK-DEV-t4g.medium-80"
	zookeeperNodeSize := "KDZ-DEV-t4g.small-30"
	kafkaVersion := "apache-kafka:3.0.0"

	createClusterConfig := fmt.Sprintf(string(configBytes1), username, apiKey, hostname, kafkaNodeSize, kafkaVersion, zookeeperNodeSize)
	createKafkaTopicConfig := fmt.Sprintf(string(configBytes2), username, apiKey, hostname, kafkaNodeSize, kafkaVersion, zookeeperNodeSize,
		topic1,
		topic2)
	createKafkaTopicListConfig := fmt.Sprintf(string(configBytes3), username, apiKey, hostname, kafkaNodeSize, kafkaVersion, zookeeperNodeSize,
		topic1,
		topic2)
	invalidKafkaTopicCreateConfigDuplicate := fmt.Sprintf(string(configBytes4), username, apiKey, hostname, kafkaNodeSize, kafkaVersion, zookeeperNodeSize,
		topic1,
		topic2,
		topic1)
	updateKafkaTopicConfig := fmt.Sprintf(string(configBytes5), username, apiKey, hostname, kafkaNodeSize, kafkaVersion, zookeeperNodeSize,
		topic1,
		topic2)

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
				Config: createKafkaTopicConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValidKafka("instaclustr_kafka_topic.kafka_topic_test"),
					checkClusterRunning("kafka_cluster", hostname, username, apiKey),
					checkKafkaTopicCreated(hostname, username, apiKey),
				),
			},
			{
				Config: createKafkaTopicListConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceValidKafka("data.instaclustr_kafka_topic_list.kafka_topic_list"),
					checkKafkaTopicListCreated(hostname, username, apiKey),
				),
			},
			{
				Config:      invalidKafkaTopicCreateConfigDuplicate,
				ExpectError: regexp.MustCompile("Topic 'test1' already exists.."),
			},
			{
				Config: updateKafkaTopicConfig,
				Check:  checkKafkaTopicUpdated(hostname, username, apiKey),
			},
			// Can't rely on the resource destruction because we need the destruction to happen in order and checked,
			// i.e., we need to destroy the kafka topics resources first.
			{
				Config: createClusterConfig,
				Check: resource.ComposeTestCheckFunc(checkKafkaTopicDeleted(topic1, hostname, username, apiKey),
					checkKafkaTopicDeleted(topic2, hostname, username, apiKey)),
			},
		},
	})
}

func checkKafkaTopicCreated(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		topicResources := [2]string{
			"instaclustr_kafka_topic.kafka_topic_test",
			"instaclustr_kafka_topic.kafka_topic_test2",
		}

	OUTER:
		for _, resourceName := range topicResources {
			resourceState := s.Modules[0].Resources[resourceName]

			client := new(instaclustr.APIClient)
			client.InitClient(hostname, username, apiKey)
			kafka_topic_name := resourceState.Primary.Attributes["topic"]
			clusterId := resourceState.Primary.Attributes["cluster_id"]

			topicList, err := client.ReadKafkaTopicList(clusterId)
			if err != nil {
				return fmt.Errorf("Failed to read Kafka topic list from %s: %s", clusterId, err)
			}
			for _, str := range topicList.Topics {
				if kafka_topic_name == str {
					continue OUTER
				}
			}
			return fmt.Errorf("Topic %s is not found within the topic list of %s", kafka_topic_name, clusterId)
		}
		return nil
	}
}

func checkKafkaTopicListCreated(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["data.instaclustr_kafka_topic_list.kafka_topic_list"]

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		clusterId := resourceState.Primary.Attributes["cluster_id"]

		kafkaTopics, err := client.ReadKafkaTopicList(clusterId)
		if err != nil {
			return fmt.Errorf("Failed to read Kafka topic list from %s: %s", clusterId, err)
		}
		topicList := kafkaTopics.Topics

		resourceListLen, _ := strconv.Atoi(resourceState.Primary.Attributes["topics.#"])
		if resourceListLen != len(topicList) {
			return fmt.Errorf("List of Kafka topics of the Kafka cluster and resource are different (Length %d != %d). ", resourceListLen, len(topicList))
		}

		for index, kafka_topic := range topicList {
			resourceTopic := resourceState.Primary.Attributes[fmt.Sprintf("topics.%d", index)]
			if resourceTopic != kafka_topic {
				return fmt.Errorf("List of Kafka topics of the Kafka cluster and resource are different (Index %d: %s != %s). ", index, resourceTopic, kafka_topic)
			}
		}

		return nil
	}
}

func checkKafkaTopicUpdated(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		topicResources := [2]string{
			"instaclustr_kafka_topic.kafka_topic_test",
			"instaclustr_kafka_topic.kafka_topic_test2",
		}
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)

		for _, resourceName := range topicResources {
			resourceState := s.Modules[0].Resources[resourceName]
			if resourceState == nil {
				return fmt.Errorf("%s resource not found in state.", resourceName)
			}
			clusterId := resourceState.Primary.Attributes["cluster_id"]
			instanceState := resourceState.Primary
			if instanceState == nil {
				return fmt.Errorf("resource has no primary instance")
			}
			kafka_topic_name := instanceState.Attributes["topic"]
			kafkaTopicConfig, err := client.ReadKafkaTopicConfig(clusterId, kafka_topic_name)
			if err != nil {
				return fmt.Errorf("Failed to read Kafka topic %s's config: %s", kafka_topic_name, err)
			}

			if kafkaTopicConfig.Config.MinInsyncReplicas != 2 || *kafkaTopicConfig.Config.MessageDownconversionEnable != false ||
				*kafkaTopicConfig.Config.UncleanLeaderElectionEnable != true {
				return fmt.Errorf("The topic %s's configs in the cluster are not updated as expected.", kafka_topic_name)
			}

			if instanceState.Attributes["config.0.min_insync_replicas"] != "2" || instanceState.Attributes["config.0.message_downconversion_enable"] != "false" ||
				instanceState.Attributes["config.0.unclean_leader_election_enable"] != "true" {
				return fmt.Errorf("The topic %s's configs in the terraform state are not updated as expected.", kafka_topic_name)
			}
		}
		return nil
	}
}

func checkKafkaTopicDeleted(kafka_topic_name, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// the resource for the kafka topic has been deleted, therefore we need to get the cluster id from the cluster resource
		resourceState := s.Modules[0].Resources["instaclustr_cluster.kafka_cluster"]
		clusterId := resourceState.Primary.Attributes["cluster_id"]

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)

		topicList, err := client.ReadKafkaTopicList(clusterId)
		if err != nil {
			return fmt.Errorf("Failed to read Kafka topic list from %s: %s", clusterId, err)
		}
		for _, str := range topicList.Topics {
			if kafka_topic_name == str {
				return fmt.Errorf("Kafka topic %s still exists in %s", kafka_topic_name, clusterId)
			}
		}
		return nil
	}
}
