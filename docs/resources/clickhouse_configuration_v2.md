---
page_title: "instaclustr_clickhouse_configuration_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_clickhouse_configuration_v2 (Resource)
ClickHouse configuration overrides
## Example Usage
```
resource "instaclustr_clickhouse_configuration_v2" "example" {
  override {
    name = "max_threads"
    value = "16"
  }

  override {
    name = "log_queries"
    value = "1"
  }

  cluster_id = "b997a00d-5bd4-4774-9bd7-5c0ad6189246"
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
### Input attributes - Optional
*___cluster_id___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>ID of the ClickHouse cluster<br><br>
*___override___*<br>
<ins>Type</ins>: list of objects, optional, updatable<br>
<ins>Constraints</ins>: minimum items: 1<br><br>List of configuration overrides<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the cluster's configuration. This should be of the form: 'cfg-<cluster uuid>'.<br><br>
<a id="nested--override"></a>
## Nested schema for `override`
List of configuration overrides<br>
### Input attributes - Required
*___name___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: allowed values: [ `max_threads`, `max_insert_threads`, `use_skip_indexes`, `insert_quorum`, `insert_quorum_timeout`, `insert_quorum_parallel`, `distributed_ddl_task_timeout`, `log_queries`, `max_execution_time`, `max_bytes_before_external_group_by`, `max_bytes_before_external_sort`, `optimize_on_insert`, `max_partition_size_to_drop`, `allow_experimental_json_type`, `http_connection_timeout`, `http_send_timeout`, `http_receive_timeout`, `keep_alive_timeout`, `tcp_keep_alive_timeout` ]<br><br>Name of the configuration property.<br><br>
### Input attributes - Optional
*___value___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Override value for the configuration property.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_clickhouse_configuration_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
