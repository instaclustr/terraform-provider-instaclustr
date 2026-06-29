---
page_title: "instaclustr_mcp_gateway_mcp_tool_http_server_operation_v1 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_mcp_tool_http_server_operation_v1 (Resource)
Configuration for an HTTP Server Operation tool.
## Example Usage
```
resource "instaclustr_mcp_gateway_mcp_tool_http_server_operation_v1" "example" {
  path = "/clusters/{clusterId}"
  method = "get"
  name = "my-http-operation"
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
*___method___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: allowed values: [ `get`, `post`, `put`, `delete` ]<br><br>HTTP method for the HTTP Server operation.<br><br>
*___path___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: pattern: `[^"'\\]+`<br><br>Path template for the HTTP operation.<br><br>
### Input attributes - Optional
*___description___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Description of the tool.<br><br>
*___name___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]+$`<br><br>Name of the tool.<br><br>
*___request_body___*<br>
<ins>Type</ins>: list of objects, optional, updatable<br>
<br>
*___parameter___*<br>
<ins>Type</ins>: list of objects, optional, updatable<br>
<br>Operation parameters (JSON).<br><br>
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
<a id="nested--request_body"></a>
## Nested schema for `request_body`

### Input attributes - Required
*___schema___*<br>
<ins>Type</ins>: string (json), required, updatable<br>
<br>
### Input attributes - Optional
*___description___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Description of the request body<br><br>
<a id="nested--parameter"></a>
## Nested schema for `parameter`
Operation parameters (JSON).<br>
### Input attributes - Required
*___schema___*<br>
<ins>Type</ins>: string (json), required, updatable<br>
<br>
*___in___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: allowed values: [ `path`, `header`, `query` ]<br><br>Which part of the URL the parameter is specified<br><br>
*___name___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Name of the parameter<br><br>
### Input attributes - Optional
*___description___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Description of the parameter<br><br>
*___required___*<br>
<ins>Type</ins>: boolean, optional, updatable<br>
<br>Is the parameter required or optional<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_mcp_gateway_mcp_tool_http_server_operation_v1.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
