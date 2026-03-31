---
page_title: "instaclustr_kafka_client_metric_subscription Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_client_metric_subscription (Resource)
Configuration to subscribe to Kafka client telemetry.
## Example Usage
```
resource "instaclustr_kafka_client_metric_subscription" "example" {
  clients = "client_software_name=apache-kafka-java"
  interval = 300000
  cluster_id = "c1af59c6-ba0e-4cc2-a0f3-65cee17a5f37"
  metrics = "org.apache.kafka.consumer.node.request.latency.avg"
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
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), required, updatable<br>
<br>ID of the Kafka cluster<br><br>
*___clients___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Required. Identifiers of clients to collect metrics from. Must include at least one of: client_id, client_instance_id, client_software_name, client_software_version, client_source_address, client_source_port. Wildcard characters (.) at the end are not permitted.<br><br>
*___metrics___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Required. A single telemetry metric name to subscribe to. Wildcard characters (.) at the end are not permitted.<br><br>
### Input attributes - Optional
*___interval___*<br>
<ins>Type</ins>: integer, optional, updatable<br>
<ins>Constraints</ins>: minimum: 2E+4<br><br>The interval to collect client telemetry in milliseconds. Default value for this is 5 minutes. The minimum value is 20000 milliseconds.<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>Instaclustr identifier for the Kafka client telemetry subscription.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_kafka_client_metric_subscription.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
