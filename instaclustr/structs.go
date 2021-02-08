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

type OmitEmptyBool struct {
	value bool
}

//func (b *OmitEmptyBool) UnmarshalJSON(data []byte) error {
//
//}

type BundleOptions struct {
	AuthnAuthz                    *bool   `json:"authnAuthz,omitempty" mapstructure:"auth_n_authz,omitempty"`
	ClientEncryption              *bool   `json:"clientEncryption,omitempty" mapstructure:"client_encryption,omitempty"`
	DedicatedMasterNodes          *bool   `json:"dedicatedMasterNodes,omitempty" mapstructure:"dedicated_master_nodes,omitempty"`
	MasterNodeSize                string `json:"masterNodeSize,omitempty" mapstructure:"master_node_size,omitempty"`
	SecurityPlugin                *bool   `json:"securityPlugin,omitempty" mapstructure:"security_plugin,omitempty"`
	UsePrivateBroadcastRpcAddress *bool   `json:"usePrivateBroadcastRPCAddress,omitempty" mapstructure:"use_private_broadcast_rpc_address,omitempty"`
	LuceneEnabled                 *bool   `json:"luceneEnabled,omitempty" mapstructure:"lucene_enabled,omitempty"`
	ContinuousBackupEnabled       *bool   `json:"continuousBackupEnabled,omitempty" mapstructure:"continuous_backup_enabled,omitempty"`
	NumberPartitions              string `json:"numberPartitions,omitempty" mapstructure:"number_partitions,omitempty"`
	AutoCreateTopics              *bool   `json:"autoCreateTopics,omitempty" mapstructure:"auto_create_topics,omitempty"`
	DeleteTopics                  *bool   `json:"deleteTopics,omitempty" mapstructure:"delete_topics,omitempty"`
	PasswordAuthentication        *bool   `json:"passwordAuthentication,omitempty" mapstructure:"password_authentication,omitempty"`
	TargetKafkaClusterId          string `json:"targetKafkaClusterId,omitempty" mapstructure:"target_kafka_cluster_id,omitempty"`
	VPCType                       string `json:"vpcType,omitempty" mapstructure:"vpc_type,omitempty"`
	AWSAccessKeyId                string `json:"aws.access.key.id,omitempty" mapstructure:"aws_access_key,omitempty"`
	AWSSecretKey                  string `json:"aws.secret.access.key,omitempty" mapstructure:"aws_secret_key,omitempty"`
	S3BucketName                  string `json:"s3.bucket.name,omitempty" mapstructure:"s3_bucket_name,omitempty"`
	AzureStorageAccountName       string `json:"azure.storage.account.name,omitempty" mapstructure:"azure_storage_account_name,omitempty"`
	AzureStorageAccountKey        string `json:"azure.storage.account.key,omitempty" mapstructure:"azure_storage_account_key,omitempty"`
	AzureStorageContainerName     string `json:"azure.storage.container.name,omitempty" mapstructure:"azure_storage_container_name,omitempty"`
	SslEnabledProtocols           string `json:"ssl.enabled.protocols,omitempty" mapstructure:"ssl_enabled_protocols,omitempty"`
	SslTruststorePassword         string `json:"ssl.truststore.password,omitempty" mapstructure:"ssl_truststore_password,omitempty"`
	SslProtocol                   string `json:"ssl.protocol,omitempty" mapstructure:"ssl_protocol,omitempty"`
	SecurityProtocol              string `json:"security.protocol,omitempty" mapstructure:"security_protocol,omitempty"`
	SaslMechanism                 string `json:"sasl.mechanism,omitempty" mapstructure:"sasl_mechanism,omitempty"`
	SaslJaasConfig                string `json:"sasl.jaas.config,omitempty" mapstructure:"sasl_jaas_config,omitempty"`
	BootstrapServers              string `json:"bootstrap.servers,omitempty" mapstructure:"bootstrap_servers,omitempty"`
	Truststore                    string `json:"truststore,omitempty" mapstructure:"truststore,omitempty"`
	RedisMasterNodes              string `json:"masterNodes,omitempty" mapstructure:"master_nodes,omitempty"`
	RedisReplicaNodes             string `json:"replicaNodes,omitempty" mapstructure:"replica_nodes,omitempty"`
	DedicatedZookeeper            *bool   `json:"dedicatedZookeeper,omitempty" mapstructure:"dedicated_zookeeper,omitempty"`
	ZookeeperNodeSize             string `json:"zookeeperNodeSize,omitempty" mapstructure:"zookeeper_node_size,omitempty"`
	ZookeeperNodeCount            string `json:"zookeeperNodeCount,omitempty" mapstructure:"zookeeper_node_count,omitempty"`
}

type ClusterProvider struct {
	Name                   *string                `json:"name" mapstructure:"name"`
	AccountName            *string                `json:"accountName, omitempty" mapstructure:"account_name"`
	CustomVirtualNetworkId *string                `json:"customVirtualNetworkId, omitempty" mapstructure:"custom_virtual_network_id"`
	Tags                   map[string]interface{} `json:"tags,omitempty"`
	ResourceGroup          *string                `json:"resourceGroup,omitempty" mapstructure:"resource_group"`
	DiskEncryptionKey      *string                `json:"diskEncryptionKey,omitempty" mapstructure:"disk_encryption_key"`
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

type AddonBundles struct {
	Bundle						string		`json:"bundle"`
	Version 					string		`json:"version"`
}

type Cluster struct {
	ID                         string        `json:"id"`
	ClusterName                string        `json:"clusterName"`
	ClusterStatus              string        `json:"clusterStatus"`
	BundleType                 string        `json:"bundleType"`
	BundleVersion              string        `json:"bundleVersion"`
	AddonBundles			   map[string]interface{}	 `json:"addonBundles"`
	Username                   string        `json:"username"`
	InstaclustrUserPassword    string        `json:"instaclustrUserPassword"`
	SlaTier                    string        `json:"slaTier"`
	ClusterCertificateDownload string        `json:"clusterCertificateDownload"`
	PciCompliance              string        `json:"pciCompliance"`
	BundleOption               BundleOptions `json:"bundleOptions"`
	DataCentres                []DataCentre  `json:"dataCentres"`
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
	ID             string   `json:"id"`
	Size           string   `json:"size"`
	Rack           string   `json:"rack"`
	PublicAddress  []string `json:"publicAddress"`
	PrivateAddress []string `json:"privateAddress"`
	NodeStatus     string   `json:"nodeStatus"`
	SparkMaster    bool     `json:"sparkMaster"`
	SparkJobserver bool     `json:"sparkJobserver"`
	Zeppelin       bool     `json:"zeppelin"`
}

type CreateVPCPeeringRequest struct {
	PeerVpcID     string `json:"peerVpcId"`
	PeerAccountID string `json:"peerAccountId"`
	PeerSubnet    string `json:"peerSubnet"`
	PeerRegion    string `json:"peerRegion,omitempty"`
}

type VPCPeering struct {
	ID                 string `json:"id"`
	AWSVpcConnectionID string `json:"aws_vpc_connection_id"`
	ClusterDataCentre  string `json:"clusterDataCentre"`
	VpcID              string `json:"vpcId"`
	PeerVpcID          string `json:"peerVpcId"`
	PeerAccountID      string `json:"peerAccountId"`
	PeerSubnet         string `json:"peerSubnet"`
	StatusCode         string `json:"statusCode"`
	PeerRegion         string `json:"peerRegion"`
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
	ID       string `json:"id,omitempty"`
	Alias    string `json:"alias,omitempty"`
	ARN      string `json:"arn,omitempty"`
	Provider string `json:"provider,omitempty"`
}

type UpdateBundleUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

type BundleConfig struct {
	IsKafkaCluster    bool
	HasRestProxy      bool
	HasSchemaRegistry bool
}
