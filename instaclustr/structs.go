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
	Bundle  string         `json:"bundle" mapstructure:"bundle"`
	Version string         `json:"version" mapstructure:"version"`
	Options *BundleOptions `json:"options,omitempty" mapstructure:"options"`
}

type OmitEmptyBool struct {
	value bool
}

type BundleOptions struct {
	AuthnAuthz                    		*bool  `json:"authnAuthz,omitempty" mapstructure:"auth_n_authz,omitempty"`
	ClientEncryption              		*bool  `json:"clientEncryption,omitempty" mapstructure:"client_encryption,omitempty"`
	DedicatedMasterNodes          		*bool  `json:"dedicatedMasterNodes,omitempty" mapstructure:"dedicated_master_nodes,omitempty"`
	MasterNodeSize                		string `json:"masterNodeSize,omitempty" mapstructure:"master_node_size,omitempty"`
	KibanaNodeSize                		string `json:"kibanaNodeSize,omitempty" mapstructure:"kibana_node_size,omitempty"`
	OpenSearchDashboardsNodeSize 		string `json:"openSearchDashboardsNodeSize,omitempty" mapstructure:"opensearch_dashboards_node_size,omitempty"`
	DataNodeSize                  		string `json:"dataNodeSize,omitempty" mapstructure:"data_node_size,omitempty"`
	SecurityPlugin                		*bool  `json:"securityPlugin,omitempty" mapstructure:"security_plugin,omitempty"`
	UsePrivateBroadcastRpcAddress 		*bool  `json:"usePrivateBroadcastRPCAddress,omitempty" mapstructure:"use_private_broadcast_rpc_address,omitempty"`
	LuceneEnabled                 		*bool  `json:"luceneEnabled,omitempty" mapstructure:"lucene_enabled,omitempty"`
	ContinuousBackupEnabled       		*bool  `json:"continuousBackupEnabled,omitempty" mapstructure:"continuous_backup_enabled,omitempty"`
	NumberPartitions              		int    `json:"numberPartitions,omitempty" mapstructure:"number_partitions,omitempty"`
	AutoCreateTopics              		*bool  `json:"autoCreateTopics,omitempty" mapstructure:"auto_create_topics,omitempty"`
	DeleteTopics                  		*bool  `json:"deleteTopics,omitempty" mapstructure:"delete_topics,omitempty"`
	PasswordAuthentication        		*bool  `json:"passwordAuthentication,omitempty" mapstructure:"password_authentication,omitempty"`
	TargetKafkaClusterId          		string `json:"targetKafkaClusterId,omitempty" mapstructure:"target_kafka_cluster_id,omitempty"`
	VPCType                       		string `json:"vpcType,omitempty" mapstructure:"vpc_type,omitempty"`
	AWSAccessKeyId                		string `json:"aws.access.key.id,omitempty" mapstructure:"aws_access_key,omitempty"`
	AWSSecretKey                  		string `json:"aws.secret.access.key,omitempty" mapstructure:"aws_secret_key,omitempty"`
	S3BucketName                  		string `json:"s3.bucket.name,omitempty" mapstructure:"s3_bucket_name,omitempty"`
	AzureStorageAccountName       		string `json:"azure.storage.account.name,omitempty" mapstructure:"azure_storage_account_name,omitempty"`
	AzureStorageAccountKey        		string `json:"azure.storage.account.key,omitempty" mapstructure:"azure_storage_account_key,omitempty"`
	AzureStorageContainerName     		string `json:"azure.storage.container.name,omitempty" mapstructure:"azure_storage_container_name,omitempty"`
	SslEnabledProtocols           		string `json:"ssl.enabled.protocols,omitempty" mapstructure:"ssl_enabled_protocols,omitempty"`
	SslTruststorePassword         		string `json:"ssl.truststore.password,omitempty" mapstructure:"ssl_truststore_password,omitempty"`
	SslProtocol                   		string `json:"ssl.protocol,omitempty" mapstructure:"ssl_protocol,omitempty"`
	SecurityProtocol              		string `json:"security.protocol,omitempty" mapstructure:"security_protocol,omitempty"`
	SaslMechanism                 		string `json:"sasl.mechanism,omitempty" mapstructure:"sasl_mechanism,omitempty"`
	SaslJaasConfig                		string `json:"sasl.jaas.config,omitempty" mapstructure:"sasl_jaas_config,omitempty"`
	BootstrapServers              		string `json:"bootstrap.servers,omitempty" mapstructure:"bootstrap_servers,omitempty"`
	Truststore                    		string `json:"truststore,omitempty" mapstructure:"truststore,omitempty"`
	RedisMasterNodes              		int    `json:"masterNodes,omitempty" mapstructure:"master_nodes,omitempty"`
	RedisReplicaNodes             		int    `json:"replicaNodes,omitempty" mapstructure:"replica_nodes,omitempty"`
	RedisPasswordAuth             		*bool  `json:"passwordAuth,omitempty" mapstructure:"password_auth,omitempty"`
	DedicatedZookeeper            		*bool  `json:"dedicatedZookeeper,omitempty" mapstructure:"dedicated_zookeeper,omitempty"`
	ZookeeperNodeSize             		string `json:"zookeeperNodeSize,omitempty" mapstructure:"zookeeper_node_size,omitempty"`
	ZookeeperNodeCount            		int    `json:"zookeeperNodeCount,omitempty" mapstructure:"zookeeper_node_count,omitempty"`
	PostgresqlNodeCount           		int    `json:"postgresqlNodeCount,omitempty" mapstructure:"postgresql_node_count,omitempty"`
	PostgresqlReplicationMode     		string `json:"replicationMode,omitempty" mapstructure:"replication_mode,omitempty"`
	PostgresqlSynchronousModeStrict     *bool `json:"synchronousModeStrict,omitempty" mapstructure:"synchronous_mode_strict,omitempty"`
	CadenceAdvancedVisibility           *bool  `json:"useAdvancedVisibility,omitempty" mapstructure:"advanced_visibility,omitempty"`
	CadenceTargetCassandraDataCentreID  string `json:"targetCassandraCdcId,omitempty" mapstructure:"target_cassandra_data_centre_id,omitempty"`
	CadenceTargetCassandraVPCType       string `json:"targetCassandraVpcType,omitempty" mapstructure:"target_cassandra_vpc_type,omitempty"`
	CadenceTargetOpensearchDataCentreID string `json:"targetOpenSearchCdcId,omitempty" mapstructure:"target_opensearch_data_centre_id,omitempty"`
	CadenceTargetOpensearchVPCType      string `json:"targetOpenSearchVpcType,omitempty" mapstructure:"target_opensearch_vpc_type,omitempty"`
	CadenceTargetKafkaDataCentreID      string `json:"targetKafkaCdcId,omitempty" mapstructure:"target_kafka_data_centre_id,omitempty"`
	CadenceTargetKafkaVPCType           string `json:"targetKafkaVpcType,omitempty" mapstructure:"target_kafka_vpc_type,omitempty"`
}

type ClusterProvider struct {
	Name                   *string                `json:"name" mapstructure:"name"`
	AccountName            *string                `json:"accountName,omitempty" mapstructure:"account_name"`
	CustomVirtualNetworkId *string                `json:"customVirtualNetworkId,omitempty" mapstructure:"custom_virtual_network_id"`
	Tags                   map[string]interface{} `json:"tags,omitempty"`
	ResourceGroup          *string                `json:"resourceGroup,omitempty" mapstructure:"resource_group"`
	DiskEncryptionKey      *string                `json:"diskEncryptionKey,omitempty" mapstructure:"disk_encryption_key"`
}

type RackAllocation struct {
	NumberOfRacks string `json:"numberOfRacks" mapstructure:"number_of_racks"`
	NodesPerRack  string `json:"nodesPerRack" mapstructure:"nodes_per_rack"`
}

type CreateRequest struct {
	ClusterName           string                    `json:"clusterName"`
	Bundles               []Bundle                  `json:"bundles,omitempty"`
	Provider              *ClusterProvider          `json:"provider,omitempty"`
	SlaTier               string                    `json:"slaTier,omitempty"`
	NodeSize              string                    `json:"nodeSize,omitempty"`
	DataCentre            string                    `json:"dataCentre,omitempty"`
	DataCentreCustomName  string                    `json:"dataCentreCustomName,omitempty"`
	DataCentres           []DataCentreCreateRequest `json:"dataCentres,omitempty"`
	ClusterNetwork        string                    `json:"clusterNetwork,omitempty"`
	PrivateNetworkCluster string                    `json:"privateNetworkCluster,omitempty"`
	PCICompliantCluster   string                    `json:"pciCompliantCluster,omitempty"`
	RackAllocation        *RackAllocation           `json:"rackAllocation,omitempty"`
}

type DataCentreCreateRequest struct {
	Name           string           `json:"name" mapstructure:"name"`
	Network        string           `json:"network" mapstructure:"network"`
	DataCentre     string           `json:"dataCentre" mapstructure:"data_centre"`
	Provider       *ClusterProvider `json:"provider,omitempty" mapstructure:"provider,omitempty"`
	NodeSize       string           `json:"nodeSize,omitempty" mapstructure:"node_size,omitempty"`
	Bundles        []Bundle         `json:"bundles,omitempty" mapstructure:"bundles,omitempty"`
	RackAllocation *RackAllocation  `json:"rackAllocation,omitempty" mapstructure:"rack_allocation,omitempty"`
}

type AddonBundles struct {
	Bundle  string `json:"bundle"`
	Version string `json:"version"`
}

type Cluster struct {
	ID                         string                   `json:"id"`
	ClusterName                string                   `json:"clusterName"`
	ClusterStatus              string                   `json:"clusterStatus"`
	CdcId                      string                   `json:"cdcId"`
	BundleType                 string                   `json:"bundleType"`
	BundleVersion              string                   `json:"bundleVersion"`
	AddonBundles               []map[string]interface{} `json:"addonBundles"`
	Username                   string                   `json:"username"`
	InstaclustrUserPassword    string                   `json:"instaclustrUserPassword"`
	SlaTier                    string                   `json:"slaTier"`
	ClusterCertificateDownload string                   `json:"clusterCertificateDownload"`
	PciCompliance              string                   `json:"pciCompliance"`
	BundleOption               *BundleOptions           `json:"bundleOptions"`
	DataCentre                 string                   `json:"dataCentre"`
	DataCentres                []DataCentre             `json:"dataCentres"`
	Provider                   []ClusterProvider        `json:"clusterProvider"`
}

type ClusterListItem struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	NodeCount        int    `json:"nodeCount,omitempty"`
	RunningNodeCount int    `json:"runningNodeCount,omitempty"`
	DerivedStatus    string `json:"derivedStatus,omitempty"`
	SlaTier          string `json:"slaTier,omitempty"`
	PciCompliance    string `json:"pciCompliance,omitempty"`
}

type DataCentre struct {
	ID                            string          `json:"id,omitempty"`
	Name                          string          `json:"name" mapstructure:"name"`
	CdcName                       string          `json:"cdcName,omitempty" mapstructure:"cdcName"`
	Provider                      string          `json:"provider,omitempty"`
	CdcNetwork                    string          `json:"cdcNetwork,omitempty"`
	Bundles                       []string        `json:"bundles,omitempty"`
	ClientEncryption              bool            `json:"clientEncryption,omitempty"`
	PasswordAuthentication        bool            `json:"passwordAuthentication,omitempty"`
	UserAuthorization             bool            `json:"userAuthorization,omitempty"`
	UsePrivateBroadcastRPCAddress bool            `json:"usePrivateBroadcastRPCAddress,omitempty"`
	PrivateIPOnly                 bool            `json:"privateIPOnly,omitempty"`
	Nodes                         []Node          `json:"nodes,omitempty"`
	NodeCount                     int             `json:"nodeCount,omitempty"`
	EncryptionKeyId               string          `json:"encryptionKeyId,omitempty"`
	ResizeTargetNodeSize          string          `json:"resizeTargetNodeSize,omitempty"`
	DataCentreRegion              string          `json:"dataCentre,omitempty" mapstructure:"data_centre_region"`
	CdcStatus                     string          `json:"cdcStatus,omitempty"`
	RackAllocation                *RackAllocation `json:"rackAllocation,omitempty" mapstructure:"rack_allocation,omitempty"`
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
	PeerVpcID     string        `json:"peerVpcId"`
	PeerAccountID string        `json:"peerAccountId"`
	PeerSubnet    string        `json:"peerSubnet"`
	PeerSubnets   []interface{} `json:"peerSubnets"`
	PeerRegion    string        `json:"peerRegion,omitempty"`
}

type CreateGCPVPCPeeringRequest struct {
	PeerProjectID      string        `json:"peerProjectId"`
	PeerVPCNetworkName string        `json:"peerVpcNetworkName"`
	PeerSubnets        []interface{} `json:"peerSubnets"`
}

type VPCPeering struct {
	ID                 string        `json:"id"`
	AWSVpcConnectionID string        `json:"aws_vpc_connection_id"`
	ClusterDataCentre  string        `json:"clusterDataCentre"`
	VpcID              string        `json:"vpcId"`
	PeerVpcID          string        `json:"peerVpcId"`
	PeerAccountID      string        `json:"peerAccountId"`
	PeerSubnet         string        `json:"peerSubnet"`
	PeerSubnets        []interface{} `json:"peerSubnets"`
	StatusCode         string        `json:"statusCode"`
	PeerRegion         string        `json:"peerRegion"`
}

type GCPVPCPeering struct {
	ID                 string `json:"id"`
	ClusterDataCentre  string `json:"clusterDataCentre"`
	VpcID              string `json:"vpcNetworkName"`
	PeerProjectID      string `json:"projectId"`
	PeerVPCNetworkName string `json:"peerVpcNetworkName"`

	PeerSubnets []interface{} `json:"peerSubnets"`
	StatusCode  string        `json:"statusCode"`
}

type ResizeClusterRequest struct {
	NewNodeSize           string       `json:"newNodeSize"`
	ConcurrentResizes     int          `json:"concurrentResizes"`
	NotifySupportContacts string       `json:"notifySupportContacts"`
	NodePurpose           *NodePurpose `json:"nodePurpose"`
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

type BundleConfig struct {
	IsKafkaCluster    bool
	HasRestProxy      bool
	HasSchemaRegistry bool
}
