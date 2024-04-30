---
page_title: "instaclustr_zookeeper_cluster_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_zookeeper_cluster_v2_instance (Data Source)
Definition of a managed Apache Zookeeper cluster that can be provisioned in Instaclustr.
## Example Usage
```
data "instaclustr_zookeeper_cluster_v2_instance" "example" { 
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
*___data_centre___*<br>
<ins>Type</ins>: nested block, read-only, see [data_centre](#nested--data_centre) for nested schema<br>
<ins>Constraints</ins>: minimum items: 1<br><br>List of data centre settings.<br><br>
*___description___*<br>
<ins>Type</ins>: string, read-only<br>
<br>A description of the cluster<br><br>
*___sla_tier___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `PRODUCTION`, `NON_PRODUCTION` ]<br><br>SLA Tier of the cluster. Non-production clusters may receive lower priority support and reduced SLAs. Production tier is not available when using Developer class nodes. See [SLA Tier](https://www.instaclustr.com/support/documentation/useful-information/sla-tier/) for more information.<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the cluster.<br><br>
*___zookeeper_version___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `[0-9]+\.[0-9]+\.[0-9]+`<br><br>Version of Apache Zookeeper to run on the cluster. Available versions: <ul> <li>`3.7.2`</li> <li>`3.8.4`</li> <li>`3.8.2`</li> </ul><br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the cluster.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`<br><br>Name of the cluster.<br><br>
*___private_network_cluster___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Creates the cluster with private network only, see [Private Network Clusters](https://www.instaclustr.com/support/documentation/useful-information/private-network-clusters/).<br><br>
*___two_factor_delete___*<br>
<ins>Type</ins>: nested block, read-only, see [two_factor_delete](#nested--two_factor_delete) for nested schema<br>
<br>
*___current_cluster_operation_status___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `NO_OPERATION`, `OPERATION_IN_PROGRESS`, `OPERATION_FAILED` ]<br><br>Indicates if the cluster is currently performing any restructuring operation such as being created or resized<br><br>
<a id="nested--data_centre"></a>
## Nested schema for `data_centre`
List of data centre settings.<br>
### Read-only attributes
*___custom_subject_alternative_names___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<br>List of Subject Alternative Names FQDNs as per RFC 1035.  Used by the applications with self signed certificates in keystores of nodes in the datacenter.<br><br>
*___cloud_provider___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `AWS_VPC`, `GCP`, `AZURE`, `AZURE_AZ`, `ONPREMISES` ]<br><br>Name of a cloud provider service.<br><br>
*___number_of_nodes___*<br>
<ins>Type</ins>: integer, read-only<br>
<br>Total number of Zookeeper nodes in the Data Centre.<br><br>
*___enforce_auth_schemes___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<ins>Constraints</ins>: allowed values: [ `NONE`, `SASL` ]<br><br>A list of authentication schemes to enforce when enforce.auth.enabled=true.<br><br>
*___region___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Region of the Data Centre. See the description for node size for a compatible Data Centre for a given node size.<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the Data Centre.<br><br>
*___azure_settings___*<br>
<ins>Type</ins>: nested block, read-only, see [azure_settings](#nested--azure_settings) for nested schema<br>
<br>Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br><br>
*___deleted_nodes___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [deleted_nodes](#nested--deleted_nodes) for nested schema<br>
<br>List of deleted nodes in the data centre<br><br>
*___gcp_settings___*<br>
<ins>Type</ins>: nested block, read-only, see [gcp_settings](#nested--gcp_settings) for nested schema<br>
<br>GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br><br>
*___enforce_auth_enabled___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Enables Enforced SASL Authentication.<br><br>
*___client_to_server_encryption___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Enables Client ⇄ Node Encryption.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the Cluster Data Centre.<br><br>
*___tag___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [tag](#nested--tag) for nested schema<br>
<br>List of tags to apply to the Data Centre. Tags are metadata labels which  allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require. Note `tag` is not supported in terraform lifecycle `ignore_changes`.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>A logical name for the data centre within a cluster. These names must be unique in the cluster.<br><br>
*___nodes___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [nodes](#nested--nodes) for nested schema<br>
<br>List of non-deleted nodes in the data centre<br><br>
*___node_size___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Size of the nodes provisioned in the Data Centre. Available node sizes: <details> <summary>*Amazon Web Services* [__AWS_VPC__]</summary> <br> <details> <summary>*Africa (Cape Town)* [__AF_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Hong Kong)* [__AP_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Jakarta)* [__AP_SOUTHEAST_3__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Melbourne)* [__AP_SOUTHEAST_4__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Mumbai)* [__AP_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Osaka)* [__AP_NORTHEAST_3__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Seoul)* [__AP_NORTHEAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Singapore)* [__AP_SOUTHEAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Sydney)* [__AP_SOUTHEAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Tokyo)* [__AP_NORTHEAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Canada (Central)* [__CA_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU Central (Frankfurt)* [__EU_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU Central (Zurich)* [__EU_CENTRAL_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU North (Stockholm)* [__EU_NORTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU South (Milan)* [__EU_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU South (Spain)* [__EU_SOUTH_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU West (Ireland)* [__EU_WEST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU West (London)* [__EU_WEST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU West (Paris)* [__EU_WEST_3__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Middle East (Bahrain)* [__ME_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Middle East (UAE)* [__ME_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*South America (São Paulo)* [__SA_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*US East (Northern Virginia)* [__US_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*US East (Ohio)* [__US_EAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*US West (Northern California)* [__US_WEST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*US West (Oregon)* [__US_WEST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>MZK-DEV-t4g.small-20 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.large-60 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6g.xlarge-120 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.large-75 </td> <td>General Availability</td> </tr> <tr> <td>MZK-PRD-m6gd.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.large-75 </td> <td>General Availability</td> </tr> <tr> <td>ZKR-PRD-m5d.xlarge-150 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-developer-t3.small-20 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.large-60 </td> <td>General Availability</td> </tr> <tr> <td>zookeeper-production-m5.xlarge-120 </td> <td>General Availability</td> </tr> </table> <br> </details> <br> </details><br><br>
*___aws_settings___*<br>
<ins>Type</ins>: nested block, read-only, see [aws_settings](#nested--aws_settings) for nested schema<br>
<br>AWS specific settings for the Data Centre. Cannot be provided with GCP or Azure settings.<br><br>
*___network___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The private network address block for the Data Centre specified using CIDR address notation. The network must have a prefix length between `/12` and `/22` and must be part of a private address space.<br><br>
*___provider_account_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.<br><br>
<a id="nested--azure_settings"></a>
## Nested schema for `azure_settings`
Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br>
### Read-only attributes
*___storage_network___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The private network address block to be used for the storage network. This is only used for certain node sizes, currently limited to those which use Azure NetApp Files: for all other node sizes, this field should not be provided. The network must have a prefix length between /16 and /28, and must be part of a private address range.<br><br>
*___custom_virtual_network_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>VNet ID into which the Data Centre will be provisioned. The VNet must have an available address space for the Data Centre's network allocation to be appended to the VNet. Currently supported for PostgreSQL clusters only.<br><br>
*___resource_group___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The name of the Azure Resource Group into which the Data Centre will be provisioned.<br><br>
<a id="nested--deleted_nodes"></a>
## Nested schema for `deleted_nodes`
List of deleted nodes in the data centre<br>
### Read-only attributes
*___start_time___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Start time of the node as a UTC timestamp<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Provisioning status of the node.<br><br>
*___deletion_time___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Deletion time of the node as a UTC timestamp<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the node.<br><br>
*___rack___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Rack name in which the node is located.<br><br>
*___node_size___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Size of the node.<br><br>
*___private_address___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Private IP address of the node.<br><br>
*___node_roles___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<ins>Constraints</ins>: allowed values: [ `CASSANDRA`, `SPARK_MASTER`, `SPARK_JOBSERVER`, `KAFKA_BROKER`, `KAFKA_DEDICATED_ZOOKEEPER`, `KAFKA_DEDICATED_KRAFT_CONTROLLER`, `KAFKA_ZOOKEEPER`, `KAFKA_SCHEMA_REGISTRY`, `KAFKA_REST_PROXY`, `APACHE_ZOOKEEPER`, `POSTGRESQL`, `PGBOUNCER`, `KAFKA_CONNECT`, `KAFKA_KARAPACE_SCHEMA_REGISTRY`, `KAFKA_KARAPACE_REST_PROXY`, `CADENCE`, `CLICKHOUSE_SERVER`, `CLICKHOUSE_KEEPER`, `CLICKHOUSE_SERVER_AND_KEEPER`, `COUCHBASE_DATA`, `COUCHBASE_INDEX`, `COUCHBASE_QUERY`, `COUCHBASE_SEARCH`, `COUCHBASE_EVENTING`, `COUCHBASE_ANALYTICS`, `MONGODB`, `REDIS_MASTER`, `REDIS_REPLICA`, `OPENSEARCH_DASHBOARDS`, `OPENSEARCH_COORDINATOR`, `OPENSEARCH_MASTER`, `OPENSEARCH_DATA`, `OPENSEARCH_INGEST`, `OPENSEARCH_DATA_AND_INGEST` ]<br><br>The roles or purposes of the node. Useful for filtering for nodes that have a specific role.<br><br>
*___public_address___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Public IP address of the node.<br><br>
<a id="nested--gcp_settings"></a>
## Nested schema for `gcp_settings`
GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br>
### Read-only attributes
*___custom_virtual_network_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Network name or a relative Network or Subnetwork URI.
The Data Centre's network allocation must match the IPv4 CIDR block of the specified subnet.

Examples:
- Network URI: <code>projects/{riyoa-gcp-project-name}/global/networks/{network-name}</code>.
- Network name: <code>{network-name}</code>, equivalent to <code>projects/{riyoa-gcp-project-name}/global/networks/{network-name}</code>.
- Same-project subnetwork URI: <code>projects/{riyoa-gcp-project-name}/regions/{region-id}/subnetworks/{subnetwork-name}</code>.
- Shared VPC subnetwork URI: <code>projects/{riyoa-gcp-host-project-name}/regions/{region-id}/subnetworks/{subnetwork-name}</code>.
<br><br>
<a id="nested--tag"></a>
## Nested schema for `tag`
List of tags to apply to the Data Centre. Tags are metadata labels which  allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require. Note `tag` is not supported in terraform lifecycle `ignore_changes`.<br>
### Read-only attributes
*___key___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Key of the tag to be added to the Data Centre.<br><br>
*___value___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Value of the tag to be added to the Data Centre.<br><br>
<a id="nested--nodes"></a>
## Nested schema for `nodes`
List of non-deleted nodes in the data centre<br>
### Read-only attributes
*___start_time___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Start time of the node as a UTC timestamp<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Provisioning status of the node.<br><br>
*___deletion_time___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Deletion time of the node as a UTC timestamp<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the node.<br><br>
*___rack___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Rack name in which the node is located.<br><br>
*___node_size___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Size of the node.<br><br>
*___private_address___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Private IP address of the node.<br><br>
*___node_roles___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<ins>Constraints</ins>: allowed values: [ `CASSANDRA`, `SPARK_MASTER`, `SPARK_JOBSERVER`, `KAFKA_BROKER`, `KAFKA_DEDICATED_ZOOKEEPER`, `KAFKA_DEDICATED_KRAFT_CONTROLLER`, `KAFKA_ZOOKEEPER`, `KAFKA_SCHEMA_REGISTRY`, `KAFKA_REST_PROXY`, `APACHE_ZOOKEEPER`, `POSTGRESQL`, `PGBOUNCER`, `KAFKA_CONNECT`, `KAFKA_KARAPACE_SCHEMA_REGISTRY`, `KAFKA_KARAPACE_REST_PROXY`, `CADENCE`, `CLICKHOUSE_SERVER`, `CLICKHOUSE_KEEPER`, `CLICKHOUSE_SERVER_AND_KEEPER`, `COUCHBASE_DATA`, `COUCHBASE_INDEX`, `COUCHBASE_QUERY`, `COUCHBASE_SEARCH`, `COUCHBASE_EVENTING`, `COUCHBASE_ANALYTICS`, `MONGODB`, `REDIS_MASTER`, `REDIS_REPLICA`, `OPENSEARCH_DASHBOARDS`, `OPENSEARCH_COORDINATOR`, `OPENSEARCH_MASTER`, `OPENSEARCH_DATA`, `OPENSEARCH_INGEST`, `OPENSEARCH_DATA_AND_INGEST` ]<br><br>The roles or purposes of the node. Useful for filtering for nodes that have a specific role.<br><br>
*___public_address___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Public IP address of the node.<br><br>
<a id="nested--aws_settings"></a>
## Nested schema for `aws_settings`
AWS specific settings for the Data Centre. Cannot be provided with GCP or Azure settings.<br>
### Read-only attributes
*___custom_virtual_network_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>VPC ID into which the Data Centre will be provisioned. The Data Centre's network allocation must match the IPv4 CIDR block of the specified VPC.<br><br>
*___ebs_encryption_key___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of a KMS encryption key to encrypt data on nodes. KMS encryption key must be set in Cluster Resources through the Instaclustr Console before provisioning an encrypted Data Centre.<br><br>
<a id="nested--two_factor_delete"></a>
## Nested schema for `two_factor_delete`

### Read-only attributes
*___confirmation_phone_number___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The phone number which will be contacted when the cluster is requested to be delete.<br><br>
*___confirmation_email___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The email address which will be contacted when the cluster is requested to be deleted.<br><br>
