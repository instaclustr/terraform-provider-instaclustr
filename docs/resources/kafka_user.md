---
page_title: "instaclustr_kafka_user Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Resource  `instaclustr_kafka_user` and Data Source `instaclustr_kafka_user_list`                             
Resources for managing Kafka user for a Kafka cluster. 
Kafka user list is a read-only data source used to get the list of kafka user in a cluster, 
while Kafka user is a resource used to create, update password, and delete Kafka users.

Note: There is a possibility that someone else delete the Kafka user through other means, e.g., instaclustr console and API.
This will cause the Kafka user resource to be not in sync because it does not have read.
In such case, to recreate the user just remove the Kafka user resource, and then create a new one.

#### Properties
`instaclustr_kafka_user`

Property | Description | Default
---------|-------------|--------
cluster_id|The ID of an existing Instaclustr Kafka managed cluster. |Required
username|User name for the Kafka user|Required
password|Password for the Kafka user|Required
initial_permissions|Initial permission set (ACL) associated with this user. Possible values are: `standard`, `read-only`, and `none`. | `none`

`instaclustr_kafka_user_list`

Property | Description | Default
---------|-------------|--------
cluster_id|The ID of an existing Instaclustr Kafka managed cluster. |Required

#### Example
```
resource "instaclustr_kafka_user" "kafka_user_charlie" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  username = "charlie"
  password = "charlie1!"
  initial_permissions = "none"
}


data "instaclustr_kafka_user_list" "kafka_user_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}
```
