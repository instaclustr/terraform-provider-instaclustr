---
page_title: "instaclustr_aws_security_group_firewall_rule_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_security_group_firewall_rule_v2 (Resource)
Definition of an AWS Security Group based firewall rule to be applied to a cluster.
## Example Usage
```
resource "instaclustr_aws_security_group_firewall_rule_v2" "example" {
  security_group_id = "sg-01abc2345defg6789"
  cluster_id = "c1af59c6-ba0e-4cc2-a0f3-65cee17a5f37"
  type = "CASSANDRA"
}
```
## Glossary
The following terms are used to describe properties in the schema of this resource:
- **_read-only_** - These are properties that can only be read and not provided as an input to the resource.<br><br>
- **_required_** - These properties must be provided for the resource to be created.<br><br>
- **_optional_** - These input properties can be omitted, and doing so may result in a default value being used.<br><br>
- **_immutable_** - These are input properties that cannot be changed after the resource is created. The resource will be destroyed and re-created on `terraform apply` if Terraform detects a change in such properties.<br><br>
- **_updatable_** - These input properties can be updated to a different value if needed, and doing so will trigger an update operation.<br><br>
- **_nested block_** - These properties use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Properties with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested properties. When reading nested block properties, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_property[0].nested_property`.
## Schema
### security_group_id<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>The security group ID of the AWS security group firewall rule.
### deferred_reason<br>
<ins>Type</ins>: string, read-only<br>
<br>The reason (if needed) for the deferred status of the AWS security group firewall rule.
### cluster_id<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the cluster for the AWS security group firewall rule.
### id<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the AWS security group firewall rule.
### type<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: allowed values: [ `APACHE_ZOOKEEPER`, `CADENCE`, `CADENCE_GRPC`, `CADENCE_WEB`, `CASSANDRA`, `CASSANDRA_CQL`, `ELASTICSEARCH`, `KAFKA`, `KAFKA_CONNECT`, `KAFKA_REST_PROXY`, `KAFKA_SCHEMA_REGISTRY`, `KARAPACE_REST_PROXY`, `KARAPACE_SCHEMA_REGISTRY`, `OPENSEARCH`, `OPENSEARCH_DASHBOARDS`, `PGBOUNCER`, `POSTGRESQL`, `REDIS`, `SEARCH_DASHBOARDS`, `SECURE_APACHE_ZOOKEEPER`, `SPARK`, `SPARK_JOBSERVER` ]<br><br>The type of firewall rule.
### status<br>
<ins>Type</ins>: string, read-only<br>
<br>The status of the AWS security group firewall rule.

## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_aws_security_group_firewall_rule_v2.<resource-name> "<resource-id>"
```
`<resource-id>` is the unique identifier for this resource matching the value of the `id` property defined above.
