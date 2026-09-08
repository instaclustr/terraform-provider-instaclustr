---
page_title: "instaclustr_mcp_gateway_mcp_backend_instaclustr_api_v1 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_mcp_backend_instaclustr_api_v1 (Resource)
Configuration for an Instaclustr API backend.
## Example Usage
```
resource "instaclustr_mcp_gateway_mcp_backend_instaclustr_api_v1" "example" {
  virtual_server_id = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
  components = [ "CLUSTER_MANAGEMENT_API_CORE", "MONITORING_API_CORE" ]
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
*___components___*<br>
<ins>Type</ins>: list of strings, required, updatable<br>
<ins>Constraints</ins>: minimum items: 1, allowed values: [ `CLUSTER_MANAGEMENT_API_CORE`, `MONITORING_API_CORE` ]<br><br>Instaclustr API groups to expose as MCP tools. At least one group is required. CLUSTER_MANAGEMENT_API_CORE covers cluster details and operational state. MONITORING_API_CORE covers metrics.<br><br>
*___api_user___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Username of an account-linked user or Service User whose Instaclustr API keys this backend uses. A MONITORING key is required when MONITORING_API_CORE is selected, and a PROVISIONING_READONLY key is required when CLUSTER_MANAGEMENT_API_CORE is selected. After you rotate a key, call the Instaclustr API backend sync endpoint to update configuration<br><br>
*___virtual_server_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the virtual server this backend belongs to.<br><br>
### Read-only attributes
*___server_url___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Host URL of the Instaclustr API. Assigned by the platform.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the backend.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the MCP Gateway cluster.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_mcp_gateway_mcp_backend_instaclustr_api_v1.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
