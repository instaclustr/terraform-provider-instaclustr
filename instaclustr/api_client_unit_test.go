package instaclustr

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAPIClientRead(t *testing.T) {
	id := "77b5a4e1-c422-4a78-b551-d8fa5c42ad95"
	request := fmt.Sprintf("%s/terraform-description", id)
	client := SetupMock(t, request, fmt.Sprintf(`{"id":"%s"}`, id), 202)
	cluster, err := client.ReadCluster(id)
	if err != nil {
		t.Fatalf("Failed to read cluster %s: %s", id, err)
	}
	if cluster.ID != id {
		t.Fatalf("Cluster expected %s but got %s", id, cluster.ID)
	}
}

func TestAPIClientReadNull(t *testing.T) {
	id := "Invalid_ID"
	request := fmt.Sprintf("%s/terraform-description", id)
	client := SetupMock(t, request, "", 404)
	var _, err = client.ReadCluster(id)
	if err == nil {
		t.Fatalf("Read a cluster expected error but got null")
	}
}

func TestAPIClientDelete(t *testing.T) {
	id := "77b5a4e1-c422-4a78-b551-d8fa5c42ad95"
	client := SetupMock(t, id, "", 202)
	var err = client.DeleteCluster(id)
	if err != nil {
		t.Fatalf("Failed to delete cluster %s: %s", id, err)
	}
}

func TestAPIClientGCPRead(t *testing.T) {
	id := "77b5a4e1-c422-4a78-b551-d8fa5c42ad95"
	vpcpeeringid := "3467890"
	request := fmt.Sprintf("vpc-peering/%s/%s", id, vpcpeeringid)
	client := SetupMock(t, request, "", 404)
	var _, err = client.GCPReadVpcPeering(id, vpcpeeringid)
	if err == nil {
		t.Fatalf("Read a cluster expected error but got null")
	}
}
func TestAPIClientDeleteNull(t *testing.T) {
	id := "Invalid_ID"
	client := SetupMock(t, id, "", 404)
	var err = client.DeleteCluster(id)
	if err == nil {
		t.Fatalf("Delete a cluster expected error but got null")
	}
}

func TestAPIClientCreate(t *testing.T) {
	filename := "data/valid_create.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "extended/", `{"id":"should-be-uuid"}`, 202)
	id, err := client.CreateCluster(jsonStr)
	if err != nil {
		t.Fatalf("Failed to create cluster: %s", err)
	}
	if id == "" {
		t.Fatalf("Failed to fetch cluster id")
	}
}

func TestAPIClientCreateInvalid(t *testing.T) {
	filename := "data/invalid_create.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "extended/", ``, 401)
	_, err = client.CreateCluster(jsonStr)
	if err == nil {
		t.Fatalf("Create a cluster expected error but got null")
	}
}

func TestAPIClientCreateNoNetwork(t *testing.T) {
	filename := "data/valid_create.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}

	client := new(APIClient)
	client.InitClient("http://127.0.0.1:5000", "Trisolaris", "DoNotAnswer!")
	_, err = client.CreateCluster(jsonStr)
	if err == nil {
		t.Fatalf("Create a cluster expected error but got null")
	}
}

func TestAPIClientReadNoNetwork(t *testing.T) {
	id := "77b5a4e1-c422-4a78-b551-d8fa5c42ad95"
	client := new(APIClient)
	client.InitClient("http://127.0.0.1:5000", "Trisolaris", "DoNotAnswer!")
	_, err := client.ReadCluster(id)
	if err == nil {
		t.Fatalf("Read a cluster expected error but got null")
	}
}

func TestAPIClientDeleteNoNetwork(t *testing.T) {
	id := "77b5a4e1-c422-4a78-b551-d8fa5c42ad95"
	client := new(APIClient)
	client.InitClient("http://127.0.0.1:5000", "Trisolaris", "DoNotAnswer!")
	err := client.DeleteCluster(id)
	if err == nil {
		t.Fatalf("Delete a cluster expected error but got null")
	}
}

func TestAPIClientCreateSgFirewall(t *testing.T) {
	filename := "data/valid_sg_firewall.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "should-be-uuid/firewallRules/", `{"id":"should-be-uuid"}`, 202)
	err2 := client.CreateFirewallRule(jsonStr, "should-be-uuid")
	if err2 != nil {
		t.Fatalf("Failed to create firewall rule: %s", err2)
	}
}

func TestAPIClientReadSgFirewall(t *testing.T) {
	filename := "data/valid_sg_firewall.json"
	parse_file, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	jsonStr := fmt.Sprintf("[%s]", parse_file)
	client := SetupMock(t, "should-be-uuid/firewallRules/", jsonStr, 200)
	values, err2 := client.ReadFirewallRules("should-be-uuid")
	if err2 != nil {
		t.Fatalf("Failed to read firewall rule: %s", err2)
	}
	if (*values)[0].SecurityGroupId != "sg-0123abcde456ffabc" {
		t.Fatalf("Values do not match, expected %s but got %s", values, "should-be-uuid")
	}
}

func TestAPIDeleteSgFirewall(t *testing.T) {
	filename := "data/valid_sg_firewall.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "should-be-uuid/firewallRules/", `{"id":"should-be-uuid"}`, 202)
	err2 := client.DeleteFirewallRule(jsonStr, "should-be-uuid")
	if err2 != nil {
		t.Fatalf("Failed to create firewall rule: %s", err2)
	}
}

func TestAPIClientCreateNullFirewall(t *testing.T) {
	filename := "data/null_firewall.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "should-be-uuid/firewallRules/", `{"id":"should-be-uuid"}`, 400)
	err2 := client.CreateFirewallRule(jsonStr, "should-be-uuid")
	if err2 == nil {
		t.Fatalf("Firewall creation expected error but got null: %s", err2)
	}
}

func TestAPIClientCreateInvalidFirewall(t *testing.T) {
	filename := "data/invalid_firewall.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "should-be-uuid/firewallRules/", `{"id":"should-be-uuid"}`, 400)
	err2 := client.CreateFirewallRule(jsonStr, "should-be-uuid")
	if err2 == nil {
		t.Fatalf("Firewall creation expected error but got null: %s", err2)
	}
}
