package instaclustr

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAPIClientCreateKafkaTopic(t *testing.T) {
	filename := "data/valid_kafka_topic.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "should-be-uuid/kafka/topics", `{"id":"should-be-uuid"}`, 201)
	err2 := client.CreateKafkaTopic("should-be-uuid", jsonStr)
	if err2 != nil {
		t.Fatalf("Failed to create kafka topic: %s", err2)
	}
}

func TestAPIDeleteKafkaTopic(t *testing.T) {
	client := SetupMock(t, "should-be-uuid/kafka/topics/test", `{"id":"should-be-uuid"}`, 200)
	err := client.DeleteKafkaTopic("should-be-uuid", "test")
	if err != nil {
		t.Fatalf("Failed to delete kafka topic: %s", err)
	}
}

func TestAPIClientReadKafkaTopicConfig(t *testing.T) {
	filename := "data/valid_kafka_topic_config.json"
	parseFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	jsonStr := fmt.Sprintf("%s", parseFile)
	client := SetupMock(t, "should-be-uuid/kafka/topics/test/config", jsonStr, 200)
	values, err2 := client.ReadKafkaTopicConfig("should-be-uuid", "test")
	if err2 != nil {
		t.Fatalf("Failed to read Kafka topic config: %s", err2)
	}
	if (*values).Config.CompressionType != "producer" || *(*values).Config.MessageDownconversionEnable != true ||
		(*values).Config.MinInsyncReplicas != 2 {
		t.Fatalf("Values do not match.")
	}
}

func TestAPIClientCreateKafkaTopicList(t *testing.T) {
	filename := "data/valid_kafka_topic_list.json"
	parseFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	jsonStr := fmt.Sprintf("%s", parseFile)
	client := SetupMock(t, "should-be-uuid/kafka/topics", jsonStr, 200)
	topicList, err2 := client.ReadKafkaTopicList("should-be-uuid")
	if err2 != nil {
		t.Fatalf("Failed to create kafka topic: %s", err2)
	}
	if topicList.Topics[0] != "test1" || topicList.Topics[1] != "test2" {
		t.Fatalf("Values do not match.")
	}
}

func TestAPIClientCreateKafkaUser(t *testing.T) {
	filename := "data/valid_kafka_user.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "should-be-uuid/kafka/users", `{"id":"should-be-uuid"}`, 201)
	err2 := client.CreateKafkaUser("should-be-uuid", jsonStr)
	if err2 != nil {
		t.Fatalf("Failed to create kafka user: %s", err2)
	}
}
