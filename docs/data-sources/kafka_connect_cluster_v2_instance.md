---
page_title: "instaclustr_kafka_connect_cluster_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_connect_cluster_v2_instance (Data Source)
Data type for kafka connect specific properties
## Example Usage
```
data "instaclustr_kafka_connect_cluster_v2_instance" "example" { 
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
*___sla_tier___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `PRODUCTION`, `NON_PRODUCTION` ]<br><br>SLA Tier of the cluster. Non-production clusters may receive lower priority support and reduced SLAs. Production tier is not available when using Developer class nodes. See [SLA Tier](https://www.instaclustr.com/support/documentation/useful-information/sla-tier/) for more information.<br><br>
*___custom_connectors___*<br>
<ins>Type</ins>: nested block, read-only, see [custom_connectors](#nested--custom_connectors) for nested schema<br>
<br>Defines the location for custom connector storage and access info.<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the cluster.<br><br>
*___kafka_connect_version___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `[0-9]+\.[0-9]+\.[0-9]+`<br><br>Version of Kafka connect to run on the cluster. Available versions: <ul> <li>`3.1.2`</li> <li>`3.0.2`</li> <li>`2.8.2`</li> <li>`2.7.1`</li> </ul><br><br>
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
*___target_cluster___*<br>
<ins>Type</ins>: nested block, read-only, see [target_cluster](#nested--target_cluster) for nested schema<br>
<br>Details to connect to a target Kafka Cluster cluster.<br><br>
<a id="nested--data_centre"></a>
## Nested schema for `data_centre`
List of data centre settings.<br>
### Read-only attributes
*___cloud_provider___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `AWS_VPC`, `GCP`, `AZURE`, `AZURE_AZ` ]<br><br>Name of the cloud provider service in which the Data Centre will be provisioned.<br><br>
*___number_of_nodes___*<br>
<ins>Type</ins>: integer, read-only<br>
<br>Total number of nodes in the Data Centre. Must be a multiple of `replicationFactor`.<br><br>
*___replication_factor___*<br>
<ins>Type</ins>: integer, read-only<br>
<br>Number of racks to use when allocating nodes.<br><br>
*___region___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Region of the Data Centre. See the description for node size for a compatible Data Centre for a given node size.<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the Data Centre.<br><br>
*___azure_settings___*<br>
<ins>Type</ins>: nested block, read-only, see [azure_settings](#nested--azure_settings) for nested schema<br>
<br>Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br><br>
*___gcp_settings___*<br>
<ins>Type</ins>: nested block, read-only, see [gcp_settings](#nested--gcp_settings) for nested schema<br>
<br>GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the Cluster Data Centre.<br><br>
*___tag___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [tag](#nested--tag) for nested schema<br>
<br>List of tags to apply to the Data Centre. Tags are metadata labels which  allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>A logical name for the data centre within a cluster. These names must be unique in the cluster.<br><br>
*___nodes___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [nodes](#nested--nodes) for nested schema<br>
<br>
*___node_size___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Size of the nodes provisioned in the Data Centre. Available node sizes: <details> <summary>*Amazon Web Services* [__AWS_VPC__]</summary> <br> <details> <summary>*Africa (Cape Town)* [__AF_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>c5.2xlarge-20-gp2</td> </tr> <tr> <td>c5.4xlarge-40-gp2</td> </tr> <tr> <td>c5.xlarge-10-gp2</td> </tr> <tr> <td>r5.2xlarge-20-gp2</td> </tr> <tr> <td>r5.4xlarge-40-gp2</td> </tr> <tr> <td>r5.xlarge-10-gp2</td> </tr> <tr> <td>t3.medium-10-gp2</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Hong Kong)* [__AP_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Mumbai)* [__AP_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Seoul)* [__AP_NORTHEAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Singapore)* [__AP_SOUTHEAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Sydney)* [__AP_SOUTHEAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Tokyo)* [__AP_NORTHEAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*Canada (Central)* [__CA_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*EU Central (Frankfurt)* [__EU_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*EU North (Stockholm)* [__EU_NORTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*EU South (Milan)* [__EU_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*EU West (Ireland)* [__EU_WEST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*EU West (London)* [__EU_WEST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*EU West (Paris)* [__EU_WEST_3__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*Middle East (Bahrain)* [__ME_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>c5.2xlarge-20-gp2</td> </tr> <tr> <td>c5.4xlarge-40-gp2</td> </tr> <tr> <td>c5.xlarge-10-gp2</td> </tr> <tr> <td>r5.2xlarge-20-gp2</td> </tr> <tr> <td>r5.4xlarge-40-gp2</td> </tr> <tr> <td>r5.xlarge-10-gp2</td> </tr> <tr> <td>t3.medium-10-gp2</td> </tr> </table> <br> </details> <details> <summary>*South America (São Paulo)* [__SA_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*US East (Northern Virginia)* [__US_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*US East (Ohio)* [__US_EAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*US West (Northern California)* [__US_WEST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <details> <summary>*US West (Oregon)* [__US_WEST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>KCN-DEV-t4g.medium-30</td> </tr> <tr> <td>KCN-PRD-c6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-c6g.large-80</td> </tr> <tr> <td>KCN-PRD-c6g.xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.2xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.4xlarge-80</td> </tr> <tr> <td>KCN-PRD-r6g.large-80</td> </tr> <tr> <td>KCN-PRD-r6g.xlarge-80</td> </tr> <tr> <td>c5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>c5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>c5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>r5.2xlarge-20-gp2 (deprecated)</td> </tr> <tr> <td>r5.4xlarge-40-gp2 (deprecated)</td> </tr> <tr> <td>r5.xlarge-10-gp2 (deprecated)</td> </tr> <tr> <td>t3.medium-10-gp2 (deprecated)</td> </tr> </table> <br> </details> <br> </details> <details> <summary>*Microsoft Azure* [__AZURE_AZ__]</summary> <br> <details> <summary>*Australia East (NSW)* [__AUSTRALIA_EAST__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D16s_v3-40</td> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> <tr> <td>Standard_D8s_v3-20</td> </tr> </table> <br> </details> <details> <summary>*Canada Central (Toronto)* [__CANADA_CENTRAL__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D16s_v3-40</td> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> <tr> <td>Standard_D8s_v3-20</td> </tr> </table> <br> </details> <details> <summary>*Central US (Iowa)* [__CENTRAL_US__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D16s_v3-40</td> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> <tr> <td>Standard_D8s_v3-20</td> </tr> </table> <br> </details> <details> <summary>*East US (Virginia)* [__EAST_US__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D16s_v3-40</td> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> <tr> <td>Standard_D8s_v3-20</td> </tr> </table> <br> </details> <details> <summary>*East US 2 (Virginia)* [__EAST_US_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> </table> <br> </details> <details> <summary>*North Europe (Ireland)* [__NORTH_EUROPE__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D16s_v3-40</td> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> <tr> <td>Standard_D8s_v3-20</td> </tr> </table> <br> </details> <details> <summary>*South Central US (Texas)* [__SOUTH_CENTRAL_US__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D16s_v3-40</td> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> <tr> <td>Standard_D8s_v3-20</td> </tr> </table> <br> </details> <details> <summary>*Southeast Asia (Singapore)* [__SOUTHEAST_ASIA__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D16s_v3-40</td> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> <tr> <td>Standard_D8s_v3-20</td> </tr> </table> <br> </details> <details> <summary>*West Europe (Netherlands)* [__WEST_EUROPE__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D16s_v3-40</td> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> <tr> <td>Standard_D8s_v3-20</td> </tr> </table> <br> </details> <details> <summary>*West US 2 (Washington)* [__WEST_US_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_D2s_v3-10</td> </tr> <tr> <td>Standard_D4s_v3-10</td> </tr> </table> <br> </details> <br> </details> <details> <summary>*Google Cloud Platform* [__GCP__]</summary> <br> <details> <summary>*Central US (Iowa)* [__us-central1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Eastern Asia-Pacific (Taiwan)* [__asia-east1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Eastern South America (Brazil)* [__southamerica-east1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Eastern US (North Virginia)* [__us-east4__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Eastern US (South Carolina)* [__us-east1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Northeastern Asia-pacific (Japan)* [__asia-northeast1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Northeastern North America (Canada)* [__northamerica-northeast1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Northern Europe (Finland)* [__europe-north1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Southeastern Asia (Singapore)* [__asia-southeast1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Southeastern Australia (Sydney)* [__australia-southeast1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Southern Asia (India)* [__asia-south1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (Belgium)* [__europe-west1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (England)* [__europe-west2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (Germany)* [__europe-west3__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (Netherlands)* [__europe-west4__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (Zurich)* [__europe-west6__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Western US (California)* [__us-west2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <details> <summary>*Western US (Oregon)* [__us-west1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-standard-16-40</td> </tr> <tr> <td>n1-standard-2-10</td> </tr> <tr> <td>n1-standard-4-10</td> </tr> <tr> <td>n1-standard-8-20</td> </tr> </table> <br> </details> <br> </details><br><br>
*___aws_settings___*<br>
<ins>Type</ins>: nested block, read-only, see [aws_settings](#nested--aws_settings) for nested schema<br>
<br>AWS specific settings for the Data Centre. Cannot be provided with GCP or Azure settings.<br><br>
*___network___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The private network address block for the Data Centre specified using CIDR address notation. The network must have a prefix length between `/12` and `/22` and must be part of a private address space.<br><br>
*___provider_account_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.<br><br>
<a id="nested--azure_connector_settings"></a>
## Nested schema for `azure_connector_settings`
Defines the information to access custom connectors located in an azure storage container. Cannot be provided if custom connectors are stored in GCP or AWS.<br>
### Read-only attributes
*___storage_container_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Azure storage container name for Kafka Connect custom connector.<br><br>
*___storage_account_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Azure storage account name to access your Azure bucket for Kafka Connect custom connector.<br><br>
*___storage_account_key___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Azure storage account key to access your Azure bucket for Kafka Connect custom connector.<br><br>
<a id="nested--gcp_connector_settings"></a>
## Nested schema for `gcp_connector_settings`
Defines the information to access custom connectors located in a gcp storage container. Cannot be provided if custom connectors are stored in AWS or AZURE.<br>
### Read-only attributes
*___project_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___client_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___client_email___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___storage_bucket_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___private_key_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___private_key___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
<a id="nested--custom_connectors"></a>
## Nested schema for `custom_connectors`
Defines the location for custom connector storage and access info.<br>
### Read-only attributes
*___azure_connector_settings___*<br>
<ins>Type</ins>: nested block, read-only, see [azure_connector_settings](#nested--azure_connector_settings) for nested schema<br>
<br>Defines the information to access custom connectors located in an azure storage container. Cannot be provided if custom connectors are stored in GCP or AWS.<br><br>
*___gcp_connector_settings___*<br>
<ins>Type</ins>: nested block, read-only, see [gcp_connector_settings](#nested--gcp_connector_settings) for nested schema<br>
<br>Defines the information to access custom connectors located in a gcp storage container. Cannot be provided if custom connectors are stored in AWS or AZURE.<br><br>
*___aws_connector_settings___*<br>
<ins>Type</ins>: nested block, read-only, see [aws_connector_settings](#nested--aws_connector_settings) for nested schema<br>
<br>Defines the information to access custom connectors located in a S3 bucket. Cannot be provided if custom connectors are stored in GCP or AZURE.<br><br>
<a id="nested--azure_settings"></a>
## Nested schema for `azure_settings`
Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br>
### Read-only attributes
*___resource_group___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The name of the Azure Resource Group into which the Data Centre will be provisioned.<br><br>
<a id="nested--external_cluster"></a>
## Nested schema for `external_cluster`
Details to connect to a Non-Instaclustr managed cluster. Cannot be provided if targeting an Instaclustr managed cluster.<br>
### Read-only attributes
*___ssl_truststore_password___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___sasl_jaas_config___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___truststore___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Base64 encoded version of the TLS trust store (in JKS format) used to connect to your Kafka Cluster. Only required if connecting to a Non-Instaclustr managed Kafka Cluster with TLS enabled.<br><br>
*___ssl_enabled_protocols___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___sasl_mechanism___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___ssl_protocol___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___security_protocol___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___bootstrap_servers___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
<a id="nested--gcp_settings"></a>
## Nested schema for `gcp_settings`
GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br>
### Read-only attributes
*___custom_virtual_network_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Network name or a relative Network or Subnetwork URI e.g. projects/my-project/regions/us-central1/subnetworks/my-subnet. The Data Centre's network allocation must match the IPv4 CIDR block of the specified subnet.<br><br>
<a id="nested--tag"></a>
## Nested schema for `tag`
List of tags to apply to the Data Centre. Tags are metadata labels which  allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require.<br>
### Read-only attributes
*___key___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Key of the tag to be added to the Data Centre.<br><br>
*___value___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Value of the tag to be added to the Data Centre.<br><br>
<a id="nested--nodes"></a>
## Nested schema for `nodes`

### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Provisioning status of the node.<br><br>
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
<ins>Constraints</ins>: allowed values: [ `CASSANDRA`, `SPARK_MASTER`, `SPARK_JOBSERVER`, `KAFKA_BROKER`, `KAFKA_DEDICATED_ZOOKEEPER`, `KAFKA_ZOOKEEPER`, `KAFKA_SCHEMA_REGISTRY`, `KAFKA_REST_PROXY`, `APACHE_ZOOKEEPER`, `POSTGRESQL`, `PGBOUNCER`, `KAFKA_CONNECT`, `KAFKA_KARAPACE_SCHEMA_REGISTRY`, `KAFKA_KARAPACE_REST_PROXY`, `CADENCE`, `MONGODB`, `REDIS_MASTER`, `REDIS_REPLICA`, `OPENSEARCH_DASHBOARDS`, `OPENSEARCH_COORDINATOR`, `OPENSEARCH_MASTER`, `OPENSEARCH_DATA_AND_INGEST` ]<br><br>The roles or purposes of the node. Useful for filtering for nodes that have a specific role.<br><br>
*___public_address___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Public IP address of the node.<br><br>
<a id="nested--managed_cluster"></a>
## Nested schema for `managed_cluster`
Details to connect to a Instaclustr managed cluster. Cannot be provided if targeting an external cluster.<br>
### Read-only attributes
*___target_kafka_cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Target kafka cluster id.<br><br>
*___kafka_connect_vpc_type___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Available options are KAFKA_VPC, VPC_PEERED, SEPARATE_VPC<br><br>
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
<a id="nested--aws_connector_settings"></a>
## Nested schema for `aws_connector_settings`
Defines the information to access custom connectors located in a S3 bucket. Cannot be provided if custom connectors are stored in GCP or AZURE.<br>
### Read-only attributes
*___secret_key___*<br>
<ins>Type</ins>: string, read-only<br>
<br>AWS Secret Key associated with the Access Key id that can access your specified S3 bucket for Kafka Connect custom connector.<br><br>
*___access_key___*<br>
<ins>Type</ins>: string, read-only<br>
<br>AWS Access Key id that can access your specified S3 bucket for Kafka Connect custom connector.<br><br>
*___s3_bucket_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>S3 bucket name for Kafka Connect custom connector.<br><br>
<a id="nested--two_factor_delete"></a>
## Nested schema for `two_factor_delete`

### Read-only attributes
*___confirmation_phone_number___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The phone number which will be contacted when the cluster is requested to be delete.<br><br>
*___confirmation_email___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The email address which will be contacted when the cluster is requested to be deleted.<br><br>
<a id="nested--target_cluster"></a>
## Nested schema for `target_cluster`
Details to connect to a target Kafka Cluster cluster.<br>
### Read-only attributes
*___external_cluster___*<br>
<ins>Type</ins>: nested block, read-only, see [external_cluster](#nested--external_cluster) for nested schema<br>
<br>Details to connect to a Non-Instaclustr managed cluster. Cannot be provided if targeting an Instaclustr managed cluster.<br><br>
*___managed_cluster___*<br>
<ins>Type</ins>: nested block, read-only, see [managed_cluster](#nested--managed_cluster) for nested schema<br>
<br>Details to connect to a Instaclustr managed cluster. Cannot be provided if targeting an external cluster.<br><br>
