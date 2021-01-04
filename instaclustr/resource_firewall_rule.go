package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallRuleCreate,
		Read:   resourceFirewallRuleRead,
		Update: resourceFirewallRuleUpdate,
		Delete: resourceFirewallRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirewallRuleStateImport,
		},

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"rule_cidr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_security_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: schema.TypeString,
				},
			},
		},
	}
}

func resourceFirewallRuleCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating firewall rule.")
	client := meta.(*Config).Client

	ruleTarget, ruleTargetError := getRuleTarget(d)

	if ruleTargetError != nil {
		return fmt.Errorf("[Error] Error creating firewall rule: %s", ruleTargetError)
	}

	rules := make([]RuleType, 0)

	for _, rule := range d.Get("rules").([]interface{}) {
		aRule := ""

		for _, value := range rule.(map[string]interface{}) {
			aRule = fmt.Sprintf("%v", value)
		}

		rules = append(rules, RuleType{Type: aRule})
	}

	rule := FirewallRule{
		Network:         d.Get("rule_cidr").(string),
		SecurityGroupId: d.Get("rule_security_group_id").(string),
		Rules:           rules,
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(rule)
	if err != nil {
		return fmt.Errorf("[Error] Error creating firewall rule: %s", err)
	}

	err = client.CreateFirewallRule(jsonStr, d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error creating firewall fule: %s", err)
	}
	log.Printf("[INFO] Firewall rule %s has been created.", d.Get("cluster_id").(string))
	d.SetId(ruleTarget)
	return nil

}

func resourceFirewallRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("cluster_id").(string)

	ruleTarget, ruleTargetError := getRuleTarget(d)

	if ruleTargetError != nil {
		return fmt.Errorf("[Error] Error reading firewall rule: %s", ruleTargetError)
	}

	log.Printf("[INFO] Reading the status of cluster %s.", id)
	firewallRules, err := client.ReadFirewallRules(id)
	if err != nil {
		return fmt.Errorf("[Error] Error reading firewall rules: %s", err)
	}
	for _, value := range *firewallRules {
		if value.Network == ruleTarget || value.SecurityGroupId == ruleTarget {
			log.Printf("[INFO] Read rule %s from cluster %s", ruleTarget, id)
			d.Set("cluster_id", id)
			d.Set("rule_cidr", value.Network)
			d.Set("rule_security_group_id", value.SecurityGroupId)
			d.SetId(ruleTarget)

			rules := make([]map[string]interface{}, 0)
			for _, rule := range value.Rules {
				ruleMapStruct := &RuleType{Type: rule.Type}
				ruleMap, _ := StructToMap(ruleMapStruct)
				rules = append(rules, ruleMap)
			}

			if err := d.Set("rules", rules); err != nil {
				return fmt.Errorf("error setting rules: %s", err)
			}
		}
	}
	return nil
}

func resourceFirewallRuleStateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "&")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("Unexpected format of ID (%q), expected CLUSTER-ID&RULE-CIDR", d.Id())
	}
	cluster_id := idParts[0]
	rule_cidr := idParts[1]
	d.Set("cluster_id", cluster_id)
	d.Set("rule_cidr", rule_cidr)
	return []*schema.ResourceData{d}, nil
}

func resourceFirewallRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFirewallRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("cluster_id").(string)
	log.Printf("[INFO] Deleting rule %s.", d.Get("rule_cidr").(string))

	rules := make([]RuleType, 0)

	for _, rule := range d.Get("rules").([]interface{}) {
		aRule := ""

		for _, value := range rule.(map[string]interface{}) {
			aRule = fmt.Sprintf("%v", value)
		}

		rules = append(rules, RuleType{Type: aRule})
	}

	rule := FirewallRule{
		SecurityGroupId: d.Get("rule_security_group_id").(string),
		Network:         d.Get("rule_cidr").(string),
		Rules:           rules,
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(rule)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting firewall rule: %s", err)
	}

	err = client.DeleteFirewallRule(jsonStr, id)

	if err != nil {
		return fmt.Errorf("[Error] Error deleting firewall request: %s : %s", jsonStr, err)
	}

	log.Printf("[INFO] Firewall rule %s has been deleted.", rule)
	d.SetId("")
	return nil
}

func getRuleTarget(d *schema.ResourceData) (string, error) {
	cidrRuleTarget := d.Get("rule_cidr").(string)
	securityGroupRuleTarget := d.Get("rule_security_group_id").(string)

	if len(cidrRuleTarget) == 0 && len(securityGroupRuleTarget) == 0 {
		return "", fmt.Errorf("Either one of Security Group or Rule Cidr is required.")
	}

	if len(cidrRuleTarget) > 0 && len(securityGroupRuleTarget) > 0 {
		return "", fmt.Errorf("Only one of Security Group or Rule Cidr can be provided per rule.")
	}

	if len(cidrRuleTarget) > 0 {
		return cidrRuleTarget, nil
	}

	return securityGroupRuleTarget, nil
}
