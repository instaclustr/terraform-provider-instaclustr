---
page_title: "instaclustr_cluster_network_firewall_rules_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_cluster_network_firewall_rules_v2 (Data Source)

## Schema
### cluster_id<br>
<ins>Type</ins>: string<br>

### firewall_rules<br>
<ins>Type</ins>: block list, see [firewall_rules](#nested--firewall_rules) for nested schema<br>

<a id="nested--firewall_rules"></a>
## Nested schema for `firewall_rules`<br>

### deferred_reason<br>
<ins>Type</ins>: string<br>
<br>The reason (if needed) for the deferred status of the cluster network firewall rule.
### cluster_id<br>
<ins>Type</ins>: string (uuid)<br>
<br>ID of the cluster for the cluster network firewall rule.
### id<br>
<ins>Type</ins>: string (uuid)<br>
<br>ID of the cluster network firewall rule.
### type<br>
<ins>Type</ins>: string<br>
<ins>Constraints</ins>: allowed values: [ `APACHE_ZOOKEEPER`, `CADENCE`, `CADENCE_GRPC`, `CADENCE_WEB`, `CASSANDRA`, `CASSANDRA_CQL`, `ELASTICSEARCH`, `KAFKA`, `KAFKA_CONNECT`, `KAFKA_REST_PROXY`, `KAFKA_SCHEMA_REGISTRY`, `KARAPACE_REST_PROXY`, `KARAPACE_SCHEMA_REGISTRY`, `OPENSEARCH`, `OPENSEARCH_DASHBOARDS`, `PGBOUNCER`, `POSTGRESQL`, `REDIS`, `SEARCH_DASHBOARDS`, `SECURE_APACHE_ZOOKEEPER`, `SPARK`, `SPARK_JOBSERVER` ]<br><br>The type of firewall rule.
### network<br>
<ins>Type</ins>: string<br>
<br>The network of the cluster network firewall rule.
### status<br>
<ins>Type</ins>: string<br>
<br>The status of the cluster network firewall rule.

