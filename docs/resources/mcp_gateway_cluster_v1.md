---
page_title: "instaclustr_mcp_gateway_cluster_v1 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_mcp_gateway_cluster_v1 (Resource)
Definition of a managed MCP Gateway cluster that can be provisioned in Instaclustr.
## Example Usage
```
resource "instaclustr_mcp_gateway_cluster_v1" "example" {
  data_centre {
    cloud_provider = "AWS_VPC"
    name = "AWS_VPC_US_EAST_1"
    network = "10.0.0.0/16"
    node_size = "MCP-DEV-t4g.small-30"
    number_of_nodes = 1
    region = "US_EAST_1"
  }

  mcp_gateway_version = "[x.y.z]"
  private_network_cluster = false
  name = "MyMCPGatewayCluster"
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
*___mcp_gateway_version___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `[0-9]+\.[0-9]+\.[0-9]+`<br><br>Version of MCP Gateway to run on the cluster. Available versions: <ul> <li>`1.0.1`</li> </ul><br><br>
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
*___number_of_nodes___*<br>
<ins>Type</ins>: integer, required, updatable<br>
<ins>Constraints</ins>: minimum: 1<br><br>Total number of MCP gateway nodes in the Data Centre.<br><br>
*___region___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Region of the Data Centre. See the description for node size for a compatible Data Centre for a given node size.<br><br>
*___name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>A logical name for the data centre within a cluster. These names must be unique in the cluster.<br><br>
*___node_size___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Size of the nodes provisioned in the Data Centre. Available versions: <ul> <li>`1.0.1`</li> </ul><br><br>
*___network___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The private network address block for the Data Centre specified using CIDR address notation. The network must have a prefix length between `/16` and `/26` and must be part of a private address space.<br><br>
### Input attributes - Optional
*___zero_inbound_access___*<br>
<ins>Type</ins>: boolean, optional, immutable<br>
<br>Zero Inbound Access gateways eliminate the requirement for any public IP addresses in cluster deployment.<br><br>
*___tag_management_enabled___*<br>
<ins>Type</ins>: boolean, optional, updatable<br>
<br>(Optional) Enable tag management for the data centre, allowing you to create, update and delete custom tags on the data centre via Instaclustr Terraform Provider v2, Cluster Management API or Management Console. Tag management is only available for RIYOA clusters and cannot be disabled once enabled. If not specified, the current value will remain unchanged.<br><br>
*___azure_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [azure_settings](#nested--azure_settings) for nested schema<br>
<br>Azure specific settings for the Data Centre. Cannot be provided with AWS or GCP settings.<br><br>
*___gcp_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [gcp_settings](#nested--gcp_settings) for nested schema<br>
<br>GCP specific settings for the Data Centre. Cannot be provided with AWS or Azure settings.<br><br>
*___tag___*<br>
<ins>Type</ins>: repeatable nested block, optional, updatable, see [tag](#nested--tag) for nested schema<br>
<br>List of tags to apply to the Data Centre. Tags are metadata labels which allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require. Note: Tags will be returned sorted by key in alphabetical order regardless of input order. Terraform users: `tag` is not supported in terraform lifecycle `ignore_changes`.<br><br>
*___aws_settings___*<br>
<ins>Type</ins>: nested block, optional, immutable, see [aws_settings](#nested--aws_settings) for nested schema<br>
<br>AWS specific settings for the Data Centre. Cannot be provided with GCP or Azure settings.<br><br>
*___provider_account_name___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.<br><br>
### Read-only attributes
*___current_operations___*<br>
<ins>Type</ins>: nested block, read-only, see [current_operations](#nested--current_operations) for nested schema<br>
<br>Active operations in the data centre.<br><br>
*___data_centre_domain___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Domain of the Data Centre. Balances requests against nodes.<br><br>
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
<ins>Constraints</ins>: allowed values: [ `CASSANDRA`, `SPARK_MASTER`, `SPARK_JOBSERVER`, `KAFKA_BROKER`, `KAFKA_DEDICATED_ZOOKEEPER`, `KAFKA_DEDICATED_KRAFT_CONTROLLER`, `KAFKA_ZOOKEEPER`, `KAFKA_KRAFT_CONTROLLER`, `KAFKA_SCHEMA_REGISTRY`, `KAFKA_REST_PROXY`, `APACHE_ZOOKEEPER`, `POSTGRESQL`, `PGBOUNCER`, `KAFKA_CONNECT`, `KAFKA_KARAPACE_SCHEMA_REGISTRY`, `KAFKA_KARAPACE_REST_PROXY`, `CADENCE`, `CLICKHOUSE_SERVER`, `CLICKHOUSE_KEEPER`, `CLICKHOUSE_SERVER_AND_KEEPER`, `REDIS_MASTER`, `REDIS_REPLICA`, `VALKEY_MASTER`, `VALKEY_REPLICA`, `OPENSEARCH_DASHBOARDS`, `OPENSEARCH_COORDINATOR`, `OPENSEARCH_MASTER`, `OPENSEARCH_DATA`, `OPENSEARCH_INGEST`, `OPENSEARCH_DATA_AND_INGEST`, `KAFKA_SHOTOVER_PROXY`, `MCP_GATEWAY_SERVER`, `KAFKA_DISKLESS` ]<br><br>The roles or purposes of the node. Useful for filtering for nodes that have a specific role.<br><br>
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
List of tags to apply to the Data Centre. Tags are metadata labels which allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require. Note: Tags will be returned sorted by key in alphabetical order regardless of input order. Terraform users: `tag` is not supported in terraform lifecycle `ignore_changes`.<br>
### Input attributes - Required
*___key___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Key of the custom tag.<br><br>
### Input attributes - Optional
*___value___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Value of the custom tag.<br><br>
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
<ins>Constraints</ins>: allowed values: [ `CASSANDRA`, `SPARK_MASTER`, `SPARK_JOBSERVER`, `KAFKA_BROKER`, `KAFKA_DEDICATED_ZOOKEEPER`, `KAFKA_DEDICATED_KRAFT_CONTROLLER`, `KAFKA_ZOOKEEPER`, `KAFKA_KRAFT_CONTROLLER`, `KAFKA_SCHEMA_REGISTRY`, `KAFKA_REST_PROXY`, `APACHE_ZOOKEEPER`, `POSTGRESQL`, `PGBOUNCER`, `KAFKA_CONNECT`, `KAFKA_KARAPACE_SCHEMA_REGISTRY`, `KAFKA_KARAPACE_REST_PROXY`, `CADENCE`, `CLICKHOUSE_SERVER`, `CLICKHOUSE_KEEPER`, `CLICKHOUSE_SERVER_AND_KEEPER`, `REDIS_MASTER`, `REDIS_REPLICA`, `VALKEY_MASTER`, `VALKEY_REPLICA`, `OPENSEARCH_DASHBOARDS`, `OPENSEARCH_COORDINATOR`, `OPENSEARCH_MASTER`, `OPENSEARCH_DATA`, `OPENSEARCH_INGEST`, `OPENSEARCH_DATA_AND_INGEST`, `KAFKA_SHOTOVER_PROXY`, `MCP_GATEWAY_SERVER`, `KAFKA_DISKLESS` ]<br><br>The roles or purposes of the node. Useful for filtering for nodes that have a specific role.<br><br>
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
*___storage_network___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>The private network address block to be used for the storage network. This is only used for certain node sizes, currently limited to those which use AWS FSx ONTAP: for all other node sizes, this field should not be provided. The network must have a prefix length between /16 and /28, and must be part of a private address range.<br><br>
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
<ins>Constraints</ins>: pattern: `^(([\s]*[^<>()\[\]\\.,;:@\s"]+(\.[^<>()\[\]\\.,;:\s@"]+)*))@((\[\d{1,3}(\.\d{1,3}){3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}[\s]*))$`<br><br>The email address which will be contacted when the cluster is requested to be deleted.<br><br>
### Input attributes - Optional
*___confirmation_phone_number___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<ins>Constraints</ins>: pattern: `^(?![\s])[\-\s\(\)\+0-9]*$`<br><br>The phone number which will be contacted when the cluster is requested to be delete.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_mcp_gateway_cluster_v1.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
