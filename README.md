# Terraform Instaclustr Provider

A [Terraform](http://terraform.io) provider for managing Instaclustr Platform resources.  

It provides a flexible set of resources for provisioning and managing [Instaclustr based clusters](http://instaclustr.com/) via the use of Terraform.  

For general information about Terraform, visit the [official website](https://terraform.io/) and [GitHub project page](https://github.com/hashicorp/terraform).

For further information about Instaclustr, please see [FAQ](https://www.instaclustr.com/resources/faqs/) and [Support](https://support.instaclustr.com/hc/en-us) 

## Key benefits

- Removes the need to write custom code integration directly with the Instaclustr API
- Instaclustr based infrastructure as code deployments with minimal effort
- Ease of integration into existing terraform or automated CI/CD based workflows
- Ease of customisation and configuration in order to meet operational requirements
- Use of existing Instaclustr authentication methodologies

## Requirements

- Terraform 0.10.x or higher
- Go 1.8 or higher

## Building The Provider

Clone the provider source code

```sh
$ mkdir -p $GOPATH/src/github.com/instaclustr; cd $GOPATH/src/github.com/instaclustr
$ git clone hhttps://github.com/instaclustr/terraform-provider-instaclustr.git
```

Build the source into executable binary

```sh
$ cd $GOPATH/src/github.com/instaclustr/terraform-provider-instaclustr
$ make build
```

Install the provider

```sh
$ cd $GOPATH/src/github.com/instaclustr/terraform-provider-instaclustr
$ make install
```

For further details on Provider installation please see the [Terraform installation guide](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin).

## Authentication

This provider requires an API Key in order to provision Instaclustr resources. To create an API key, please log into the [Instaclustr Console](https://console.instaclustr.com) or signup for an account [here](https://console.instaclustr.com/user/signup) if you dont have one.  Navigate to `Account` -> `API Keys` page, locate the `Provisioning` role and click `Generate Key`.  This username and API key combination should then be placed into the provider configuration:

```
provider "instaclustr" {
    username = "<Your instaclustr username here>"
    api_key = "<Your provisioning API key here>"
}
```
or just input them when Terraform indicates you.

## Example Usage

```
resource "instaclustr_cluster" "example" {
    cluster_name = "testcluster"
    node_size = "t3.small"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }
    bundles = [
        {
            bundle = "APACHE_CASSANDRA"
            version = "apache-cassandra-3.0.18"
        },
        {
            bundle = "SPARK"
            version = "apache-spark:2.3.2"
        },
        {
            bundle = "ZEPPELIN"
            version = "apache-zeppelin:0.8.0-spark-2.3.2"
        }
    ]
}
```

## Configuration
### Resources
`instaclustr_cluster`  
A resource for managing clusters on Instaclustr Managed Platform. A cluster contains a base application and several add-ons.

#### Properties
Property | Description | Default
---------|-------------|--------
cluster_name|The name of new cluster. May contain a combination of letters, numbers and underscores with a maximum length of 32 characters.|Required
node_size|Desired node size. See [here](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/#section-reference-data-data-centres-and-node-sizes) for more details.|Required
data_centre|Desired data centre. See [here](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/#section-reference-data-data-centres-and-node-sizes) for more details.|Required
sla_tier|Accepts PRODUCTION/NON_PRODUCTION. The SLA Tier feature on the Instaclustr console is used to classify clusters as either production and non_production. See [here](https://www.instaclustr.com/support/documentation/useful-information/sla-tier/) for more details.|NON_PRODUCTION
cluster_network|The private network address block for the cluster specified using CIDR address notation. The network must have a prefix length between /12 and /22 and must be part of a private address space.|10.224.0.0/12
private_network_cluster|Accepts true/false. Creates the cluster with private network only.|false
cluster_provider|The information of infrastructure provider. See below for its options.|Required
rack_allocation|The number of resources to use. See below for its options.|Required
bundles|Array of bundle information. See below for bundle options.|Required

`cluster_provider`

Property | Description | Default
---------|-------------|--------
name|Accepts AWS_VPC now. The new cluster will be deployed on Amazon Web Service.|Required
account_name|For customers running in their own account. Your provider account can be found on the ‘Account’ tab on the console, or the “Provider Account” property on any existing cluster.|""
tags|If specified, the value is a map from tag key to value. For restrictions, refer to the [AWS User Guide](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#tag-restrictions). Tags are defined per cluster and will be applied to every instance in the cluster.|""

`rack_allocation`

Property | Description | Default
---------|-------------|--------
number_of_racks|Number of racks to use when allocating nodes. Max allowed is 5|Required
nodes_per_rack|Number of nodes per rack. Max allowed is 10|Required

`bundles`

Property | Description | Default
---------|-------------|--------
bundle|Accepts APACHE_CASSANDRA. Compatible bundles: SPARK and/or ZEPPELIN.|Required
version|For Cassandra: apache-cassandra-2.1.19, apache-cassandra-2.2.13, apache-cassandra-3.0.14, apache-cassandra-3.0.17, apache-cassandra-3.0.18, apache-cassandra-3.11, apache-cassandra-3.11.3, apache-cassandra-3.11.4.ic1. For Spark: apache-spark:2.1.3, apache-spark:2.1.3.ic1, apache-spark:2.3.2. For Zeppelin: apache-zeppelin:0.8.0-spark-2.3.2, apache-zeppelin:0.7.1-spark-2.1.1|Required

## Contributing

Firstly thanks!  We value your time and will do our best to review the PR as soon as possible. 

1. [Install golang](https://golang.org/doc/install#install)
2. Clone repository to: $GOPATH/src/github.com/instaclustr/terraform-provider-instaclust
3. Build the provider by `$ make build`
4. Run the tests by `$ make test`
5. Setup environment variable `TF_ACC` to enable online acceptance test cases by `$ export TF_ACC=1`
6. Setup environment variables `IC_USERNAME` and `IC_API_KEY` of your provisioning API to grant online acceptance test cases permission by `$ export IC_USERNAME=<your instaclustr username>` and `$ export IC_API_KEY=<your provisioning API key>`
7. Run the acceptance tests `$ make testacc`
8. Create a PR and send it our way :)

## Further information and documentation

This provider makes use of the Instaclustr API.  For further information including latest updates and value definitions, please see [the provisioning API documentation](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/).

# License

Apache2 - See the included LICENSE file for more details.