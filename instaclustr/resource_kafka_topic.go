package instaclustr

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
	"log"
	"reflect"
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
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compression_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"leader_replication_throttled_replicas": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"min_insync_replicas": {
							Type:     schema.TypeInt,
							Optional: true,
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
							Type:     schema.TypeString,
							Optional: true,
						},
						"flush_ms": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"follower_replication_throttled_replicas": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"retention_ms": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"segment_bytes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"flush_messages": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"message_format_version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"file_delete_delay_ms": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_compaction_lag_ms": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_message_bytes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_compaction_lag_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"message_timestamp_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"preallocate": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"index_interval_bytes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_cleanable_dirty_ratio": {
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"unclean_leader_election_enable": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"delete_retention_ms": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"retention_bytes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"segment_ms": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"message_timestamp_difference_max_ms": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"segment_index_bytes": {
							Type:     schema.TypeInt,
							Optional: true,
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
	log.Printf("[INFO] Reading the config of topic %s.", topic)
	kafkaTopicConfig, err := client.ReadKafkaTopicConfig(cluster_id, topic)
	if err != nil {
		return fmt.Errorf("[Error] Error reading Kafka topic %s's config: %w", topic, err)
	}

	// It won't read the topic's config if the config{} block is missing in the resource
	if len(d.Get("config").([]interface{})) == 0 {
		log.Printf("[INFO] A block of config{} needs to be provided in the resource to read the topic's config.")
		log.Printf("[INFO] Topic %s's config is not read.", topic)
		return nil
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
// Terraform always detects that there is an update in case 2 if you run the plan again, because resourceKafkaTopicRead
// reads all the configs, but there are only some of them provided in resource. For the missing configs, Terraform just
// sets them as their default values.
//
// Although Terraform detects an update, func getChangedConfigMap is used to handle case 2. It makes sure that there
// will not be a real update if the configs in the resources are not changed. So, it's safe to keep only partial of
// configs in the config{} block, there won't be an update request if their values are same as what have been read.
func resourceKafkaTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	cluster_id := d.Get("cluster_id").(string)
	topic := d.Get("topic").(string)
	log.Printf("[INFO] Updating the config of Kafka topic %s in %s.", topic, cluster_id)

	// This is for when only replication_factor or partitions are changed, but we don't support changing them currently
	if !d.HasChange("config") {
		log.Printf("[INFO] Currently we only support updating topic's config. There are no changes in topic %s's config.", topic)
		return nil
	}
	// This is for when there is no config{} block in the resource, or the whole config{} is removed from the resource.
	if len(d.Get("config").([]interface{})) == 0 {
		log.Printf("[INFO] There are no configs to be updated.")
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

	var kafkaTopicConfigOptions KafkaTopicConfigOptions
	err = mapstructure.Decode(changedConfigMap, &kafkaTopicConfigOptions)
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
		valueOfV := reflect.ValueOf(decodedConfigMap[k])
		if d.HasChange(key) {
			before, after := d.GetChange(key)
			isBeforeZeroValue := before == reflect.Zero(reflect.TypeOf(before)).Interface()
			isAfterZeroValue := after == reflect.Zero(reflect.TypeOf(after)).Interface()
			if !isBeforeZeroValue && !isAfterZeroValue {
				changedConfigMap[k] = v
			}
			if !isBeforeZeroValue && isAfterZeroValue && (valueOfV.Kind() == reflect.Ptr) {
				changedConfigMap[k] = v
			}
			if isBeforeZeroValue && !isAfterZeroValue {
				changedConfigMap[k] = v
			}
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
