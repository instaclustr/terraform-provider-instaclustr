---
page_title: "instaclustr_kafka_acl_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_acl_v2_instance (Data Source)
List of access control lists for a Kafka cluster.
## Example Usage
```
data "instaclustr_kafka_acl_v2_instance" "example" { 
  id = "<id>" // the value of the `id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the ACL list for a principal. The value of this property has the form: [clusterId]_[principalUserQuery]
The user query is the principal value without the leading "User:".<br><br>
*___acl___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [acl](#nested--acl) for nested schema<br>
<br>List of ACLs for the given principal.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>UUID of the Kafka cluster.<br><br>
*___user_query___*<br>
<ins>Type</ins>: string, read-only<br>
<br>This is the principal without the "User:" prefix.<br><br>
<a id="nested--acl"></a>
## Nested schema for `acl`
List of ACLs for the given principal.<br>
### Read-only attributes
*___principal___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^User:.*$`<br><br>Valid values must start with "User:" including the wildcard, e.g., "User:*"<br><br>
*___permission_type___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `ALLOW`, `DENY` ]<br><br>Valid values for permissionType: "ALLOW", "DENY"<br><br>
*___pattern_type___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `LITERAL`, `PREFIXED` ]<br><br>Valid values for patternType: "LITERAL", "PREFIXED"<br><br>
*___host___*<br>
<ins>Type</ins>: string, read-only<br>
<br>It takes any string including the wildcard "*"<br><br>
*___resource_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>It takes any string including the wildcard "*"<br><br>
*___resource_type___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `CLUSTER`, `TOPIC`, `GROUP`, `DELEGATION_TOKEN`, `TRANSACTIONAL_ID` ]<br><br>Valid values for resourceType: "CLUSTER", "TOPIC", "GROUP", "DELEGATION_TOKEN", "TRANSACTIONAL_ID"
<br><br>
*___operation___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS`, `ALTER_CONFIGS`, `IDEMPOTENT_WRITE` ]<br><br>Valid values for operation: "ALL", "READ", "WRITE", "CREATE", "DELETE", "ALTER", "DESCRIBE", "CLUSTER_ACTION", "DESCRIBE_CONFIGS", "ALTER_CONFIGS", "IDEMPOTENT_WRITE"<br><br>
