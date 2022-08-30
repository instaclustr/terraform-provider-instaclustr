---
page_title: "instaclustr_clusters_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_clusters_v2 (Data Source)

## Schema
### account_id<br>
<ins>Type</ins>: string<br>

### clusters<br>
<ins>Type</ins>: block list, see [clusters](#nested--clusters) for nested schema<br>

<a id="nested--clusters"></a>
## Nested schema for `clusters`<br>

### application<br>
<ins>Type</ins>: string<br>
<ins>Constraints</ins>: allowed values: [ `APACHE_CASSANDRA`, `KAFKA`, `UNKNOWN` ]<br>
### id<br>
<ins>Type</ins>: string<br>


