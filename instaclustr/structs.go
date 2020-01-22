package instaclustr

type FirewallRule struct {
	Network string     `json:"network"`
	Rules   []RuleType `json:"rules"`
}

type RuleType struct {
	Type string `json:"type"`
}

type Bundle struct {
	Bundle  string `json:"bundle"`
	Version string `json:"version"`
}

type ClusterProvider struct {
	Name        *string `json:"name"`
	AccountName *string `json:"accountName"`
}

type RackAllocation struct {
	NumberOfRacks string `json:"numberOfRacks"`
	NodesPerRack  string `json:"nodesPerRack"`
}

type CreateRequest struct {
	ClusterName    string          `json:"clusterName"`
	Bundles        []Bundle        `json:"bundles"`
	Provider       ClusterProvider `json:"provider"`
	SlaTier        string          `json:"slaTier"`
	NodeSize       string          `json:"nodeSize"`
	DataCentre     string          `json:"dataCentre"`
	ClusterNetwork string          `json:"clusterNetwork"`
	RackAllocation RackAllocation  `json:"rackAllocation"`
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
