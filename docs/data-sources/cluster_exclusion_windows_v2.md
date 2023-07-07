---
page_title: "instaclustr_cluster_exclusion_windows_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_cluster_exclusion_windows_v2 (Data Source)
Definition to the Cluster exclusion windows
## Example Usage
```
data "instaclustr_cluster_exclusion_windows_v2" "example" { 
  cluster_id = "<cluster_id>" // the value of the `cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the cluster<br><br>
### Read-only attributes
*___exclusion_windows___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [exclusion_windows](#nested--exclusion_windows) for nested schema<br>
<br>List of cluster exclusion windows<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the cluster<br><br>
<a id="nested--exclusion_windows"></a>
## Nested schema for `exclusion_windows`
List of cluster exclusion windows<br>
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
