---
page_title: "instaclustr_mcp_gateway_cluster_mcp_backend_http_server_v1 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_cluster_mcp_backend_http_server_v1 (Data Source)
HTTP Server backend configuration.
## Example Usage
```
data "instaclustr_mcp_gateway_cluster_mcp_backend_http_server_v1" "example" { 
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
*___server_url___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Host URL of the HTTP Server backend.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the backend.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^[a-z0-9_-]+$`<br><br>Name of the backend.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the MCP Gateway cluster.<br><br>
*___authentication___*<br>
<ins>Type</ins>: nested block, read-only, see [authentication](#nested--authentication) for nested schema<br>
<br>
*___virtual_server_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the virtual server this backend belongs to.<br><br>
<a id="nested--header"></a>
## Nested schema for `header`
HTTP header configuration for backend authentication. Defaults to Authorization: Bearer {KEY}<br>
### Read-only attributes
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>HTTP header name for backend authentication. Defaults to Authorization.<br><br>
*___key_prefix___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Key prefix, usually specifying authentication scheme.<br><br>
<a id="nested--authentication"></a>
## Nested schema for `authentication`

### Read-only attributes
*___cookie_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>HTTP cookie name for backend authentication.<br><br>
*___header___*<br>
<ins>Type</ins>: list of objects, read-only<br>
<br>HTTP header configuration for backend authentication. Defaults to Authorization: Bearer {KEY}<br><br>
*___scheme___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `PASSTHROUGH`, `KEY` ]<br><br>HTTP Authentication scheme for the backend.<br><br>
*___query_parameter_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>HTTP query parameter name for backend authentication.<br><br>
*___key___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The authentication credentials for the backend.<br><br>
