---
page_title: "instaclustr_mcp_gateway_mcp_tool_cassandra_query_v1 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_mcp_tool_cassandra_query_v1 (Resource)
Configuration for a Cassandra Query tool.
## Example Usage
```
resource "instaclustr_mcp_gateway_mcp_tool_cassandra_query_v1" "example" {
  query = "SELECT key, value FROM keyspace.table WHERE key = :key"
  name = "my-cassandra-query-tool"
  backend_id = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
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
*___backend_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the backend this tool is associated with.<br><br>
*___query___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>CQL executed by this tool.<br><br>
### Input attributes - Optional
*___description___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Description of the tool.<br><br>
*___name___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]+$`<br><br>Name of the tool.<br><br>
*___result_field___*<br>
<ins>Type</ins>: repeatable nested block, optional, updatable, see [result_field](#nested--result_field) for nested schema<br>
<br>Expected fields in the result set.<br><br>
*___parameter___*<br>
<ins>Type</ins>: repeatable nested block, optional, updatable, see [parameter](#nested--parameter) for nested schema<br>
<br>Bind variables accepted by the CQL query, in order.<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the tool.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the MCP Gateway cluster.<br><br>
*___virtual_server_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the virtual server this backend belongs to.<br><br>
<a id="nested--result_field"></a>
## Nested schema for `result_field`
Expected fields in the result set.<br>
### Input attributes - Required
*___name___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Name of the bind variable or result column.<br><br>
*___type___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Cassandra type name (for example uuid, text, int).<br><br>
### Input attributes - Optional
*___description___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Description of the parameter or column.<br><br>
<a id="nested--parameter"></a>
## Nested schema for `parameter`
Bind variables accepted by the CQL query, in order.<br>
### Input attributes - Required
*___name___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Name of the bind variable or result column.<br><br>
*___type___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Cassandra type name (for example uuid, text, int).<br><br>
### Input attributes - Optional
*___description___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Description of the parameter or column.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_mcp_gateway_mcp_tool_cassandra_query_v1.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
