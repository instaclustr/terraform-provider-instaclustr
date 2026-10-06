---
page_title: "instaclustr_cadence_global_domain_rate_limits_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_cadence_global_domain_rate_limits_v2_instance (Data Source)
Domain request rate limits for a Cadence cluster.
## Example Usage
```
data "instaclustr_cadence_global_domain_rate_limits_v2_instance" "example" { 
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
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Cadence global domain request-rate-limits resource (the target cluster ID).<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Cadence cluster.<br><br>
*___default_requests_per_second___*<br>
<ins>Type</ins>: integer (int32), read-only<br>
<ins>Constraints</ins>: minimum: 1<br><br>Requests per second for all domains without a domain-specific limit. Written as the final unconstrained Cadence rule.<br><br>
*___domain_rate_limits___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [domain_rate_limits](#nested--domain_rate_limits) for nested schema<br>
<ins>Constraints</ins>: maximum items: 999<br><br>Domain request rate limits.<br><br>
<a id="nested--domain_rate_limits"></a>
## Nested schema for `domain_rate_limits`
Domain request rate limits.<br>
### Read-only attributes
*___requests_per_second___*<br>
<ins>Type</ins>: integer (int32), read-only<br>
<ins>Constraints</ins>: minimum: 1<br><br>Requests per second for the domain.<br><br>
*___domain_name___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^[A-Za-z][A-Za-z0-9._-]*$`<br><br>Exact Cadence domain name. It must start with a letter and may contain letters, numbers, periods, underscores, and hyphens. Cadence determines the supported maximum length.<br><br>
