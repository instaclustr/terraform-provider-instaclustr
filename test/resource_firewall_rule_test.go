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

func TestAccFirewallRuleResource(t *testing.T) {
	testProviders := map[string]terraform.ResourceProvider{
		"instaclustr": instaclustr.Provider(),
	}
	tfFile, _ := ioutil.ReadFile("data/valid_with_firewall_rule.tf")
	username := os.Getenv("IC_USERNAME")
	apiKey := os.Getenv("IC_API_KEY")
	config := fmt.Sprintf(string(tfFile), username, apiKey)

	hostname := getenv("IC_API_URL", instaclustr.DefaultApiHostname)
	resource.Test(t, resource.TestCase{
		Providers:    testProviders,
		CheckDestroy: checkFirewallRuleDeleted(hostname, username, apiKey),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					checkFirewallRuleState,
					checkFirewallRuleCreated(hostname, username, apiKey),
				),
			},
		},
	})
}

func checkFirewallRuleState(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["instaclustr_firewall_rule.valid_with_firewall_rule"]
	if resourceState == nil {
		return fmt.Errorf("valid: resource not found in state")
	}

	instanceState := resourceState.Primary
	if instanceState == nil {
		return fmt.Errorf("resource has no primary instance")
	}
	return nil
}

func checkFirewallRuleCreated(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_firewall_rule.valid_with_firewall_rule"]

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		clusterID := resourceState.Primary.Attributes["cluster_id"]

		firewallRules, err := client.ReadFirewallRules(clusterID)
		if err != nil {
			return fmt.Errorf("Failed to read firewall rules for %s: %s", clusterID, err)
		}
		if len(*firewallRules) == 0 {
			return fmt.Errorf("Expected firewall rules to be created for %s", clusterID)
		}
		return nil
	}
}

func checkFirewallRuleDeleted(hostname, username, apiKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState := s.Modules[0].Resources["instaclustr_firewall_rule.valid_with_firewall_rule"]
		clusterID := resourceState.Primary.Attributes["cluster_id"]

		client := new(instaclustr.APIClient)
		client.InitClient(hostname, username, apiKey)
		_, err := client.ReadFirewallRules(clusterID)
		if err == nil {
			return fmt.Errorf("Expected firewall rules to be deleted for %s", clusterID)
		}
		return nil
	}
}
