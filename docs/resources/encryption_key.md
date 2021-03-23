---
page_title: "instaclustr_encryption_key Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Resource: `instaclustr_encryption_key`  
A resource for managing EBS encryption of nodes with KMS keys. This is only avaliable for clusters hosted with the AWS provider.

#### Properties
Property | Description | Default
---------|-------------|--------
`key_id`|Internal ID of the KMS encryption key. Can be found via GET to `https://api.instaclustr.com/provisioning/v1/encryption-keys`|""
`alias`|KMS key alias, a human-readibly identifier specified alongside your KMS ARN|""
`arn`|KMS ARN, identifier specifying provider, location and key in a ':' value seperated string|""
`key_provider`|For customers running in their own account. Value specifying the provider account’s name, similar to `instaclustr_cluster.cluster_provider.account_name`|INSTACLUSTR

#### Example
```
resource "instaclustr_encryption_key" "example_encryption_key" {
    alias = "virginia 1"
    arn = "arn:aws:kms:us-east-1:123456789012:key12345678-1234-1234-1234-123456789abc"
}
```