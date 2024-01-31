---
layout: ""
page_title: "Provider: Instaclustr"
description: |-
  A Terraform provider for managing Instaclustr Platform resources.
---

# Instaclustr Terraform Provider v2

A Terraform provider for managing resources on the [Instaclustr Platform](https://instaclustr.com).

It provides a flexible set of resources for provisioning and managing Instaclustr based clusters with the use of Terraform.


For further information about Instaclustr, please see [FAQ](https://www.instaclustr.com/faqs/) and [Support](https://support.instaclustr.com/).

## Support Contact
Please reach out to support@instaclustr.com for issues with this Terraform Provider. Please note that we've disabled the issues feature on this repository and transferred open issues to our Zendesk system.

## Example Usage

```terraform
terraform {
  required_providers {
    instaclustr = {
      source  = "instaclustr/instaclustr"
      version = ">= 2.0.0, < 3.0.0"
    }
  }
}

provider "instaclustr" {
  terraform_key = "Instaclustr-Terraform johndoe:a1b2c3def45g6hij789k0l1m2n3opq45"
}
```

## Schema

### Required

- **terraform_key** (String) Key to authorize API requests from the Terraform Provider. `Instaclustr-Terraform <USERNAME>:<PROVISIONING_API_KEY>`


## Migrating from v1 to v2 of the Instaclustr Terraform Provider

With the v2 version of the Instaclustr Terraform Provider, new resources have been introduced and schemas of existing resources have been changed. While the [Terraform Registry](https://registry.terraform.io/providers/instaclustr/instaclustr/latest/docs) contains the schema definitions of each resource, for a tool assisted approach of migrating to the v2 version of the Terraform Provider, see our support article on [importing Terraform resources](https://www.instaclustr.com/support/api-integrations/integrations/terraform-code-generation/).

## Known Limitations

### Ignore_changes lifecycle feature inconsistency

The `ignore_changes` lifecycle feature in Terraform does not work as expected in the Instaclustr Terraform Provider.

#### Case 1
When all changes in the Terraform configuration file (.tf) are ignored using `ignore_changes`, `ignore_changes` works correctly. The plan stage shows no changes and the apply stage is not triggered.

**Example**

Initial Terraform Configuration:
```terraform
resource "project" "foo" {
  name = "foo"
  role {
    name         = "owner 1"
  }
  role {
    name         = "owner 2"
  }
}
```
Updated Terraform Configuration:
```terraform
resource "project" "foo" {
  name = "foo"
  role {
    name         = "owner 1 update"
  }
  role {
    name         = "owner 2"
  }
  Life_cycle {
    Ignore_changes= role[0]
  }
}
```
Terraform plan won’t show any changes and apply won’t be triggered.

#### Case 2
If only certain changes in the Terraform configuration are meant to be ignored using `ignore_changes`, the functionality does not work as expected: during the plan stage, the plan show changes are ignored as expected. However, during the apply stage, the changes that should be ignored according to `ignore_changes` are not actually ignored due to a custom method implemented in our provider.


**Example**

Initial Terraform Configuration:
```terraform
resource "project" "foo" {
  name = "foo"
  role {
    name         = "owner 1"
  }
  role {
    name         = "owner 2"
  }
}
```
Updated Terraform Configuration:
```terraform
resource "project" "foo" {
  name = "foo"
  role {
    name         = "owner 1 update"
  }
  role {
    name         = "owner 2 update"
  }
  Life_cycle {
    Ignore_changes= role[0]
  }
}
```
Terraform plan will show:
```terraform
~resource "project" "foo" {
  ~role {
     ~name         = "owner 2" => “owner 2 update”
  }
}
```
However, in apply stage, the PUT API request will include a payload that both roles are updated.

Please carefully review Terraform outputs and avoid using `ignore_changes` if possible.

### Ordering of nested block types

Altering the order of repeatable nested block types in the Terraform configuration can lead to discrepancies between the Terraform plan and the subsequent apply phase. This problem arises due to our use of Terraform's SDK `TypeList` for an array of objects, where the order of resources is maintained, so naturally the resources in an array cannot reorder.

To overcome this, we did some customize implementation on terraform provider. Our terraform apply/plan adjust the order of entries in TF state to match order in TF configuration without changing the actual resource.

**Example**

Initial Terraform Configuration:
```terraform
resource "project" "foo" {
  name = "foo"
  role {
    name         = "owner 1"
  }
  role {
    name         = "owner 2"
  }
}
```
Updated Terraform Configuration:
```terraform
resource "project" "foo" {
  name = "foo"
  role {
    name         = "owner 2"
  }
  role {
    name         = "owner 1"
  }
}
```
Terraform plan will show:
```terraform
~resource "project" "foo" {
  name = "foo"
  ~role {
    ~name         = "owner 2"  => "owner 1"
  }
  ~role {
    ~name         = "owner 1"  => "owner 2"
  }
}
```
However, the apply won’t cause any actual resource changes.
We advise users to review their Terraform plans and outputs. If possible, avoid reordering data centers and tags to prevent these inconsistencies.

We are currently exploring options to address these limitations, which includes updates to our Terraform Provider with [latest SDK](https://github.com/hashicorp/terraform-plugin-framework). We understand that the limitations can pose a challenge, and we are committed to resolving it. If you are facing these issues or have any further questions, please contact [our friendly team](mailto:support@instaclustr.com) at any time. Your input is invaluable in helping us improve.