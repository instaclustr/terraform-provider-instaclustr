package instaclustr

import (
        "testing"
	"fmt"
)

func TestDataSourceAclRead(t *testing.T) {
	acl, _, mockResourceData, mockClient := setupKafkaAclMocks(true)

	if err := doDataSourceKafkaAclListRead(mockResourceData, mockClient); err != nil {
		t.Fatalf("Reading Kafka ACL data list should have succeeded but got: %s", err)
	}	

	if mockResourceData.Get("id").(string) != "mock-acl-list" {
		t.Fatalf("ID set in the resource does not match the expected")
	}
	resourceAcls := mockResourceData.Get("acls").([]string)
	expectedAcl := fmt.Sprintf("(principal=%s, host=%s, resourceType=%s, resourceName=%s, operation=%s, permissionType=%s, patternType=%s)", 
			acl.Principal, acl.Host, acl.ResourceType, acl.ResourceName, acl.Operation, acl.PermissionType, acl.PatternType)
	if len(resourceAcls) != 1 && resourceAcls[0] != expectedAcl {
		t.Fatalf("Incorrect request returned.\nExpected:%s\nActual:%s", expectedAcl, resourceAcls)
	}
}
