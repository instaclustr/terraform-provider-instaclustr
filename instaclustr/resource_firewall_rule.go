package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallRuleCreate,
		Read:   resourceFirewallRuleRead,
		Update: resourceFirewallRuleUpdate,
		Delete: resourceFirewallRuleDelete,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"rule_cidr": {
				Type:     schema.TypeString,
				Required: true,
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

	rules := make([]RuleType, 0)

	for _, rule := range d.Get("rules").([]interface{}) {
		aRule := ""

		for _, value := range rule.(map[string]interface{}) {
			aRule = fmt.Sprintf("%v", value)
		}

		rules = append(rules, RuleType{Type: aRule})
	}

	rule := FirewallRule{
		Network: d.Get("rule_cidr").(string),
		Rules:   rules,
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
	d.SetId(d.Get("rule_cidr").(string))
	return nil

}

func resourceFirewallRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("cluster_id").(string)
	rule := d.Get("rule_cidr").(string)
	log.Printf("[INFO] Reading the status of cluster %s.", id)
	firewallRules, err := client.ReadFirewallRules(id)
	if err != nil {
		return fmt.Errorf("[Error] Error reading firewall rules: %s", err)
	}
	for _, value := range *firewallRules {
		if value.Network == rule {
			log.Printf("[INFO] Read rule %s from cluster %s", value.Network, id)
			d.Set("cluster_id", id)
			d.Set("rule_cidr", value.Network)
			d.SetId(value.Network)
			d.Set("rules", value.Rules)
		}
	}
	return nil
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
		Network: d.Get("rule_cidr").(string),
		Rules:   rules,
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
