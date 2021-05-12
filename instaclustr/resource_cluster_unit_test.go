package instaclustr

import (
	"fmt"
	"testing"
)


func TestFormatCreateErrMsg(t *testing.T) {
	testError := fmt.Errorf("test error")
	formattedError := formatCreateErrMsg(testError)
	expectedError := fmt.Sprintf("[Error] Error creating cluster: %s", testError.Error())

	if formattedError.Error() != expectedError {
		t.Fatalf("Incorrectly formatted error.\nExpected: %s\nActual: %s", expectedError, formattedError)
	}
}

// (Here)
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
