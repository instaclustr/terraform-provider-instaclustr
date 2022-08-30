---
page_title: "instaclustr_aws_encryption_keys_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_encryption_keys_v2 (Data Source)

## Schema
### account_id<br>
<ins>Type</ins>: string<br>

### encryption_keys<br>
<ins>Type</ins>: block list, see [encryption_keys](#nested--encryption_keys) for nested schema<br>

<a id="nested--encryption_keys"></a>
## Nested schema for `encryption_keys`<br>

### in_use<br>
<ins>Type</ins>: boolean<br>
<br>Whether the encryption key is used by a cluster.
### alias<br>
<ins>Type</ins>: string<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]{1}[a-zA-Z0-9 _-]*$`<br><br>Encryption key alias for display purposes.
### id<br>
<ins>Type</ins>: string (uuid)<br>
<br>ID of the encryption key.
### arn<br>
<ins>Type</ins>: string<br>
<br>AWS ARN for the encryption key.
### provider_account_name<br>
<ins>Type</ins>: string<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.

