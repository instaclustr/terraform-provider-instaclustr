---
page_title: "instaclustr_kafka_client_metric_subscription_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_client_metric_subscription_instance (Data Source)
Configuration to subscribe to Kafka client metrics.
## Example Usage
```
data "instaclustr_kafka_client_metric_subscription_instance" "example" { 
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
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>Instaclustr identifier for the Kafka client metric subscription.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Kafka cluster<br><br>
*___interval___*<br>
<ins>Type</ins>: integer, read-only<br>
<ins>Constraints</ins>: minimum: 2E+4<br><br>The interval to collect client metrics in milliseconds. Default value for this is 5 minutes. The minimum value is 20000 milliseconds.<br><br>
*___clients___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Identifiers of clients to collect metrics from. Wild card values are supported here.<br><br>
*___metrics___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The metrics to subscribe to. Wild card values are supported here.<br><br>
