---
page_title: "instaclustr_kafka_acl_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_acl_v2 (Resource)
List of access control lists for a Kafka cluster.
## Example Usage
```
resource "instaclustr_kafka_acl_v2" "example" {
  acl {
    host = "*"
    operation = "DESCRIBE"
    pattern_type = "LITERAL"
    permission_type = "ALLOW"
    principal = "User:test"
    resource_name = "kafka-cluster"
    resource_type = "CLUSTER"
  }

}
```
## Glossary
The following terms are used to describe attributes in the schema of this resource:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the resource.
- **_required_** - These attributes must be provided for the resource to be created.
- **_optional_** - These input attributes can be omitted, and doing so may result in a default value being used.
- **_immutable_** - These are input attributes that cannot be changed after the resource is created.
- **_updatable_** - These input attributes can be updated to a different value if needed, and doing so will trigger an update operation.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___acl___*<br>
<ins>Type</ins>: repeatable nested block, required, updatable, see [acl](#nested--acl) for nested schema<br>
<br>List of ACLs for the given principal.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), required, updatable<br>
<br>UUID of the Kafka cluster.<br><br>
*___user_query___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>This is the principal without the "User:" prefix.<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the ACL list for a principal. The value of this property has the form: [clusterId]_[principalUserQuery]
The user query is the principal value without the leading "User:".<br><br>
<a id="nested--acl"></a>
## Nested schema for `acl`
List of ACLs for the given principal.<br>
### Input attributes - Required
*___principal___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: pattern: `^User:.*$`<br><br>Valid values must start with "User:" including the wildcard, e.g., "User:*"<br><br>
*___permission_type___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: allowed values: [ `ALLOW`, `DENY` ]<br><br>Valid values for permissionType: "ALLOW", "DENY"<br><br>
*___pattern_type___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: allowed values: [ `LITERAL`, `PREFIXED` ]<br><br>Valid values for patternType: "LITERAL", "PREFIXED"<br><br>
*___host___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>It takes any string including the wildcard "*"<br><br>
*___resource_name___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>It takes any string including the wildcard "*"<br><br>
*___resource_type___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: allowed values: [ `CLUSTER`, `TOPIC`, `GROUP`, `DELEGATION_TOKEN`, `TRANSACTIONAL_ID` ]<br><br>Valid values for resourceType: "CLUSTER", "TOPIC", "GROUP", "DELEGATION_TOKEN", "TRANSACTIONAL_ID"
<br><br>
*___operation___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: allowed values: [ `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS`, `ALTER_CONFIGS`, `IDEMPOTENT_WRITE` ]<br><br>Valid values for operation: "ALL", "READ", "WRITE", "CREATE", "DELETE", "ALTER", "DESCRIBE", "CLUSTER_ACTION", "DESCRIBE_CONFIGS", "ALTER_CONFIGS", "IDEMPOTENT_WRITE"<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_kafka_acl_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
