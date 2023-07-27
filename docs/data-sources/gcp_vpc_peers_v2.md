---
page_title: "instaclustr_gcp_vpc_peers_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_gcp_vpc_peers_v2 (Data Source)
A listable data source of all GCP VPC Peering requests in an Instaclustr Account.
## Example Usage
```
data "instaclustr_gcp_vpc_peers_v2" "example" { }
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Read-only attributes
*___peering_requests___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [peering_requests](#nested--peering_requests) for nested schema<br>
<br>
*___account_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>UUID of the Instaclustr Account.<br><br>
<a id="nested--peering_requests"></a>
## Nested schema for `peering_requests`

### Read-only attributes
*___data_centre_vpc_network_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Vpc Network Name of the Data Centre VPC.<br><br>
*___peer_subnets___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<br>The subnets for the peering VPC.<br><br>
*___peer_vpc_network_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The name of the VPC Network you wish to peer to.<br><br>
*___cdc_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Cluster Data Centre.<br><br>
*___peer_project_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The project ID of the owner of the accepter VPC.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the VPC peering connection.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the Peering Connection.<br><br>
*___data_centre_project_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>GCP Project ID of the Data Centre.<br><br>
