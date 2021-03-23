---
page_title: "instaclustr_cluster_credentials Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

### Data Source `instaclustr_cluster_credentials`
A read-only data source used to get the password and certificate download link of a cluster.

Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr cluster.|Required
`cluster_password`|The password of the existing Instaclustr cluster.|Computed
`certificate_download`|The certificate download link of the existing Instaclustr cluster.|Computed
