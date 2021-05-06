---
page_title: "instaclustr_firewall_rule Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Resource:  `instaclustr_firewall_rule`                             
A resource for managing cluster firewall rules on Instaclustr Managed Platform. A firewall rule allows access to your Instaclustr cluster.
Note: Either `rule_cidr` OR `rule_security_group_id` must be provided per rule (but not both)

#### Properties
Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr managed cluster|Required
`rule_cidr`|The network to add to your cluster firewall rule. Must be a valid IPv4 CIDR (e.g. 123.4.5.67/32) |Optional
`rule_security_group_id`|The Peered AWS VPC Security Group id to your cluster firewall rule (e.g. sg-12345678) |Optional
`rules`|List of rule types that the specified network is allowed access to. See below for rule options.|Required

#### rules

Property | Description | Default
---------|-------------|--------
`type`|Accepts CASSANDRA, SPARK, SPARK_JOBSERVER, APACHE_ZOOKEEPER, KAFKA, KAFKA_CONNECT, ELASTICSEARCH, REDIS|Required

#### Example
```
resource "instaclustr_firewall_rule" "example" {
    cluster_id = "${instaclustr_cluster.example.id}"
    rule_cidr = "10.1.0.0/16"
    rules = [
        { 
            type = "CASSANDRA"
        }
    ]
}
```

