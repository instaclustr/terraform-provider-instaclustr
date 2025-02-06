---
page_title: "instaclustr_clickhouse_configuration_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_clickhouse_configuration_v2_instance (Data Source)
ClickHouse configuration overrides
## Example Usage
```
data "instaclustr_clickhouse_configuration_v2_instance" "example" { 
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
<br>ID of the cluster's configuration. This should be of the form: 'cfg-<cluster uuid>'.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the ClickHouse cluster<br><br>
*___override___*<br>
<ins>Type</ins>: list of objects, read-only<br>
<ins>Constraints</ins>: minimum items: 1<br><br>List of configuration overrides<br><br>
<a id="nested--override"></a>
## Nested schema for `override`
List of configuration overrides<br>
### Read-only attributes
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `max_threads`, `max_insert_threads`, `use_skip_indexes`, `insert_quorum`, `insert_quorum_timeout`, `insert_quorum_parallel`, `distributed_ddl_task_timeout`, `log_queries`, `max_execution_time`, `max_bytes_before_external_group_by`, `max_bytes_before_external_sort`, `optimize_on_insert`, `max_partition_size_to_drop` ]<br><br>Name of the configuration property.<br><br>
*___value___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Override value for the configuration property.<br><br>
