package instaclustr

type CreateKafkaUserRequest struct {
	Username           string                 `json:"username"`
	Password           string                 `json:"password"`
	InitialPermissions string                 `json:"initial-permissions"`
	Options            KafkaUserCreateOptions `json:"options,omitempty"`
}

type KafkaUserCreateOptions struct {
	AuthenticationMechanism string `json:"sasl-scram-mechanism,omitempty" mapstructure:"sasl-scram-mechanism"`
	OverrideExistingUser    bool   `json:"override-existing-user" mapstructure:"override-existing-user"`
}

type UpdateKafkaUserRequest struct {
	Username string                        `json:"username"`
	Password string                        `json:"password"`
	Options  KafkaUserResetPasswordOptions `json:"options,omitempty"`
}

type KafkaUserResetPasswordOptions struct {
	AuthenticationMechanism string `json:"sasl-scram-mechanism,omitempty" mapstructure:"sasl-scram-mechanism"`
}

type DeleteKafkaUserRequest struct {
	Username string `json:"username"`
}

type CreateKafkaTopicRequest struct {
	Topic             string `json:"topic"`
	ReplicationFactor int    `json:"replication-factor" mapstructure:"replication-factor"`
	Partitions        int    `json:"partitions"`
}

type KafkaTopicConfigOptions struct {
	CompressionType                      string  `json:"compression.type,omitempty" mapstructure:"compression_type,omitempty"`
	LeaderReplicationThrottledReplicas   *string `json:"leader.replication.throttled.replicas,omitempty" mapstructure:"leader_replication_throttled_replicas,omitempty"`
	MinInsyncReplicas                    int64   `json:"min.insync.replicas,omitempty,string" mapstructure:"min_insync_replicas,omitempty"`
	MessageDownconversionEnable          *bool   `json:"message.downconversion.enable,omitempty,string" mapstructure:"message_downconversion_enable,omitempty"`
	SegmentJitterMs                      *int64  `json:"segment.jitter.ms,omitempty,string" mapstructure:"segment_jitter_ms,omitempty"`
	CleanupPolicy                        string  `json:"cleanup.policy,omitempty" mapstructure:"cleanup_policy,omitempty"`
	FlushMs                              string  `json:"flush.ms,omitempty" mapstructure:"flush_ms,omitempty"` //Using string because terraform has difficulties parsing big int, it will lose precision.
	FollowerReplicationThrottledReplicas string  `json:"follower.replication.throttled.replicas,omitempty" mapstructure:"follower_replication_throttled_replicas,omitempty"`
	RetentionMs                          int64   `json:"retention.ms,omitempty,string" mapstructure:"retention_ms,omitempty"`
	SegmentBytes                         int64   `json:"segment.bytes,omitempty,string" mapstructure:"segment_bytes,omitempty"`
	FlushMessages                        string  `json:"flush.messages,omitempty" mapstructure:"flush_messages,omitempty"` //Using string because terraform has difficulties parsing big int, it will lose precision.
	MessageFormatVersion                 string  `json:"message.format.version,omitempty" mapstructure:"message_format_version,omitempty"`
	FileDeleteDelayMs                    int64   `json:"file.delete.delay.ms,omitempty,string" mapstructure:"file_delete_delay_ms,omitempty"`
	MaxCompactionLagMs                   string  `json:"max.compaction.lag.ms,omitempty" mapstructure:"max_compaction_lag_ms,omitempty"` //Using string because terraform has difficulties parsing big int, it will lose precision.
	MaxMessageBytes                      int64   `json:"max.message.bytes,omitempty,string" mapstructure:"max_message_bytes,omitempty"`
	MinCompactionLagMs                   *int64  `json:"min.compaction.lag.ms,omitempty,string" mapstructure:"min_compaction_lag_ms,omitempty"`
	MessageTimestampType                 string  `json:"message.timestamp.type,omitempty" mapstructure:"message_timestamp_type,omitempty"`
	Preallocate                          *bool   `json:"preallocate,omitempty,string" mapstructure:"preallocate,omitempty"`
	IndexIntervalBytes                   int64   `json:"index.interval.bytes,omitempty,string" mapstructure:"index_interval_bytes,omitempty"`
	MinCleanableDirtyRatio               float32 `json:"min.cleanable.dirty.ratio,omitempty,string" mapstructure:"min_cleanable_dirty_ratio,omitempty"`
	UncleanLeaderElectionEnable          *bool   `json:"unclean.leader.election.enable,omitempty,string" mapstructure:"unclean_leader_election_enable,omitempty"`
	DeleteRetentionMs                    int64   `json:"delete.retention.ms,omitempty,string" mapstructure:"delete_retention_ms,omitempty"`
	RetentionBytes                       int64   `json:"retention.bytes,omitempty,string" mapstructure:"retention_bytes,omitempty"`
	SegmentMs                            int64   `json:"segment.ms,omitempty,string" mapstructure:"segment_ms,omitempty"`
	MessageTimestampDifferenceMaxMs      string  `json:"message.timestamp.difference.max.ms,omitempty" mapstructure:"message_timestamp_difference_max_ms,omitempty"` //Using string because terraform has difficulties parsing big int, it will lose precision.
	SegmentIndexBytes                    int64   `json:"segment.index.bytes,omitempty,string" mapstructure:"segment_index_bytes,omitempty"`
}

type UpdateKafkaTopicRequest struct {
	Config *KafkaTopicConfigOptions `json:"config,omitempty" mapstructure:"config,omitempty"`
}

type KafkaTopicConfig struct {
	Topic  string                   `json:"topic"`
	Config *KafkaTopicConfigOptions `json:"config,omitempty" mapstructure:"config,omitempty"`
}

type KafkaTopics struct {
	Topics []string `json:"topics,omitempty" mapstructure:"topics,omitempty"`
}
