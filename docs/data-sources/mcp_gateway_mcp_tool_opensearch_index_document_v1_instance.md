---
page_title: "instaclustr_mcp_gateway_mcp_tool_opensearch_index_document_v1_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_mcp_tool_opensearch_index_document_v1_instance (Data Source)
Configuration for an OpenSearch Index Document tool.
## Example Usage
```
data "instaclustr_mcp_gateway_mcp_tool_opensearch_index_document_v1_instance" "example" { 
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
*___description___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Description of the tool.<br><br>
*___backend_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the backend this tool is associated with.<br><br>
*___document_schema___*<br>
<ins>Type</ins>: string, read-only<br>
<br>JSON Schema describing the structure of document to be ingested.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the tool.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]+$`<br><br>Name of the tool.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the MCP Gateway cluster.<br><br>
*___index_name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `[a-z0-9][a-z0-9_\-\.]*`<br><br>The OpenSearch index to write documents into.<br><br>
*___virtual_server_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the virtual server this backend belongs to.<br><br>
