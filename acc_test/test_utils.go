package test

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
	"io/ioutil"
	"os"
	"time"
)

func getOptionalEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func checkClusterRunning(resourceName, hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster."+resourceName]
		id := resourceState.Primary.Attributes["cluster_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)

		const ClusterReadInterval = 5
		const WaitForClusterTimeout = 40 * 60
		var latestStatus string
		timePassed := 0
		fmt.Print("\033[s")
		for {
			cluster, err := client.ReadCluster(id)
			if err != nil {
				fmt.Printf("\n")
				return fmt.Errorf("[Error] Error retrieving cluster info: %s", err)
			}
			latestStatus = cluster.ClusterStatus
			if cluster.ClusterStatus == "RUNNING" {
				break
			}
			if timePassed > WaitForClusterTimeout {
				fmt.Printf("\n")
				return fmt.Errorf("[Error] Timed out waiting for cluster to have the status 'RUNNING'. Current cluster status is '%s'", latestStatus)
			}
			timePassed += ClusterReadInterval
			fmt.Printf("\033[u\033[K%ds has elapsed while waiting for the cluster to reach RUNNING.\n", timePassed)
			time.Sleep(ClusterReadInterval * time.Second)
		}
		fmt.Printf("\n")
		// wait another minute after cluster goes to RUNNING to make sure all operations will work
		// sometimes says cluster is not ready for resizing
		time.Sleep(60 * time.Second)
		return nil
	}
}

func addDCtoCluster(resourceName, hostname, username, apiKey, requestBody string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_cluster."+resourceName]
		id := resourceState.Primary.Attributes["cluster_id"]
		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		body, _ := ioutil.ReadFile(requestBody)

		// Adding a new DC to an existing cluster
		url := fmt.Sprintf("%s/provisioning/v1/%s/cluster-data-centres", hostname, id)
		resp, err := client.MakeRequest(url, "POST", body)

		if err != nil {
			return err
		}
		bodyText, err := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != 202 {
			return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
		}
		return nil
	}
}
