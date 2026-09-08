---
page_title: "instaclustr_mcp_gateway_mcp_backend_instaclustr_api_v1_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_mcp_backend_instaclustr_api_v1_instance (Data Source)
Configuration for an Instaclustr API backend.
## Example Usage
```
data "instaclustr_mcp_gateway_mcp_backend_instaclustr_api_v1_instance" "example" { 
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
*___server_url___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Host URL of the Instaclustr API. Assigned by the platform.<br><br>
*___components___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<ins>Constraints</ins>: minimum items: 1, allowed values: [ `CLUSTER_MANAGEMENT_API_CORE`, `MONITORING_API_CORE` ]<br><br>Instaclustr API groups to expose as MCP tools. At least one group is required. CLUSTER_MANAGEMENT_API_CORE covers cluster details and operational state. MONITORING_API_CORE covers metrics.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the backend.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the MCP Gateway cluster.<br><br>
*___api_user___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Username of an account-linked user or Service User whose Instaclustr API keys this backend uses. A MONITORING key is required when MONITORING_API_CORE is selected, and a PROVISIONING_READONLY key is required when CLUSTER_MANAGEMENT_API_CORE is selected. After you rotate a key, call the Instaclustr API backend sync endpoint to update configuration<br><br>
*___virtual_server_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the virtual server this backend belongs to.<br><br>
