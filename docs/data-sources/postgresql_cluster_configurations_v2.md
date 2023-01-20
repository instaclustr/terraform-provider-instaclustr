---
page_title: "instaclustr_postgresql_cluster_configurations_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_postgresql_cluster_configurations_v2 (Data Source)
PostgreSQL configuration properties
## Example Usage
```
data "instaclustr_postgresql_cluster_configurations_v2" "example" { 
  postgresql_cluster_id = "<postgresql_cluster_id>" // the value of the `postgresql_cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___postgresql_cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the PostgreSQL cluster.<br><br>
### Read-only attributes
*___configuration_properties___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [configuration_properties](#nested--configuration_properties) for nested schema<br>
<br>List of configuration properties<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the PostgreSQL cluster<br><br>
<a id="nested--configuration_properties"></a>
## Nested schema for `configuration_properties`
List of configuration properties<br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the PostgreSQL configuration property. The value of this property has the form: [cluster-id]|[configuration_name]<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the configuration property.<br><br>
*___value___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Value of the configuration property.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Id of the PostgreSQL cluster.<br><br>
