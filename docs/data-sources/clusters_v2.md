---
page_title: "instaclustr_clusters_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_clusters_v2 (Data Source)
A listable data source of all cluster IDs in an Instaclustr Account.
## Example Usage
```
data "instaclustr_clusters_v2" "example" { }
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Read-only attributes
*___account_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>UUID of the Instaclustr Account.<br><br>
*___clusters___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [clusters](#nested--clusters) for nested schema<br>
<br>
<a id="nested--clusters"></a>
## Nested schema for `clusters`

### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>
*___application___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `APACHE_CASSANDRA`, `KAFKA`, `APACHE_ZOOKEEPER`, `KAFKA_CONNECT`, `CADENCE`, `CLICKHOUSE`, `REDIS`, `VALKEY`, `OPENSEARCH`, `POSTGRESQL`, `MCP_GATEWAY`, `KAFKA_DISKLESS`, `UNKNOWN` ]<br><br>
