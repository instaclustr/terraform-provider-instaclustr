---
page_title: "instaclustr_azure_private_link_service_subscription_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_azure_private_link_service_subscription_v2_instance (Data Source)
A listable data source of all Azure Subscriptions allowed to access Azure Private Link Service in an Instaclustr managed cluster data centre.
## Example Usage
```
data "instaclustr_azure_private_link_service_subscription_v2_instance" "example" { 
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
*___reason___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Message to provide additional details of status<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The status of subscription for the cluster data centre<br><br>
*___azure_subscription_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>Resource Id in Azure for subscription for associated access to the private link service<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID used to manage the subscription for the cluster data centre<br><br>
*___cluster_data_centre_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the cluster data centre<br><br>
