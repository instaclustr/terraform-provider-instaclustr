---
page_title: "instaclustr_kafka_topic_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_topic_v2 (Resource)
Definition of a Kafka Topic to be applied to a Kafka cluster.
## Example Usage
```
resource "instaclustr_kafka_topic_v2" "example" {
  partitions = 3
  configs {
    key = "compression.type"
    value = "producer"
  }

  configs {
    key = "leader.replication.throttled.replicas"
    value = ""
  }

  configs {
    key = "message.downconversion.enable"
    value = "false"
  }

  configs {
    key = "min.insync.replicas"
    value = "1"
  }

  configs {
    key = "segment.jitter.ms"
    value = "0"
  }

  configs {
    key = "cleanup.policy"
    value = "delete"
  }

  configs {
    key = "flush.ms"
    value = "9223372036854775807"
  }

  configs {
    key = "follower.replication.throttled.replicas"
    value = ""
  }

  configs {
    key = "segment.bytes"
    value = "1073741824"
  }

  configs {
    key = "retention.ms"
    value = "604800000"
  }

  configs {
    key = "flush.messages"
    value = "9223372036854775807"
  }

  configs {
    key = "message.format.version"
    value = "3.0-IV1"
  }

  configs {
    key = "max.compaction.lag.ms"
    value = "9223372036854775807"
  }

  configs {
    key = "file.delete.delay.ms"
    value = "60000"
  }

  configs {
    key = "max.message.bytes"
    value = "1048588"
  }

  configs {
    key = "min.compaction.lag.ms"
    value = "0"
  }

  configs {
    key = "message.timestamp.type"
    value = "CreateTime"
  }

  configs {
    key = "preallocate"
    value = "false"
  }

  configs {
    key = "min.cleanable.dirty.ratio"
    value = "0.5"
  }

  configs {
    key = "index.interval.bytes"
    value = "4096"
  }

  configs {
    key = "unclean.leader.election.enable"
    value = "false"
  }

  configs {
    key = "retention.bytes"
    value = "-1"
  }

  configs {
    key = "delete.retention.ms"
    value = "86400000"
  }

  configs {
    key = "segment.ms"
    value = "604800000"
  }

  configs {
    key = "message.timestamp.difference.max.ms"
    value = "9223372036854775807"
  }

  configs {
    key = "segment.index.bytes"
    value = "10485760"
  }

  replication_factor = 3
  topic = "topic-test"
  cluster_id = "c1af59c6-ba0e-4cc2-a0f3-65cee17a5f37"
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
*___replication_factor___*<br>
<ins>Type</ins>: integer, required, immutable<br>
<br>Replication factor for Topic<br><br>
*___partitions___*<br>
<ins>Type</ins>: integer, required, immutable<br>
<br>Topic partition count<br><br>
*___topic___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Kafka Topic name<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the Kafka cluster<br><br>
### Input attributes - Optional
*___configs___*<br>
<ins>Type</ins>: repeatable nested block, optional, updatable, see [configs](#nested--configs) for nested schema<br>
<br>List of Kafka topic configs which have non-default values. These could be set by terraform or other methods like kafka cli etc.<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the Kafka topic. The value of this property has the form: [cluster-id]_[kafka-topic]<br><br>
<a id="nested--configs"></a>
## Nested schema for `configs`
List of Kafka topic configs which have non-default values. These could be set by terraform or other methods like kafka cli etc.<br>
### Input attributes - Required
*___key___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Kafka Topic config key<br><br>
*___value___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Kafka Topic config value<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_kafka_topic_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
