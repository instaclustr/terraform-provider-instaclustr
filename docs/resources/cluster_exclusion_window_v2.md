---
page_title: "instaclustr_cluster_exclusion_window_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_cluster_exclusion_window_v2 (Resource)
Definition to Cluster exclusion window
## Example Usage
```
resource "instaclustr_cluster_exclusion_window_v2" "example" {
  day_of_week = "TUESDAY"
  start_hour = 14
  duration_in_hours = 6
  cluster_id = "cf4fccf3-2ac0-494b-9f40-e95288dd752d"
  id = "af4fccf3-2ac0-494b-9f40-e95288dd752d"
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
### Input attributes - Optional
*___start_hour___*<br>
<ins>Type</ins>: integer (int32), optional, updatable<br>
<ins>Constraints</ins>: minimum: 0, maximum: 23<br><br>The hour of the day that this exclusion window starts on<br><br>
*___duration_in_hours___*<br>
<ins>Type</ins>: integer (int32), optional, updatable<br>
<ins>Constraints</ins>: minimum: 1<br><br>The duration (in hours) of this exclusion window<br><br>
*___day_of_week___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<ins>Constraints</ins>: allowed values: [ `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY` ]<br><br>The day of the week that this exclusion window starts on<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), optional, updatable<br>
<br>ID of the Cluster exclusion window<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), optional, updatable<br>
<br>Cluster Id for the cluster that this exclusion window relates to<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_cluster_exclusion_window_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
