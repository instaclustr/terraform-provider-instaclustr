package instaclustr

type NodePurpose string

const (
	ELASTICSEARCH_DATA_AND_INGEST NodePurpose = "ELASTICSEARCH_DATA_AND_INGEST"
	ELASTICSEARCH_MASTER                      = "ELASTICSEARCH_MASTER"
	ELASTICSEARCH_KIBANA                      = "ELASTICSEARCH_KIBANA"
	ELASTICSEARCH_COORDINATOR                 = "ELASTICSEARCH_COORDINATOR"
	KAFKA_BROKER                              = "KAFKA_BROKER"
	KAFKA_DEDICATED_ZOOKEEPER                 = "KAFKA_DEDICATED_ZOOKEEPER"
)

func (p NodePurpose) String() string {
	return string(p)
}
