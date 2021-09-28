package instaclustr

import (
	"github.com/hashicorp/terraform/terraform"

	"testing"
)
//https://github.com/samsara-dev/terraform-provider-aws/blob/d5ae201faa90c78e709d8525778929da61e4154e/aws/resource_aws_route53_record_migrate_test.go
func TestClusterMigrateState(t *testing.T) {
	testCases := map[string]struct {
		StateVersion int
		Attributes   map[string]string
		Expected     map[string]string
		Meta         interface{}
	}{
		"v0_to_v1": {
			StateVersion: 0,
			Attributes: map[string]string{
				//"node_size": "fam-test-name",
				//"bundle.0.bundle": "REDIS",
				//"bundle.0.version" : "redis:6.0.9",
				//"bundle.0.options.0.master_nodes" : "3",
				//"bundle.0.options.0.password_auth": "false",
				//"bundle.0.options.0.client_encryption" : "false",
				//"data_centre" : "US_WEST_2",
				"node_size" : "r5.large-100-r",
				//"cluster_provider.0.name": "AWS_VPC",
				//"cluster_network": "192.168.0.0/18",
				//"data_centres": "",
			},
			Expected: map[string]string{
				//"bundle.0.bundle": "REDIS",
				//"bundle.0.version" : "redis:6.0.9",
				//"bundle.0.options.0.master_nodes" : "3",
				//"bundle.0.options.0.password_auth": "false",
				//"bundle.0.options.0.client_encryption" : "false",
				//"data_centres.0.data_centre" : "US_WEST_2",
				//"data_centres.0.name" : "US_WEST_2",
				//"data_centres.0.node_size" : "r5.large-100-r",
				//"data_centres.0.provider.0.name": "AWS_VPC",
				//"data_centres.0.network": "192.168.0.0/18",
				//"data_centres.0.bundles.0.bundle": "REDIS",
				//"data_centres.0.bundles.0.version" : "redis:6.0.9",
				//"data_centres.0.bundles.0.options.0.master_nodes" : "3",
				//"data_centres.0.bundles.0.options.0.password_auth": "false",
				//"data_centres.0.bundles.0.options.0.client_encryption" : "false",
				"node_size": "r5.large-100-r.",
			},
		},
	}

	for testName, testCase := range testCases {
		instanceState := &terraform.InstanceState{
			ID:         "some_id",
			Attributes: testCase.Attributes,
		}

		tfResource := resourceCluster()

		if tfResource.MigrateState == nil {
			t.Fatalf("bad: %s, err: missing MigrateState function in resource", testName)
		}

		instanceState, err := tfResource.MigrateState(testCase.StateVersion, instanceState, testCase.Meta)
		if err != nil {
			t.Fatalf("bad: %s, err: %#v", testName, err)
		}

		for key, expectedValue := range testCase.Expected {
			if instanceState.Attributes[key] != expectedValue {
				t.Fatalf(
					"bad: %s\n\n expected: %#v -> %#v\n got: %#v -> %#v\n in: %#v",
					testName, key, expectedValue, key, instanceState.Attributes[key], instanceState.Attributes)
			}
		}
	}
}