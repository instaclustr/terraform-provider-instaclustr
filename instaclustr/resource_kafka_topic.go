package instaclustr

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
	"log"
	"strings"
)

func resourceKafkaTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourceKafkaTopicCreate,
		Read:   resourceKafkaTopicRead,
		Update: resourceKafkaTopicUpdate,
		Delete: resourceKafkaTopicDelete,

		Importer: &schema.ResourceImporter{
			State: resourceKafkaTopicStateImport,
		},

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"topic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"replication_factor": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"partitions": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"config": {
				Type:             schema.TypeList,
				Optional:         true,
				MaxItems:         1,
				DiffSuppressFunc: customIntDiffSuppressFunc,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compression_type": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: customStringDiffSuppressFunc,
						},
						"leader_replication_throttled_replicas": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"min_insync_replicas": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"message_downconversion_enable": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"segment_jitter_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"cleanup_policy": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: customStringDiffSuppressFunc,
						},
						"flush_ms": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: customStringDiffSuppressFunc,
						},
						"follower_replication_throttled_replicas": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: customStringDiffSuppressFunc,
						},
						"retention_ms": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"segment_bytes": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"flush_messages": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: customStringDiffSuppressFunc,
						},
						"message_format_version": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: customStringDiffSuppressFunc,
						},
						"file_delete_delay_ms": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"max_compaction_lag_ms": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: customStringDiffSuppressFunc,
						},
						"max_message_bytes": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"min_compaction_lag_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"message_timestamp_type": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: customStringDiffSuppressFunc,
						},
						"preallocate": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"index_interval_bytes": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"min_cleanable_dirty_ratio": {
							Type:             schema.TypeFloat,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"unclean_leader_election_enable": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"delete_retention_ms": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"retention_bytes": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"segment_ms": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
						"message_timestamp_difference_max_ms": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: customStringDiffSuppressFunc,
						},
						"segment_index_bytes": {
							Type:             schema.TypeInt,
							Optional:         true,
							DiffSuppressFunc: customIntDiffSuppressFunc,
						},
					},
				},
			},
		},
	}
}

func resourceKafkaTopicCreate(d *schema.ResourceData, meta interface{}) error {
	cluster_id := d.Get("cluster_id").(string)
	topic := d.Get("topic").(string)
	client := meta.(*Config).Client

	// Cluster has to reach running state first
	cluster, err := client.ReadCluster(cluster_id)
	if err != nil {
		return fmt.Errorf("[Error] Error in getting the status of the cluster: %w", err)
	}
	if cluster.ClusterStatus != "RUNNING" {
		return fmt.Errorf("[Error] Cluster %s is not RUNNING. Currently in %s state", cluster_id, cluster.ClusterStatus)
	}
	createKafkaTopicRequest := CreateKafkaTopicRequest{
		Topic:             topic,
		ReplicationFactor: d.Get("replication_factor").(int),
		Partitions:        d.Get("partitions").(int),
	}
	var jsonStr []byte
	jsonStr, err = json.Marshal(createKafkaTopicRequest)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka topic creation request: %w", err)
	}

	err = client.CreateKafkaTopic(cluster_id, jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka topic: %w", err)
	}

	d.SetId(fmt.Sprintf("%s&%s", cluster_id, topic))

	log.Printf("[INFO] Kafka topic %s has been created.", topic)

	configList := d.Get("config").([]interface{})
	if len(configList) == 1 {
		log.Printf("[INFO] Updating the config of Kafka topic %s in %s.", topic, cluster_id)
		updateConfig := d.Get("config").([]interface{})[0].(map[string]interface{})

		err = updateUtilities(d, meta, updateConfig, topic, client, cluster_id)
		if err != nil {
			return fmt.Errorf("[Error] Error updating the config of Kafka topic %s: %w", topic, err)
		}
	}
	return nil
}

func updateUtilities(d *schema.ResourceData, meta interface{}, configMap map[string]interface{}, topic string, client *APIClient, cluster_id string) error {
	var kafkaTopicConfigOptions KafkaTopicConfigOptions
	err := mapstructure.Decode(configMap, &kafkaTopicConfigOptions)
	if err != nil {
		return fmt.Errorf("[Error] Error decoding the changed config map to KafkaTopicConfigOptions for topic %s: %w", topic, err)
	}

	updateKafkaTopicRequest := UpdateKafkaTopicRequest{
		Config: &kafkaTopicConfigOptions,
	}

	jsonStr, err := json.Marshal(updateKafkaTopicRequest)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka topic update request: %w", err)
	}

	err = client.UpdateKafkaTopic(cluster_id, topic, jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error updating the config of Kafka topic %s: %w", topic, err)
	}

	//Reading updated configs
	err = resourceKafkaTopicRead(d, meta)
	if err != nil {
		return fmt.Errorf("[Error] Error reading the updated topic config for topic %s: %w", topic, err)
	}

	log.Printf("[INFO] The configs of Kafka topic %s has been updated.", topic)
	return nil
}

func resourceKafkaTopicRead(d *schema.ResourceData, meta interface{}) error {
	cluster_id := d.Get("cluster_id").(string)
	topic := d.Get("topic").(string)
	client := meta.(*Config).Client

	topicList, err := client.ReadKafkaTopicList(cluster_id)
	if err != nil {
		return fmt.Errorf("[Error] Error fetching the kafka user list: %w", err)
	}
	var topicExist bool
	for _, t := range topicList.Topics {
		if t == topic {
			topicExist = true
		}
	}
	if !topicExist {
		d.SetId("")
		log.Printf("[INFO] Topic %s not found in cluster %s.", topic, cluster_id)
		return nil
	}

	// Cluster has to reach running state first
	cluster, err := client.ReadCluster(cluster_id)
	if err != nil {
		return fmt.Errorf("[Error] Error in getting the status of the cluster: %w", err)
	}
	if cluster.ClusterStatus != "RUNNING" {
		return fmt.Errorf("[Error] Cluster %s is not RUNNING. Currently in %s state", cluster_id, cluster.ClusterStatus)
	}

	log.Printf("[INFO] Reading the replication factor and number of partitions of topic %s.", topic)
	kafkaTopic, err := client.ReadKafkaTopic(cluster_id, topic)
	if err != nil {
		return fmt.Errorf("[Error] Error reading Kafka topic %s's replication facotr and number of partitions: %w", topic, err)
	}
	d.Set("replication_factor", kafkaTopic.ReplicationFactor)
	d.Set("partitions", kafkaTopic.Partitions)

	log.Printf("[INFO] Reading the config of topic %s.", topic)
	kafkaTopicConfig, err := client.ReadKafkaTopicConfig(cluster_id, topic)
	if err != nil {
		return fmt.Errorf("[Error] Error reading Kafka topic %s's config: %w", topic, err)
	}

	config := make(map[string]interface{}, 0)
	configList := make([]map[string]interface{}, 0)
	err = mapstructure.Decode(kafkaTopicConfig.Config, &config)
	if err != nil {
		return fmt.Errorf("[Error] Error decoding the fetched kafka topic %s's config to a map: %w", topic, err)
	}
	configList = append(configList, config)
	err = d.Set("config", configList)
	if err != nil {
		return fmt.Errorf("[Error] Error setting the kafka topic %s's config to Terraform state: %w", topic, err)
	}

	log.Printf("[INFO] Successfully fetch the config of topic %s.", topic)
	return nil
}

// There are two use cases of resourceKafkaTopicUpdate:
// 1. Providing all the possible configs in resource, and modify the ones you want to update. Then keeping all the
// configs in the resource.
// 2. Providing some configs (not all) you want to update, and then keep only these configs in the resource.
// customStringDiffSuppressFunc is used to handle case 2, it makes sure that Terraform won't detect a change if some configs
// are not provided in the resource.
func resourceKafkaTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	cluster_id := d.Get("cluster_id").(string)
	topic := d.Get("topic").(string)
	log.Printf("[INFO] Updating the config of Kafka topic %s in %s.", topic, cluster_id)

	// This is for when only replication_factor or partitions are changed, but we don't support changing them currently
	if !d.HasChange("config") {
		log.Printf("[INFO] Currently we only support updating topic's config. There are no changes in topic %s's config.", topic)
		return nil
	}
	client := meta.(*Config).Client

	changedConfigMap, err := getChangedConfigMap(d)
	if err != nil {
		return fmt.Errorf("[Error] Error getting the changed config map for topic %s: %w", topic, err)
	}
	if len(changedConfigMap) == 0 {
		log.Printf("[INFO] There are no configs to be updated.")
		err = resourceKafkaTopicRead(d, meta)
		if err != nil {
			return fmt.Errorf("[Error] Error reading the updated topic config for topic %s: %w", topic, err)
		}
		return nil
	}

	err = updateUtilities(d, meta, changedConfigMap, topic, client, cluster_id)
	if err != nil {
		return fmt.Errorf("[Error] Error updating the config of Kafka topic %s: %w", topic, err)
	}
	return nil
}

// Although customStringDiffSuppressFunc can suppress the diff, but when getting the change it still gets all the configs that are
// not provided in the resource. This function is to filter the configs that have a real change in their values, it returns
// a map of these configs and their new values.
func getChangedConfigMap(d *schema.ResourceData) (map[string]interface{}, error) {
	newConfig := d.Get("config").([]interface{})[0].(map[string]interface{})
	changedConfigMap := make(map[string]interface{}, 0)
	var kafkaTopicConfig KafkaTopicConfigOptions
	decodedConfigMap := make(map[string]interface{}, 0) // This map stores all non-zero values and all the pointers

	// Decode twice here because zero-value non-pointers need to be removed
	err := mapstructure.Decode(newConfig, &kafkaTopicConfig)
	if err != nil {
		return nil, fmt.Errorf("[Error] Error decoding the config {} block to KafkaTopicConfigOptions: %w", err)
	}
	err = mapstructure.Decode(kafkaTopicConfig, &decodedConfigMap)
	if err != nil {
		return nil, fmt.Errorf("[Error] Error decoding KafkaTopicConfigOptions to a map: %w", err)
	}

	// The following block is to get what configs are indeed changed
	for k, v := range newConfig {
		key := fmt.Sprintf("config.0.%s", k)
		if d.HasChange(key) {
			changedConfigMap[k] = v
		}
	}
	return changedConfigMap, nil
}

func resourceKafkaTopicDelete(d *schema.ResourceData, meta interface{}) error {
	cluster_id := d.Get("cluster_id").(string)
	topic := d.Get("topic").(string)
	log.Printf("[INFO] Deleting Kafka topic %s in %s.", topic, cluster_id)
	client := meta.(*Config).Client

	err := client.DeleteKafkaTopic(cluster_id, topic)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting Kafka topic %s: %w", topic, err)
	}
	d.SetId("")
	d.Set("cluster_id", "")
	d.Set("topic", "")

	log.Printf("[INFO] Kafka topic %s has been deleted.", topic)
	return nil
}

func resourceKafkaTopicStateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "&")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("[Error] Unexpected format of ID (%q), expected <CLUSTER-ID>&<TOPIC>", d.Id())
	}
	d.Set("cluster_id", idParts[0])
	d.Set("topic", idParts[1])
	return []*schema.ResourceData{d}, nil
}

//customStringDiffSuppressFunc is used to suppress the diff if a string config is not provided in the resource.
func customStringDiffSuppressFunc(k, old, new string, d *schema.ResourceData) bool {
	return new == ""
}

//customIntDiffSuppressFunc is used to suppress the diff if an int config is not provided in the resource.
func customIntDiffSuppressFunc(k, old, new string, d *schema.ResourceData) bool {
	return new == "0"
}
