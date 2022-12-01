---
layout: ""
page_title: "Provider: Instaclustr"
description: |-
  A Terraform provider for managing Instaclustr Platform resources.
---

# Instaclustr Terraform Provider v2

A Terraform provider for managing resources on the [Instaclustr Platform](https://instaclustr.com).

It provides a flexible set of resources for provisioning and managing Instaclustr based clusters with the use of Terraform.

This provider is in __General Availability__ with support for the following offerings:

| Application | Support Status |
| ----------- |----------------|
| Apache Cassandra | Complete       |
| Apache Kafka | Complete       |
| Kafka Connect | Complete    |
| Apache ZooKeeper | Complete    |
| OpenSearch | In Progress    |
| Redis | In Progress    |
| PostgreSQL | In Progress    |
| Cadence | Complete    |

Support for other offerings will be added progressively through to the end of 2022. If a resource you are looking to manage through Terraform is not yet supported by the Instaclustr Terraform Provider v2, you may use v1.x of the Instaclustr Terraform Provider instead. 

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


## Migrating from v1 to v2 of the Instaclustr Terraform Provider

With the v2 version of the Instaclustr Terraform Provider, new resources have been introduced and schemas of existing resources have been changed. While the [Terraform Registry](https://registry.terraform.io/providers/instaclustr/instaclustr/latest/docs) contains the schema definitions of each resource, for a tool assisted approach of migrating to the v2 version of the Terraform Provider, see our support article on [importing Terraform resources](https://www.instaclustr.com/support/api-integrations/integrations/terraform-code-generation/).
