---
page_title: "instaclustr_cluster_exclusion_window_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_cluster_exclusion_window_v2_instance (Data Source)
Definition to Cluster exclusion window
## Example Usage
```
data "instaclustr_cluster_exclusion_window_v2_instance" "example" { 
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
*___start_hour___*<br>
<ins>Type</ins>: integer (int32), read-only<br>
<ins>Constraints</ins>: minimum: 0, maximum: 23<br><br>The hour of the day that this exclusion window starts on<br><br>
*___duration_in_hours___*<br>
<ins>Type</ins>: integer (int32), read-only<br>
<ins>Constraints</ins>: minimum: 1<br><br>The duration (in hours) of this exclusion window<br><br>
*___day_of_week___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY` ]<br><br>The day of the week that this exclusion window starts on<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Cluster exclusion window<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>Cluster Id for the cluster that this exclusion window relates to<br><br>
