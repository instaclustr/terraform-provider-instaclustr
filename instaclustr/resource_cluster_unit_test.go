package instaclustr

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

func TestCreateBundleUserUpdateRequest(t *testing.T) {
	testUsername := "hail"
	testPassword := "reallySecure123"
	testBundleRequest := createBundleUserUpdateRequest(testUsername, testPassword)

	expectedOutput := []byte(fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", testUsername, testPassword))

	if !reflect.DeepEqual(testBundleRequest, expectedOutput) {
		t.Fatalf("Incorrect request returned.\nExpected:%s\nActual:%s", expectedOutput, testBundleRequest)
	}
}

func TestGetBundleConfig(t *testing.T) {
	var testBundles []Bundle
	var testBundleConfig BundleConfig

	testBundles = append(testBundles, Bundle{Bundle: "KAFKA"})
	testBundleConfig = getBundleConfig(testBundles)
	expectedOutput := BundleConfig{
		IsKafkaCluster:    true,
		HasRestProxy:      false,
		HasSchemaRegistry: false}

	if testBundleConfig != expectedOutput {
		t.Fatalf("Incorrect Bundle Config returned.\nExpected: %+v\nActual: %+v", expectedOutput, testBundleConfig)
	}

	testBundles = append(testBundles, Bundle{Bundle: "KAFKA_REST_PROXY"})
	testBundleConfig = getBundleConfig(testBundles)
	expectedOutput = BundleConfig{
		IsKafkaCluster:    true,
		HasRestProxy:      true,
		HasSchemaRegistry: false}

	if testBundleConfig != expectedOutput {
		t.Fatalf("Incorrect Bundle Config returned.\nExpected: %+v\nActual: %+v", expectedOutput, testBundleConfig)
	}

	testBundles = append(testBundles, Bundle{Bundle: "KAFKA_SCHEMA_REGISTRY"})
	testBundleConfig = getBundleConfig(testBundles)
	expectedOutput = BundleConfig{
		IsKafkaCluster:    true,
		HasRestProxy:      true,
		HasSchemaRegistry: true}

	if testBundleConfig != expectedOutput {
		t.Fatalf("Incorrect Bundle Config returned.\nExpected: %+v\nActual: %+v", expectedOutput, testBundleConfig)
	}
}

func TestAppendIfMissing(t *testing.T) {
	var testSlice []string
	testSlice = append(testSlice, "1", "2")
	testString := "test"

	expectedSlice := append(testSlice, testString)
	appendedSlice := appendIfMissing(testSlice, testString)

	if !reflect.DeepEqual(expectedSlice, appendedSlice) {
		t.Fatalf("Value appened incorrectly to the slice.\nExpected: %s\nActual: %s", expectedSlice, appendedSlice)
	}

	appendedSlice = appendIfMissing(testSlice, testString)
	if !reflect.DeepEqual(expectedSlice, appendedSlice) {
		t.Fatalf("Value appened incorrectly to the slice.\nExpected: %s\nActual: %s", expectedSlice, appendedSlice)
	}
}

func TestFormatCreateErrMsg(t *testing.T) {
	testError := fmt.Errorf("test error")
	formattedError := formatCreateErrMsg(testError)
	expectedError := fmt.Sprintf("[Error] Error creating cluster: %s", testError.Error())

	if formattedError.Error() != expectedError {
		t.Fatalf("Incorrectly formatted error.\nExpected: %s\nActual: %s", expectedError, formattedError)
	}
}

func TestCheckIfBundleRequiresRackAllocation(t *testing.T) {
	bundles := []Bundle{{Bundle: "REDIS"}}
	isRackAllocationRequired := checkIfBundleRequiresRackAllocation(bundles)

	if isRackAllocationRequired == true {
		t.Fatalf("Incorrect check performed for REDIS bundle.\nExpected: %v\nActual: %v\n", false, true)
	}

	bundles = []Bundle{{Bundle: "APACHE_ZOOKEEPER"}}
	isRackAllocationRequired = checkIfBundleRequiresRackAllocation(bundles)

	if isRackAllocationRequired == true {
		t.Fatalf("Incorrect check performed for APACHE_ZOOKEEPER bundle.\nExpected: %v\nActual: %v\n", false, true)
	}

	bundles = []Bundle{{Bundle: "POSTGRESQL"}}
	isRackAllocationRequired = checkIfBundleRequiresRackAllocation(bundles)

	if isRackAllocationRequired == true {
		t.Fatalf("Incorrect check performed for POSTGRESQL bundle.\nExpected: %v\nActual: %v\n", false, true)
	}

	bundles = []Bundle{{Bundle: "APACHE_CASSANDRA"}}
	isRackAllocationRequired = checkIfBundleRequiresRackAllocation(bundles)

	if isRackAllocationRequired == false {
		t.Fatalf("Incorrect check performed for APACHE_CASSANDRA.\nExpected: %v\nActual: %v\n", true, false)
	}
}

func TestIsElasticsearchSizeAllChange(t *testing.T) {
	helper := func(kibanaSize, masterSize, dataSize, expectedNewSize string, kibana, dedicatedMaster, expectedIsAll bool) {
		newSize, isAllChange := isSearchSizeAllChange(kibanaSize, masterSize, dataSize, kibana, dedicatedMaster)
		if isAllChange != expectedIsAll {
			t.Fatalf("changeAll should be %t when using kibanaSize: %s, masterSize: %s, dataSize: %s, kibana: %t, dedicatedMaster: %t", expectedIsAll, kibanaSize, masterSize, dataSize, kibana, dedicatedMaster)
		}
		if isAllChange && newSize != expectedNewSize {
			t.Fatalf("newSize should be %s when using kibanaSize: %s, masterSize: %s, dataSize: %s, kibana: %t, dedicatedMaster:%t", expectedNewSize, kibanaSize, masterSize, dataSize, kibana, dedicatedMaster)
		}
	}
	helper("t3.small-v2", "t3.small-v2", "t3.small-v2", "t3.small-v2", true, true, true)
	helper("t3.small-v2", "t3.small-v2", "", "t3.small-v2", true, false, true)
	helper("", "t3.small-v2", "t3.small-v2", "t3.small-v2", false, true, true)
	helper("", "t3.small-v2", "", "t3.small-v2", false, false, true)

	helper("t3.small-v2", "m5l-400-v2", "t3.small-v2", "t3.small-v2", true, true, false)
	helper("t3.small-v2", "t3.small-v2", "m5l-400-v2", "t3.small-v2", true, true, false)
	helper("m5l-400-v2", "t3.small-v2", "t3.small-v2", "t3.small-v2", true, true, false)

	helper("", "t3.small-v2", "", "t3.small-v2", false, true, false)
	helper("", "t3.small-v2", "t3.small-v2", "t3.small-v2", true, false, false)
	helper("t3.small-v2", "", "t3.small-v2", "t3.small-v2", false, false, false)
}

func TestIsKafkaSizeAllChange(t *testing.T) {
	helper := func(brokerSize, zookeeperSize, expectedNewSize string, dedicatedZookeeper, expectedIsAll bool) {
		newSize, isAllChange := isKafkaSizeAllChange(brokerSize, zookeeperSize, dedicatedZookeeper)
		if isAllChange != expectedIsAll {
			t.Fatalf("changeAll should be %t when using brokerSize: %s, zookeeperSize: %s, dedicatedZookeeper: %t", expectedIsAll, brokerSize, zookeeperSize, dedicatedZookeeper)
		}
		if isAllChange && newSize != expectedNewSize {
			t.Fatalf("newSize should be %s when using brokerSize: %s, zookeeperSize: %s, dedicatedZookeeper: %t c", expectedNewSize, brokerSize, zookeeperSize, dedicatedZookeeper)
		}
	}
	helper("t3.small-v2", "t3.small-v2", "t3.small-v2", true, true)
	helper("t3.small-v2", "t3.small-v2", "t3.small-v2", false, true)
	helper("t3.small-v2", "", "t3.small-v2", true, false)
	helper("t3.small-v2", "m5l-400-v2", "t3.small-v2", true, false)
}

func TestGetSingleChangedElasticsearchSizeAndPurpose(t *testing.T) {
	helper := func(kibanaSize, masterSize, dataSize, expectedNewSize string, kibana, dedicatedMaster, expectErr bool, expectedNodePurpose NodePurpose) {
		newSize, nodePurpose, err := getSingleChangedElasticsearchSizeAndPurpose(kibanaSize, masterSize, dataSize, kibana, dedicatedMaster)
		if expectErr {
			if err == nil {
				t.Fatalf("expect error when using kibanaSize: %s, masterSize: %s, dataSize: %s, kibana: %t, dedicatedMaster: %t", kibanaSize, masterSize, dataSize, kibana, dedicatedMaster)
			} else {
				return
			}
		}
		if err != nil {
			t.Fatalf("got unexpected error: %s when using kibanaSize: %s, masterSize: %s, dataSize: %s, kibana: %t, dedicatedMaster: %t", err.Error(), kibanaSize, masterSize, dataSize, kibana, dedicatedMaster)
		}
		if newSize != expectedNewSize {
			t.Fatalf("newSize should be %s when using kibanaSize: %s, masterSize: %s, dataSize: %s, kibana: %t, dedicatedMaster:%t", expectedNewSize, kibanaSize, masterSize, dataSize, kibana, dedicatedMaster)
		}
		if nodePurpose.String() != expectedNodePurpose.String() {
			t.Fatalf("nodePurpose should be %s when using kibanaSize: %s, masterSize: %s, dataSize: %s, kibana: %t, dedicatedMaster:%t", expectedNodePurpose, kibanaSize, masterSize, dataSize, kibana, dedicatedMaster)
		}
	}
	helper("t3.small-v2", "t3.small-v2", "t3.small-v2", "t3.small-v2", true, true, true, ELASTICSEARCH_KIBANA)
	helper("", "", "t3.small-v2", "t3.small-v2", true, false, true, ELASTICSEARCH_KIBANA)
	helper("t3.small-v2", "", "", "t3.small-v2", false, false, true, ELASTICSEARCH_KIBANA)
	helper("t3.small-v2", "", "t3.small-v2", "t3.small-v2", false, false, true, ELASTICSEARCH_KIBANA)
	helper("t3.small-v2", "", "t3.small-v2", "t3.small-v2", false, false, true, ELASTICSEARCH_KIBANA)

	helper("t3.small-v2", "", "", "t3.small-v2", true, false, false, ELASTICSEARCH_KIBANA)
	helper("t3.small-v2", "", "", "t3.small-v2", true, true, false, ELASTICSEARCH_KIBANA)
	helper("", "t3.small-v2", "", "t3.small-v2", true, true, false, ELASTICSEARCH_MASTER)
	helper("", "t3.small-v2", "", "t3.small-v2", true, false, false, ELASTICSEARCH_MASTER)
	helper("", "", "t3.small-v2", "t3.small-v2", true, true, false, ELASTICSEARCH_DATA_AND_INGEST)
}

func TestGetSingleChangedOpenSearchSizeAndPurpose(t *testing.T) {
	helper := func(openSearchDashboardsSize, masterSize, dataSize, expectedNewSize string, openSearchDashboards, dedicatedMaster, expectErr bool, expectedNodePurpose NodePurpose) {
		newSize, nodePurpose, err := getSingleChangedOpenSearchSizeAndPurpose(openSearchDashboardsSize, masterSize, dataSize, openSearchDashboards, dedicatedMaster)
		if expectErr {
			if err == nil {
				t.Fatalf("expect error when using openSearchDashboardsSize: %s, masterSize: %s, dataSize: %s, openSearchDashboards: %t, dedicatedMaster: %t", openSearchDashboardsSize, masterSize, dataSize, openSearchDashboards, dedicatedMaster)
			} else {
				return
			}
		}
		if err != nil {
			t.Fatalf("got unexpected error: %s when using openSearchDashboardsSize: %s, masterSize: %s, dataSize: %s, openSearchDashboards: %t, dedicatedMaster: %t", err.Error(), openSearchDashboardsSize, masterSize, dataSize, openSearchDashboards, dedicatedMaster)
		}
		if newSize != expectedNewSize {
			t.Fatalf("newSize should be %s when using openSearchDashboardsSize: %s, masterSize: %s, dataSize: %s, openSearchDashboards: %t, dedicatedMaster:%t", expectedNewSize, openSearchDashboardsSize, masterSize, dataSize, openSearchDashboards, dedicatedMaster)
		}
		if nodePurpose.String() != expectedNodePurpose.String() {
			t.Fatalf("nodePurpose should be %s when using openSearchDashboardsSize: %s, masterSize: %s, dataSize: %s, openSearchDashboards: %t, dedicatedMaster:%t", expectedNodePurpose, openSearchDashboardsSize, masterSize, dataSize, openSearchDashboards, dedicatedMaster)
		}
	}
	helper("t3.small-v2", "t3.small-v2", "t3.small-v2", "t3.small-v2", true, true, true, OPENSEARCH_DASHBOARDS)
	helper("", "", "t3.small-v2", "t3.small-v2", true, false, true, OPENSEARCH_DASHBOARDS)
	helper("t3.small-v2", "", "", "t3.small-v2", false, false, true, OPENSEARCH_DASHBOARDS)
	helper("t3.small-v2", "", "t3.small-v2", "t3.small-v2", false, false, true, OPENSEARCH_DASHBOARDS)
	helper("t3.small-v2", "", "t3.small-v2", "t3.small-v2", false, false, true, OPENSEARCH_DASHBOARDS)

	helper("t3.small-v2", "", "", "t3.small-v2", true, false, false, OPENSEARCH_DASHBOARDS)
	helper("t3.small-v2", "", "", "t3.small-v2", true, true, false, OPENSEARCH_DASHBOARDS)
	helper("", "t3.small-v2", "", "t3.small-v2", true, true, false, OPENSEARCH_MASTER)
	helper("", "t3.small-v2", "", "t3.small-v2", true, false, false, OPENSEARCH_MASTER)
	helper("", "", "t3.small-v2", "t3.small-v2", true, true, false, OPENSEARCH_DATA_AND_INGEST)
}

func TestGetSingleChangedKafkaSizeAndPurpose(t *testing.T) {
	helper := func(brokerSize, zookeeperSize, expectedNewSize string, dedicatedZookeeper, expectErr bool, expectedNodePurpose NodePurpose) {
		newSize, nodePurpose, err := getSingleChangedKafkaSizeAndPurpose(brokerSize, zookeeperSize, dedicatedZookeeper)
		if expectErr {
			if err == nil {
				t.Fatalf("expect error when using brokerSize: %s, zookeeperSize: %s, dedicatedZookeeper: %t", brokerSize, zookeeperSize, dedicatedZookeeper)
			} else {
				return
			}
		}
		if err != nil {
			t.Fatalf("got unexpected error: %s when using brokerSize: %s, zookeeperSize: %s, dedicatedZookeeper: %t", err.Error(), brokerSize, zookeeperSize, dedicatedZookeeper)
		}
		if newSize != expectedNewSize {
			t.Fatalf("newSize should be %s when using brokerSize: %s, zookeeperSize: %s, dedicatedZookeeper: %t", expectedNewSize, brokerSize, zookeeperSize, dedicatedZookeeper)
		}
		if nodePurpose.String() != expectedNodePurpose.String() {
			t.Fatalf("nodePurpose should be %s when using brokerSize: %s, zookeeperSize: %s, dedicatedZookeeper: %t", expectedNodePurpose, brokerSize, zookeeperSize, dedicatedZookeeper)
		}
	}

	helper("t3.small-v2", "t3.small-v2", "t3.small-v2", false, true, KAFKA_BROKER)
	helper("t3.small-v2", "t3.small-v2", "t3.small-v2", true, true, KAFKA_BROKER)
	helper("t3.small-v2", "", "t3.small-v2", false, false, KAFKA_BROKER)
	helper("t3.small-v2", "", "t3.small-v2", true, false, KAFKA_BROKER)
	helper("", "t3.small-v2", "t3.small-v2", true, false, KAFKA_DEDICATED_ZOOKEEPER)
}

func TestGetBundleOptionKey(t *testing.T) {
	helper := func(bundleIndex int, option, expect string) {
		if getBundleOptionKey(bundleIndex, option) != expect {
			t.Fatalf("With parameter %d, %s, should return %s", bundleIndex, option, expect)
		}
	}
	helper(0, "test", "bundle.0.options.test")
	helper(2, "kibana_node_size", "bundle.2.options.kibana_node_size")
	helper(-1, "kibana_node_size", "bundle.-1.options.kibana_node_size")
}

func TestGetNodeSize(t *testing.T) {
	helper := func(data resourceDataInterface, bundles []Bundle, expectedErrMsg, expectedSize string) {
		size, err := getNodeSize(data, bundles)
		if len(expectedErrMsg) > 0 {
			if err == nil || err.Error() != expectedErrMsg {
				t.Fatalf("Expect error %s but got %s", expectedErrMsg, err)
			}
		} else {
			if err != nil {
				t.Fatalf("Expect error to be nil but got %s", err)
			}
			if size != expectedSize {
				t.Fatalf("Expect size %s but got %s", expectedSize, size)
			}
		}
	}
	data := MockResourceData{
		map[string]MockChange{},
	}
	bundles := []Bundle{
		{
			Bundle: "ELASTICSEARCH",
			Options: &BundleOptions{
				DedicatedMasterNodes: nil,
				MasterNodeSize:       "",
				KibanaNodeSize:       "",
				DataNodeSize:         "",
			},
		},
	}
	helper(data, bundles, "[ERROR] 'master_node_size' is required in the bundle option.", "")

	bundles = []Bundle{
		{
			Bundle:  "OPENSEARCH",
			Options: &BundleOptions{},
		},
	}
	helper(data, bundles, "[ERROR] 'master_node_size' is required in the bundle option.", "")

	bundles = []Bundle{
		{
			Bundle: "CASSANDRA",
			Options: &BundleOptions{
				DedicatedMasterNodes: nil,
				MasterNodeSize:       "",
				KibanaNodeSize:       "",
				DataNodeSize:         "",
			},
		},
	}
	data.changes["node_size"] = MockChange{
		before: "",
		after:  "t3.small",
	}
	helper(&data, bundles, "", "t3.small")

	bundles = []Bundle{
		{
			Bundle: "Kafka",
			Options: &BundleOptions{
				DedicatedMasterNodes: nil,
				MasterNodeSize:       "",
				KibanaNodeSize:       "",
				DataNodeSize:         "",
			},
		},
	}
	data.changes["node_size"] = MockChange{
		before: "",
		after:  "t3.small",
	}
	helper(&data, bundles, "", "t3.small")

	dedicatedMaster := true
	bundles = []Bundle{
		{
			Bundle: "ELASTICSEARCH",
			Options: &BundleOptions{
				DedicatedMasterNodes: &dedicatedMaster,
				MasterNodeSize:       "t3.small",
				KibanaNodeSize:       "",
				DataNodeSize:         "",
			},
		},
	}
	helper(&data, bundles, "[ERROR] dedicated master is enabled, 'data_node_size' is required in the bundle option.", "")

	bundles = []Bundle{
		{
			Bundle: "OPENSEARCH",
			Options: &BundleOptions{
				DedicatedMasterNodes: &dedicatedMaster,
				MasterNodeSize:       "t3.small",
			},
		},
	}
	helper(&data, bundles, "[ERROR] dedicated master is enabled, 'data_node_size' is required in the bundle option.", "")

	bundles = []Bundle{
		{
			Bundle: "ELASTICSEARCH",
			Options: &BundleOptions{
				DedicatedMasterNodes: &dedicatedMaster,
				MasterNodeSize:       "t3.small",
				KibanaNodeSize:       "",
				DataNodeSize:         "t3.small-v2",
			},
		},
	}
	helper(&data, bundles, "", "t3.small-v2")

	bundles = []Bundle{
		{
			Bundle: "OPENSEARCH",
			Options: &BundleOptions{
				DedicatedMasterNodes: &dedicatedMaster,
				MasterNodeSize:       "t3.small",
				DataNodeSize:         "t3.small-v2",
			},
		},
	}
	helper(&data, bundles, "", "t3.small-v2")
	dedicatedMaster = false
	bundles = []Bundle{
		{
			Bundle: "ELASTICSEARCH",
			Options: &BundleOptions{
				DedicatedMasterNodes: &dedicatedMaster,
				MasterNodeSize:       "t3.small",
				KibanaNodeSize:       "",
				DataNodeSize:         "t3.small-v2",
			},
		},
	}
	helper(&data, bundles, "[ERROR] When 'dedicated_master_nodes' is not true , data_node_size can be either null or equal to master_node_size.", "")
	bundles = []Bundle{
		{
			Bundle: "OPENSEARCH",
			Options: &BundleOptions{
				DedicatedMasterNodes: &dedicatedMaster,
				MasterNodeSize:       "t3.small",
				DataNodeSize:         "t3.small-v2",
			},
		},
	}
	helper(&data, bundles, "[ERROR] When 'dedicated_master_nodes' is not true , data_node_size can be either null or equal to master_node_size.", "")
	bundles = []Bundle{
		{
			Bundle: "ELASTICSEARCH",
			Options: &BundleOptions{
				DedicatedMasterNodes: &dedicatedMaster,
				MasterNodeSize:       "t3.small",
				KibanaNodeSize:       "",
				DataNodeSize:         "t3.small",
			},
		},
	}
	helper(&data, bundles, "", "t3.small")
	bundles = []Bundle{
		{
			Bundle: "OPENSEARCH",
			Options: &BundleOptions{
				DedicatedMasterNodes: &dedicatedMaster,
				MasterNodeSize:       "t3.small",
				DataNodeSize:         "t3.small",
			},
		},
	}
	helper(&data, bundles, "", "t3.small")
}

func TestGetBundleIndex(t *testing.T) {
	if index, err := getBundleIndex("ELASTICSEARCH", []Bundle{
		{Bundle: "LOG_SHIPPER"},
		{Bundle: "ELASTICSEARCH"},
	}); err != nil || index != 1 {
		t.Fatalf("Expect no error and 1, got %v and %v", err, index)
	}

	if index, err := getBundleIndex("OPENSEARCH", []Bundle{
		{Bundle: "LOG_SHIPPER"},
		{Bundle: "OPENSEARCH"},
	}); err != nil || index != 1 {
		t.Fatalf("Expect no error and 1, got %v and %v", err, index)
	}

	if index, err := getBundleIndex("ELASTICSEARCH", []Bundle{
		{Bundle: "ELASTICSEARCH"},
	}); err != nil || index != 0 {
		t.Fatalf("Expect no error and 0, got %v and %v", err, index)
	}

	if index, err := getBundleIndex("OPENSEARCH", []Bundle{
		{Bundle: "OPENSEARCH"},
	}); err != nil || index != 0 {
		t.Fatalf("Expect no error and 0, got %v and %v", err, index)
	}
}

func TestGetNewSizeOrEmpty(t *testing.T) {
	data := schema.ResourceData{}
	if size := getNewSizeOrEmpty(&data, "node_size"); size != "" {
		t.Fatalf("Expect empty string but got %v", size)
	}
}

func TestHasElasticsearchSizeChanges(t *testing.T) {
	data := schema.ResourceData{}
	if hasChange := hasElasticsearchSizeChanges(0, &data); hasChange {
		t.Fatalf("Expect false but got %v", hasChange)
	}
}

func TestHasOpenSearchSizeChanges(t *testing.T) {
	data := schema.ResourceData{}
	if hasChange := hasOpenSearchSizeChanges(0, &data); hasChange {
		t.Fatalf("Expect false but got %v", hasChange)
	}
}

func TestHasKafkaSizeChanges(t *testing.T) {
	data := schema.ResourceData{}
	if hasChange := hasKafkaSizeChanges(0, &data); hasChange {
		t.Fatalf("Expect false but got %v", hasChange)
	}
}

func TestHasCassandraSizeChanges(t *testing.T) {
	data := schema.ResourceData{}
	if hasChange := hasCassandraSizeChanges(&data); hasChange {
		t.Fatalf("Expect false but got %v", hasChange)
	}
}

func TestDoClusterResizeDefault(t *testing.T) {
	err := doClusterResize(MockApiClient{
		cluster: Cluster{
			ID:         "CADENCE",
			BundleType: "CADENCE",
		},
	}, "mock", MockResourceData{}, []Bundle{
		{Bundle: "CADENCE"},
	})
	if err == nil || !strings.Contains(err.Error(), "CDC resize does not support:") {
		t.Fatalf("Expect err with  'CDC resize does not support:' but got %v", err)
	}
}

func TestDoClusterResizeES(t *testing.T) {
	client := MockApiClient{
		cluster: Cluster{
			ID:           "mock",
			BundleType:   "ELASTICSEARCH",
			BundleOption: &BundleOptions{},
			DataCentres: []DataCentre{
				{ID: "test"},
			},
		},
	}
	data := MockResourceData{
		changes: map[string]MockChange{"bundle.0.options.master_node_size": {before: "t3.small", after: "t3.small-v2"}},
	}
	bundles := []Bundle{
		{Bundle: "ELASTICSEARCH"},
	}
	err := doClusterResize(client, "mock", data, bundles)
	if err != nil {
		t.Fatalf("Expect nil err but got %v", err)
	}
	delete(data.changes, "bundle.0.options.master_node_size")
	err = doClusterResize(client, "mock", data, bundles)
	if err != nil {
		t.Fatalf("Expect nil err but got %v", err)
	}
}

func TestDoClusterResizeOpenSearch(t *testing.T) {
	client := MockApiClient{
		cluster: Cluster{
			ID:           "mock",
			BundleType:   "OPENSEARCH",
			BundleOption: &BundleOptions{},
			DataCentres: []DataCentre{
				{ID: "test"},
			},
		},
	}
	data := MockResourceData{
		changes: map[string]MockChange{"bundle.0.options.master_node_size": {before: "t3.small", after: "t3.small-v2"}},
	}
	bundles := []Bundle{
		{Bundle: "OPENSEARCH"},
	}
	err := doClusterResize(client, "mock", data, bundles)
	if err != nil {
		t.Fatalf("Expect nil err but got %v", err)
	}
	delete(data.changes, "bundle.0.options.master_node_size")
	err = doClusterResize(client, "mock", data, bundles)
	if err != nil {
		t.Fatalf("Expect nil err but got %v", err)
	}
}

func TestDoClusterResizeKA(t *testing.T) {
	client := MockApiClient{
		cluster: Cluster{
			ID:           "mock",
			BundleType:   "KAFKA",
			BundleOption: &BundleOptions{},
			DataCentres: []DataCentre{
				{ID: "test"},
			},
		},
	}
	data := MockResourceData{
		changes: map[string]MockChange{"node_size": {before: "t3.small", after: "t3.small-v2"}},
	}
	bundles := []Bundle{
		{Bundle: "KAFKA"},
	}
	err := doClusterResize(client, "mock", data, bundles)
	if err != nil {
		t.Fatalf("Expect nil err but got %v", err)
	}
	delete(data.changes, "node_size")
	err = doClusterResize(client, "mock", data, bundles)
	if err != nil {
		t.Fatalf("Expect nil err but got %v", err)
	}
}

func TestDoClusterResizeCA(t *testing.T) {
	client := MockApiClient{
		cluster: Cluster{
			ID:           "mock",
			BundleType:   "APACHE_CASSANDRA",
			BundleOption: &BundleOptions{},
			DataCentres: []DataCentre{
				{ID: "test"},
			},
		},
	}
	data := MockResourceData{
		changes: map[string]MockChange{"node_size": {before: "t3.small", after: "t3.small-v2"}},
	}
	bundles := []Bundle{
		{Bundle: "APACHE_CASSANDRA"},
	}
	err := doClusterResize(client, "mock", data, bundles)
	if err == nil || err.Error() != "[Error] Cannot resize nodes from t3.small to t3.small-v2" {
		t.Fatalf("Expect err to be '[Error] Cannot resize nodes from t3.small to t3.small-v2' but got %v", err)
	}
	delete(data.changes, "node_size")
	err = doClusterResize(client, "mock", data, bundles)
	if err != nil {
		t.Fatalf("Expect nil err but got %v", err)
	}
}

func TestDoClusterResizeRedis(t *testing.T) {
	client := MockApiClient{
		cluster: Cluster{
			ID:           "mock",
			BundleType:   "REDIS",
			BundleOption: &BundleOptions{},
			DataCentres: []DataCentre{
				{ID: "test"},
			},
		},
	}
	data := MockResourceData{
		changes: map[string]MockChange{"node_size": {before: "t3.small", after: "t3.small-v2"}},
	}
	bundles := []Bundle{
		{Bundle: "REDIS"},
	}
	err := doClusterResize(client, "mock", data, bundles)
	if err != nil {
		t.Fatalf("Expect nil err but got %v", err)
	}
	delete(data.changes, "node_size")
	err = doClusterResize(client, "mock", data, bundles)
	if err != nil {
		t.Fatalf("Expect nil err but got %v", err)
	}
}

func TestCreateVpcPeeringRequest(t *testing.T) {
	resourceSchema := map[string]*schema.Schema{
		"peer_vpc_id": {
			Type: schema.TypeString,
		},
		"peer_account_id": {
			Type: schema.TypeString,
		},
		"peer_subnets": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"peer_region": {
			Type: schema.TypeString,
		},
	}

	peerSubnets := schema.NewSet(schema.HashString, []interface{}{"10.20.0.0/16", "10.21.0.0/16"})
	resourceDataMap := map[string]interface{}{
		"peer_vpc_id":     "vpc-12345678",
		"peer_account_id": "494111121110",
		"peer_subnets":    peerSubnets.List(),
		"peer_region":     "",
	}
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	if _, err := createVpcPeeringRequest(resourceLocalData); err != nil {
		t.Fatalf("Expected nil error but got %v", err)
	}
}

func TestGCPCreateVpcPeeringRequest(t *testing.T) {
	resourceSchema := map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
		},
		"peer_vpc_network_name": {
			Type: schema.TypeString,
		},
		"peer_project_id": {
			Type: schema.TypeString,
		},
		"peer_subnets": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	peerSubnets := schema.NewSet(schema.HashString, []interface{}{"10.20.0.0/16", "10.21.0.0/16"})
	resourceDataMap := map[string]interface{}{
		"name":                  "Kaka",
		"peer_vpc_network_name": "my-vpc1",
		"peer_project_id":       "instaclustr-dev",
		"peer_subnets":          peerSubnets.List(),
	}
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	if _, err := GCPcreateVpcPeeringRequest(resourceLocalData); err != nil {
		t.Fatalf("Expected nil error but got %v", err)
	}
}

func TestGCPReadVpcPeeringRequest(t *testing.T) {
	resourceSchema := map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
		},
		"peer_vpc_network_name": {
			Type: schema.TypeString,
		},
		"peer_project_id": {
			Type: schema.TypeString,
		},
		"peer_subnets": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	peerSubnets := schema.NewSet(schema.HashString, []interface{}{"10.20.0.0/16", "10.21.0.0/16"})
	resourceDataMap := map[string]interface{}{
		"name":                  "Kaka",
		"peer_vpc_network_name": "my-vpc1",
		"peer_project_id":       "instaclustr-dev",
		"peer_subnets":          peerSubnets.List(),
	}
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	if _, err := GCPcreateVpcPeeringRequest(resourceLocalData); err != nil {
		t.Fatalf("Expected nil error but got %v", err)
	}
}
func TestGCPReadVpcPeeringRequest2(t *testing.T) {
	resourceSchema := map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
		},
		"peer_vpc_network_name": {
			Type: schema.TypeString,
		},
		"peer_project_id": {
			Type: schema.TypeString,
		},
		"peer_subnets": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	peerSubnets := schema.NewSet(schema.HashString, []interface{}{"10.20.0.0/16", "10.21.0.0/16"})
	resourceDataMap := map[string]interface{}{
		"name":                  "Kaka",
		"peer_vpc_network_name": "my-vpc1",
		"peer_project_id":       "instaclustr-dev",
		"peer_subnets":          peerSubnets.List(),
	}
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	if _, err := GCPcreateVpcPeeringDelete(resourceLocalData); err != nil {
		t.Fatalf("Expected nil error but got %v", err)
	}
}

func TestCreateVpcPeeringRequestLegacy(t *testing.T) {
	resourceSchema := map[string]*schema.Schema{
		"peer_vpc_id": {
			Type: schema.TypeString,
		},
		"peer_account_id": {
			Type: schema.TypeString,
		},
		"peer_subnet": {
			Type: schema.TypeString,
		},
		"peer_region": {
			Type: schema.TypeString,
		},
	}

	resourceDataMap := map[string]interface{}{
		"peer_vpc_id":     "vpc-12345678",
		"peer_account_id": "494111121110",
		"peer_subnet":     "10.20.0.0/16",
		"peer_region":     "",
	}
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	if _, err := createVpcPeeringRequest(resourceLocalData); err != nil {
		t.Fatalf("Expected nil error but got %v", err)
	}
}

func TestDeleteAttributesConflict(t *testing.T) {
	clusterSchema := map[string]*schema.Schema{
		"attributeA": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"attributeB": {
			Type:          schema.TypeString,
			Required:      true,
			ConflictsWith: []string{"data_centres"},
			ForceNew:      true,
		},

		"attributeC": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"data_centre"},
			ForceNew:      true,
		},
	}

	resourceDataMap := map[string]interface{}{
		"attributeA": "A",
		"attributeB": "B",
		"attributeC": "C",
	}

	d := schema.TestResourceDataRaw(t, clusterSchema, resourceDataMap)

	if err := deleteAttributesConflict(clusterSchema, d, "data_centres"); err != nil {
		t.Fatalf("Unexpected error occured during deletion %s", err)
	}

	checkAttributeValue := func(attribute string, expected interface{}) {
		if value, _ := d.GetOk(attribute); value != expected {
			t.Fatalf("%s not modified properly", attribute)
		}
	}

	checkAttributeValue("attributeA", "A")
	checkAttributeValue("attributeB", schema.TypeString.Zero())
	checkAttributeValue("attributeC", "C")
}

type MockApiClient struct {
	cluster Cluster
	err     error
}

func (m MockApiClient) ReadCluster(clusterID string) (*Cluster, error) {
	return &m.cluster, m.err
}

func (m MockApiClient) ResizeCluster(clusterID string, cdcID string, newNodeSize string, nodePurpose *NodePurpose) error {
	return m.err
}

type MockChange struct {
	before interface{}
	after  interface{}
}

type MockResourceData struct {
	changes map[string]MockChange
}

func (m MockResourceData) HasChange(key string) bool {
	_, ok := m.changes[key]
	return ok
}

func (m MockResourceData) GetChange(key string) (interface{}, interface{}) {
	return m.changes[key].before, m.changes[key].after
}

func (m MockResourceData) GetOk(key string) (interface{}, bool) {
	change, ok := m.changes[key]
	if ok {
		return change.after, ok
	} else {
		return nil, ok
	}
}

func (m MockResourceData) Get(key string) interface{} {
	change, ok := m.changes[key]
	if ok {
		return change.after
	} else {
		return nil
	}
}
