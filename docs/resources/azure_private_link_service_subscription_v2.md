---
page_title: "instaclustr_azure_private_link_service_subscription_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_azure_private_link_service_subscription_v2 (Resource)
A listable data source of all Azure Subscriptions allowed to access Azure Private Link Service in an Instaclustr managed cluster data centre.
## Example Usage
```
resource "instaclustr_azure_private_link_service_subscription_v2" "example" {
  reason = ""
  cluster_data_centre_id = "91346037-b969-4c4d-8112-59e1d2aa67e9"
  id = "a1f46d37-d9f9-4c4d-8112-e9e8d2826792"
  azure_subscription_id = "d13ae037-ba12-13ed-8bb2-a9e1d2fa67eb"
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
*___azure_subscription_id___*<br>
<ins>Type</ins>: string (uuid), required, updatable<br>
<br>Resource Id in Azure for subscription for associated access to the private link service<br><br>
*___cluster_data_centre_id___*<br>
<ins>Type</ins>: string (uuid), required, updatable<br>
<br>ID of the cluster data centre<br><br>
### Input attributes - Optional
*___reason___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Message to provide additional details of status<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), optional, updatable<br>
<br>ID used to manage the subscription for the cluster data centre<br><br>
### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The status of subscription for the cluster data centre<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_azure_private_link_service_subscription_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
