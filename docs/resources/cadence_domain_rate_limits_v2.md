---
page_title: "instaclustr_cadence_domain_rate_limits_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_cadence_domain_rate_limits_v2 (Resource)
Per-frontend-instance domain request rate limits for a Cadence cluster.
## Example Usage
```
resource "instaclustr_cadence_domain_rate_limits_v2" "example" {
  cluster_id = "b997a00d-5bd4-4774-9bd7-5c0ad6189246"
  default_requests_per_second = 250
  domain_rate_limits {
    domain_name = "payments"
    requests_per_second = 1000
  }

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
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the Cadence cluster.<br><br>
### Input attributes - Optional
*___default_requests_per_second___*<br>
<ins>Type</ins>: integer (int32), optional, updatable<br>
<ins>Constraints</ins>: minimum: 1<br><br>Requests per second for all domains without a domain-specific limit. Written as the final unconstrained Cadence rule.<br><br>
*___domain_rate_limits___*<br>
<ins>Type</ins>: repeatable nested block, optional, updatable, see [domain_rate_limits](#nested--domain_rate_limits) for nested schema<br>
<ins>Constraints</ins>: maximum items: 999<br><br>Per-frontend-instance domain request rate limits.<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Cadence domain request-rate-limits resource (the target cluster ID).<br><br>
<a id="nested--domain_rate_limits"></a>
## Nested schema for `domain_rate_limits`
Per-frontend-instance domain request rate limits.<br>
### Input attributes - Required
*___requests_per_second___*<br>
<ins>Type</ins>: integer (int32), required, updatable<br>
<ins>Constraints</ins>: minimum: 1<br><br>Requests per second for the domain.<br><br>
*___domain_name___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: pattern: `^[A-Za-z][A-Za-z0-9._-]*$`<br><br>Exact Cadence domain name. It must start with a letter and may contain letters, numbers, periods, underscores, and hyphens. Cadence determines the supported maximum length.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_cadence_domain_rate_limits_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
