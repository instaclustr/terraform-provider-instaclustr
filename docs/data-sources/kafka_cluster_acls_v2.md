---
page_title: "instaclustr_kafka_cluster_acls_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_cluster_acls_v2 (Data Source)
List of access control lists for a Kafka cluster.
## Example Usage
```
data "instaclustr_kafka_cluster_acls_v2" "example" { 
  kafka_cluster_id = "<kafka_cluster_id>" // the value of the `kafka_cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___kafka_cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the Kafka cluster.<br><br>
### Read-only attributes
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>UUID of the Kafka cluster<br><br>
*___acl_lists___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [acl_lists](#nested--acl_lists) for nested schema<br>
<br>List of all ACLs in the Kafka cluster.<br><br>
<a id="nested--acl"></a>
## Nested schema for `acl`
List of ACLs for the given principal.<br>
### Read-only attributes
*___principal___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^User:.*$`<br><br>Specifies the users(s) for which this ACL applies and can include the wildcard '*'. Valid values must start with "User:" including the wildcard.<br><br>
*___permission_type___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `ALLOW`, `DENY` ]<br><br>Specifies whether to allow or deny the operation.<br><br>
*___pattern_type___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `LITERAL`, `PREFIXED` ]<br><br>Indicates the resource-pattern-type<br><br>
*___host___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The IP address to which this ACL applies. It takes any string including the wildcard "*" for all IP addresses.<br><br>
*___resource_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Any string that fits the resource name, e.g. topic name if the resource type is TOPIC<br><br>
*___resource_type___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `CLUSTER`, `TOPIC`, `GROUP`, `DELEGATION_TOKEN`, `TRANSACTIONAL_ID` ]<br><br>Specifies the type of resource.<br><br>
*___operation___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS`, `ALTER_CONFIGS`, `IDEMPOTENT_WRITE` ]<br><br>The operation that will be allowed or denied.<br><br>
<a id="nested--acl_lists"></a>
## Nested schema for `acl_lists`
List of all ACLs in the Kafka cluster.<br>
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
