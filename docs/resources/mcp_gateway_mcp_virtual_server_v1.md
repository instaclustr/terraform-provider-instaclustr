---
page_title: "instaclustr_mcp_gateway_mcp_virtual_server_v1 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_mcp_virtual_server_v1 (Resource)
Definition of an MCP Gateway virtual server that can be managed.
## Example Usage
```
resource "instaclustr_mcp_gateway_mcp_virtual_server_v1" "example" {
  endpoint_url = "/my-mcp-server"
  name = "my-virtual-server"
  cluster_id = "4650aa98-e894-4675-9614-558915d5837a"
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
*___endpoint_url___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>URL path pattern for the virtual server endpoint.<br><br>
*___name___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]+$`<br><br>Name of the virtual server.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the MCP Gateway cluster.<br><br>
### Input attributes - Optional
*___authentication___*<br>
<ins>Type</ins>: nested block, optional, updatable, see [authentication](#nested--authentication) for nested schema<br>
<br>Authentication configuration for the virtual server.<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the virtual server.<br><br>
<a id="nested--authentication"></a>
## Nested schema for `authentication`
Authentication configuration for the virtual server.<br>
### Input attributes - Required
*___issuer___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>JWT token issuer.<br><br>
*___audiences___*<br>
<ins>Type</ins>: list of strings, required, updatable<br>
<ins>Constraints</ins>: minimum items: 1<br><br>Audiences that will be accepted by the MCP Gateway, defaults to the route URL.<br><br>
*___jwks_url___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>JWKS URL for JWT token validation.<br><br>
### Input attributes - Optional
*___scopes_supported___*<br>
<ins>Type</ins>: list of strings, optional, updatable<br>
<br>List of supported scopes for the virtual server. 'openid' is required and will be added automatically.<br><br>
*___roles_claim_name___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<ins>Constraints</ins>: pattern: `[^"'\\]*`<br><br>JWT role claim name field uses for Authorization, defaults to "roles"<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_mcp_gateway_mcp_virtual_server_v1.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
