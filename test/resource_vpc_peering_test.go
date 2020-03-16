package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
)

func TestAccVpcPeeringResource(t *testing.T) {
	testProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	tfFile, _ := ioutil.ReadFile("data/valid_with_vpc_peering.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	config := fmt.Sprintf(string(tfFile), username, apiKey)

	hostname := instaclustr.ApiHostname
	resource.Test(t, resource.TestCase{
		Providers:    testProviders,
		CheckDestroy: checkVpcPeeringDeleted(hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					checkVpcPeeringState,
					checkVpcPeeringCreated(hostname, username, apiKey),
				),
			},
		},
	})
}

func checkVpcPeeringState(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["instaclustr_vpc_peering.valid_with_vpc_peering"]
	if resourceState == nil {
		return fmt.Errorf("valid: resource not found in state")
	}

	instanceState := resourceState.Primary
	if instanceState == nil {
		return fmt.Errorf("resource has no primary instance")
	}
	return nil
}

func checkVpcPeeringCreated(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_vpc_peering.valid_with_vpc_peering"]

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		cdcID := resourceState.Primary.Attributes["cdc_id"]
		vpcPeeringID := resourceState.Primary.Attributes["vpc_peering_id"]

		vpcPeering, err := client.ReadVpcPeering(cdcID, vpcPeeringID)
		if err != nil {
			return fmt.Errorf("Failed to read VPC peering %s: %s", vpcPeeringID, err)
		}
		if vpcPeering.ID != vpcPeeringID {
			return fmt.Errorf("VPC peering connection expected %s but got %s", vpcPeeringID, vpcPeering.ID)
		}
		return nil
	}
}

func checkVpcPeeringDeleted(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_vpc_peering.valid_with_vpc_peering"]
		cdcID := resourceState.Primary.Attributes["cdc_id"]
		vpcPeeringID := resourceState.Primary.Attributes["vpc_peering_id"]

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		err := client.DeleteVpcPeering(cdcID, vpcPeeringID)
		if err == nil {
			return fmt.Errorf("VPC peering connection %s still exists", vpcPeeringID)
		}
		return nil
	}
}
