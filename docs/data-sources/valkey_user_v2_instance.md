---
page_title: "instaclustr_valkey_user_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_valkey_user_v2_instance (Data Source)
Definition of a Valkey User to be applied to a Valkey cluster.
## Example Usage
```
data "instaclustr_valkey_user_v2_instance" "example" { 
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
*___username___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^(?![zZ][oO][oO][kK][eE][eE][pP][eE][rR]$)[a-zA-Z0-9][a-zA-Z0-9_-]*$`<br><br>Username of the Valkey user.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the Valkey user. The value of this property has the form: [cluster-id]_[valkey-username]<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Valkey cluster.<br><br>
*___password___*<br>
<ins>Type</ins>: string (password), read-only<br>
<br>Password for the Valkey user.<br><br>
*___initial_permissions___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `standard`, `read-only`, `none` ]<br><br>Permissions initially granted to Valkey user upon creation.<br><br>
