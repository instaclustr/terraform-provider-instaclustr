package instaclustr

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAPIClientDeleteKafkaAcl(t *testing.T) {
	client := SetupMock(t, "should-be-uuid/kafka/acls", `{"id":"should-be-uuid"}`, 200)
	err := client.DeleteKafkaAcl("should-be-uuid", nil)
	if err != nil {
		t.Fatalf("Failed to delete kafka ACL: %s", err)
	}
}

func TestAPIClientCreateKafkaAcl(t *testing.T) {
	client := SetupMock(t, "should-be-uuid/kafka/acls", `{"id":"should-be-uuid"}`, 200)
	err2 := client.CreateKafkaAcl("should-be-uuid", nil)
	if err2 != nil {
		t.Fatalf("Failed to create kafka ACL: %s", err2)
	}
}

func TestAPIClientReadKafkaAcls(t *testing.T) {
	filename := "data/valid_kafka_acls.json"
	parseFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	jsonStr := fmt.Sprintf("%s", parseFile)
	client := SetupMock(t, "should-be-uuid/kafka/acls/searches", jsonStr, 200)
	acls, err2 := client.ReadKafkaAcls("should-be-uuid", nil)
	if err2 != nil {
		t.Fatalf("Failed to list Kafka ACL: %s", err2)
	}
	if acls[0].Principal != "User:test1" || acls[1].Principal != "User:test2" {
		t.Fatalf("Values do not match.")
	}
}
