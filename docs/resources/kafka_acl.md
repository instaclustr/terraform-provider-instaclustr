---
page_title: "instaclustr_kafka_acl Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Resource  `instaclustr_kafka_acl`
A resource to manage Kafka ACL for a Kafka cluster.

## Import
You can import the existing ACLs of a Kafka cluster created via other ways (e.g., CLI) to your terraform resource.

### Import an ACL
```shell
terraform import instaclustr_kafka_acl.<resource-name> "<cluster-id>&<principal>&<host>&<resourceType>&<resourceName>&<operation>&<permissionType>&<patternType>"
```

Replace `<resource-name>` with the terraform resource name. Replace `<principal>`, `<host>`, `<resourceType>`, `<resourceName>`, `<operation>`, `<permissionType>`, 
and `<patternType>` with the corresponding ACL. Keep this "&" symbol iin the command.

### Example

```shell
terraform import instaclustr_kafka_acl.kafka_acl_test "f4cb7a63-8217-4dc8-ab12-28e54efc00d1&User:test&*&TOPIC&*&ALL&ALLOW&LITERAL"
```

After importing, you should see the imported ACL in ***terraform.tfstate***.

## Usage

### Properties
`instaclustr_kafka_acl`

See https://kafka.apache.org/documentation/#security_authz_primitives for more details on the `resource_type` and `operation`

Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr Kafka managed cluster. |Required
`principal`| The principal part of the ACL, e.g. `User:test` |Required
`host`| The host part of the ACL, e.g., "*" |Required
`resource_type`| The resource type part of the ACL, e.g., `TOPIC` |Required
`resource_name`| The resource name part of the ACL, e.g., if `resource_type` is `TOPIC`, then the resource name corresponds to the topic name |Required
`operation`| The operation part of the ACL, e.g., `DESCRIBE` |Required
`permission_type`| The permission of the ACL: either `ALLOW` or `DENY` |Required
`pattern_type`| The pattern of the ACL: either `LITERAL` or `PREFIXED` |Required

### Example
```
resource "instaclustr_kafka_acl" "kafka_acl_example" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  principal = "User:test"
  host = "*"
  resource_type = "TOPIC"
  resource_name = "*"
  operation = "ALL"
  permission_type = "ALLOW" 
  pattern_type = "LITERAL"
}
```

