---
page_title: "instaclustr_mcp_gateway_cluster_mcp_tool_http_server_operation_v1 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_cluster_mcp_tool_http_server_operation_v1 (Data Source)
Configuration for an HTTP Server Operation tool.
## Example Usage
```
data "instaclustr_mcp_gateway_cluster_mcp_tool_http_server_operation_v1" "example" { 
  mcp_gateway_cluster_id = "<mcp_gateway_cluster_id>" // the value of the `mcp_gateway_cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___mcp_gateway_cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the MCP Gateway cluster.<br><br>
### Read-only attributes
*___description___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Description of the tool.<br><br>
*___backend_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the backend this tool is associated with.<br><br>
*___method___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `get`, `post`, `put`, `delete` ]<br><br>HTTP method for the HTTP Server operation.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the tool.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]+$`<br><br>Name of the tool.<br><br>
*___path___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `[^"'\\]+`<br><br>Path template for the HTTP operation.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the MCP Gateway cluster.<br><br>
*___request_body___*<br>
<ins>Type</ins>: list of objects, read-only<br>
<br>
*___virtual_server_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the virtual server this backend belongs to.<br><br>
*___parameter___*<br>
<ins>Type</ins>: list of objects, read-only<br>
<br>Operation parameters (JSON).<br><br>
<a id="nested--request_body"></a>
## Nested schema for `request_body`

### Read-only attributes
*___description___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Description of the request body<br><br>
*___schema___*<br>
<ins>Type</ins>: string (json), read-only<br>
<br>
<a id="nested--parameter"></a>
## Nested schema for `parameter`
Operation parameters (JSON).<br>
### Read-only attributes
*___description___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Description of the parameter<br><br>
*___schema___*<br>
<ins>Type</ins>: string (json), read-only<br>
<br>
*___required___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Is the parameter required or optional<br><br>
*___in___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `path`, `header`, `query` ]<br><br>Which part of the URL the parameter is specified<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the parameter<br><br>
