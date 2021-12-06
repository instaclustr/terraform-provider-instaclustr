package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/validation"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

var (
	validClusterStates = map[string]bool{
		"RUNNING":     true,
		"PROVISIONED": true,
	}
	semanticVersioningPattern, _ = regexp.Compile("([0-9]*\\.){2}[0-9]")
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceClusterStateImport,
		},

		CustomizeDiff: resourceClusterCustomizeDiff,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"cluster_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"node_size": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"data_centres"},
			},

			"data_centre": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"data_centres"},
				ForceNew:      true,
			},

			"data_centres": {
				Type:          schema.TypeSet,
				Optional:      true,
				ConflictsWith: []string{"data_centre"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						"data_centre": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"network": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						"node_size": {
							Type:     schema.TypeString,
							Required: true,
						},

						"rack_allocation": {
							Type:     schema.TypeMap,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"number_of_racks": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"nodes_per_rack": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},

						"provider": {
							Type:     schema.TypeMap,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"account_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"custom_virtual_network_id": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"resource_group": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"disk_encryption_key": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},

						"bundles": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bundle": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										ValidateFunc: validation.StringInSlice([]string{
											"APACHE_CASSANDRA",
											"SPARK",
										}, false),
									},
									"version": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										DiffSuppressFunc: versionDiffSuppressFunc,
									},
									"options": {
										Type:     schema.TypeMap,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auth_n_authz": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"client_encryption": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"password_authentication": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"use_private_broadcast_rpc_address": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"lucene_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"continuous_backup_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},

			"sla_tier": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NON_PRODUCTION",
				ForceNew: true,
			},

			"cluster_network": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"private_network_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"pci_compliant_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"public_contact_point": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"private_contact_point": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"cluster_provider": {
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{"data_centres"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"account_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"custom_virtual_network_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"disk_encryption_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},

			"rack_allocation": {
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{"data_centres"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"number_of_racks": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"nodes_per_rack": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			"bundles": {
				Type:          schema.TypeSet,
				Optional:      true,
				ConflictsWith: []string{"data_centres"},
				Elem: &schema.Schema{
					Type:     schema.TypeMap,
					Elem:     schema.TypeString,
					ForceNew: true,
				},
				Removed: "Please change bundles argument -> bundle blocks (example under example/main.tf), and to avoid causing an update to the existing tfstate - replace all keys named 'bundles' with 'bundle' in resources with the provider 'provider.instaclustr'",
			},

			"bundle": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"data_centres"},
				MinItems:      1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bundle": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							DiffSuppressFunc: versionDiffSuppressFunc,
						},
						"options": {
							// This type is not correct. TypeMaps cannot have complex structures defined in the same way that TypeLists and TypeSets can
							// See https://www.terraform.io/docs/extend/schemas/schema-types.html
							// Essentially everything in the Elem property here is being ignored, terraform assumes the element type is string
							// This should have been implemented as a TypeSet. Unfortunately changing it now would change the syntax
							// required in the terraform file and so would be a breaking change for existing configurations.
							// As such, changing this will wait until a major version change.
							Type:     schema.TypeMap,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// Cover up for the API that has optional arguments that get given default values
								// and returns the defaults in subsequent calls
								return old == "false" && new == ""
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_n_authz": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"client_encryption": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"use_private_broadcast_rpc_address": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"lucene_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"continuous_backup_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"number_partitions": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
									"auto_create_topics": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"delete_topics": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"password_authentication": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"target_kafka_cluster_id": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"vpc_type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"aws_access_key": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"aws_secret_key": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"s3_bucket_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"azure_storage_account_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"azure_storage_account_key": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"azure_storage_container_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"ssl_enabled_protocols": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"ssl_truststore_password": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"ssl_protocol": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"security_protocol": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"sasl_mechanism": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"sasl_jaas_config": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"bootstrap_servers": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"truststore": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"dedicated_zookeeper": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"zookeeper_node_size": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"zookeeper_node_count": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
									"master_nodes": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
									"replica_nodes": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
									"password_auth": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"dedicated_master_nodes": {
										Type:     schema.TypeBool,
										Optional: false,
										ForceNew: true,
									},
									"master_node_size": {
										Type:     schema.TypeString,
										Optional: false,
										ForceNew: true,
									},
									"data_node_size": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"kibana_node_size": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"opensearch_dashboards_node_size": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"postgresql_node_count": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
			"kafka_rest_proxy_user_password": {
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
			"kafka_schema_registry_user_password": {
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
			"wait_for_state": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",

				ValidateFunc: func(i interface{}, s string) (ws []string, errors []error) {
					state := i.(string)

					if len(state) != 0 && !validClusterStates[strings.ToUpper(state)] {
						errors = append(errors, fmt.Errorf("%s is not valid cluster state. Use RUNNING or PROVISIONED.", state))
					}

					return
				},
			},
		},
	}
}

func versionDiffSuppressFunc(k, old string, new string, d *schema.ResourceData) bool {
	oldSemVer := semanticVersioningPattern.FindString(old)
	newSemVer := semanticVersioningPattern.FindString(new)
	return oldSemVer == newSemVer
}

func resourceClusterCustomizeDiff(diff *schema.ResourceDiff, i interface{}) error {

	if _, isBundle := diff.GetOk("bundle"); isBundle {
		bundle := diff.Get("bundle").([]interface{})
		bundleMap := bundle[0].(map[string]interface{})

		// Mainly check Single DC Redis Cluster
		if bundleMap["bundle"] == "REDIS" {
			if _, isRackAllocationSet := diff.GetOk("rack_allocation"); isRackAllocationSet {
				return fmt.Errorf("[Error] 'rack_allocation' is not supported in REDIS")
			}
			// Remove this logic once INS-13970 is implemented
			if diff.Id() != "" && (diff.HasChange("bundle.0.options")) {
				err := diff.ForceNew("bundle.0.options")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func getNodeSize(d resourceDataInterface, bundles []Bundle) (string, error) {
	for i, bundle := range bundles {
		if bundle.Bundle == "ELASTICSEARCH" || bundle.Bundle == "OPENSEARCH" {
			if len(bundle.Options.MasterNodeSize) == 0 {
				return "", fmt.Errorf("[ERROR] 'master_node_size' is required in the bundle option.")
			}
			dedicatedMaster := bundle.Options.DedicatedMasterNodes != nil && *bundle.Options.DedicatedMasterNodes
			if dedicatedMaster {
				if len(bundle.Options.DataNodeSize) == 0 {
					return "", fmt.Errorf("[ERROR] dedicated master is enabled, 'data_node_size' is required in the bundle option.")
				}
				return bundle.Options.DataNodeSize, nil
			} else {
				if len(bundle.Options.DataNodeSize) != 0 && bundle.Options.DataNodeSize != bundle.Options.MasterNodeSize {
					return "", fmt.Errorf("[ERROR] When 'dedicated_master_nodes' is not true , data_node_size can be either null or equal to master_node_size.")
				}
				size := bundle.Options.MasterNodeSize
				bundles[i].Options.MasterNodeSize = ""
				return size, nil
			}
		}
	}
	if size, ok := d.Get("node_size").(string); ok {
		return size, nil
	}
	return "", nil
}

func resourceClusterCreate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("[INFO] Creating cluster.")
	client := meta.(*Config).Client

	bundles, err := getBundles(d)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	var clusterProvider ClusterProvider
	err = mapstructure.Decode(d.Get("cluster_provider").(map[string]interface{}), &clusterProvider)
	if err != nil {
		return err
	}

	clusterProvider.Tags = d.Get("tags").(map[string]interface{})

	size, err := getNodeSize(d, bundles)
	if err != nil {
		return err
	}
	var createData = CreateRequest{
		ClusterName:           d.Get("cluster_name").(string),
		Bundles:               bundles,
		Provider:              &clusterProvider,
		SlaTier:               d.Get("sla_tier").(string),
		NodeSize:              size,
		PrivateNetworkCluster: fmt.Sprintf("%v", d.Get("private_network_cluster")),
		PCICompliantCluster:   fmt.Sprintf("%v", d.Get("pci_compliant_cluster")),
	}

	dataCentre := d.Get("data_centre").(string)
	dataCentres, err := getDataCentres(d)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	// we will throw an error if neither data_centre nor data_centres found
	if dataCentre == "" && len(dataCentres) == 0 {
		return fmt.Errorf("[Error] Error creating cluster: either data_centre or data_centres should be provided")
	}

	var isSingleDCCluster = dataCentre != "" && len(dataCentres) == 0

	// for a single DC cluster
	if isSingleDCCluster {
		clusterNetwork := d.Get("cluster_network").(string)
		createData.DataCentre = dataCentre
		createData.ClusterNetwork = clusterNetwork
	} else {
		createData.DataCentres = dataCentres
	}

	kafkaSchemaRegistryUserPassword := d.Get("kafka_schema_registry_user_password").(string)
	kafkaRestProxyUserPassword := d.Get("kafka_rest_proxy_user_password").(string)
	waitForClusterState := d.Get("wait_for_state").(string)

	bundleConfig := getBundleConfig(bundles)

	if (len(kafkaSchemaRegistryUserPassword) > 0 || len(kafkaRestProxyUserPassword) > 0) && waitForClusterState != "RUNNING" {
		return fmt.Errorf("[Error] wait_for_state must be set as RUNNING when providing the kafka-schema-registry or kafka-rest-proxy user password")
	}

	if !bundleConfig.IsKafkaCluster && (len(kafkaSchemaRegistryUserPassword) > 0 || len(kafkaRestProxyUserPassword) > 0) {
		return fmt.Errorf("[Error] kafka-schema-registry or kafka-rest-proxy user passwords may only be provided for Kafka clusters")
	}

	// Some Bundles do not use Rack Allocation so add that separately if needed. (Redis for example)
	if isSingleDCCluster && checkIfBundleRequiresRackAllocation(bundles) {
		var rackAllocation RackAllocation
		err = mapstructure.Decode(d.Get("rack_allocation").(map[string]interface{}), &rackAllocation)
		if err != nil {
			return err
		}

		createData.RackAllocation = &rackAllocation
	}
	// for multi-DC cluster
	if len(dataCentres) > 1 {
		createData.RackAllocation = createData.DataCentres[0].RackAllocation
		createData.NodeSize = createData.DataCentres[0].NodeSize
		createData.Provider = createData.DataCentres[0].Provider
		createData.Bundles = createData.DataCentres[0].Bundles
	}

	var jsonStrCreate []byte
	jsonStrCreate, err = json.Marshal(createData)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	id, err := client.CreateCluster(jsonStrCreate)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	log.Printf("[DEBUG] Instaclustr REST API request: %s", jsonStrCreate)

	d.SetId(id)
	d.Set("cluster_id", id)
	log.Printf("[INFO] Cluster %s has been created.", id)

	if len(waitForClusterState) == 0 {
		return nil
	}

	return waitForClusterStateAndDoUpdate(client, waitForClusterState, bundleConfig, kafkaRestProxyUserPassword, kafkaSchemaRegistryUserPassword, d, id, meta)
}

func waitForClusterStateAndDoUpdate(client *APIClient,
	waitForClusterState string,
	bundleConfig BundleConfig,
	kafkaRestProxyUserPassword string,
	kafkaSchemaRegistryUserPassword string,
	d *schema.ResourceData,
	id string,
	meta interface{}) error {
	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		fmt.Printf("waiting for cluster to reach %s\n", waitForClusterState)
		//reading cluster details
		cluster, err := client.ReadCluster(id)
		resourceClusterRead(d, meta)

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("[Error] Error retrieving cluster info: %s", err))
		}

		if cluster.ClusterStatus != waitForClusterState {
			return resource.RetryableError(fmt.Errorf("[DEBUG] Cluster is in state %s, waiting for it to reach state %s", cluster.ClusterStatus, waitForClusterState))
		}

		if bundleConfig.IsKafkaCluster && bundleConfig.HasRestProxy && (len(kafkaRestProxyUserPassword) > 0) {
			err = client.UpdateBundleUser(d.Get("cluster_id").(string), "kafka_rest_proxy", createBundleUserUpdateRequest("ickafkarest", d.Get("kafka_rest_proxy_user_password").(string)))
			if err != nil {
				return resource.RetryableError(fmt.Errorf("[DEBUG] Error updating the kafka rest proxy bundle user password : %s", err))
			}
		}

		if bundleConfig.IsKafkaCluster && bundleConfig.HasSchemaRegistry && (len(kafkaSchemaRegistryUserPassword) > 0) {
			err = client.UpdateBundleUser(d.Get("cluster_id").(string), "kafka_schema_registry", createBundleUserUpdateRequest("ickafkaschema", d.Get("kafka_schema_registry_user_password").(string)))
			if err != nil {
				return resource.RetryableError(fmt.Errorf("[DEBUG] Error updating the kafka schema registry bundle user password : %s", err))
			}
		}

		return nil
	})
}

func resourceClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	// currently only cluster resize, kafka-schema-registry user password update and kafka-rest-proxy user password update are supported

	client := meta.(*Config).Client
	clusterID := d.Get("cluster_id").(string)

	log.Printf("[INFO] Updating cluster %s.", clusterID)

	kafkaSchemaRegistryUserUpdate := d.HasChange("kafka_schema_registry_user_password")
	kafkaRestProxyUserUpdate := d.HasChange("kafka_rest_proxy_user_password")

	bundles, err := getBundles(d)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	bundleConfig := getBundleConfig(bundles)

	if bundleConfig.IsKafkaCluster && bundleConfig.HasSchemaRegistry && kafkaSchemaRegistryUserUpdate {
		//updating the bundle user
		err = client.UpdateBundleUser(clusterID, "kafka_schema_registry", createBundleUserUpdateRequest("ickafkaschema", d.Get("kafka_schema_registry_user_password").(string)))
		if err != nil {
			return fmt.Errorf("[Error] Error updating the password for kafka schema registry user : %s", err)
		}
	}

	if bundleConfig.IsKafkaCluster && bundleConfig.HasRestProxy && kafkaRestProxyUserUpdate {
		//updating the bundle user
		err = client.UpdateBundleUser(clusterID, "kafka_rest_proxy", createBundleUserUpdateRequest("ickafkarest", d.Get("kafka_rest_proxy_user_password").(string)))
		if err != nil {
			return fmt.Errorf("[Error] Error updating the password for kafka rest proxy user : %s", err)
		}
	}

	err = doClusterResize(client, clusterID, d, bundles)
	if err != nil {
		return fmt.Errorf("[Error] Error resizing the cluster : %s", err)
	}

	if !bundleConfig.IsKafkaCluster && (kafkaSchemaRegistryUserUpdate || kafkaRestProxyUserUpdate) {
		return fmt.Errorf("[Error] Error updating the bundle user passwords, because it should be a KAFKA cluster in order to update the schema-registry or rest-proxy users")
	}

	d.SetPartial("node_size")
	d.SetPartial("kafka_schema_registry_user_password")
	d.SetPartial("kafka_rest_proxy_user_password")
	return nil
}

func createBundleUserUpdateRequest(bundleUsername string, bundleUserPassword string) []byte {

	var err error
	//preparing the bundle user update request
	updateBundleUserData := UpdateBundleUserRequest{
		Username: bundleUsername,
		Password: bundleUserPassword,
	}
	var jsonStrUpdateBundleUser []byte
	jsonStrUpdateBundleUser, err = json.Marshal(updateBundleUserData)

	if err != nil {
		log.Printf("[ERROR] Error creating the bundle user update request : %s", err)
		return nil
	}
	return jsonStrUpdateBundleUser
}

func getBundleConfig(bundles []Bundle) BundleConfig {
	configs := BundleConfig{
		IsKafkaCluster:    false,
		HasRestProxy:      false,
		HasSchemaRegistry: false,
	}

	for i := 0; i < len(bundles); i++ {

		if bundles[i].Bundle == "KAFKA" {
			configs.IsKafkaCluster = true
		}
		if bundles[i].Bundle == "KAFKA_REST_PROXY" {
			configs.HasRestProxy = true
		}
		if bundles[i].Bundle == "KAFKA_SCHEMA_REGISTRY" {
			configs.HasSchemaRegistry = true
		}
	}
	return configs
}

func appendIfMissing(slice []string, toAppend string) []string {
	for _, element := range slice {
		if element == toAppend {
			return slice
		}
	}
	return append(slice, toAppend)
}

func getBundleIndex(bundleType string, bundles []Bundle) (int, error) {
	for i := range bundles {
		if bundles[i].Bundle == bundleType {
			return i, nil
		}
	}
	return -1, fmt.Errorf("can't find bundle %s", bundleType)
}

func hasElasticsearchSizeChanges(bundleIndex int, d resourceDataInterface) bool {
	return len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "master_node_size"))) > 0 ||
		len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "kibana_node_size"))) > 0 ||
		len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "data_node_size"))) > 0
}

func hasOpenSearchSizeChanges(bundleIndex int, d resourceDataInterface) bool {
	return len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "master_node_size"))) > 0 ||
		len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "opensearch_dashboards_node_size"))) > 0 ||
		len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "data_node_size"))) > 0
}

func hasKafkaSizeChanges(bundleIndex int, d resourceDataInterface) bool {
	return len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "zookeeper_node_size"))) > 0 ||
		len(getNewSizeOrEmpty(d, "node_size")) > 0
}

func hasCassandraSizeChanges(d resourceDataInterface) bool {
	return len(getNewSizeOrEmpty(d, "node_size")) > 0
}

func hasRedisSizeChanges(d resourceDataInterface) bool {
	return len(getNewSizeOrEmpty(d, "node_size")) > 0
}

type resourceDataInterface interface {
	HasChange(key string) bool
	GetChange(key string) (interface{}, interface{})
	GetOk(key string) (interface{}, bool)
	Get(key string) interface{}
}

func doClusterResize(client APIClientInterface, clusterID string, d resourceDataInterface, bundles []Bundle) error {
	cluster, err := client.ReadCluster(clusterID)
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}
	bundleIndex, err := getBundleIndex(cluster.BundleType, bundles)
	if err != nil {
		return err
	}
	switch cluster.BundleType {
	case "APACHE_CASSANDRA":
		if hasCassandraSizeChanges(d) {
			return doLegacyCassandraClusterResize(client, cluster, d)
		} else {
			return nil
		}
	case "ELASTICSEARCH":
		if hasElasticsearchSizeChanges(bundleIndex, d) {
			return doElasticsearchClusterResize(client, cluster, d, bundleIndex)
		} else {
			return nil
		}
	case "OPENSEARCH":
		if hasOpenSearchSizeChanges(bundleIndex, d) {
			return doOpenSearchClusterResize(client, cluster, d, bundleIndex)
		} else {
			return nil
		}
	case "KAFKA":
		if hasKafkaSizeChanges(bundleIndex, d) {
			return doKafkaClusterResize(client, cluster, d, bundleIndex)
		} else {
			return nil
		}
	case "REDIS":
		if hasRedisSizeChanges(d) {
			return doRedisClusterResize(client, cluster, d, bundleIndex)
		} else {
			return nil
		}
	default:
		return fmt.Errorf("CDC resize does not support: %s", cluster.BundleType)
	}
}

func getNewSizeOrEmpty(d resourceDataInterface, key string) string {
	if !d.HasChange(key) {
		return ""
	}
	_, after := d.GetChange(key)
	return after.(string)
}

func isSearchSizeAllChange(dashboardSize, masterSize, dataSize string, dashboard, dataNodes bool) (string, bool) {
	if len(masterSize) == 0 {
		return "", false
	}
	if dashboard && (len(dashboardSize) == 0 || dashboardSize != masterSize) {
		return "", false
	}
	if dataNodes && (len(dataSize) == 0 || dataSize != masterSize) {
		return "", false
	}
	return masterSize, true
}

func getSingleChangedElasticsearchSizeAndPurpose(kibanaSize, masterSize, dataSize string, kibana, dataNodes bool) (string, NodePurpose, error) {
	changedCount := 0
	var nodePurpose NodePurpose
	var nodeSize string
	if len(masterSize) > 0 {
		changedCount += 1
		nodePurpose = ELASTICSEARCH_MASTER
		nodeSize = masterSize
	}
	if len(dataSize) > 0 {
		if !dataNodes {
			return "", "", fmt.Errorf("[ERROR] This cluster does not have data only nodes, so data_node_size is not used for this cluster. Please use master_node_size to change the size instead")
		}
		changedCount += 1
		nodePurpose = ELASTICSEARCH_DATA_AND_INGEST
		nodeSize = dataSize
	}
	if len(kibanaSize) > 0 {
		if !kibana {
			return "", "", fmt.Errorf("[ERROR] This cluster didn't enable Kibana, kibana_node_size is not used for this cluster. Please use master_node_size to change the size instead")
		}
		changedCount += 1
		nodePurpose = ELASTICSEARCH_KIBANA
		nodeSize = kibanaSize
	}
	if changedCount > 1 {
		return "", "", fmt.Errorf("[ERROR] Please change either a single node size at a time or change all nodes to the same size at once")
	}
	return nodeSize, nodePurpose, nil
}

func getSingleChangedOpenSearchSizeAndPurpose(openSearchDashboardsSize, masterSize, dataSize string, openSearchDashboards, dataNodes bool) (string, NodePurpose, error) {
	changedCount := 0
	var nodePurpose NodePurpose
	var nodeSize string
	if len(masterSize) > 0 {
		changedCount += 1
		nodePurpose = OPENSEARCH_MASTER
		nodeSize = masterSize
	}
	if len(dataSize) > 0 {
		if !dataNodes {
			return "", "", fmt.Errorf("[ERROR] This cluster does not have data only nodes, so data_node_size is not used for this cluster. Please use master_node_size to change the size instead")
		}
		changedCount += 1
		nodePurpose = OPENSEARCH_DATA_AND_INGEST
		nodeSize = dataSize
	}
	if len(openSearchDashboardsSize) > 0 {
		if !openSearchDashboards {
			return "", "", fmt.Errorf("[ERROR] This cluster didn't enable OpenSearch Dashboards, opensearch_dashboards_node_size is not used for this cluster. Please use master_node_size to change the size instead")
		}
		changedCount += 1
		nodePurpose = OPENSEARCH_DASHBOARDS
		nodeSize = openSearchDashboardsSize
	}
	if changedCount > 1 {
		return "", "", fmt.Errorf("[ERROR] Please change either a single node size at a time or change all nodes to the same size at once")
	}
	return nodeSize, nodePurpose, nil
}

func getBundleOptionKey(bundleIndex int, option string) string {
	return fmt.Sprintf("bundle.%d.options.%s", bundleIndex, option)
}

func doElasticsearchClusterResize(client APIClientInterface, cluster *Cluster, d resourceDataInterface, bundleIndex int) error {
	kibana := len(cluster.BundleOption.KibanaNodeSize) > 0
	dataNodes := len(cluster.BundleOption.DataNodeSize) > 0
	var nodePurpose *NodePurpose
	var nodeSize string
	masterNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "master_node_size"))
	kibanaNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "kibana_node_size"))
	dataNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "data_node_size"))

	if newSize, isAllChange := isSearchSizeAllChange(kibanaNewSize, masterNewSize, dataNewSize, kibana, dataNodes); isAllChange {
		nodeSize = newSize
		nodePurpose = nil
	} else {
		newSize, purpose, err := getSingleChangedElasticsearchSizeAndPurpose(kibanaNewSize, masterNewSize, dataNewSize, kibana, dataNodes)
		if err != nil {
			return err
		}
		nodeSize = newSize
		nodePurpose = &purpose
	}
	log.Printf("[INFO] Resizing Elasticsearch cluster. nodePurpose: %s, newSize: %s", nodePurpose, nodeSize)
	err := client.ResizeCluster(cluster.ID, cluster.DataCentres[0].ID, nodeSize, nodePurpose)
	if err != nil {
		return fmt.Errorf("[Error] Error resizing cluster %s with error %s", cluster.ID, err)
	}
	return nil
}

func doOpenSearchClusterResize(client APIClientInterface, cluster *Cluster, d resourceDataInterface, bundleIndex int) error {
	openSearchDashboards := len(cluster.BundleOption.OpenSearchDashboardsNodeSize) > 0
	dataNodes := len(cluster.BundleOption.DataNodeSize) > 0
	var nodePurpose *NodePurpose
	var nodeSize string
	masterNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "master_node_size"))
	openSearchDashboardsNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "opensearch_dashboards_node_size"))
	dataNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "data_node_size"))

	if newSize, isAllChange := isSearchSizeAllChange(openSearchDashboardsNewSize, masterNewSize, dataNewSize, openSearchDashboards, dataNodes); isAllChange {
		nodeSize = newSize
		nodePurpose = nil
	} else {
		newSize, purpose, err := getSingleChangedOpenSearchSizeAndPurpose(openSearchDashboardsNewSize, masterNewSize, dataNewSize, openSearchDashboards, dataNodes)
		if err != nil {
			return err
		}
		nodeSize = newSize
		nodePurpose = &purpose
	}
	log.Printf("[INFO] Resizing OpenSearch cluster. nodePurpose: %s, newSize: %s", nodePurpose, nodeSize)
	err := client.ResizeCluster(cluster.ID, cluster.DataCentres[0].ID, nodeSize, nodePurpose)
	if err != nil {
		return fmt.Errorf("[ERROR] Error resizing cluster %s with error %s", cluster.ID, err)
	}
	return nil
}

func isKafkaSizeAllChange(brokerSize, zookeeperSize string, dedicatedZookeeper bool) (string, bool) {
	if len(brokerSize) == 0 {
		return "", false
	}
	if dedicatedZookeeper && (len(zookeeperSize) == 0 || zookeeperSize != brokerSize) {
		return "", false
	}
	return brokerSize, true
}

func getSingleChangedKafkaSizeAndPurpose(brokerSize, zookeeperSize string, dedicatedZookeeper bool) (string, NodePurpose, error) {
	changedCount := 0
	var nodePurpose NodePurpose
	var nodeSize string
	if len(brokerSize) > 0 {
		changedCount += 1
		nodePurpose = KAFKA_BROKER
		nodeSize = brokerSize
	}
	if len(zookeeperSize) > 0 {
		if !dedicatedZookeeper {
			return "", "", fmt.Errorf("[ERROR] This cluster didn't enable Dedicated Zookeeper, zookeeper_node_size is not used for this cluster. Please use master_node_size to change the size instead")
		}
		nodePurpose = KAFKA_DEDICATED_ZOOKEEPER
		nodeSize = zookeeperSize
		changedCount += 1
	}
	if changedCount > 1 {
		return "", "", fmt.Errorf("[ERROR] Please change either a single node size at a time or change all nodes to the same size at once")
	}
	return nodeSize, nodePurpose, nil
}

func doKafkaClusterResize(client APIClientInterface, cluster *Cluster, d resourceDataInterface, bundleIndex int) error {
	dedicatedZookeeper := cluster.BundleOption.DedicatedZookeeper != nil && *cluster.BundleOption.DedicatedZookeeper

	var nodePurpose *NodePurpose
	var nodeSize string
	zookeeperNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "zookeeper_node_size"))
	brokerNewSize := getNewSizeOrEmpty(d, "node_size")

	if newSize, isAllChange := isKafkaSizeAllChange(brokerNewSize, zookeeperNewSize, dedicatedZookeeper); isAllChange {
		nodeSize = newSize
		nodePurpose = nil
	} else {
		newSize, purpose, err := getSingleChangedKafkaSizeAndPurpose(brokerNewSize, zookeeperNewSize, dedicatedZookeeper)
		if err != nil {
			return err
		}
		nodeSize = newSize
		nodePurpose = &purpose
	}

	log.Printf("[INFO] Resizing Kafka cluster. nodePurpose: %s, newSize: %s", nodePurpose, nodeSize)
	err := client.ResizeCluster(cluster.ID, cluster.DataCentres[0].ID, nodeSize, nodePurpose)
	if err != nil {
		return fmt.Errorf("[Error] Error resizing cluster %s with error %s", cluster.ID, err)
	}
	return nil
}

func doLegacyCassandraClusterResize(client APIClientInterface, cluster *Cluster, d resourceDataInterface) error {
	before, after := d.GetChange("node_size")
	regex := regexp.MustCompile(`resizeable-(small|large)`)
	oldNodeClass := regex.FindString(before.(string))
	newNodeClass := regex.FindString(after.(string))

	isNotResizable := oldNodeClass == ""
	isNotSameSizeClass := newNodeClass != oldNodeClass
	if isNotResizable || isNotSameSizeClass {
		return fmt.Errorf("[Error] Cannot resize nodes from %s to %s", before, after)
	}

	err := client.ResizeCluster(cluster.ID, cluster.DataCentres[0].ID, after.(string), nil)
	if err != nil {
		return fmt.Errorf("[Error] Error resizing cluster %s with error %s", cluster.ID, err)
	}
	return nil
}

func getChangedRedisSizeAndPurpose(nodeSize string) (string, NodePurpose, error) {
	if len(nodeSize) > 0 {
		return nodeSize, REDIS, nil
	}
	return "", "", fmt.Errorf("[ERROR] Please change node size before resize")
}

func doRedisClusterResize(client APIClientInterface, cluster *Cluster, d resourceDataInterface, bundleIndex int) error {
	var nodePurpose *NodePurpose
	var nodeSize string
	nodeNewSize := getNewSizeOrEmpty(d, "node_size")

	newSize, purpose, err := getChangedRedisSizeAndPurpose(nodeNewSize)
	if err != nil {
		return err
	}
	nodeSize = newSize
	nodePurpose = &purpose

	log.Printf("[INFO] Resizing Redis cluster. nodePurpose: %s, newSize: %s", nodePurpose, nodeSize)
	err = client.ResizeCluster(cluster.ID, cluster.DataCentres[0].ID, nodeSize, nodePurpose)
	if err != nil {
		return fmt.Errorf("[Error] Error resizing cluster %s with error %s", cluster.ID, err)
	}
	return nil
}

func resourceClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("cluster_id").(string)
	log.Printf("[INFO] Reading status of cluster %s.", id)
	cluster, err := client.ReadCluster(id)
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}
	d.SetId(cluster.ID)
	d.Set("cluster_id", cluster.ID)
	d.Set("cluster_name", cluster.ClusterName)

	if isClusterSingleDataCentre(*cluster) {
		bundles, err := getBundlesFromCluster(cluster)
		if err != nil {
			return err
		}
		if err := d.Set("bundle", bundles); err != nil {
			return fmt.Errorf("[Error] Error reading cluster: %s", err)
		}

		clusterProvider := make(map[string]interface{}, 0)
		mapstructure.Decode(cluster.Provider[0], &clusterProvider)
		processedClusterProvider := processProvider(d, clusterProvider)
		d.Set("cluster_provider", processedClusterProvider)

		nodeSize := ""
		/*
		*  Ideally, we would like this information to be coming directly from the API cluster status.
		*  Hence, this is a slightly hacky way of ignoring zookeeper node sizes (Kafka bundles specific).
		 */
		for _, node := range cluster.DataCentres[0].Nodes {
			nodeSize = node.Size
			if !isDedicatedZookeeperNodeSize(nodeSize) {
				break
			}
		}

		//loop over each data centre to determine nodes and racks for rack allocation
		nodeCount := 0
		rackCount := 0
		rackList := make([]string, 0)
		for _, node := range cluster.DataCentres[0].Nodes {
			if !isDedicatedZookeeperNodeSize(node.Size) {
				nodeCount += 1
			}
			rackList = appendIfMissing(rackList, node.Rack)
		}
		rackCount += len(rackList)

		nodesPerRack := nodeCount / rackCount
		rackAllocation := make(map[string]interface{}, 0)
		rackAllocation["number_of_racks"] = strconv.Itoa(rackCount)
		rackAllocation["nodes_per_rack"] = strconv.Itoa(nodesPerRack)

		if cluster.BundleType == "REDIS" {
			rackAllocation = nil
		}

		if err := d.Set("rack_allocation", rackAllocation); err != nil {
			return fmt.Errorf("[Error] Error reading cluster, rack allocation could not be derived: %s", err)
		}

		if len(cluster.DataCentres[0].ResizeTargetNodeSize) > 0 {
			nodeSize = cluster.DataCentres[0].ResizeTargetNodeSize
		}
		if cluster.BundleType != "ELASTICSEARCH" && cluster.BundleType != "OPENSEARCH" {
			d.Set("node_size", nodeSize)
		}
		d.Set("data_centre", cluster.DataCentres[0].Name)
		d.Set("cluster_network", cluster.DataCentres[0].CdcNetwork)

		if err := deleteAttributesConflict(resourceCluster().Schema, d, "data_centre"); err != nil {
			return err
		}
	} else {

		dataCentres, err := getDataCentresFromCluster(cluster)
		if err != nil {
			return err
		}
		// set data centres
		if len(dataCentres) > 1 {
			if err := d.Set("data_centres", dataCentres); err != nil {
				return fmt.Errorf("[Error] Error setting data centres into terraform state, data centres could not be derived: %s", err)
			}
		}

		if err := deleteAttributesConflict(resourceCluster().Schema, d, "data_centres"); err != nil {
			return err
		}
	}

	d.Set("sla_tier", strings.ToUpper(cluster.SlaTier))
	d.Set("private_network_cluster", cluster.DataCentres[0].PrivateIPOnly)
	d.Set("pci_compliant_cluster", cluster.PciCompliance == "ENABLED")

	azList := make([]string, 0)
	publicContactPointList := make([]string, 0)
	privateContactPointList := make([]string, 0)

	for _, dataCentre := range cluster.DataCentres {
		for _, node := range dataCentre.Nodes {
			if !stringInSlice(node.Rack, azList) {
				if !isDedicatedZookeeperNodeSize(node.Size) {
					azList = appendIfMissing(azList, node.Rack)
					privateContactPointList = appendIfMissing(privateContactPointList, node.PrivateAddress)
					publicContactPointList = appendIfMissing(publicContactPointList, node.PublicAddress)
				}
			}
		}
	}

	if !cluster.DataCentres[0].PrivateIPOnly && len(publicContactPointList) > 0 {
		err = d.Set("public_contact_point", publicContactPointList)
	} else {
		err = d.Set("public_contact_point", nil)
	}

	if len(privateContactPointList) > 0 {
		err = d.Set("private_contact_point", privateContactPointList)
	} else {
		err = d.Set("public_contact_point", nil)
	}

	toCheck := [2]string{"cluster_provider", "rack_allocation"}
	for _, changing := range toCheck {

		if !d.HasChange(changing) {
			continue
		}
		_, after := d.GetChange(changing)
		d.Set(changing, after)
	}

	log.Printf("[INFO] Fetched cluster %s info from the remote server.", cluster.ID)
	return nil
}

func deleteAttributesConflict(schema map[string]*schema.Schema, d *schema.ResourceData, conflictAttr string) error {
	for key, value := range schema {
		if _, exist := d.GetOk(key); exist {
			for _, conflictsWith := range value.ConflictsWith {
				if conflictsWith == conflictAttr {
					if err := d.Set(key, value.Type.Zero()); err != nil {
						return err
					}
					break
				}
			}
		}
	}
	return nil
}

func getBaseBundlesFromCluster(cluster *Cluster) ([]map[string]interface{}, error) {
	baseBundle := make(map[string]interface{}, 3)
	baseBundle["bundle"] = cluster.BundleType

	baseBundleOptions := make(map[string]interface{}, 0)
	err := mapstructure.Decode(cluster.BundleOption, &baseBundleOptions)
	if err != nil {
		return nil, fmt.Errorf("[Error] Error decoding bundles option: %s", err)
	}

	convertedBundleOptions := dereferencePointerInStruct(baseBundleOptions)

	baseBundle["options"] = convertedBundleOptions
	baseBundle["version"] = cluster.BundleVersion

	bundles := make([]map[string]interface{}, 0)
	bundles = append(bundles, baseBundle)

	return bundles, nil
}

func getBundlesFromCluster(cluster *Cluster) ([]map[string]interface{}, error) {
	bundles, error := getBaseBundlesFromCluster(cluster)
	if error != nil {
		return nil, error
	}

	addonBundles := make([]map[string]interface{}, 0)
	for _, addOnBundle := range cluster.AddonBundles {
		if addOnBundle != nil {
			addonBundles = append(addonBundles, addOnBundle)
		}
	}

	if addonBundles == nil {
		return nil, nil
	}

	sort.Slice(addonBundles, func(i, j int) bool { return addonBundles[i]["bundle"].(string) > addonBundles[j]["bundle"].(string) })
	for _, addonBundle := range addonBundles {
		if len(addonBundle) != 0 {
			bundles = append(bundles, addonBundle)
		}
	}

	return bundles, nil
}

func getDataCentresFromCluster(cluster *Cluster) ([]map[string]interface{}, error) {
	dataCentres := make([]map[string]interface{}, 0)
	for _, dataCentre := range cluster.DataCentres {
		dataCentreMap := make(map[string]interface{})
		dataCentreMap["name"] = dataCentre.CdcName
		dataCentreMap["data_centre"] = dataCentre.Name
		dataCentreMap["network"] = dataCentre.CdcNetwork

		// find the node size for this data centre
		dataCentreMap["node_size"] = dataCentre.Nodes[0].Size

		// find rack allocation for each data centre
		nodeCount := 0
		rackCount := 0
		rackList := make([]string, 0)
		for _, node := range dataCentre.Nodes {
			nodeCount += 1
			rackList = appendIfMissing(rackList, node.Rack)
		}
		rackCount += len(rackList)
		nodesPerRack := nodeCount / rackCount
		rackAllocation := make(map[string]interface{}, 0)
		rackAllocation["number_of_racks"] = strconv.Itoa(rackCount)
		rackAllocation["nodes_per_rack"] = strconv.Itoa(nodesPerRack)
		dataCentreMap["rack_allocation"] = rackAllocation

		// find provider for each data centre
		provider := make(map[string]interface{})
		provider["name"] = dataCentre.Provider
		dataCentreMap["provider"] = provider

		// find bundles for each data centre
		thisDataCentreBundles, _ := getBaseBundlesFromCluster(cluster)
		if dataCentre.Bundles != nil && len(dataCentre.Bundles) != 0 {
			for _, thisDataCentreBundle := range dataCentre.Bundles {
				for _, addOnBundle := range cluster.AddonBundles {
					if addOnBundle != nil && addOnBundle["bundle"].(string) == thisDataCentreBundle {
						thisDataCentreBundles = append(thisDataCentreBundles, addOnBundle)
					}
				}
			}
		}
		dataCentreMap["bundles"] = thisDataCentreBundles

		dataCentres = append(dataCentres, dataCentreMap)
	}
	return dataCentres, nil
}

func dereferencePointerInStruct(data map[string]interface{}) map[string]interface{} {
	for k, v := range data {
		// Terraform expects strings for everything
		// This block iterates through the bundle options map and checks for an interface of a pointer
		// For each interface{*type} value, it changes: interface{*type} -> *type -> type -> interface{type} -> String
		// For a non-pointer, it directly formats to a string
		valueOfV := reflect.ValueOf(v)
		if valueOfV.Kind() == reflect.Ptr {
			data[k] = fmt.Sprintf("%v", reflect.Indirect(valueOfV.Elem()).Interface())
		} else {
			data[k] = fmt.Sprintf("%v", valueOfV)
		}
	}

	return data
}

func processProvider(d *schema.ResourceData, data map[string]interface{}) (newData map[string]interface{}) {
	// This is used to ignore values that are not set in the resource
	// Otherwise terraform will store the default and generated values (e.g. Provider Account Name)
	// into the state and result a diff between the state and resource (plan)
	newData = make(map[string]interface{})
	for k, v := range data {
		resource := d.Get(fmt.Sprintf("cluster_provider.%s", k))
		// Store the required field "name"
		// There doesn't seem to be a way to programmatically get the "required" schemas with ResourceData
		if resource != "" || k == "name" {
			newData[k] = v
		}
	}

	return newData
}

func resourceClusterDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("cluster_id").(string)
	log.Printf("[INFO] Deleting cluster %s.", id)
	err := client.DeleteCluster(id)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting cluster: %s", err)
	}
	d.SetId("")
	d.Set("cluster_id", "")
	log.Printf("[INFO] Cluster %s has been marked for deletion.", id)
	return nil
}

func resourceClusterStateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	clusterId := d.Id()
	d.Set("cluster_id", clusterId)
	return []*schema.ResourceData{d}, nil
}

func getBundles(d *schema.ResourceData) ([]Bundle, error) {
	bundles := make([]Bundle, 0)
	for _, inBundle := range d.Get("bundle").([]interface{}) {
		var bundle Bundle
		inBundleMap := inBundle.(map[string]interface{})
		if len(inBundleMap["options"].(map[string]interface{})) == 0 {
			inBundleMap["options"] = nil
		}
		err := mapstructure.WeakDecode(inBundleMap, &bundle)
		if err != nil {
			return nil, err
		}
		bundles = append(bundles, bundle)
	}

	return bundles, nil
}

func getDataCentres(d *schema.ResourceData) ([]DataCentreCreateRequest, error) {
	dataCentres_ := d.Get("data_centres").(*schema.Set)
	dataCentres := make([]DataCentreCreateRequest, 0)
	for _, inDataCentre := range dataCentres_.List() {
		var dataCentre DataCentreCreateRequest
		inDataCentreMap := inDataCentre.(map[string]interface{})
		inDataCentreBundles := inDataCentreMap["bundles"].(*schema.Set).List()
		inDataCentreMap["bundles"] = inDataCentreBundles
		err := mapstructure.WeakDecode(inDataCentreMap, &dataCentre)
		if err != nil {
			return nil, err
		}
		dataCentres = append(dataCentres, dataCentre)
	}
	return dataCentres, nil
}

func formatCreateErrMsg(err error) error {
	return fmt.Errorf("[Error] Error creating cluster: %s", err)
}

func checkIfBundleRequiresRackAllocation(bundles []Bundle) bool {
	var noRackAllocationBundles = []string{
		"REDIS",
		"APACHE_ZOOKEEPER",
		"POSTGRESQL",
	}

	for i := 0; i < len(bundles); i++ {
		for j := 0; j < len(noRackAllocationBundles); j++ {
			if strings.ToLower(bundles[i].Bundle) == strings.ToLower(noRackAllocationBundles[j]) {
				return false
			}
		}
	}

	return true
}

func isClusterSingleDataCentre(cluster Cluster) bool {
	if len(cluster.DataCentres) == 1 {
		return true
	}
	return false
}

// Currently, there is no API to tell if a node should be included as the main contact point
// or calculated in the rack allocation scheme (that is not returned by the API).
// Dedicated ZooKeeper nodes fall into this category and require a specific handling by the provider
func isDedicatedZookeeperNodeSize(nodeSize string) bool {
	return strings.HasPrefix(nodeSize, "zk-") || strings.HasPrefix(nodeSize, "KDZ-")
}
