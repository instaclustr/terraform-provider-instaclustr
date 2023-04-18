---
page_title: "instaclustr_opensearch_cluster_egress_rules_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_opensearch_cluster_egress_rules_v2 (Data Source)
List of OpenSearch egress rules
## Example Usage
```
data "instaclustr_opensearch_cluster_egress_rules_v2" "example" { 
  opensearch_cluster_id = "<opensearch_cluster_id>" // the value of the `opensearch_cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___opensearch_cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>Id of the OpenSearch cluster<br><br>
### Read-only attributes
*___rules___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [rules](#nested--rules) for nested schema<br>
<br>List of OpenSearch egress rules<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>OpenSearch cluster id<br><br>
<a id="nested--rules"></a>
## Nested schema for `rules`
List of OpenSearch egress rules<br>
### Read-only attributes
*___source___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `ALERTING`, `NOTIFICATIONS` ]<br><br>Source OpenSearch plugin that manages the channel/destination<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `[a-zA-Z\d-]+~\w+~[\w-]+`<br><br>Instaclustr id of the egress rule in the format `{clusterId}~{source}~{bindingId}`<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of channel/desination assosciated with webhook<br><br>
*___type___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `SLACK`, `WEBHOOK`, `CUSTOM_WEBHOOK`, `CHIME` ]<br><br>Type of the channel/destination<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>OpenSearch cluster Id<br><br>
*___open_search_binding_id___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `[\w-]+`<br><br>OpenSearch ID for alerting/notifications channel/destination for webhook<br><br>
