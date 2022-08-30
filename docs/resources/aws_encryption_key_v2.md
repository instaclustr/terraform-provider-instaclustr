---
page_title: "instaclustr_aws_encryption_key_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_encryption_key_v2 (Resource)

## Schema
### in_use<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Whether the encryption key is used by a cluster.
### alias<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]{1}[a-zA-Z0-9 _-]*$`<br><br>Encryption key alias for display purposes.
### id<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the encryption key.
### arn<br>
<ins>Type</ins>: string, required, immutable<br>
<br>AWS ARN for the encryption key.
### provider_account_name<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.

