---
page_title: "instaclustr_kafka_cluster_client_metric_subscriptions Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_cluster_client_metric_subscriptions (Data Source)
Configuration to subscribe to Kafka client telemetry.
## Example Usage
```
data "instaclustr_kafka_cluster_client_metric_subscriptions" "example" { 
  kafka_cluster_id = "<kafka_cluster_id>" // the value of the `kafka_cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___kafka_cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the Kafka cluster.<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>Instaclustr identifier for the Kafka client telemetry subscription.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Kafka cluster<br><br>
*___interval___*<br>
<ins>Type</ins>: integer, read-only<br>
<ins>Constraints</ins>: minimum: 2E+4<br><br>The interval to collect client telemetry in milliseconds. Default value for this is 5 minutes. The minimum value is 20000 milliseconds.<br><br>
*___clients___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Required. Identifiers of clients to collect metrics from. Must include at least one of: client_id, client_instance_id, client_software_name, client_software_version, client_source_address, client_source_port. Wildcard characters (.) at the end are not permitted.<br><br>
*___metrics___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Required. A single telemetry metric name to subscribe to. Wildcard characters (.) at the end are not permitted.<br><br>
