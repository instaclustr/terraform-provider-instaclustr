package instaclustr

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

func TestFirewallRuleUpdate(t *testing.T) {
	resourceSchema := resourceFirewallRule().Schema
	resourceDataMap := map[string]interface{}{
		"cluster_id": "something",
		"rule_cidr":  "10.0.0.6/32",
	}
	config := &Config{
		Client: SetupMock(t, "something/firewallRules/", "", 202),
	}
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)
	resourceLocalData.Set("rules", []map[string]interface{}{{"type": "CASSANDRA"}, {"type": "CASSANDRA_CQL"}})
	err := resourceFirewallRuleUpdate(resourceLocalData, config)
	if err != nil {
		t.Fatal("update firewall rule errored out")
	}
}
