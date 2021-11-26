package instaclustr

type KafkaAcl struct {
	Principal	string	`json:"principal,omitempty"`
	Host		string	`json:"host,omitempty"`
	ResourceType	string	`json:"resourceType,omitempty"`
	ResourceName	string	`json:"resourceName,omitempty"`
	Operation	string	`json:"operation,omitempty"`
	PermissionType	string	`json:"permissionType,omitempty"`
	PatternType	string	`json:"patternType,omitempty"`
}

type KafkaAclList struct {
	Acls		[]KafkaAcl	`json:"acls"`
} 
