---
page_title: "instaclustr_aws_security_group_firewall_rules_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_security_group_firewall_rules_v2 (Resource)
A listable data source of all AWS Security Group based firewall rules in an Instaclustr managed cluster.
## Example Usage
```
resource "instaclustr_aws_security_group_firewall_rules_v2" "example" {
  cluster_id = "91346037-b969-4c4d-8112-59e1d2aa67e9"
  firewall_rule {
    security_group_id = "sg-903004f0"
    type = "CASSANDRA_CQL"
  }

  firewall_rule {
    security_group_id = "sg-903004f6"
    type = "CASSANDRA_CQL"
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
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), required, updatable<br>
<br>ID of the cluster<br><br>
*___firewall_rule___*<br>
<ins>Type</ins>: repeatable nested block, required, updatable, see [firewall_rule](#nested--firewall_rule) for nested schema<br>
<br>Cluster AWS Security Group firewall rules<br><br>
### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The status of the cluster<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the cluster security group firewall rules resource<br><br>
<a id="nested--firewall_rule"></a>
## Nested schema for `firewall_rule`
Cluster AWS Security Group firewall rules<br>
### Input attributes - Required
*___security_group_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The security group ID of the AWS security group firewall rule.<br><br>
*___type___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: allowed values: [ `APACHE_ZOOKEEPER`, `CADENCE`, `CADENCE_GRPC`, `CADENCE_HTTP_API`, `CADENCE_WEB`, `CASSANDRA`, `CASSANDRA_CQL`, `CLICKHOUSE`, `CLICKHOUSE_WEB`, `ELASTICSEARCH`, `KAFKA`, `KAFKA_CONNECT`, `KAFKA_ENCRYPTION`, `KAFKA_MTLS`, `KAFKA_NO_ENCRYPTION`, `KAFKA_PRIVATE_SASL_PLAINTEXT_LISTENER`, `KAFKA_PUBLIC_SASL_PLAINTEXT_LISTENER`, `KAFKA_PRIVATE_PLAINTEXT_LISTENER`, `KAFKA_PUBLIC_PLAINTEXT_LISTENER`, `KAFKA_PRIVATE_SSL_LISTENER`, `KAFKA_PUBLIC_SSL_LISTENER`, `KAFKA_REST_PROXY`, `KAFKA_SCHEMA_REGISTRY`, `KARAPACE_REST_PROXY`, `KARAPACE_SCHEMA_REGISTRY`, `OPENSEARCH`, `OPENSEARCH_DASHBOARDS`, `PGBOUNCER`, `POSTGRESQL`, `REDIS`, `VALKEY`, `SEARCH_DASHBOARDS`, `SECURE_APACHE_ZOOKEEPER`, `SPARK`, `SPARK_JOBSERVER`, `SHOTOVER_PROXY`, `DEBEZIUM_CONNECTOR_CASSANDRA_KAFKA`, `DEBEZIUM_CONNECTOR_CASSANDRA_SCHEMA` ]<br><br>The type of firewall rule.<br><br>
### Read-only attributes
*___deferred_reason___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The reason (if needed) for the deferred status of the AWS security group firewall rule.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the AWS security group firewall rule.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_aws_security_group_firewall_rules_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
