---
page_title: "instaclustr_clickhouse_cluster_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_clickhouse_cluster_v2 (Resource)
Definition of a managed ClickHouse cluster that can be provisioned in Instaclustr.
## Example Usage
```
resource "instaclustr_clickhouse_cluster_v2" "example" {
  data_centre {
    cloud_provider = "AWS_VPC"
    dedicated_click_house_keeper {
      node_count = 3
      node_size = "CLK-DEV-m7i.large-50"
    }

    load_balancer_enabled = false
    name = "AWS_VPC_US_EAST_1"
    network = "10.0.0.0/16"
    node_size = "CLK-DEV-m7i.large-50"
    region = "US_EAST_1"
    replicas = 3
    shards = 1
    tiered_storage {
      s3_settings {
        s3_bucket_name = "sample-s3-bucket"
      }

    }

  }

  private_network_cluster = false
  name = "MyClickHouseCluster"
  clickhouse_version = "[x.y.z]"
  sla_tier = "NON_PRODUCTION"
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
*___clickhouse_version___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `[0-9]+\.[0-9]+\.[0-9]+`<br><br>Version of ClickHouse to run on the cluster. Available versions: <ul> <li>`25.3.3`</li> <li>`23.8.16`</li> <li>`24.3.18`</li> <li>`24.8.14`</li> </ul><br><br>
*___name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`<br><br>Name of the cluster.<br><br>
*___private_network_cluster___*<br>
<ins>Type</ins>: boolean, required, immutable<br>
<br>Creates the cluster with private network only, see [Private Network Clusters](https://www.instaclustr.com/support/documentation/useful-information/private-network-clusters/).<br><br>
### Input attributes - Optional
*___description___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>A description of the cluster<br><br>
*___resize_settings___*<br>
<ins>Type</ins>: nested block, optional, updatable, see [resize_settings](#nested--resize_settings) for nested schema<br>
<br>Settings to determine how resize requests will be performed for the cluster.<br><br>
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
<ins>Constraints</ins>: allowed values: [ `AWS_VPC`, `GCP`, `AZURE`, `AZURE_AZ`, `ONPREMISES` ]<br><br>Name of a cloud provider service.<br><br>
*___region___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Region of the Data Centre. See the description for node size for a compatible Data Centre for a given node size.<br><br>
*___shards___*<br>
<ins>Type</ins>: integer, required, updatable<br>
<ins>Constraints</ins>: minimum: 1, maximum: 1E+2<br><br>Total number of shards in the Data Centre.<br><br>
*___replicas___*<br>
<ins>Type</ins>: integer, required, updatable<br>
<ins>Constraints</ins>: minimum: 1, maximum: 3<br><br>Total number of replicas of data in the Data Centre<br><br>
*___name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>A logical name for the data centre within a cluster. These names must be unique in the cluster.<br><br>
*___node_size___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Size of the nodes provisioned in the Data Centre. Available node sizes: <details> <summary>*Amazon Web Services* [__AWS_VPC__]</summary> <br> <details> <summary>*Africa (Cape Town)* [__AF_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Hong Kong)* [__AP_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Hyderabad)* [__AP_SOUTH_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Jakarta)* [__AP_SOUTHEAST_3__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Malaysia)* [__AP_SOUTHEAST_5__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Melbourne)* [__AP_SOUTHEAST_4__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Mumbai)* [__AP_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Osaka)* [__AP_NORTHEAST_3__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Seoul)* [__AP_NORTHEAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Singapore)* [__AP_SOUTHEAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Sydney)* [__AP_SOUTHEAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Thailand)* [__AP_SOUTHEAST_7__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Asia Pacific (Tokyo)* [__AP_NORTHEAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Canada (Central)* [__CA_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Canada West (Calgary)* [__CA_WEST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU Central (Frankfurt)* [__EU_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU Central (Zurich)* [__EU_CENTRAL_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU North (Stockholm)* [__EU_NORTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU South (Milan)* [__EU_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU South (Spain)* [__EU_SOUTH_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU West (Ireland)* [__EU_WEST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU West (London)* [__EU_WEST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*EU West (Paris)* [__EU_WEST_3__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Israel (Tel Aviv)* [__IL_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Middle East (Bahrain)* [__ME_SOUTH_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Middle East (UAE)* [__ME_CENTRAL_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*South America (São Paulo)* [__SA_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*US East (Northern Virginia)* [__US_EAST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*US East (Ohio)* [__US_EAST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*US West (Northern California)* [__US_WEST_1__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*US West (Oregon)* [__US_WEST_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-t3.small-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-m7i.large-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.2xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.4xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.8xlarge-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-m7i.xlarge-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <br> </details> <details> <summary>*Microsoft Azure* [__AZURE_AZ__]</summary> <br> <details> <summary>*Australia East (NSW)* [__AUSTRALIA_EAST__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Canada Central (Toronto)* [__CANADA_CENTRAL__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Central US (Iowa)* [__CENTRAL_US__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*East US (Virginia)* [__EAST_US__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*East US 2 (Virginia)* [__EAST_US_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*North Europe (Ireland)* [__NORTH_EUROPE__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*South Central US (Texas)* [__SOUTH_CENTRAL_US__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Southeast Asia (Singapore)* [__SOUTHEAST_ASIA__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*Switzerland North (Zurich)* [__SWITZERLAND_NORTH__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*UK South (London)* [__UK_SOUTH__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*West Europe (Netherlands)* [__WEST_EUROPE__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <details> <summary>*West US 2 (Washington)* [__WEST_US_2__]</summary> <br> <table> <tr> <th>Node Size</th> <th>Lifecycle State</th> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DEV-Standard_D2s_v5-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-DEV-Standard_B2ls_v2-30 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-100 </td> <td>General Availability</td> </tr> <tr> <td>CLK-DK-PRD-Standard_B2s_v2-50 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D16s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-1000 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D32s_v5-750 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D4s_v5-500 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-250 </td> <td>General Availability</td> </tr> <tr> <td>CLK-PRD-Standard_D8s_v5-500 </td> <td>General Availability</td> </tr> </table> <br> </details> <br> </details><br><br>
*___network___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The private network address block for the Data Centre specified using CIDR address notation. The network must have a prefix length between `/16` and `/26` and must be part of a private address space.<br><br>
### Input attributes - Optional
*___azure_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [azure_settings](#nested--azure_settings) for nested schema<br>
<br>Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br><br>
*___gcp_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [gcp_settings](#nested--gcp_settings) for nested schema<br>
<br>GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br><br>
*___tiered_storage___*<br>
<ins>Type</ins>: nested block, optional, updatable, see [tiered_storage](#nested--tiered_storage) for nested schema<br>
<br>Enable Tiered Storage for ClickHouse<br><br>
*___load_balancer_enabled___*<br>
<ins>Type</ins>: boolean, optional, immutable<br>
<br>Enable Load Balancer for ClickHouse<br><br>
*___tag___*<br>
<ins>Type</ins>: repeatable nested block, optional, immutable, see [tag](#nested--tag) for nested schema<br>
<br>List of tags to apply to the Data Centre. Tags are metadata labels which  allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require. Note `tag` is not supported in terraform lifecycle `ignore_changes`.<br><br>
*___aws_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [aws_settings](#nested--aws_settings) for nested schema<br>
<br>AWS specific settings for the Data Centre. Cannot be provided with GCP or Azure settings.<br><br>
*___dedicated_click_house_keeper___*<br>
<ins>Type</ins>: nested block, optional, updatable, see [dedicated_click_house_keeper](#nested--dedicated_click_house_keeper) for nested schema<br>
<br>Provision additional dedicated nodes for ClickHouse Keeper to run on. ClickHouse Keeper will be co-located with ClickHouse Server if this is not provided<br><br>
*___provider_account_name___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.<br><br>
### Read-only attributes
*___load_balancer_domain___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Domain of the Network Load Balancer if enabled. Balances requests against nodes.<br><br>
*___current_operations___*<br>
<ins>Type</ins>: nested block, read-only, see [current_operations](#nested--current_operations) for nested schema<br>
<br>Active operations in the data centre.<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the Data Centre.<br><br>
*___deleted_nodes___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [deleted_nodes](#nested--deleted_nodes) for nested schema<br>
<br>List of deleted nodes in the data centre<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the Cluster Data Centre.<br><br>
*___nodes___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [nodes](#nested--nodes) for nested schema<br>
<br>List of non-deleted nodes in the data centre<br><br>
<a id="nested--current_operations"></a>
## Nested schema for `current_operations`
Active operations in the data centre.<br>
### Read-only attributes
*___resize___*<br>
<ins>Type</ins>: nested block, read-only, see [resize](#nested--resize) for nested schema<br>
<br>Active node resize operations<br><br>
*___delete_nodes___*<br>
<ins>Type</ins>: nested block, read-only, see [delete_nodes](#nested--delete_nodes) for nested schema<br>
<br>Latest active delete nodes operation<br><br>
<a id="nested--resize"></a>
## Nested schema for `resize`
Active node resize operations<br>
### Read-only attributes
*___operations___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [operations](#nested--operations) for nested schema<br>
<br>
<a id="nested--azure_settings"></a>
## Nested schema for `azure_settings`
Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br>
### Input attributes - Optional
*___storage_network___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>The private network address block to be used for the storage network. This is only used for certain node sizes, currently limited to those which use Azure NetApp Files: for all other node sizes, this field should not be provided. The network must have a prefix length between /16 and /28, and must be part of a private address range.<br><br>
*___custom_virtual_network_id___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>VNet ID into which the Data Centre will be provisioned. The VNet must have an available address space for the Data Centre's network allocation to be appended to the VNet. Currently supported for PostgreSQL clusters only.<br><br>
*___resource_group___*<br>
<ins>Type</ins>: string, optional, immutable<br>
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
<ins>Constraints</ins>: allowed values: [ `CASSANDRA`, `SPARK_MASTER`, `SPARK_JOBSERVER`, `KAFKA_BROKER`, `KAFKA_DEDICATED_ZOOKEEPER`, `KAFKA_DEDICATED_KRAFT_CONTROLLER`, `KAFKA_ZOOKEEPER`, `KAFKA_SCHEMA_REGISTRY`, `KAFKA_REST_PROXY`, `APACHE_ZOOKEEPER`, `POSTGRESQL`, `PGBOUNCER`, `KAFKA_CONNECT`, `KAFKA_KARAPACE_SCHEMA_REGISTRY`, `KAFKA_KARAPACE_REST_PROXY`, `CADENCE`, `CLICKHOUSE_SERVER`, `CLICKHOUSE_KEEPER`, `CLICKHOUSE_SERVER_AND_KEEPER`, `REDIS_MASTER`, `REDIS_REPLICA`, `VALKEY_MASTER`, `VALKEY_REPLICA`, `OPENSEARCH_DASHBOARDS`, `OPENSEARCH_COORDINATOR`, `OPENSEARCH_MASTER`, `OPENSEARCH_DATA`, `OPENSEARCH_INGEST`, `OPENSEARCH_DATA_AND_INGEST`, `KAFKA_SHOTOVER_PROXY` ]<br><br>The roles or purposes of the node. Useful for filtering for nodes that have a specific role.<br><br>
*___public_address___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Public IP address of the node.<br><br>
<a id="nested--delete_nodes"></a>
## Nested schema for `delete_nodes`
Latest active delete nodes operation<br>
### Input attributes - Optional
*___status___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<ins>Constraints</ins>: allowed values: [ `GENESIS`, `RUNNING`, `COMPLETED`, `CANCELLED`, `FAILED` ]<br><br>
*___number_of_nodes_to_delete___*<br>
<ins>Type</ins>: integer, optional, updatable<br>
<br>Number of nodes set to delete in the operation.<br><br>
### Read-only attributes
*___cdc_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Cluster Data Centre.<br><br>
*___modified___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Timestamp of the last modification of the operation<br><br>
*___delete_node_operations___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [delete_node_operations](#nested--delete_node_operations) for nested schema<br>
<br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>Operation id<br><br>
*___created___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Timestamp of the creation of the operation<br><br>
<a id="nested--gcp_settings"></a>
## Nested schema for `gcp_settings`
GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br>
### Input attributes - Optional
*___custom_virtual_network_id___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Network name or a relative Network or Subnetwork URI.
The Data Centre's network allocation must match the IPv4 CIDR block of the specified subnet.

Examples:
- Network URI: <code>projects/{riyoa-gcp-project-name}/global/networks/{network-name}</code>.
- Network name: <code>{network-name}</code>, equivalent to <code>projects/{riyoa-gcp-project-name}/global/networks/{network-name}</code>.
- Same-project subnetwork URI: <code>projects/{riyoa-gcp-project-name}/regions/{region-id}/subnetworks/{subnetwork-name}</code>.
- Shared VPC subnetwork URI: <code>projects/{riyoa-gcp-host-project-name}/regions/{region-id}/subnetworks/{subnetwork-name}</code>.

<br><br>
<a id="nested--s3_settings"></a>
## Nested schema for `s3_settings`
Defines information about the S3 bucket to be used for remote storage.<br>
### Input attributes - Required
*___s3_bucket_name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>S3 bucket name for ClickHouse remote storage<br><br>
### Input attributes - Optional
*___prefix___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z\d\-_]{0,100}$`<br><br>By default data in the S3 bucket will be stored in a folder named after the cluster's ID. If a prefix is provided, data will be stored in `<prefix>/<cluster_id>` instead<br><br>
<a id="nested--tiered_storage"></a>
## Nested schema for `tiered_storage`
Enable Tiered Storage for ClickHouse<br>
### Input attributes - Optional
*___s3_settings___*<br>
<ins>Type</ins>: nested block, optional, updatable, see [s3_settings](#nested--s3_settings) for nested schema<br>
<br>Defines information about the S3 bucket to be used for remote storage.<br><br>
*___azure_blob_storage_settings___*<br>
<ins>Type</ins>: nested block, optional, updatable, see [azure_blob_storage_settings](#nested--azure_blob_storage_settings) for nested schema<br>
<br>Defines information about the blob storage container to be used for remote storage.<br><br>
<a id="nested--delete_node_operations"></a>
## Nested schema for `delete_node_operations`

### Input attributes - Optional
*___status___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<ins>Constraints</ins>: allowed values: [ `GENESIS`, `RUNNING`, `COMPLETED`, `CANCELLED`, `FAILED` ]<br><br>
### Read-only attributes
*___modified___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Timestamp of the last modification of the operation<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>Operation id<br><br>
*___created___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Timestamp of the creation of the operation<br><br>
*___node_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the node being replaced.<br><br>
<a id="nested--tag"></a>
## Nested schema for `tag`
List of tags to apply to the Data Centre. Tags are metadata labels which  allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require. Note `tag` is not supported in terraform lifecycle `ignore_changes`.<br>
### Input attributes - Required
*___key___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Key of the tag for the Data Centre.<br><br>
### Input attributes - Optional
*___value___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Value of the tag for the Data Centre.<br><br>
<a id="nested--operations"></a>
## Nested schema for `operations`

### Read-only attributes
*___completed___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Timestamp of the completion of the operation.<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `COMPLETED`, `FAILED`, `DELETED`, `IN_PROGRESS`, `UNKNOWN` ]<br><br>Status of the operation<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the operation.<br><br>
*___replace_operations___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [replace_operations](#nested--replace_operations) for nested schema<br>
<br>
*___concurrent_resizes___*<br>
<ins>Type</ins>: integer, read-only<br>
<br>Number of nodes that can be concurrently resized at a given time.<br><br>
*___created___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Timestamp of the creation of the operation<br><br>
*___node_purpose___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Purpose of the node<br><br>
*___instaclustr_support_alerted___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Timestamp of when Instaclustr Support has been alerted to the resize operation.<br><br>
*___new_node_size___*<br>
<ins>Type</ins>: string, read-only<br>
<br>New size of the node.<br><br>
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
<ins>Constraints</ins>: allowed values: [ `CASSANDRA`, `SPARK_MASTER`, `SPARK_JOBSERVER`, `KAFKA_BROKER`, `KAFKA_DEDICATED_ZOOKEEPER`, `KAFKA_DEDICATED_KRAFT_CONTROLLER`, `KAFKA_ZOOKEEPER`, `KAFKA_SCHEMA_REGISTRY`, `KAFKA_REST_PROXY`, `APACHE_ZOOKEEPER`, `POSTGRESQL`, `PGBOUNCER`, `KAFKA_CONNECT`, `KAFKA_KARAPACE_SCHEMA_REGISTRY`, `KAFKA_KARAPACE_REST_PROXY`, `CADENCE`, `CLICKHOUSE_SERVER`, `CLICKHOUSE_KEEPER`, `CLICKHOUSE_SERVER_AND_KEEPER`, `REDIS_MASTER`, `REDIS_REPLICA`, `VALKEY_MASTER`, `VALKEY_REPLICA`, `OPENSEARCH_DASHBOARDS`, `OPENSEARCH_COORDINATOR`, `OPENSEARCH_MASTER`, `OPENSEARCH_DATA`, `OPENSEARCH_INGEST`, `OPENSEARCH_DATA_AND_INGEST`, `KAFKA_SHOTOVER_PROXY` ]<br><br>The roles or purposes of the node. Useful for filtering for nodes that have a specific role.<br><br>
*___public_address___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Public IP address of the node.<br><br>
<a id="nested--replace_operations"></a>
## Nested schema for `replace_operations`

### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `GENESIS`, `RESIZING_DISK`, `RESIZED_DISK`, `EXPANDED_FILESYSTEM`, `GRACEFULLY_SHUTTING_DOWN`, `CREATING_REPLACEMENT`, `PROVISIONING`, `PROVISIONED`, `BACKEDUP`, `RESTORING`, `FLUSHING`, `FLUSHED`, `SWAPPING`, `SWAPPED`, `CLEARING_INSTALLED_BUNDLES`, `CLEARED_INSTALLED_BUNDLES`, `POST_BUNDLE_PROCESSING`, `RESTARTING`, `REPLACED`, `CANCELLED`, `FAILED`, `UNKNOWN` ]<br><br>Status of the node replacement operation.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the node replacement operation.<br><br>
*___created___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Timestamp of the creation of the node replacement operation.<br><br>
*___new_node_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the new node in the replacement operation.<br><br>
*___node_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the node being replaced.<br><br>
<a id="nested--azure_blob_storage_settings"></a>
## Nested schema for `azure_blob_storage_settings`
Defines information about the blob storage container to be used for remote storage.<br>
### Input attributes - Required
*___container_name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Blob storage container name for ClickHouse remote storage<br><br>
<a id="nested--resize_settings"></a>
## Nested schema for `resize_settings`
Settings to determine how resize requests will be performed for the cluster.<br>
### Input attributes - Optional
*___concurrency___*<br>
<ins>Type</ins>: integer, optional, updatable<br>
<br>Number of concurrent nodes to resize during a resize operation.<br><br>
*___notify_support_contacts___*<br>
<ins>Type</ins>: boolean, optional, updatable<br>
<br>Setting this property to `true` will notify the Instaclustr Account's designated support contacts on resize completion.<br><br>
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
<ins>Type</ins>: string, required, updatable<br>
<br>The email address which will be contacted when the cluster is requested to be deleted.<br><br>
### Input attributes - Optional
*___confirmation_phone_number___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>The phone number which will be contacted when the cluster is requested to be delete.<br><br>
<a id="nested--dedicated_click_house_keeper"></a>
## Nested schema for `dedicated_click_house_keeper`
Provision additional dedicated nodes for ClickHouse Keeper to run on. ClickHouse Keeper will be co-located with ClickHouse Server if this is not provided<br>
### Input attributes - Required
*___node_size___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Size of the nodes provisioned as dedicated ClickHouse Keeper nodes.<br><br>
*___node_count___*<br>
<ins>Type</ins>: integer, required, immutable<br>
<ins>Constraints</ins>: minimum: 3, maximum: 3<br><br>Dedicated ClickHouse Keeper node count, it must be 3.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_clickhouse_cluster_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
