package instaclustr

type FirewallRule struct {
	Network         string     `json:"network,omitempty"`
	SecurityGroupId string     `json:"securityGroupId,omitempty"`
	Rules           []RuleType `json:"rules"`
}

type RuleType struct {
	Type string `json:"type"`
}

type Bundle struct {
	Bundle  string        `json:"bundle" mapstructure:"bundle"`
	Version string        `json:"version" mapstructure:"version"`
	Options BundleOptions `json:"options,omitempty" mapstructure:"options"`
}

type BundleOptions struct {
	AuthnAuthz                    string `json:"authnAuthz,omitempty" mapstructure:"auth_n_authz"`
	ClientEncryption              string `json:"clientEncryption,omitempty" mapstructure:"client_encryption"`
	DedicatedMasterNodes          string `json:"dedicatedMasterNodes,omitempty" mapstructure:"dedicated_master_nodes"`
	MasterNodeSize                string `json:"masterNodeSize,omitempty" mapstructure:"master_node_size"`
	SecurityPlugin                string `json:"securityPlugin,omitempty" mapstructure:"security_plugin"`
	UsePrivateBroadcastRpcAddress string `json:"usePrivateBroadcastRPCAddress,omitempty" mapstructure:"use_private_broadcast_rpc_address"`
	LuceneEnabled                 string `json:"luceneEnabled,omitempty" mapstructure:"lucene_enabled"`
	ContinuousBackupEnabled       string `json:"continuousBackupEnabled,omitempty" mapstructure:"continuous_backup_enabled"`
	NumberPartitions              string `json:"numberPartitions,omitempty" mapstructure:"number_partitions"`
	AutoCreateTopics              string `json:"autoCreateTopics,omitempty" mapstructure:"auto_create_topics"`
	DeleteTopics                  string `json:"deleteTopics,omitempty" mapstructure:"delete_topics"`
	PasswordAuthentication        string `json:"passwordAuthentication,omitempty" mapstructure:"password_authentication"`
	TargetKafkaClusterId          string `json:"targetKafkaClusterId,omitempty" mapstructure:"target_kafka_cluster_id"`
	VPCType                       string `json:"vpcType,omitempty" mapstructure:"vpc_type"`
	AWSAccessKeyId                string `json:"aws.access.key.id,omitempty" mapstructure:"aws_access_key"`
	AWSSecretKey                  string `json:"aws.secret.access.key,omitempty" mapstructure:"aws_secret_key"`
	S3BucketName                  string `json:"s3.bucket.name,omitempty" mapstructure:"s3_bucket_name"`
	AzureStorageAccountName       string `json:"azure.storage.account.name,omitempty" mapstructure:"azure_storage_account_name"`
	AzureStorageAccountKey        string `json:"azure.storage.account.key,omitempty" mapstructure:"azure_storage_account_key"`
	AzureStorageContainerName     string `json:"azure.storage.container.name,omitempty" mapstructure:"azure_storage_container_name"`
	SslEnabledProtocols           string `json:"ssl.enabled.protocols,omitempty" mapstructure:"ssl_enabled_protocols"`
	SslTruststorePassword         string `json:"ssl.truststore.password,omitempty" mapstructure:"ssl_truststore_password"`
	SslProtocol                   string `json:"ssl.protocol,omitempty" mapstructure:"ssl_protocol"`
	SecurityProtocol              string `json:"security.protocol,omitempty" mapstructure:"security_protocol"`
	SaslMechanism                 string `json:"sasl.mechanism,omitempty" mapstructure:"sasl_mechanism"`
	SaslJaasConfig                string `json:"sasl.jaas.config,omitempty" mapstructure:"sasl_jaas_config"`
	BootstrapServers              string `json:"bootstrap.servers,omitempty" mapstructure:"bootstrap_servers"`
	Truststore                    string `json:"truststore,omitempty" mapstructure:"truststore"`
	RedisMasterNodes              string `json:"masterNodes,omitempty" mapstructure:"master_nodes"`
	RedisReplicaNodes             string `json:"replicaNodes,omitempty" mapstructure:"replica_nodes"`
	DedicatedZookeeper            string `json:"dedicatedZookeeper,omitempty" mapstructure:"dedicated_zookeeper"`
	ZookeeperNodeSize             string `json:"zookeeperNodeSize,omitempty" mapstructure:"zookeeper_node_size"`
	ZookeeperNodeCount            string `json:"zookeeperNodeCount,omitempty" mapstructure:"zookeeper_node_count"`
}

type ClusterProvider struct {
	Name                   *string `json:"name" mapstructure:"name"`
	AccountName            *string `json:"accountName, omitempty" mapstructure:"account_name"`
	CustomVirtualNetworkId *string `json:"customVirtualNetworkId, omitempty" mapstructure:"custom_virtual_network_id"`
	Tags                   map[string]interface{} `json:"tags,omitempty"`
	ResourceGroup          *string `json:"resourceGroup,omitempty" mapstructure:"resource_group"`
	DiskEncryptionKey      *string `json:"diskEncryptionKey,omitempty" mapstructure:"disk_encryption_key"`
}

type RackAllocation struct {
	NumberOfRacks string `json:"numberOfRacks" mapstructure:"number_of_racks"`
	NodesPerRack  string `json:"nodesPerRack" mapstructure:"nodes_per_rack"`
}

type CreateRequest struct {
	ClusterName           string          `json:"clusterName"`
	Bundles               []Bundle        `json:"bundles"`
	Provider              ClusterProvider `json:"provider"`
	SlaTier               string          `json:"slaTier"`
	NodeSize              string          `json:"nodeSize"`
	DataCentre            string          `json:"dataCentre"`
	ClusterNetwork        string          `json:"clusterNetwork"`
	PrivateNetworkCluster string          `json:"privateNetworkCluster"`
	PCICompliantCluster   string          `json:"pciCompliantCluster"`
	RackAllocation        *RackAllocation `json:"rackAllocation,omitempty"`
}

type Cluster struct {
	ID                         string       `json:"id"`
	ClusterName                string       `json:"clusterName"`
	ClusterStatus              string       `json:"clusterStatus"`
	CassandraVersion           string       `json:"cassandraVersion"`
	Username                   string       `json:"username"`
	InstaclustrUserPassword    string       `json:"instaclustrUserPassword"`
	SlaTier                    string       `json:"slaTier"`
	ClusterCertificateDownload string       `json:"clusterCertificateDownload"`
	PciCompliance              string       `json:"pciCompliance"`
	DataCentres                []DataCentre `json:"dataCentres"`
}

type DataCentre struct {
	ID                            string   `json:"id"`
	Name                          string   `json:"name"`
	Provider                      string   `json:"provider"`
	CdcNetwork                    string   `json:"cdcNetwork"`
	Bundles                       []string `json:"bundles"`
	ClientEncryption              bool     `json:"clientEncryption"`
	PasswordAuthentication        bool     `json:"passwordAuthentication"`
	UserAuthorization             bool     `json:"userAuthorization"`
	UsePrivateBroadcastRPCAddress bool     `json:"usePrivateBroadcastRPCAddress"`
	PrivateIPOnly                 bool     `json:"privateIPOnly"`
	Nodes                         []Node   `json:"nodes"`
	NodeCount                     int      `json:"nodeCount"`
	EncryptionKeyId               []string `json:"encryptionKeyId"`
	ResizeTargetNodeSize          string   `json:"resizeTargetNodeSize"`
}

type Node struct {
	ID             string `json:"id"`
	Size           string `json:"size"`
	Rack           string `json:"rack"`
	PublicAddress  string `json:"publicAddress"`
	PrivateAddress string `json:"privateAddress"`
	NodeStatus     string `json:"nodeStatus"`
	SparkMaster    bool   `json:"sparkMaster"`
	SparkJobserver bool   `json:"sparkJobserver"`
	Zeppelin       bool   `json:"zeppelin"`
}

type CreateVPCPeeringRequest struct {
	PeerVpcID     string `json:"peerVpcId"`
	PeerAccountID string `json:"peerAccountId"`
	PeerSubnet    string `json:"peerSubnet"`
	PeerRegion    string `json:"peerRegion,omitempty"`
}

type VPCPeering struct {
	ID                 string           `json:"id"`
	AWSVpcConnectionID string           `json:"aws_vpc_connection_id"`
	ClusterDataCentre  string           `json:"clusterDataCentre"`
	VpcID              string           `json:"vpcId"`
	PeerVpcID          string           `json:"peerVpcId"`
	PeerAccountID      string           `json:"peerAccountId"`
	PeerSubnet         VPCPeeringSubnet `json:"peerSubnet"`
	StatusCode         string           `json:"statusCode"`
	PeerRegion         string           `json:"peerRegion"`
}

type VPCPeeringSubnet struct {
	Network      string `json:"network"`
	PrefixLength string `json:"prefixLength"`
}

type ResizeClusterRequest struct {
	NewNodeSize           string `json:"newNodeSize"`
	ConcurrentResizes     int    `json:"concurrentResizes"`
	NotifySupportContacts string `json:"notifySupportContacts"`
}

type EncryptionKey struct {
	ID    string `json:"id,omitempty"`
	Alias string `json:"alias,omitempty"`
	ARN   string `json:"arn,omitempty"`
	Provider string `json:"provider,omitempty"`
}

type CreateKafkaUserRequest struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	InitialPermissions string `json:"initial-permissions"`
}

type UpdateKafkaUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeleteKafkaUserRequest struct {
	Username string `json:"username"`
}
