---
page_title: "instaclustr_mcp_gateway_mcp_backend_cassandra_v1 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_mcp_backend_cassandra_v1 (Resource)
Configuration for a Cassandra backend.
## Example Usage
```
resource "instaclustr_mcp_gateway_mcp_backend_cassandra_v1" "example" {
  virtual_server_id = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
  name = "my-cassandra-backend"
  cassandra_data_centre_id = "d4e5f6a7-b8c9-0123-def0-234567890123"
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
*___name___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: pattern: `^[a-z0-9_-]+$`<br><br>Name of the backend.<br><br>
*___cassandra_data_centre_id___*<br>
<ins>Type</ins>: string (uuid), required, updatable<br>
<br>ID of the Cassandra Data Centre<br><br>
*___virtual_server_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the virtual server this backend belongs to.<br><br>
### Input attributes - Optional
*___authentication_sasl_plain___*<br>
<ins>Type</ins>: nested block, optional, updatable, see [authentication_sasl_plain](#nested--authentication_sasl_plain) for nested schema<br>
<br>
### Read-only attributes
*___cassandra_cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Cassandra Cluster<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the backend.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the MCP Gateway cluster.<br><br>
<a id="nested--authentication_sasl_plain"></a>
## Nested schema for `authentication_sasl_plain`

### Input attributes - Required
*___username___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>
*___password___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_mcp_gateway_mcp_backend_cassandra_v1.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
