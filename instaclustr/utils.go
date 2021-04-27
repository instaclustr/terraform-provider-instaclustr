package instaclustr

import (
	"encoding/json"
	"strings"
)

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap)
	return
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getContactPointIPs(cluster *Cluster) (publicContactPointList []string, privateContactPointList []string){
	azList := make([]string, 0)

	for _, dataCentre := range cluster.DataCentres {
		for _, node := range dataCentre.Nodes {
			if !stringInSlice(node.Rack, azList) {
				if !strings.HasPrefix(node.Size, "zk-") {
					azList = appendIfMissing(azList, node.Rack)
					privateContactPointList = appendIfMissing(privateContactPointList, node.PrivateAddress)
					publicContactPointList = appendIfMissing(publicContactPointList, node.PublicAddress)
				}
			}
		}
	}
	return publicContactPointList, privateContactPointList
}