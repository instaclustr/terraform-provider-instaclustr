---
page_title: "instaclustr_valkey_cluster_users_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_valkey_cluster_users_v2 (Data Source)
A listable data source of all Valkey users within a Valkey Cluster.
## Example Usage
```
data "instaclustr_valkey_cluster_users_v2" "example" { 
  valkey_cluster_id = "<valkey_cluster_id>" // the value of the `valkey_cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___valkey_cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the cluster.<br><br>
### Read-only attributes
*___valkey_users___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [valkey_users](#nested--valkey_users) for nested schema<br>
<br>List of all Valkey users in the cluster.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the valkey cluster<br><br>
<a id="nested--valkey_users"></a>
## Nested schema for `valkey_users`
List of all Valkey users in the cluster.<br>
### Read-only attributes
*___username___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^(?![zZ][oO][oO][kK][eE][eE][pP][eE][rR]$)[a-zA-Z0-9][a-zA-Z0-9_-]*$`<br><br>Username of the Valkey user.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the Valkey user. The value of this property has the form: [cluster-id]_[valkey-username]<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Valkey cluster.<br><br>