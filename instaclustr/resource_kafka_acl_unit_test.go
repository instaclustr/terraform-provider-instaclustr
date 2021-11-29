package instaclustr

import (
        "testing"
	"fmt"
)

func TestResourceAclCreate_noDuplicate(t *testing.T) {
	acl, _, mockResourceData, mockClient := setupKafkaAclMocks(false)

	setMockResourceData(mockResourceData, acl)

	if err := doResourceKafkaAclCreate(mockResourceData, mockClient); err != nil {
		t.Fatalf("Create ACL should have succeeed but got: %s", err)
	}	

	expectedId := fmt.Sprintf("mock&%s&%s&%s&%s&%s&%s&%s", acl.Principal, acl.Host, acl.ResourceType, 
		acl.ResourceName, acl.Operation, acl.PermissionType, acl.PatternType)
	if mockResourceData.Get("id").(string) != expectedId {
		t.Fatalf("ID set in the resource does not match the expected")
	}
}

func TestResourceAclCreate_duplicate(t *testing.T) {
	acl, mockData, mockResourceData, mockClient := setupKafkaAclMocks(true)

	setMockResourceData(mockResourceData, acl)

	if err := doResourceKafkaAclCreate(mockResourceData, mockClient); err == nil {
		t.Fatalf("Create ACL should have failed due to duplicate ACL, but succeeed")
	}	

	if mockData["id"] != nil {
		t.Fatalf("ID shouldn't be set in the failed case")
	}
}

func TestResourceAclRead_exists(t *testing.T) {
	acl, mockData, mockResourceData, mockClient := setupKafkaAclMocks(true)

	setMockResourceData(mockResourceData, acl)

	if err := doResourceKafkaAclRead(mockResourceData, mockClient); err != nil {
		t.Fatalf("Read ACL should have succeeded, but failed")
	}	

	if mockData["id"] == "" || mockData["principal"] == "" || mockData["host"] == "" || mockData["resource_type"] == "" ||
		mockData["resource_name"] == "" || mockData["operation"] == "" || mockData["permission_type"] == "" || mockData["pattern_type"] == "" {
		t.Fatalf("Resource should not be reset because the ACL exists in the server")
	}
}

func TestResourceAclRead_notExists(t *testing.T) {
	acl, mockData, mockResourceData, mockClient := setupKafkaAclMocks(false)

	setMockResourceData(mockResourceData, acl)

	if err := doResourceKafkaAclRead(mockResourceData, mockClient); err != nil {
		t.Fatalf("Read ACL should have succeeded, but failed")
	}	

	if !(mockData["id"] == "" || mockData["principal"] == "" || mockData["host"] == "" || mockData["resource_type"] == "" ||
		mockData["resource_name"] == "" || mockData["operation"] == "" || mockData["permission_type"] == "" || mockData["pattern_type"] == "") {
		t.Fatalf("Resource should be reset because the ACL does exist in the server")
	}
}

func TestResourceAclDelete(t *testing.T) {
	acl, mockData, mockResourceData, mockClient := setupKafkaAclMocks(false)

	setMockResourceData(mockResourceData, acl)

	if err := doResourceKafkaAclDelete(mockResourceData, mockClient); err != nil {
		t.Fatalf("Read ACL should have succeeded, but failed")
	}	

	if !(mockData["id"] == "" || mockData["principal"] == "" || mockData["host"] == "" || mockData["resource_type"] == "" ||
		mockData["resource_name"] == "" || mockData["operation"] == "" || mockData["permission_type"] == "" || mockData["pattern_type"] == "") {
		t.Fatalf("Resource should be reset because the ACL is deleted")
	}
}

func setupKafkaAclMocks(aclIncluded bool) (KafkaAcl, map[string]interface{}, KafkaAclResourceDataInterface, KafkaAclAPIClientInterface) {
	mockData := make(map[string]interface{})
	mockResourceData := MockKafkaAclResourceData {
		data:		mockData,
	}
	mockResourceData.Set("cluster_id", "mock")

	var acls []KafkaAcl
	acl := KafkaAcl {
		Principal:      "User:test",
		Host:           "*",
		ResourceType:   "TOPIC",
		ResourceName:   "*",
		Operation:      "ALL",
		PermissionType: "ALLOW",
		PatternType:    "LITERAL",
	}

	if aclIncluded {
		acls = append(acls, acl)
	}

	mockClient := MockKafkaAclApiClient {
		acls:		acls,
		cluster:	Cluster {
			ClusterStatus: "RUNNING",
		},
		err:		nil,
	}
	
	return acl, mockData, mockResourceData, mockClient
}

func setMockResourceData(mockResourceData KafkaAclResourceDataInterface, acl KafkaAcl) {
	mockResourceData.Set("principal", acl.Principal)
	mockResourceData.Set("host", acl.Host)
	mockResourceData.Set("resource_type", acl.ResourceType)
	mockResourceData.Set("resource_name", acl.ResourceName)
	mockResourceData.Set("operation", acl.Operation)
	mockResourceData.Set("permission_type", acl.PermissionType)
	mockResourceData.Set("pattern_type", acl.PatternType)
}
