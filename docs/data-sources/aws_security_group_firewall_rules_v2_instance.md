---
page_title: "instaclustr_aws_security_group_firewall_rules_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_security_group_firewall_rules_v2_instance (Data Source)
A listable data source of all AWS Security Group based firewall rules in an Instaclustr managed cluster.
## Example Usage
```
data "instaclustr_aws_security_group_firewall_rules_v2_instance" "example" { 
  cluster_id = "<cluster_id>" // the value of the `cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The status of the cluster<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>
*___firewall_rule___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [firewall_rule](#nested--firewall_rule) for nested schema<br>
<br>
<a id="nested--firewall_rule"></a>
## Nested schema for `firewall_rule`

### Read-only attributes
*___security_group_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The security group ID of the AWS security group firewall rule.<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The status of the AWS security group firewall rule.<br><br>
*___deferred_reason___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The reason (if needed) for the deferred status of the AWS security group firewall rule.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the AWS security group firewall rule.<br><br>
*___type___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `APACHE_ZOOKEEPER`, `CADENCE`, `CADENCE_GRPC`, `CADENCE_WEB`, `CASSANDRA`, `CASSANDRA_CQL`, `ELASTICSEARCH`, `KAFKA`, `KAFKA_CONNECT`, `KAFKA_ENCRYPTION`, `KAFKA_MTLS`, `KAFKA_NO_ENCRYPTION`, `KAFKA_PRIVATE_SASL_PLAINTEXT_LISTENER`, `KAFKA_PUBLIC_SASL_PLAINTEXT_LISTENER`, `KAFKA_PRIVATE_PLAINTEXT_LISTENER`, `KAFKA_PUBLIC_PLAINTEXT_LISTENER`, `KAFKA_PRIVATE_SSL_LISTENER`, `KAFKA_PUBLIC_SSL_LISTENER`, `KAFKA_REST_PROXY`, `KAFKA_SCHEMA_REGISTRY`, `KARAPACE_REST_PROXY`, `KARAPACE_SCHEMA_REGISTRY`, `MONGODB`, `OPENSEARCH`, `OPENSEARCH_DASHBOARDS`, `PGBOUNCER`, `POSTGRESQL`, `REDIS`, `SEARCH_DASHBOARDS`, `SECURE_APACHE_ZOOKEEPER`, `SPARK`, `SPARK_JOBSERVER`, `SHOTOVER_PROXY`, `DEBEZIUM_CONNECTOR_CASSANDRA_KAFKA`, `DEBEZIUM_CONNECTOR_CASSANDRA_SCHEMA` ]<br><br>The type of firewall rule.<br><br>
