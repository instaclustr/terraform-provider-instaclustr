# Terraform Instaclustr Provider

[![CircleCI](https://img.shields.io/circleci/build/gh/instaclustr/terraform-provider-instaclustr/master?style=for-the-badge)](https://app.circleci.com/pipelines/github/instaclustr/terraform-provider-instaclustr)
![GoLang Version](https://img.shields.io/github/go-mod/go-version/instaclustr/terraform-provider-instaclustr?logo=go&style=for-the-badge)
![Latest Release Version](https://img.shields.io/github/v/release/instaclustr/terraform-provider-instaclustr?logo=github&sort=semver&style=for-the-badge)
![License](https://img.shields.io/github/license/instaclustr/terraform-provider-instaclustr?style=for-the-badge)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=instaclustr_terraform-provider-instaclustr&metric=alert_status)](https://sonarcloud.io/dashboard?id=instaclustr_terraform-provider-instaclustr)

A [Terraform](http://terraform.io) provider for managing Instaclustr Platform resources.

It provides a flexible set of resources for provisioning and managing Apache Cassandra, Apache Kafka, Apache Zookeeper, OpenSearch and Redis clusters on the Instaclustr Managed Platform via Terraform.

For further information about Instaclustr, please see [the Instaclustr website](https://instaclustr.com/) and [Support Pages](https://support.instaclustr.com/hc/en-us)

For general information about Terraform, visit the [official website](https://terraform.io/) and [GitHub project page](https://github.com/hashicorp/terraform).

## Key benefits

- Removes the need to write custom code integration directly with the Instaclustr API
- Instaclustr based infrastructure as code deployments with minimal effort
- Ease of integration into existing terraform or automated CI/CD based workflows
- Ease of customisation and configuration in order to meet operational requirements
- Use of existing Instaclustr authentication methodologies

## Requirements

- Terraform v0.10.x - .v0.15.x.
- Go 1.14 or higher

## Using The Provider

To install this provider using Terraform 0.13+, copy and paste this code into your Terraform configuration. Then, run terraform init.

```
terraform {
  required_providers {
    instaclustr = {
      source = "instaclustr/instaclustr"
      version = "1.9.9"
    }
  }
}

provider "instaclustr" {
  # Configuration options
}
```

For further details on Provider installation, and installation on older versions of terraform please see the [Terraform installation guide](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin).

## Authentication

This provider requires an API Key in order to provision Instaclustr resources. To create an API key, please log into the [Instaclustr Console](https://console.instaclustr.com) or signup for an account [here](https://console.instaclustr.com/user/signup) if you don't have one.  Navigate to `Account` -> `API Keys` page, locate the `Provisioning` role and click `Generate Key`.  This username and API key combination should be placed into the provider configuration:

```
provider "instaclustr" {
    username = "<Your instaclustr username here>"
    api_key = "<Your provisioning API key here>"
}
```

If you wish to keep secrets in the ENV instead of stored in your terraform file use the following method:

In console export the desired variable:

```export api_key={instaclustrAPIkey}```

In your terraform file create a variable:
```
variable "api_key" {
 type = string
 default = "xxx"
}
```

In the provider block use the variable:
```
provider "instaclustr" {
    username= "<Your instaclustr username>"
    api_key = var.api_key
}
```
When running terraform plan/apply, pipe in the variables as follows:

```terraform apply -var= "api_key=$api_key"```

## Example Usage

It's possible to provision clusters for different cloud providers by changing the variable `cluster_provider`.
The accepted cloud providers are: `AWS`, `GCP`, `AZURE`.


AWS:
```
resource "instaclustr_cluster" "example" {
    cluster_name = "testcluster"
    node_size = "m5l-250-v2"
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
        version = "3.11.8"
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

AZURE:
```
resource "instaclustr_cluster" "azure_example" {
  cluster_name = "testcluster"
  node_size = "Standard_DS2_v2-256-an"
  data_centre = "CENTRAL_US"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = false
  cluster_provider = {
    name = "AZURE_AZ"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "3.11.8"
    options = {
      auth_n_authz = true
    }
  }
}
```

GCP:
```
resource "instaclustr_cluster" "gcp_example" {
  cluster_name = "testclustergcp"
  node_size = "n1-standard-2"
  data_centre = "us-east1"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = false
  cluster_provider = {
    name = "GCP"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "3.11.8"
    options = {
      auth_n_authz = true
    }
  }
}
```

Multi Data Centre Provisioning:

For Multi Data Centre provisioning, please specify `node_size`, `rack_allocation`, `provider` and `bundles` in each `data_centres`;  
For each `data_centres`, it requires at least one `bundles` to be the base application, e.g. `APACHE_CASSANDRA`.  
```
resource "instaclustr_cluster" "multi_DC_example" {
  cluster_name = "testcluster_multiDC"
  sla_tier     = "NON_PRODUCTION"

  data_centres {
    name        = "DC1"
    data_centre = "US_WEST_1"
    network     = "10.1.0.0/18"
    node_size    = "m5l-250-v2"
    rack_allocation = {
      number_of_racks = 2
      nodes_per_rack  = 1
    }
    provider = {
      name = "AWS_VPC"
    }
    bundles {
      bundle = "APACHE_CASSANDRA"
      version = "3.11.8"
      options = {
        auth_n_authz = true
        use_private_broadcast_rpc_address = false
        client_encryption = false
        lucene_enabled = false
        continuous_backup_enabled = true
      }
    }
    bundles {
      bundle = "SPARK"
      version = "2.3.2"
    }
  }

  data_centres {
    name        = "DC2"
    data_centre = "CENTRAL_US"
    network     = "10.0.0.0/18"
    node_size    = "Standard_DS2_v2-256-an"
    rack_allocation = {
      number_of_racks = 2
      nodes_per_rack  = 1
    }
    provider = {
      name = "AZURE_AZ"
    }
    bundles {
      bundle = "APACHE_CASSANDRA"
      version = "3.11.8"
      options = {
        auth_n_authz = true
        use_private_broadcast_rpc_address = false
        client_encryption = false
        lucene_enabled = false
        continuous_backup_enabled = true
      }
    }
  }

  data_centres {
    name = "DC3"
    data_centre = "US_WEST_2"
    network     = "192.168.0.0/18"
    node_size    = "m5l-250-v2"
    rack_allocation = {
      number_of_racks = 2
      nodes_per_rack  = 1
    }
    provider = {
      name = "AWS_VPC"
    }
    bundles {
      bundle = "APACHE_CASSANDRA"
      version = "3.11.8"
      options = {
        auth_n_authz = true
        use_private_broadcast_rpc_address = false
        client_encryption = false
        lucene_enabled = false
        continuous_backup_enabled = true
      }
    }
    bundles {
      bundle = "SPARK"
      version = "2.3.2"
    }
  }
}
```


## Configuration

Configuration documentation can be found at the [Instaclustr Terraform Registry](https://registry.terraform.io/providers/instaclustr/instaclustr/latest/docs/resources/encryption_key)

## Bundles and Versions

| Bundle                   | Versions                              | Compatible With                                                                    |
|--------------------------|---------------------------------------|------------------------------------------------------------------------------------|
| APACHE_CASSANDRA         | 2.2.18, 3.0.19, 3.11.8, 4.0 (preview) |                                                                                    |
| SPARK                    | 2.1.3, 2.3.2                          | APACHE_CASSANDRA                                                                   |
| KAFKA                    | 2.7.1, 2.8.1, 3.0.0, 3.1.1            |                                                                                    |
| KAFKA_REST_PROXY         | 5.0.0                                 | KAFKA                                                                              |
| KAFKA_SCHEMA_REGISTRY    | 5.0.0, 5.0.4                          | KAFKA                                                                              |
| KARAPACE_SCHEMA_REGISTRY | 3.2.0                                 | KAFKA <br/> **Not compatible with:** <br/> KAFKA_REST_PROXY, KAFKA_SCHEMA_REGISTRY |
| KARAPACE_REST_PROXY      | 3.2.0                                 | KAFKA <br/> **Not compatible with:** <br/> KAFKA_REST_PROXY, KAFKA_SCHEMA_REGISTRY |
| OPENSEARCH               | 1.3.4, 2.0.0                          |                                                                                    |
| ELASTICSEARCH (For Legacy Support Only)           | 1.13.3                                |                                                                                    |
| KAFKA_CONNECT            | 2.7.1, 2.8.1, 3.0.0, 3.1.1            |                                                                                    |
| REDIS                    | 6.2.7, 7.0.4                          |                                                                                    |
| APACHE_ZOOKEEPER         | 3.6.3, 3.7.1                          |                                                                                    |
| POSTGRESQL               | 13.8, 14.5                            |                                                                                    |
| PGBOUNCER                | 1.17.0                                | POSTGRESQL                                                                         |
| CADENCE                  | 0.22.4                                |                                                                                    |

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
          version = ">= 1.0.0, < 2.0.0"
        }
      }
    }
    ```
4. Run the unit tests by `$ make test`
7. Create a PR and send it our way :)

Our Circle CI pipeline will automatically run unit tests when a PR is created and new changes committed.
It is also capable of running the acceptance tests, however our staff needs to give a manual approval to run the tests.
Passing tests are a requirement to merge a PR. Please let us know when your PR is ready for acceptance tests!

Unit tests are within `instaclustr` folder with `_unit_test` suffix, and used to test the internal methods.

#### Acceptance Testing

Acceptance tests are within `acc_test` folder, and used to run end-to-end testing. We recommend using CircleCI to run your acceptance tests, however you can run them locally. Acceptance tests require end to end interaction with the Instaclustr platform and will create real (paid) infrastructure. If you wish to perform local testing you must set the variables below and run: ```make testacc``` 


| Variable                        | Command                                                                    | Description                                                                                                                   |
|---------------------------------|----------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| TF_ACC                          | `$ export TF_ACC=1`                                                        | Enables online acceptance tests.                                                                                              |
| IC_USERNAME                     | `$ export IC_USERNAME=<your instaclustr username>`                         | Authorizes Provisioning API                                                                                                   |
| IC_API_KEY                      | `$ export IC_API_KEY=<your provisioning API key>`                          | Authorizes Provisioning API                                                                                                   |
| KMS_ARN                         | `$ export KMS_ARN=<your KMS ARN>`                                          | For EBS encryption of nodes. <b><i>Note:</i></b> You cannot use an ARN previously added to your account as an encryption key. |
| IC_PROV_ACC_NAME                | `$ export IC_PROV_ACC_NAME="<your provider name>"`                         | Your "Run In Your Own Account" account name.                                                                                  |
| IC_PROV_VPC_ID                  | `$ export IC_PROV_VPC_ID="<your AWS VPC ID>"`                              | For provisioning into a custom VPC.                                                                                           |
| IC_AWS_ACCESS_KEY               | `$ export IC_PROV_VPC_ID="<access key for the AWS S3 bucket>"`             | For Kafka Connect connection information. See bundle options.                                                                 |
| IC_AWS_SECRET_KEY               | `$ export IC_PROV_VPC_ID="<secret key for the AWS S3 bucket>"`             | For Kafka Connect connection information. See bundle options.                                                                 |
| IC_S3_BUCKET_NAME               | `$ export IC_PROV_VPC_ID="<AWS S3 bucket name>"`                           | For Kafka Connect connection information. See bundle options.                                                                 |
| IC_AZURE_STORAGE_ACCOUNT_NAME   | `$ export IC_PROV_VPC_ID="<account name for the AZURE container storage>"` | For Kafka Connect connection information. See bundle options.                                                                 |
| IC_AZURE_STORAGE_ACCOUNT_KEY    | `$ export IC_PROV_VPC_ID="<account key for the AZURE container storage>"`  | For Kafka Connect connection information. See bundle options.                                                                 |
| IC_AZURE_STORAGE_CONTAINER_NAME | `$ export IC_PROV_VPC_ID="<the name of the AZURE container storage>"`      | For Kafka Connect connection information. See bundle options.                                                                 |

#### Running Specific Tests
To run a specific test, use the `testtarget` makefile goal.
```TARGET=TestName make testtarget```

## Further information and documentation

This provider makes use of the Instaclustr API.  For further information including the latest updates and value definitions, please see [the provisioning API documentation](https://www.instaclustr.com/support/api-integrations/api-reference/provisioning-api/).

Please see https://www.instaclustr.com/support/documentation/announcements/instaclustr-open-source-project-status/ for Instaclustr support status of this project.

# License

Apache2 - See the included LICENSE file for more details.
