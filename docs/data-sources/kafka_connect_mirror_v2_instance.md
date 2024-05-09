---
page_title: "instaclustr_kafka_connect_mirror_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_connect_mirror_v2_instance (Data Source)
Details of a Kafka Connect Mirror
## Example Usage
```
data "instaclustr_kafka_connect_mirror_v2_instance" "example" { 
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
*___connector_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the mirror connector. The value of this property has the form: [source-cluster].[target-cluster].[random-string]<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The overall status of this mirror.<br><br>
*___connectors___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [connectors](#nested--connectors) for nested schema<br>
<br>Detailed list of Connectors for the mirror.<br><br>
*___rename_mirrored_topics___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Whether to rename topics as they are mirrored, by prefixing the sourceCluster.alias to the topic name.<br><br>
*___max_tasks___*<br>
<ins>Type</ins>: integer, read-only<br>
<br>Maximum number of tasks for Kafka Connect to use. Should be greater than 0.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the mirror<br><br>
*___kafka_connect_cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the kafka connect cluster<br><br>
*___target_latency___*<br>
<ins>Type</ins>: integer, read-only<br>
<br>The latency in milliseconds above which this mirror will be considered out of sync. It can not be less than 1000ms. The suggested initial latency is 30000ms  for connectors to be created.<br><br>
*___topics_regex___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Regex to select which topics to mirror.<br><br>
*___mirrored_topics___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [mirrored_topics](#nested--mirrored_topics) for nested schema<br>
<br>Detailed list of Mirrored topics.<br><br>
*___source_cluster___*<br>
<ins>Type</ins>: nested block, read-only, see [source_cluster](#nested--source_cluster) for nested schema<br>
<ins>Constraints</ins>: minimum items: 1<br><br>Details to connect to the source kafka cluster<br><br>
<a id="nested--connectors"></a>
## Nested schema for `connectors`
Detailed list of Connectors for the mirror.<br>
### Read-only attributes
*___config___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Configuration of the connector.<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the connector.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the connector. Could be one of [Mirror Connector, Checkpoint Connector].<br><br>
<a id="nested--external_cluster"></a>
## Nested schema for `external_cluster`
Details to connect to a Non-Instaclustr managed cluster. Cannot be provided if targeting an Instaclustr managed cluster.<br>
### Read-only attributes
*___source_connection_properties___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Kafka connection properties string used to connect to external kafka cluster<br><br>
<a id="nested--managed_cluster"></a>
## Nested schema for `managed_cluster`
Details to connect to a Instaclustr managed cluster. Cannot be provided if targeting a non-Instaclustr managed cluster.<br>
### Read-only attributes
*___source_kafka_cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Source kafka cluster id.<br><br>
*___use_private_ips___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Whether or not to connect to source cluster's private IP addresses.<br><br>
<a id="nested--mirrored_topics"></a>
## Nested schema for `mirrored_topics`
Detailed list of Mirrored topics.<br>
### Read-only attributes
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the mirrored topic.<br><br>
*___average_rate___*<br>
<ins>Type</ins>: number, read-only<br>
<br>Average record rate for messages to travel from source to destination topics, it is 0 if there are no messages travelling in between.<br><br>
*___average_latency___*<br>
<ins>Type</ins>: number, read-only<br>
<br>Average latency in milliseconds for messages to travel from source to destination topics.<br><br>
<a id="nested--source_cluster"></a>
## Nested schema for `source_cluster`
Details to connect to the source kafka cluster<br>
### Read-only attributes
*___external_cluster___*<br>
<ins>Type</ins>: nested block, read-only, see [external_cluster](#nested--external_cluster) for nested schema<br>
<br>Details to connect to a Non-Instaclustr managed cluster. Cannot be provided if targeting an Instaclustr managed cluster.<br><br>
*___alias___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Alias to use for the source kafka cluster. This will be used to rename topics if renameMirroredTopics is turned on<br><br>
*___managed_cluster___*<br>
<ins>Type</ins>: nested block, read-only, see [managed_cluster](#nested--managed_cluster) for nested schema<br>
<br>Details to connect to a Instaclustr managed cluster. Cannot be provided if targeting a non-Instaclustr managed cluster.<br><br>
