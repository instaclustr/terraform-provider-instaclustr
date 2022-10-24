---
page_title: "instaclustr_cluster_network_firewall_rules_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_cluster_network_firewall_rules_v2 (Data Source)
A listable data source of all CIDR based firewall rules in an Instaclustr managed cluster.
## Example Usage
```
data "instaclustr_cluster_network_firewall_rules_v2" "example" { 
  cluster_id = "<cluster_id>" // the value of the `cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the cluster.<br><br>
### Read-only attributes
*___firewall_rules___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [firewall_rules](#nested--firewall_rules) for nested schema<br>
<br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>
<a id="nested--firewall_rules"></a>
## Nested schema for `firewall_rules`

### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The status of the cluster network firewall rule.<br><br>
*___deferred_reason___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The reason (if needed) for the deferred status of the cluster network firewall rule.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the cluster network firewall rule.<br><br>
*___type___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `APACHE_ZOOKEEPER`, `CADENCE`, `CADENCE_GRPC`, `CADENCE_WEB`, `CASSANDRA`, `CASSANDRA_CQL`, `ELASTICSEARCH`, `KAFKA`, `KAFKA_CONNECT`, `KAFKA_ENCRYPTION`, `KAFKA_MTLS`, `KAFKA_NO_ENCRYPTION`, `KAFKA_REST_PROXY`, `KAFKA_SCHEMA_REGISTRY`, `KARAPACE_REST_PROXY`, `KARAPACE_SCHEMA_REGISTRY`, `OPENSEARCH`, `OPENSEARCH_DASHBOARDS`, `PGBOUNCER`, `POSTGRESQL`, `REDIS`, `SEARCH_DASHBOARDS`, `SECURE_APACHE_ZOOKEEPER`, `SPARK`, `SPARK_JOBSERVER` ]<br><br>The type of firewall rule.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the cluster for the cluster network firewall rule.<br><br>
*___network___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The network of the cluster network firewall rule.<br><br>
