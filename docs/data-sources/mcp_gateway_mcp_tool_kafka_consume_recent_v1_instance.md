---
page_title: "instaclustr_mcp_gateway_mcp_tool_kafka_consume_recent_v1_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_mcp_tool_kafka_consume_recent_v1_instance (Data Source)
Configuration for a Kafka Consume Recent tool.
## Example Usage
```
data "instaclustr_mcp_gateway_mcp_tool_kafka_consume_recent_v1_instance" "example" { 
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
*___topics___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<br>Kafka topics to consume from.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the tool.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]+$`<br><br>Name of the tool.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the MCP Gateway cluster.<br><br>
*___consume_period___*<br>
<ins>Type</ins>: integer (int32), read-only<br>
<br>Period in seconds to look back when consuming messages.<br><br>
*___message_description___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Description of what consumed messages represent.<br><br>
*___virtual_server_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the virtual server this backend belongs to.<br><br>
