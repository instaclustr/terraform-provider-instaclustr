---
page_title: "instaclustr_mcp_gateway_mcp_virtual_server_v1_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_mcp_virtual_server_v1_instance (Data Source)
Definition of an MCP Gateway virtual server that can be managed.
## Example Usage
```
data "instaclustr_mcp_gateway_mcp_virtual_server_v1_instance" "example" { 
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
*___endpoint_url___*<br>
<ins>Type</ins>: string, read-only<br>
<br>URL path pattern for the virtual server endpoint.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the virtual server.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]+$`<br><br>Name of the virtual server.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the MCP Gateway cluster.<br><br>
*___authentication___*<br>
<ins>Type</ins>: nested block, read-only, see [authentication](#nested--authentication) for nested schema<br>
<br>Authentication configuration for the virtual server.<br><br>
<a id="nested--authentication"></a>
## Nested schema for `authentication`
Authentication configuration for the virtual server.<br>
### Read-only attributes
*___issuer___*<br>
<ins>Type</ins>: string, read-only<br>
<br>JWT token issuer.<br><br>
*___scopes_supported___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<br>List of supported scopes for the virtual server. 'openid' is required and will be added automatically.<br><br>
*___audiences___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<ins>Constraints</ins>: minimum items: 1<br><br>Audiences that will be accepted by the MCP Gateway, defaults to the route URL.<br><br>
*___roles_claim_name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `[^"'\\]*`<br><br>JWT role claim name field uses for Authorization, defaults to "roles"<br><br>
*___jwks_url___*<br>
<ins>Type</ins>: string, read-only<br>
<br>JWKS URL for JWT token validation.<br><br>
