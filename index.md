---
layout: ""
page_title: "Provider: Instaclustr"
description: |-
  A Terraform provider for managing Instaclustr Platform resources.
---

# Instaclustr Provider

A Terraform provider for managing resources on the [Instaclustr Platform](https://instaclustr.com).

It provides a flexible set of resources for provisioning and managing Instaclustr based clusters via the use of Terraform.

For further information about Instaclustr, please see [FAQ](https://www.instaclustr.com/faqs/) and [Support](https://support.instaclustr.com/).

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