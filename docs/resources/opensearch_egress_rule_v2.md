---
page_title: "instaclustr_opensearch_egress_rule_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_opensearch_egress_rule_v2 (Resource)
Defines an OpenSearch egress rule
## Example Usage
```
resource "instaclustr_opensearch_egress_rule_v2" "example" {
  open_search_binding_id = "qzPJmIQBGW3Cho0V3Ee_"
  cluster_id = "71e4380e-32ac-4fa7-ab42-c165fe35aa55"
  source = "NOTIFICATIONS"
  type = "WEBHOOK"
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
*___source___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: allowed values: [ `ALERTING`, `NOTIFICATIONS` ]<br><br>Source OpenSearch plugin that manages the channel/destination<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>OpenSearch cluster Id<br><br>
*___open_search_binding_id___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: pattern: `[\w-]+`<br><br>OpenSearch ID for alerting/notifications channel/destination for webhook<br><br>
### Input attributes - Optional
*___type___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<ins>Constraints</ins>: allowed values: [ `SLACK`, `WEBHOOK`, `CUSTOM_WEBHOOK`, `CHIME` ]<br><br>Type of the channel/destination<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `[a-zA-Z\d-]+~\w+~[\w-]+`<br><br>Instaclustr id of the egress rule in the format `{clusterId}~{source}~{bindingId}`<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of channel/desination assosciated with webhook<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_opensearch_egress_rule_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
