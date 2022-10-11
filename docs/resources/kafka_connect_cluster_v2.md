---
page_title: "instaclustr_kafka_connect_cluster_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_connect_cluster_v2 (Resource)
Data type for kafka connect specific properties
## Example Usage
```
resource "instaclustr_kafka_connect_cluster_v2" "example" {
  pci_compliance_mode = false
  data_centre {
    cloud_provider = "AWS_VPC"
    name = "AWS_VPC_US_EAST_1"
    network = "10.0.0.0/16"
    node_size = "KCN-PRD-c6g.4xlarge-80"
    number_of_nodes = 3
    region = "US_EAST_1"
    replication_factor = 3
  }

  private_network_cluster = false
  kafka_connect_version = "[x.y.z]"
  target_cluster {
    managed_cluster {
      kafka_connect_vpc_type = "VPC_PEERED"
      target_kafka_cluster_id = "UUID of kafka cluster"
    }

  }

  name = "MyKafkaConnectCluster"
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
*___kafka_connect_version___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `[0-9]+\.[0-9]+\.[0-9]+`<br><br>Version of Kafka connect to run on the cluster. Available versions: <ul> <li>`3.1.2`</li> <li>`3.0.2`</li> <li>`2.8.2`</li> <li>`2.7.1`</li> </ul><br><br>
*___name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`<br><br>Name of the cluster.<br><br>
*___private_network_cluster___*<br>
<ins>Type</ins>: boolean, required, immutable<br>
<br>Creates the cluster with private network only, see [Private Network Clusters](https://www.instaclustr.com/support/documentation/useful-information/private-network-clusters/).<br><br>
*___target_cluster___*<br>
<ins>Type</ins>: nested block, required, immutable, see [target_cluster](#nested--target_cluster) for nested schema<br>
<br>Details to connect to a target Kafka Cluster cluster.<br><br>
### Input attributes - Optional
*___custom_connectors___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [custom_connectors](#nested--custom_connectors) for nested schema<br>
<br>Defines the location for custom connector storage and access info.<br><br>
*___two_factor_delete___*<br>
<ins>Type</ins>: nested block, optional, updatable, see [two_factor_delete](#nested--two_factor_delete) for nested schema<br>
<br>
*___pci_compliance_mode___*<br>
<ins>Type</ins>: boolean, optional, immutable<br>
<br>Creates a PCI compliant cluster, see [PCI Compliance](https://www.instaclustr.com/support/documentation/useful-information/pci-compliance/).<br><br>
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
*___number_of_nodes___*<br>
<ins>Type</ins>: integer, required, updatable<br>
<br>Total number of nodes in the Data Centre. Must be a multiple of `replicationFactor`.<br><br>
*___replication_factor___*<br>
<ins>Type</ins>: integer, required, immutable<br>
<br>Number of racks to use when allocating nodes.<br><br>
*___region___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Region of the Data Centre. See the description for node size for a compatible Data Centre for a given node size.<br><br>
*___name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>A logical name for the data centre within a cluster. These names must be unique in the cluster.<br><br>
*___node_size___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Size of the nodes provisioned in the Data Centre. --AVAILABLE_NODE_SIZES_MARKER_V2_APACHE_KAFKA_CONNECT--<br><br>
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
<a id="nested--custom_connectors"></a>
## Nested schema for `custom_connectors`
Defines the location for custom connector storage and access info.<br>
### Input attributes - Optional
*___azure_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [azure_settings](#nested--azure_settings) for nested schema<br>
<br>Defines the information to access custom connectors located in an azure storage container.<br><br>
*___gcp_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [gcp_settings](#nested--gcp_settings) for nested schema<br>
<br>Defines the information to access custom connectors located in a gcp storage container.<br><br>
*___aws_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [aws_settings](#nested--aws_settings) for nested schema<br>
<br>Defines the information to access custom connectors located in a S3 bucket.<br><br>
<a id="nested--azure_settings"></a>
## Nested schema for `azure_settings`
Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br>
### Input attributes - Optional
*___resource_group___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>The name of the Azure Resource Group into which the Data Centre will be provisioned.<br><br>
<a id="nested--azure_settings"></a>
## Nested schema for `azure_settings`
Defines the information to access custom connectors located in an azure storage container.<br>
### Input attributes - Required
*___storage_container_name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Azure storage container name for Kafka Connect custom connector.<br><br>
*___storage_account_name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Azure storage account name to access your Azure bucket for Kafka Connect custom connector.<br><br>
*___storage_account_key___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Azure storage account key to access your Azure bucket for Kafka Connect custom connector.<br><br>
<a id="nested--external_cluster"></a>
## Nested schema for `external_cluster`
Details to connect to a Non-Instaclustr managed cluster.<br>
### Input attributes - Optional
*___ssl_truststore_password___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___sasl_jaas_config___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___truststore___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Base64 encoded version of the TLS trust store (in JKS format) used to connect to your Kafka Cluster. Only required if connecting to a Non-Instaclustr managed Kafka Cluster with TLS enabled.<br><br>
*___ssl_enabled_protocols___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___sasl_mechanism___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___ssl_protocol___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___security_protocol___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
*___bootstrap_servers___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Connection information for your Kafka Cluster. These options are analogous to the similarly named options that you would place in your Kafka Connect worker.properties file. Only required if connecting to a Non-Instaclustr managed Kafka Cluster.<br><br>
<a id="nested--gcp_settings"></a>
## Nested schema for `gcp_settings`
GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br>
### Input attributes - Optional
*___custom_virtual_network_id___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Network name or a relative Network or Subnetwork URI e.g. projects/my-project/regions/us-central1/subnetworks/my-subnet. The Data Centre's network allocation must match the IPv4 CIDR block of the specified subnet.<br><br>
<a id="nested--gcp_settings"></a>
## Nested schema for `gcp_settings`
Defines the information to access custom connectors located in a gcp storage container.<br>
### Input attributes - Required
*___project_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___client_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___client_email___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___storage_bucket_name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___private_key_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
*___private_key___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Access information for the GCP Storage bucket for kafka connect custom connectors.<br><br>
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
<ins>Constraints</ins>: allowed values: [ `CASSANDRA`, `SPARK_MASTER`, `SPARK_JOBSERVER`, `KAFKA_BROKER`, `KAFKA_DEDICATED_ZOOKEEPER`, `KAFKA_ZOOKEEPER`, `KAFKA_SCHEMA_REGISTRY`, `KAFKA_REST_PROXY`, `KAFKA_CONNECT` ]<br><br>The roles or purposes of the node. Useful for filtering for nodes that have a specific role.<br><br>
*___public_address___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Public IP address of the node.<br><br>
<a id="nested--managed_cluster"></a>
## Nested schema for `managed_cluster`
Details to connect to a Instaclustr managed cluster.<br>
### Input attributes - Required
*___target_kafka_cluster_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Target kafka cluster id.<br><br>
*___kafka_connect_vpc_type___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Available options are KAFKA_VPC, VPC_PEERED, SEPARATE_VPC<br><br>
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
<a id="nested--aws_settings"></a>
## Nested schema for `aws_settings`
Defines the information to access custom connectors located in a S3 bucket.<br>
### Input attributes - Required
*___secret_key___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>AWS Secret Key associated with the Access Key id that can access your specified S3 bucket for Kafka Connect custom connector.<br><br>
*___access_key___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>AWS Access Key id that can access your specified S3 bucket for Kafka Connect custom connector.<br><br>
*___s3_bucket_name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>S3 bucket name for Kafka Connect custom connector.<br><br>
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
<a id="nested--target_cluster"></a>
## Nested schema for `target_cluster`
Details to connect to a target Kafka Cluster cluster.<br>
### Input attributes - Optional
*___external_cluster___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [external_cluster](#nested--external_cluster) for nested schema<br>
<br>Details to connect to a Non-Instaclustr managed cluster.<br><br>
*___managed_cluster___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [managed_cluster](#nested--managed_cluster) for nested schema<br>
<br>Details to connect to a Instaclustr managed cluster.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_kafka_connect_cluster_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
