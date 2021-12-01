package instaclustr

type MockKafkaAclApiClient struct {
	acls	[]KafkaAcl
	cluster	Cluster
        err     error
}       
        
func (m MockKafkaAclApiClient) ReadKafkaAcls(clusterID string, data []byte) ([]KafkaAcl, error) {
        return m.acls, m.err
}

func (m MockKafkaAclApiClient) DeleteKafkaAcl(clusterID string, data []byte) error {
        return m.err
}

func (m MockKafkaAclApiClient) CreateKafkaAcl(clusterID string, data []byte) error {
        return m.err
}

func (m MockKafkaAclApiClient) ReadCluster(clusterID string) (*Cluster, error) {
	return &m.cluster, m.err
}

type MockKafkaAclResourceData struct {
	data map[string]interface{}
}

func (m MockKafkaAclResourceData) Get(key string) interface{} {
	return m.data[key]
}

func (m MockKafkaAclResourceData) Set(key string, v interface{}) error {
	m.data[key] = v
	return nil
}

func (m MockKafkaAclResourceData) SetId(value string) {
	m.data["id"] = value
	return
}

func (m MockKafkaAclResourceData) Id() string {
	return m.data["id"].(string)
}
