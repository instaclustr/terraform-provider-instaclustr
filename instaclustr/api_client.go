package instaclustr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type APIClientInterface interface {
	ReadCluster(clusterID string) (*Cluster, error)
	ResizeCluster(clusterID string, cdcID string, newNodeSize string, nodePurpose *NodePurpose) error
	DeleteCluster(clusterID string) error
}

type APIClient struct {
	username          string
	apiKey            string
	apiServerHostname string
	client            *http.Client
}

func (c *APIClient) InitClient(hostname string, username string, apiKey string) {
	c.apiServerHostname = hostname
	c.username = username
	c.apiKey = apiKey
	c.client = &http.Client{
		Timeout:   time.Second * 60,
		Transport: &http.Transport{},
	}
}
func (c *APIClient) SetClient(client *http.Client) {
	c.client = client
}

func (c *APIClient) MakeRequest(url string, method string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *APIClient) CreateCluster(data []byte) (string, error) {
	url := fmt.Sprintf("%s/provisioning/v1/extended/", c.apiServerHostname)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return "", errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var respJson interface{}
	var id string
	err = json.Unmarshal(bodyText, &respJson)
	if err != nil {
		return "", err
	}
	respJsonData := respJson.(map[string]interface{})
	for _, value := range respJsonData {
		id = fmt.Sprintf("%v", value)
	}
	return id, nil
}

func (c *APIClient) ListClusters() (*[]ClusterListItem, error) {
	url := fmt.Sprintf("%s/provisioning/v1", c.apiServerHostname)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 { // 404s are returned when no clusters are found
		emptyResponse := make([]ClusterListItem, 0)
		return &emptyResponse, nil
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var clusters []ClusterListItem
	json.Unmarshal(bodyText, &clusters)
	return &clusters, nil
}

func (c *APIClient) ReadCluster(clusterID string) (*Cluster, error) {
	url := fmt.Sprintf("%s/provisioning/v1/%s/terraform-description", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var cluster Cluster
	json.Unmarshal(bodyText, &cluster)
	return &cluster, nil
}

func (c *APIClient) DeleteCluster(clusterID string) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "DELETE", nil)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) ResizeCluster(clusterID string, cdcID string, newNodeSize string, nodePurpose *NodePurpose) error {
	request := ResizeClusterRequest{
		NewNodeSize:           newNodeSize,
		ConcurrentResizes:     1,
		NotifySupportContacts: "false",
		NodePurpose:           nodePurpose,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("[Error] Error creating resize cluster request: %s", err)
	}

	url := fmt.Sprintf("%s/provisioning/v1/%s/%s/resize", c.apiServerHostname, clusterID, cdcID)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) CreateFirewallRule(data []byte, clusterID string) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/firewallRules/", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 409 {
		log.Printf("Firewall rule already exists")
		return nil
	}

	if resp.StatusCode != 202 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}

	return nil
}

func (c *APIClient) ReadFirewallRules(clusterID string) (*[]FirewallRule, error) {
	url := fmt.Sprintf("%s/provisioning/v1/%s/firewallRules/", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var rules []FirewallRule
	err = json.Unmarshal(bodyText, &rules)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not unmarshal JSON - Status code: %d, message: %s", resp.StatusCode, bodyText))
	}

	return &rules, nil
}

func (c *APIClient) DeleteFirewallRule(data []byte, clusterID string) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/firewallRules/", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "DELETE", data)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) CreateVpcPeering(cdcID string, data []byte) (string, error) {
	url := fmt.Sprintf("%s/provisioning/v1/vpc-peering/%s", c.apiServerHostname, cdcID)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return "", errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var respJson interface{}
	var id string
	err = json.Unmarshal(bodyText, &respJson)
	if err != nil {
		return "", err
	}
	respJsonData := respJson.(map[string]interface{})
	for _, value := range respJsonData {
		id = fmt.Sprintf("%v", value)
	}
	return id, nil
}

func (c *APIClient) DeleteVpcPeering(cdcID string, vpcPeeringID string) error {
	url := fmt.Sprintf("%s/provisioning/v1/vpc-peering/%s/%s", c.apiServerHostname, cdcID, vpcPeeringID)
	resp, err := c.MakeRequest(url, "DELETE", nil)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) ReadVpcPeering(cdcID string, vpcPeeringID string) (*VPCPeering, error) {
	url := fmt.Sprintf("%s/provisioning/v1/vpc-peering/%s/%s", c.apiServerHostname, cdcID, vpcPeeringID)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var vpcPeering VPCPeering
	json.Unmarshal(bodyText, &vpcPeering)
	return &vpcPeering, nil
}

func (c *APIClient) CreateEncryptionKey(data []byte) (string, error) {
	url := fmt.Sprintf("%s/provisioning/v1/encryption-keys", c.apiServerHostname)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return "", err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return "", errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}

	var respJson interface{}
	var id string
	err = json.Unmarshal(bodyText, &respJson)
	if err != nil {
		return "", err
	}

	respJsonData := respJson.(map[string]interface{})
	for _, value := range respJsonData {
		id = value.(string)
	}
	return id, nil
}

func getEncryptionKeyByID(resources *[]EncryptionKey, id string) (*EncryptionKey, error) {
	for _, resource := range *resources {
		if resource.ID == id {
			return &resource, nil
		}
	}
	return nil, errors.New(id)
}

func (c *APIClient) ReadEncryptionKey(id string) (*EncryptionKey, error) {
	url := fmt.Sprintf("%s/provisioning/v1/encryption-keys", c.apiServerHostname)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var kmsKeys []EncryptionKey
	err = json.Unmarshal(bodyText, &kmsKeys)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not unmarshal JSON - Status code: %d, message: %s", resp.StatusCode, bodyText))
	}

	keyResource, err := getEncryptionKeyByID(&kmsKeys, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error encryption key %s does not exist", id))
	}

	return keyResource, nil
}

func (c *APIClient) DeleteEncryptionKey(keyID string) error {
	url := fmt.Sprintf("%s/provisioning/v1/encryption-keys/%s", c.apiServerHostname, keyID)
	resp, err := c.MakeRequest(url, "DELETE", nil)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 202 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) UpdateBundleUser(clusterID string, bundle string, data []byte) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/%s/users/reset-password", c.apiServerHostname, clusterID, bundle)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}
