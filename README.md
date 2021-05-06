# Terraform Instaclustr Provider

[![CircleCI](https://img.shields.io/circleci/build/gh/instaclustr/terraform-provider-instaclustr/master?style=for-the-badge)](https://app.circleci.com/pipelines/github/instaclustr/terraform-provider-instaclustr)
![GoLang Version](https://img.shields.io/github/go-mod/go-version/instaclustr/terraform-provider-instaclustr?logo=go&style=for-the-badge)
![Latest Release Version](https://img.shields.io/github/v/release/instaclustr/terraform-provider-instaclustr?logo=github&sort=semver&style=for-the-badge)
![License](https://img.shields.io/github/license/instaclustr/terraform-provider-instaclustr?style=for-the-badge)

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

- Terraform v0.10.x - .v0.13.x.
- Go 1.14 or higher

## Using The Provider

To install this provider, copy and paste this code into your Terraform configuration. Then, run terraform init.

Terraform 0.13+

```
terraform {
  required_providers {
    instaclustr = {
      source = "instaclustr/instaclustr"
      version = "1.9.6"
    }
  }
}

provider "instaclustr" {
  # Configuration options
}
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
    pci_compliant_cluster = false
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }
    bundle {
        bundle = "APACHE_CASSANDRA"
        version = "3.11.4"
        options = {
          auth_n_authz = true
        }
      }
      bundle {
        bundle = "SPARK"
        version = "2.3.2"
      }
}
```

## Configuration

Configuration documentation can be found at the [Instacluster Terraform Registry](https://registry.terraform.io/providers/instaclustr/instaclustr/latest/docs/resources/encryption_key)

## Bundles and Versions

Bundle | Versions | Compatible With
---------|-------------|---------------
APACHE_CASSANDRA|2.2.18, 3.0.19, 3.11.8, 4.0 (preview)|
SPARK|2.1.3, 2.3.2|APACHE_CASSANDRA
KAFKA|2.1.1, 2.3.1, 2.4.1, 2.5.1, 2.6.1|
KAFKA_REST_PROXY|5.0.0|KAFKA
KAFKA_SCHEMA_REGISTRY|5.0.0|KAFKA
ELASTICSEARCH|opendistro-for-elasticsearch:1.4.0
KAFKA_CONNECT|2.3.1, 2.4.1, 2.5.1, 2.6.1|
REDIS|6.0.9|
APACHE_ZOOKEEPER|3.5.8|

### Migrating from 0.0.1 &rarr; 1.0.0+
A schema change has been made from 0.0.1 which no longer supports the `bundles` argument and uses `bundle` blocks instead. This change can cause `terraform apply` to fail with a message that `bundles` has been removed and/or updating isn't supported. To resolve this -<br>
1. Change all usages of the `bundles` argument &rarr; `bundle` blocks (example under example/main.tf)
2. In the .tfstate files, replace all keys named `bundles` with `bundle` in resources under the Instaclustr provider.

## Contributing

Firstly thanks!  We value your time and will do our best to review the PR as soon as possible.

1. [Install golang](https://golang.org/doc/install#install)
2. Clone repository to: ```$GOPATH/src/github.com/instaclustr/terraform-provider-instaclustr```
3. Build and install the provider by `$ make build install`
1. For local testing of your changes you will need to use the local provider instead of the provider from the registry, so change your provider config to the following

    ```
    terraform {
      required_providers {
        instaclustr = {
          source = "terraform.instaclustr.com/instaclustr/instaclustr"
          version = ">= 1.0.0"
        }
      }
    }
    ```
4. Run the unit tests by `$ make test`
7. Create a PR and send it our way :)

Our Circle CI pipeline will automatically run unit tests when a PR is created and new changes are committed.
It is also capable of running the acceptance tests, however our staff needs to give a manual approval to run the tests.
Passing tests are a requirement to merge a PR. Please let us know when your PR is ready for acceptance tests!

Unit tests are within `instaclustr` folder with `_unit_test` suffix, and used to test the internal methods.

#### Acceptance Testing

Acceptance tests are within `acc_test` folder, and used to run end-to-end testing. We recommend using CircleCI to run your acceptance tests, however you can run them locally. Acceptance tests require end to end interaction with the instaclustr platform and will create real (paid) infrastructure. If you wish to perform local testing you must set the variables below and run ```make testacc``` 


Variable | Command | Description
---------|-------------|--------
TF_ACC|`$ export TF_ACC=1`|Enables online acceptance tests.
IC_USERNAME|`$ export IC_USERNAME=<your instaclustr username>`|Authorizes Provisioning API
IC_API_KEY|`$ export IC_API_KEY=<your provisioning API key>`|Authorizes Provisioning API
KMS_ARN|`$ export KMS_ARN=<your KMS ARN>`|For EBS encryption of nodes. <b><i>Note:</i></b> You cannot use an ARN previously added to your account as an encryption key.
IC_PROV_ACC_NAME|`$ export IC_PROV_ACC_NAME="<your provider name>"`|Your "Run In Your Own Account" account name.
IC_PROV_VPC_ID|`$ export IC_PROV_VPC_ID="<your AWS VPC ID>"`|For provisioning into a custom VPC.

#### Environment Variables Specific to Kafka Connect Acceptance Test

These environment variables are optional and only required when we want to do acceptance tests for Kafka Connect.
It is toggled by setting IC_TEST_KAFKA_CONNECT environment variable.

Variable | Command | Description
---------|-------------|--------
IC_TEST_KAFKA_CONNECT|`$ export IC_PROV_VPC_ID=1`|Enables acceptance tests for Kafka Connect.
IC_TARGET_KAFKA_CLUSTER_ID|`$ export IC_PROV_VPC_ID="<target kafka cluster ID>"`|For Kafka Connect connection information. See bundle options.
IC_AWS_ACCESS_KEY|`$ export IC_PROV_VPC_ID="<access key for the AWS S3 bucket>"`|For Kafka Connect connection information. See bundle options.
IC_AWS_SECRET_KEY|`$ export IC_PROV_VPC_ID="<secret key for the AWS S3 bucket>"`|For Kafka Connect connection information. See bundle options.
IC_S3_BUCKET_NAME|`$ export IC_PROV_VPC_ID="<AWS S3 bucket name>"`|For Kafka Connect connection information. See bundle options.
IC_AZURE_STORAGE_ACCOUNT_NAME|`$ export IC_PROV_VPC_ID="<account name for the AZURE container storage>"`|For Kafka Connect connection information. See bundle options.
IC_AZURE_STORAGE_ACCOUNT_KEY|`$ export IC_PROV_VPC_ID="<account key for the AZURE container storage>"`|For Kafka Connect connection information. See bundle options.
IC_AZURE_STORAGE_CONTAINER_NAME|`$ export IC_PROV_VPC_ID="<the name of the AZURE container storage>"`|For Kafka Connect connection information. See bundle options.
IC_SSL_ENABLED_PROTOCOLS|`$ export IC_PROV_VPC_ID="<SSL enabled protocols>"`|For Kafka Connect connection information. See bundle options.
IC_SSL_TRUSTSTORE_PASSWORD|`$ export IC_PROV_VPC_ID="<SSL truststore password>"`|For Kafka Connect connection information. See bundle options.
IC_SSL_PROTOCOL|`$ export IC_PROV_VPC_ID="<SSL protocol>"`|For Kafka Connect connection information. See bundle options.
IC_SECURITY_PROTOCOL|`$ export IC_PROV_VPC_ID="<Security protocol>"`|For Kafka Connect connection information. See bundle options.
IC_SASL_MECHANISM|`$ export IC_PROV_VPC_ID="<SASL mechanism>"`|For Kafka Connect connection information. See bundle options.
IC_SASL_JAAS_CONFIG|`$ export IC_PROV_VPC_ID="<SASL JAAS config>"`|For Kafka Connect connection information. See bundle options.
IC_BOOTSTRAP_SERVERS|`$ export IC_PROV_VPC_ID="<bootstrap servers>"`|For Kafka Connect connection information. See bundle options.
IC_TRUSTSTORE|`$ export IC_PROV_VPC_ID="<Base64 encoding of the truststore jks>"`|For Kafka Connect connection information. See bundle options.

#### Running Specific Tests
To run a specific test, use the `testtarget` makefile goal.
```TARGET=TestName make testtarget```

## Further information and documentation

This provider makes use of the Instaclustr API.  For further information including latest updates and value definitions, please see [the provisioning API documentation](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/).

Please see https://www.instaclustr.com/support/documentation/announcements/instaclustr-open-source-project-status/ for Instaclustr support status of this project.

# License

Apache2 - See the included LICENSE file for more details.

Test

