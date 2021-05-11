package instaclustr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateBundleUserUpdateRequest(t *testing.T) {
	testUsername := "hail"
	testPassword := "reallySecure123"
	testBundleRequest := createBundleUserUpdateRequest(testUsername, testPassword)

	expectedOutput := []byte(fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", testUsername, testPassword))

	if !reflect.DeepEqual(testBundleRequest, expectedOutput){
		t.Fatalf("Incorrect request returned.\nExpected:%s\nActual:%s", expectedOutput, testBundleRequest)
	}
}

func TestGetBundleConfig(t *testing.T) {
	var testBundles []Bundle
	var testBundleConfig BundleConfig

	testBundles = append(testBundles, Bundle{Bundle: "KAFKA"})
	testBundleConfig = getBundleConfig(testBundles)
	expectedOutput := BundleConfig{
		IsKafkaCluster: true,
		HasRestProxy: false,
		HasSchemaRegistry: false}

	if testBundleConfig != expectedOutput {
		t.Fatalf("Incorrect Bundle Config returned.\nExpected: %+v\nActual: %+v", expectedOutput, testBundleConfig)
	}

	testBundles = append(testBundles, Bundle{Bundle: "KAFKA_REST_PROXY"})
	testBundleConfig = getBundleConfig(testBundles)
	expectedOutput = BundleConfig{
		IsKafkaCluster: true,
		HasRestProxy: true,
		HasSchemaRegistry: false}

	if testBundleConfig != expectedOutput {
		t.Fatalf("Incorrect Bundle Config returned.\nExpected: %+v\nActual: %+v", expectedOutput, testBundleConfig)
	}

	testBundles = append(testBundles, Bundle{Bundle: "KAFKA_SCHEMA_REGISTRY"})
	testBundleConfig = getBundleConfig(testBundles)
	expectedOutput = BundleConfig{
		IsKafkaCluster: true,
		HasRestProxy: true,
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

func TestFormatCreateErrMsgDuplicationOne(t *testing.T) {
	testError := fmt.Errorf("test error")
	formattedError := formatCreateErrMsg(testError)
	expectedError := fmt.Sprintf("[Error] Error creating cluster: %s", testError.Error())

	if formattedError.Error() != expectedError {
		t.Fatalf("Incorrectly formatted error.\nExpected: %s\nActual: %s", expectedError, formattedError)
	}
}

func TestFormatCreateErrMsgDuplicationTwo(t *testing.T) {
	testError := fmt.Errorf("test error")
	formattedError := formatCreateErrMsg(testError)
	expectedError := fmt.Sprintf("[Error] Error creating cluster: %s", testError.Error())

	if formattedError.Error() != expectedError {
		t.Fatalf("Incorrectly formatted error.\nExpected: %s\nActual: %s", expectedError, formattedError)
	}
}

func TestFormatCreateErrMsgDuplicationThree(t *testing.T) {
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

	bundles = []Bundle{{Bundle: "APACHE_CASSANDRA"}}
	isRackAllocationRequired = checkIfBundleRequiresRackAllocation(bundles)

	if isRackAllocationRequired == false {
		t.Fatalf("Incorrect check performed for APACHE_CASSANDRA.\nExpected: %v\nActual: %v\n", true, false)
	}
}
