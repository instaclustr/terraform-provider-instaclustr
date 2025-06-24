---
page_title: "instaclustr_gcp_service_attachments_accept_list_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_gcp_service_attachments_accept_list_v2 (Resource)
List of accepted consumer projects or networks for a Cluster Data Center
## Example Usage
```
resource "instaclustr_gcp_service_attachments_accept_list_v2" "example" {
  consumer_accept_list {
    accept_network_name = "network-1"
    accept_project_id = "project-1"
    connection_limit = 10
  }

  cdc_id = "ddf992f6-96fb-4cce-b87c-f225df7f1745"
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
*___cdc_id___*<br>
<ins>Type</ins>: string (uuid), required, updatable<br>
<br>ID of the Cluster Data Centre.<br><br>
*___consumer_accept_list___*<br>
<ins>Type</ins>: repeatable nested block, required, updatable, see [consumer_accept_list](#nested--consumer_accept_list) for nested schema<br>
<br>Consumer projects or networks that are allowed to connect to the GCP Service Attachments.<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>The Instaclustr ID of the GCP Service Attachments Accept List.<br><br>
*___operation_status___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `NO_OPERATION`, `OPERATION_IN_PROGRESS`, `OPERATION_FAILED` ]<br><br>Indicates if the cluster is currently performing any operation such as being created, updated, or deleted<br><br>
<a id="nested--consumer_accept_list"></a>
## Nested schema for `consumer_accept_list`
Consumer projects or networks that are allowed to connect to the GCP Service Attachments.<br>
### Input attributes - Required
*___connection_limit___*<br>
<ins>Type</ins>: integer, required, updatable<br>
<br>The connection limit for the project or network<br><br>
*___accept_project_id___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: pattern: `^(?!.*(google|null|undefined|ssl))^[a-z][a-z0-9-]{4,28}[a-z0-9]$`<br><br>The accepted project ID<br><br>
### Input attributes - Optional
*___accept_network_name___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<ins>Constraints</ins>: pattern: `(^[a-z][a-z0-9-]{0,61}[a-z0-9]$|^$)`<br><br>The accepted network name of the parent project ID<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_gcp_service_attachments_accept_list_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
