package instaclustr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func (c *APIClient) DeleteKafkaAcl(clusterID string, data []byte) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/acls", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "DELETE", data)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) CreateKafkaAcl(clusterID string, data []byte) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/acls", c.apiServerHostname, clusterID)
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

func (c *APIClient) ReadKafkaAcls(clusterID string, data []byte) ([]KafkaAcl, error) {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/acls/searches", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}

	var acls KafkaAclList
	err = json.Unmarshal(bodyText, &acls)
	if err != nil {
		return nil, err
	}

	return acls.Acls, nil
}

