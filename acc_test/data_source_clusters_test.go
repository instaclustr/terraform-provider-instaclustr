package test

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TestClustersDataSource(t *testing.T) {
	testAccProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	validConfig, _ := ioutil.ReadFile("data/valid_multiple_with_clusters_data_source.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	hostname := getOptionalEnv("IC_API_URL", instaclustr.DefaultApiHostname)
	oriConfig := fmt.Sprintf(string(validConfig), username, apiKey, hostname)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		CheckDestroy: resource.ComposeTestCheckFunc(
			testCheckResourceDeleted("valid.0", hostname, username, apiKey),
			testCheckResourceDeleted("valid.1", hostname, username, apiKey),
		),
		Steps: []resource.TestStep{
			{
				Config:             oriConfig,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testDataSourceReturnsClusters(t, hostname, username, apiKey),
				),
			},
		},
	})
}

func testDataSourceReturnsClusters(t *testing.T, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["data.instaclustr_clusters.clusters"]
		attributes := resourceState.Primary.Attributes

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		clusterList, err := client.ListClusters()

		if err != nil {
			return fmt.Errorf("failed to read cluster list: %s", err)
		}

		numClusters, err := strconv.Atoi(attributes["cluster.#"])
		if err != nil {
			return fmt.Errorf("failed to read cluster list length: %s", err)
		}

		assert.Equal(t, numClusters, len(*clusterList), "mismatch between number of clusters from the data source and the API result")

		for i := 0; i < numClusters; i++ {
			clusterId := attributes[fmt.Sprintf("cluster.%d.cluster_id", i)]
			assert.True(t, containsCluster(*clusterList, clusterId), "failed to find matching cluster ID from API response: %s", clusterId)
		}

		return nil
	}
}

func containsCluster(clusterList []instaclustr.ClusterListItem, clusterId string) bool {
	for _, cluster := range clusterList {
		if cluster.ID == clusterId {
			return true
		}
	}
	return false
}
