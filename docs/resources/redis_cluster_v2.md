---
page_title: "instaclustr_redis_cluster_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_redis_cluster_v2 (Resource)
Definition of a managed Redis cluster that can be provisioned in Instaclustr.
## Example Usage
```
resource "instaclustr_redis_cluster_v2" "example" {
  client_to_node_encryption = true
  redis_version = "[x.y.z]"
  pci_compliance_mode = false
  data_centre {
    cloud_provider = "AWS_VPC"
    master_nodes = 3
    name = "AWS_VPC_US_EAST_1"
    network = "10.0.0.0/16"
    node_size = "r6g.large-100-r"
    region = "US_EAST_1"
    replica_nodes = 3
  }

  private_network_cluster = false
  password_and_user_auth = true
  name = "MyTestCluster"
  sla_tier = "PRODUCTION"
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
*___data_centre___*<br>
<ins>Type</ins>: nested block, required, updatable, see [data_centre](#nested--data_centre) for nested schema<br>
<ins>Constraints</ins>: minimum items: 1<br><br>List of data centre settings.<br><br>
*___sla_tier___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: allowed values: [ `PRODUCTION`, `NON_PRODUCTION` ]<br><br>SLA Tier of the cluster. Non-production clusters may receive lower priority support and reduced SLAs. Production tier is not available when using Developer class nodes. See [SLA Tier](https://www.instaclustr.com/support/documentation/useful-information/sla-tier/) for more information.<br><br>
*___redis_version___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `[0-9]+\.[0-9]+\.[0-9]+`<br><br>Version of Redis to run on the cluster. Available versions: <ul> <li>`6.2.7`</li> <li>`7.0.5`</li> </ul><br><br>
*___name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`<br><br>Name of the cluster.<br><br>
*___client_to_node_encryption___*<br>
<ins>Type</ins>: boolean, required, immutable<br>
<br>Enables Client ⇄ Node Encryption.<br><br>
*___private_network_cluster___*<br>
<ins>Type</ins>: boolean, required, immutable<br>
<br>Creates the cluster with private network only, see [Private Network Clusters](https://www.instaclustr.com/support/documentation/useful-information/private-network-clusters/).<br><br>
*___password_and_user_auth___*<br>
<ins>Type</ins>: boolean, required, immutable<br>
<br>Enables Password Authentication and User Authorization.<br><br>
*___pci_compliance_mode___*<br>
<ins>Type</ins>: boolean, required, immutable<br>
<br>Creates a PCI compliant cluster, see [PCI Compliance](https://www.instaclustr.com/support/documentation/useful-information/pci-compliance/).<br><br>
### Input attributes - Optional
*___two_factor_delete___*<br>
<ins>Type</ins>: nested block, optional, updatable, see [two_factor_delete](#nested--two_factor_delete) for nested schema<br>
<br>
### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the cluster.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the cluster.<br><br>
*___current_cluster_operation_status___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `NO_OPERATION`, `OPERATION_IN_PROGRESS`, `OPERATION_FAILED` ]<br><br>Indicates if the cluster is currently performing any restructuring operation such as being created or resized<br><br>
<a id="nested--data_centre"></a>
## Nested schema for `data_centre`
List of data centre settings.<br>
### Input attributes - Required
*___cloud_provider___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: allowed values: [ `AWS_VPC`, `GCP`, `AZURE`, `AZURE_AZ` ]<br><br>Name of the cloud provider service in which the Data Centre will be provisioned.<br><br>
*___region___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Region of the Data Centre. See the description for node size for a compatible Data Centre for a given node size.<br><br>
*___master_nodes___*<br>
<ins>Type</ins>: integer, required, updatable<br>
<ins>Constraints</ins>: minimum: 3, maximum: 1E+2<br><br>Total number of master nodes in the Data Centre.<br><br>
*___name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>A logical name for the data centre within a cluster. These names must be unique in the cluster.<br><br>
*___node_size___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Size of the nodes provisioned in the Data Centre. Available node sizes: <details> <summary>*Google Cloud Platform* [__GCP__]</summary> <br> <details> <summary>*Central US (Iowa)* [__us-central1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Eastern Asia-Pacific (Taiwan)* [__asia-east1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Eastern South America (Brazil)* [__southamerica-east1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Eastern US (North Virginia)* [__us-east4__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Eastern US (South Carolina)* [__us-east1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Northeastern Asia-pacific (Japan)* [__asia-northeast1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Northeastern North America (Canada)* [__northamerica-northeast1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Northern Europe (Finland)* [__europe-north1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Southeastern Asia (Singapore)* [__asia-southeast1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Southeastern Australia (Sydney)* [__australia-southeast1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Southern Asia (India)* [__asia-south1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (Belgium)* [__europe-west1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (England)* [__europe-west2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (Germany)* [__europe-west3__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (Netherlands)* [__europe-west4__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Western Europe (Zurich)* [__europe-west6__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Western US (California)* [__us-west2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <details> <summary>*Western US (Oregon)* [__us-west1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>n1-highmem-16-600-r</td> </tr> <tr> <td>n1-highmem-2-100-r</td> </tr> <tr> <td>n1-highmem-4-200-r</td> </tr> <tr> <td>n1-highmem-8-400-r</td> </tr> <tr> <td>n1-standard-1-30-r</td> </tr> </table> <br> </details> <br> </details> <details> <summary>*Microsoft Azure* [__AZURE_AZ__]</summary> <br> <details> <summary>*Australia East (NSW)* [__AUSTRALIA_EAST__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_DS2_v2-64-r</td> </tr> <tr> <td>Standard_E16s_v3-800-r</td> </tr> <tr> <td>Standard_E2s_v3-100-r</td> </tr> <tr> <td>Standard_E4s_v3-200-r</td> </tr> <tr> <td>Standard_E8s_v3-400-r</td> </tr> </table> <br> </details> <details> <summary>*Canada Central (Toronto)* [__CANADA_CENTRAL__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_DS2_v2-64-r</td> </tr> <tr> <td>Standard_E16s_v3-800-r</td> </tr> <tr> <td>Standard_E2s_v3-100-r</td> </tr> <tr> <td>Standard_E4s_v3-200-r</td> </tr> <tr> <td>Standard_E8s_v3-400-r</td> </tr> </table> <br> </details> <details> <summary>*Central US (Iowa)* [__CENTRAL_US__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_DS2_v2-64-r</td> </tr> <tr> <td>Standard_E16s_v3-800-r</td> </tr> <tr> <td>Standard_E2s_v3-100-r</td> </tr> <tr> <td>Standard_E4s_v3-200-r</td> </tr> <tr> <td>Standard_E8s_v3-400-r</td> </tr> </table> <br> </details> <details> <summary>*East US (Virginia)* [__EAST_US__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_DS2_v2-64-r</td> </tr> <tr> <td>Standard_E16s_v3-800-r</td> </tr> <tr> <td>Standard_E2s_v3-100-r</td> </tr> <tr> <td>Standard_E4s_v3-200-r</td> </tr> <tr> <td>Standard_E8s_v3-400-r</td> </tr> </table> <br> </details> <details> <summary>*East US 2 (Virginia)* [__EAST_US_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_DS2_v2-64-r</td> </tr> <tr> <td>Standard_E16s_v3-800-r</td> </tr> <tr> <td>Standard_E2s_v3-100-r</td> </tr> <tr> <td>Standard_E4s_v3-200-r</td> </tr> <tr> <td>Standard_E8s_v3-400-r</td> </tr> </table> <br> </details> <details> <summary>*North Europe (Ireland)* [__NORTH_EUROPE__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_DS2_v2-64-r</td> </tr> <tr> <td>Standard_E16s_v3-800-r</td> </tr> <tr> <td>Standard_E2s_v3-100-r</td> </tr> <tr> <td>Standard_E4s_v3-200-r</td> </tr> <tr> <td>Standard_E8s_v3-400-r</td> </tr> </table> <br> </details> <details> <summary>*South Central US (Texas)* [__SOUTH_CENTRAL_US__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_DS2_v2-64-r</td> </tr> <tr> <td>Standard_E16s_v3-800-r</td> </tr> <tr> <td>Standard_E2s_v3-100-r</td> </tr> <tr> <td>Standard_E4s_v3-200-r</td> </tr> <tr> <td>Standard_E8s_v3-400-r</td> </tr> </table> <br> </details> <details> <summary>*Southeast Asia (Singapore)* [__SOUTHEAST_ASIA__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_DS2_v2-64-r</td> </tr> <tr> <td>Standard_E16s_v3-800-r</td> </tr> <tr> <td>Standard_E2s_v3-100-r</td> </tr> <tr> <td>Standard_E4s_v3-200-r</td> </tr> <tr> <td>Standard_E8s_v3-400-r</td> </tr> </table> <br> </details> <details> <summary>*West Europe (Netherlands)* [__WEST_EUROPE__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>Standard_DS2_v2-64-r</td> </tr> <tr> <td>Standard_E16s_v3-800-r</td> </tr> <tr> <td>Standard_E2s_v3-100-r</td> </tr> <tr> <td>Standard_E4s_v3-200-r</td> </tr> <tr> <td>Standard_E8s_v3-400-r</td> </tr> </table> <br> </details> <br> </details> <details> <summary>*Amazon Web Services* [__AWS_VPC__]</summary> <br> <details> <summary>*Africa (Cape Town)* [__AF_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Hong Kong)* [__AP_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Mumbai)* [__AP_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Seoul)* [__AP_NORTHEAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Singapore)* [__AP_SOUTHEAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Sydney)* [__AP_SOUTHEAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Tokyo)* [__AP_NORTHEAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*Canada (Central)* [__CA_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*EU Central (Frankfurt)* [__EU_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*EU North (Stockholm)* [__EU_NORTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*EU South (Milan)* [__EU_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*EU West (Ireland)* [__EU_WEST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*EU West (London)* [__EU_WEST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*EU West (Paris)* [__EU_WEST_3__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*Middle East (Bahrain)* [__ME_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*South America (São Paulo)* [__SA_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*US East (Northern Virginia)* [__US_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*US East (Ohio)* [__US_EAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*US West (Northern California)* [__US_WEST_1__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <details> <summary>*US West (Oregon)* [__US_WEST_2__]</summary> <br> <table> <tr> <th>Node Size</th> </tr> <tr> <td>r5.2xlarge-400-r</td> </tr> <tr> <td>r5.4xlarge-600-r</td> </tr> <tr> <td>r5.8xlarge-1000-r</td> </tr> <tr> <td>r5.large-100-r</td> </tr> <tr> <td>r5.xlarge-200-r</td> </tr> <tr> <td>r6g.2xlarge-400-r</td> </tr> <tr> <td>r6g.4xlarge-600-r</td> </tr> <tr> <td>r6g.large-100-r</td> </tr> <tr> <td>r6g.xlarge-200-r</td> </tr> <tr> <td>t3.medium-80-r</td> </tr> <tr> <td>t3.small-20-r</td> </tr> </table> <br> </details> <br> </details><br><br>
*___replica_nodes___*<br>
<ins>Type</ins>: integer, required, immutable<br>
<ins>Constraints</ins>: minimum: 0, maximum: 1E+2<br><br>Total number of replica nodes in the Data Centre.<br><br>
*___network___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The private network address block for the Data Centre specified using CIDR address notation. The network must have a prefix length between `/12` and `/22` and must be part of a private address space.<br><br>
### Input attributes - Optional
*___azure_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [azure_settings](#nested--azure_settings) for nested schema<br>
<br>Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br><br>
*___gcp_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [gcp_settings](#nested--gcp_settings) for nested schema<br>
<br>GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br><br>
*___tag___*<br>
<ins>Type</ins>: repeatable nested block, optional, immutable, see [tag](#nested--tag) for nested schema<br>
<br>List of tags to apply to the Data Centre. Tags are metadata labels which  allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require.<br><br>
*___aws_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [aws_settings](#nested--aws_settings) for nested schema<br>
<br>AWS specific settings for the Data Centre. Cannot be provided with GCP or Azure settings.<br><br>
*___provider_account_name___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.<br><br>
### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the Data Centre.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the Cluster Data Centre.<br><br>
*___nodes___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [nodes](#nested--nodes) for nested schema<br>
<br>
<a id="nested--azure_settings"></a>
## Nested schema for `azure_settings`
Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br>
### Input attributes - Optional
*___resource_group___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>The name of the Azure Resource Group into which the Data Centre will be provisioned.<br><br>
<a id="nested--gcp_settings"></a>
## Nested schema for `gcp_settings`
GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br>
### Input attributes - Optional
*___custom_virtual_network_id___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Network name or a relative Network or Subnetwork URI e.g. projects/my-project/regions/us-central1/subnetworks/my-subnet. The Data Centre's network allocation must match the IPv4 CIDR block of the specified subnet.<br><br>
<a id="nested--tag"></a>
## Nested schema for `tag`
List of tags to apply to the Data Centre. Tags are metadata labels which  allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require.<br>
### Input attributes - Required
*___key___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Key of the tag to be added to the Data Centre.<br><br>
*___value___*<br>
<ins>Type</ins>: string, required, immutable<br>
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
<a id="nested--aws_settings"></a>
## Nested schema for `aws_settings`
AWS specific settings for the Data Centre. Cannot be provided with GCP or Azure settings.<br>
### Input attributes - Optional
*___custom_virtual_network_id___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>VPC ID into which the Data Centre will be provisioned. The Data Centre's network allocation must match the IPv4 CIDR block of the specified VPC.<br><br>
*___ebs_encryption_key___*<br>
<ins>Type</ins>: string (uuid), optional, immutable<br>
<br>ID of a KMS encryption key to encrypt data on nodes. KMS encryption key must be set in Cluster Resources through the Instaclustr Console before provisioning an encrypted Data Centre.<br><br>
<a id="nested--two_factor_delete"></a>
## Nested schema for `two_factor_delete`

### Input attributes - Required
*___confirmation_email___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The email address which will be contacted when the cluster is requested to be deleted.<br><br>
### Input attributes - Optional
*___confirmation_phone_number___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>The phone number which will be contacted when the cluster is requested to be delete.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_redis_cluster_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.