package instaclustr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func (c *APIClient) ReadKafkaUserList(clusterID string) ([]string, error) {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/users", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}

	usernameList := []string{}
	err = json.Unmarshal(bodyText, &usernameList)
	if err != nil {
		return nil, err
	}

	return usernameList, nil
}

func (c *APIClient) CreateKafkaUser(clusterID string, data []byte) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/users", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 201 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) UpdateKafkaUser(clusterID string, data []byte) error {
	return c.UpdateBundleUser(clusterID, "kafka", data)
}

func (c *APIClient) DeleteKafkaUser(clusterID string, data []byte) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/users", c.apiServerHostname, clusterID)
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

func (c *APIClient) ReadKafkaTopicList(clusterID string) (*KafkaTopics, error) {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/topics", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}

	var kafkaTopics KafkaTopics
	err = json.Unmarshal(bodyText, &kafkaTopics)
	if err != nil {
		return nil, err
	}

	return &kafkaTopics, nil
}

func (c *APIClient) CreateKafkaTopic(clusterID string, data []byte) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/topics", c.apiServerHostname, clusterID)
	resp, err := c.MakeRequest(url, "POST", data)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 201 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) DeleteKafkaTopic(clusterID string, topic string) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/topics/%s", c.apiServerHostname, clusterID, topic)
	resp, err := c.MakeRequest(url, "DELETE", nil)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) UpdateKafkaTopic(clusterID string, topic string, data []byte) error {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/topics/%s/config", c.apiServerHostname, clusterID, topic)
	resp, err := c.MakeRequest(url, "PUT", data)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return nil
}

func (c *APIClient) ReadKafkaTopicConfig(clusterID string, topic string) (*KafkaTopicConfig, error) {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/topics/%s/config", c.apiServerHostname, clusterID, topic)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var kafkaTopicConfig KafkaTopicConfig
	err = json.Unmarshal(bodyText, &kafkaTopicConfig)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not unmarshal JSON - Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return &kafkaTopicConfig, nil
}

func (c *APIClient) ReadKafkaTopic(clusterID string, topic string) (*CreateKafkaTopicRequest, error) {
	url := fmt.Sprintf("%s/provisioning/v1/%s/kafka/topics/%s", c.apiServerHostname, clusterID, topic)
	resp, err := c.MakeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	var kafkaTopic CreateKafkaTopicRequest
	err = json.Unmarshal(bodyText, &kafkaTopic)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not unmarshal JSON - Status code: %d, message: %s", resp.StatusCode, bodyText))
	}
	return &kafkaTopic, nil
}
