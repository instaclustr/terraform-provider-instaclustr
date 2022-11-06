---
page_title: "instaclustr_aws_endpoint_service_principal_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_endpoint_service_principal_v2 (Resource)
Definition of an IAM Principal ARN being used for Kafka PrivateLink.
## Example Usage
```
resource "instaclustr_aws_endpoint_service_principal_v2" "example" {
  cluster_data_center_id = "f3eab841-6952-430d-ba90-1bfc3f15da10"
  principal_arn = "arn:aws:iam::123456789012:role/role-name"
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
*___principal_arn___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `^arn:aws:iam::[0-9]{12}:(root$|user\/[\w+=,.@-]+|role\/[\w+=,.@-]+)$`<br><br>The IAM Principal ARN.<br><br>
*___cluster_data_center_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>The ID of the cluster data center.<br><br>
### Read-only attributes
*___end_point_service_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The Instaclustr ID of the AWS endpoint service<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>The Instaclustr ID of the IAM Principal ARN.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_aws_endpoint_service_principal_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
