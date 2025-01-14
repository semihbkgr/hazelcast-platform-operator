package config

type HazelcastWrapper struct {
	Hazelcast Hazelcast `yaml:"hazelcast"`
}

type Hazelcast struct {
	Jet                      Jet                                 `yaml:"jet,omitempty"`
	Network                  Network                             `yaml:"network,omitempty"`
	ClusterName              string                              `yaml:"cluster-name,omitempty"`
	Persistence              Persistence                         `yaml:"persistence,omitempty"`
	Map                      map[string]Map                      `yaml:"map,omitempty"`
	ExecutorService          map[string]ExecutorService          `yaml:"executor-service,omitempty"`
	DurableExecutorService   map[string]DurableExecutorService   `yaml:"durable-executor-service,omitempty"`
	ScheduledExecutorService map[string]ScheduledExecutorService `yaml:"scheduled-executor-service,omitempty"`
	UserCodeDeployment       UserCodeDeployment                  `yaml:"user-code-deployment,omitempty"`
	WanReplication           map[string]WanReplicationConfig     `yaml:"wan-replication,omitempty"`
	Properties               map[string]string                   `yaml:"properties,omitempty"`
	MultiMap                 map[string]MultiMap                 `yaml:"multimap,omitempty"`
	Topic                    map[string]Topic                    `yaml:"topic,omitempty"`
	ReplicatedMap            map[string]ReplicatedMap            `yaml:"replicatedmap,omitempty"`
	Queue                    map[string]Queue                    `yaml:"queue,omitempty"`
	Cache                    map[string]Cache                    `yaml:"cache,omitempty"`
	PartitionGroup           PartitionGroup                      `yaml:"partition-group,omitempty"`
	NativeMemory             NativeMemory                        `yaml:"native-memory,omitempty"`
}

type PartitionGroup struct {
	Enabled   *bool  `yaml:"enabled,omitempty"`
	GroupType string `yaml:"group-type,omitempty"`
}

type Jet struct {
	Enabled               *bool `yaml:"enabled,omitempty"`
	ResourceUploadEnabled *bool `yaml:"resource-upload-enabled,omitempty"`
}

type Network struct {
	Join    Join    `yaml:"join,omitempty"`
	RestAPI RestAPI `yaml:"rest-api,omitempty"`
}

type Join struct {
	Kubernetes Kubernetes `yaml:"kubernetes,omitempty"`
}

type Persistence struct {
	Enabled                   *bool  `yaml:"enabled,omitempty"`
	BaseDir                   string `yaml:"base-dir"`
	BackupDir                 string `yaml:"backup-dir,omitempty"`
	Parallelism               int32  `yaml:"parallelism"`
	ValidationTimeoutSec      int32  `yaml:"validation-timeout-seconds"`
	ClusterDataRecoveryPolicy string `yaml:"cluster-data-recovery-policy"`
	AutoRemoveStaleData       *bool  `yaml:"auto-remove-stale-data"`
}

type Kubernetes struct {
	Enabled                      *bool  `yaml:"enabled,omitempty"`
	Namespace                    string `yaml:"namespace,omitempty"`
	ServiceName                  string `yaml:"service-name,omitempty"`
	UseNodeNameAsExternalAddress *bool  `yaml:"use-node-name-as-external-address,omitempty"`
	ServicePerPodLabelName       string `yaml:"service-per-pod-label-name,omitempty"`
	ServicePerPodLabelValue      string `yaml:"service-per-pod-label-value,omitempty"`
}

type RestAPI struct {
	Enabled        *bool          `yaml:"enabled,omitempty"`
	EndpointGroups EndpointGroups `yaml:"endpoint-groups,omitempty"`
}

type EndpointGroups struct {
	HealthCheck  EndpointGroup `yaml:"HEALTH_CHECK,omitempty"`
	ClusterWrite EndpointGroup `yaml:"CLUSTER_WRITE,omitempty"`
	Persistence  EndpointGroup `yaml:"PERSISTENCE,omitempty"`
}

type EndpointGroup struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Map struct {
	BackupCount             int32                              `yaml:"backup-count"`
	AsyncBackupCount        int32                              `yaml:"async-backup-count"`
	TimeToLiveSeconds       int32                              `yaml:"time-to-live-seconds"`
	MaxIdleSeconds          int32                              `yaml:"max-idle-seconds"`
	Eviction                MapEviction                        `yaml:"eviction,omitempty"`
	ReadBackupData          bool                               `yaml:"read-backup-data"`
	InMemoryFormat          string                             `yaml:"in-memory-format"`
	StatisticsEnabled       bool                               `yaml:"statistics-enabled"`
	Indexes                 []MapIndex                         `yaml:"indexes,omitempty"`
	DataPersistence         DataPersistence                    `yaml:"data-persistence,omitempty"`
	WanReplicationReference map[string]WanReplicationReference `yaml:"wan-replication-ref,omitempty"`
	MapStoreConfig          MapStoreConfig                     `yaml:"map-store,omitempty"`
	EntryListeners          []EntryListener                    `yaml:"entry-listeners,omitempty"`
}

type EntryListener struct {
	ClassName    string `yaml:"class-name"`
	IncludeValue bool   `yaml:"include-value"`
	Local        bool   `yaml:"local"`
}

type MapEviction struct {
	Size           int32  `yaml:"size"`
	MaxSizePolicy  string `yaml:"max-size-policy,omitempty"`
	EvictionPolicy string `yaml:"eviction-policy,omitempty"`
}

type MapIndex struct {
	Name               string             `yaml:"name,omitempty"`
	Type               string             `yaml:"type"`
	Attributes         []string           `yaml:"attributes"`
	BitmapIndexOptions BitmapIndexOptions `yaml:"bitmap-index-options,omitempty"`
}

type BitmapIndexOptions struct {
	UniqueKey               string `yaml:"unique-key"`
	UniqueKeyTransformation string `yaml:"unique-key-transformation"`
}

type DataPersistence struct {
	Enabled bool `yaml:"enabled"`
	Fsync   bool `yaml:"fsync"`
}

type WanReplicationReference struct {
	MergePolicyClassName string   `yaml:"merge-policy-class-name"`
	RepublishingEnabled  bool     `yaml:"republishing-enabled"`
	Filters              []string `yaml:"filters"`
}

type ExecutorService struct {
	PoolSize      int32 `yaml:"pool-size"`
	QueueCapacity int32 `yaml:"queue-capacity"`
}

type DurableExecutorService struct {
	PoolSize   int32 `yaml:"pool-size"`
	Durability int32 `yaml:"durability"`
	Capacity   int32 `yaml:"capacity"`
}

type ScheduledExecutorService struct {
	PoolSize       int32  `yaml:"pool-size"`
	Durability     int32  `yaml:"durability"`
	Capacity       int32  `yaml:"capacity"`
	CapacityPolicy string `yaml:"capacity-policy"`
}

type MapStoreConfig struct {
	Enabled           bool              `yaml:"enabled"`
	WriteCoalescing   *bool             `yaml:"write-coalescing,omitempty"`
	WriteDelaySeconds int32             `yaml:"write-delay-seconds"`
	WriteBatchSize    int32             `yaml:"write-batch-size"`
	ClassName         string            `yaml:"class-name"`
	Properties        map[string]string `yaml:"properties"`
	InitialLoadMode   string            `yaml:"initial-mode"`
}

type Topic struct {
	GlobalOrderingEnabled bool     `yaml:"global-ordering-enabled"`
	MultiThreadingEnabled bool     `yaml:"multi-threading-enabled"`
	StatisticsEnabled     bool     `yaml:"statistics-enabled"`
	MessageListeners      []string `yaml:"message-listeners,omitempty"`
}

type UserCodeDeployment struct {
	Enabled           *bool  `yaml:"enabled,omitempty"`
	ClassCacheMode    string `yaml:"class-cache-mode,omitempty"`
	ProviderMode      string `yaml:"provider-mode,omitempty"`
	BlacklistPrefixes string `yaml:"blacklist-prefixes,omitempty"`
	WhitelistPrefixes string `yaml:"whitelist-prefixes,omitempty"`
	ProviderFilter    string `yaml:"provider-filter,omitempty"`
}

type MultiMap struct {
	BackupCount       int32       `yaml:"backup-count"`
	AsyncBackupCount  int32       `yaml:"async-backup-count"`
	Binary            bool        `yaml:"binary"`
	CollectionType    string      `yaml:"value-collection-type"`
	StatisticsEnabled bool        `yaml:"statistics-enabled"`
	MergePolicy       MergePolicy `yaml:"merge-policy"`
}

type Queue struct {
	BackupCount             int32       `yaml:"backup-count"`
	AsyncBackupCount        int32       `yaml:"async-backup-count"`
	EmptyQueueTtl           int32       `yaml:"empty-queue-ttl"`
	MaxSize                 int32       `yaml:"max-size"`
	StatisticsEnabled       bool        `yaml:"statistics-enabled"`
	MergePolicy             MergePolicy `yaml:"merge-policy"`
	PriorityComparatorClass string      `yaml:"priority-comparator-class-name"`
}

type Cache struct {
	BackupCount       int32           `yaml:"backup-count"`
	AsyncBackupCount  int32           `yaml:"async-backup-count"`
	StatisticsEnabled bool            `yaml:"statistics-enabled"`
	ManagementEnabled bool            `yaml:"management-enabled"`
	ReadThrough       bool            `yaml:"read-through"`
	WriteThrough      bool            `yaml:"write-through"`
	MergePolicy       MergePolicy     `yaml:"merge-policy"`
	KeyType           ClassType       `yaml:"key-type,omitempty"`
	ValueType         ClassType       `yaml:"value-type,omitempty"`
	InMemoryFormat    string          `yaml:"in-memory-format"`
	DataPersistence   DataPersistence `yaml:"data-persistence,omitempty"`
}

type ClassType struct {
	ClassName string `yaml:"class-name"`
}

type MergePolicy struct {
	ClassName string `yaml:"class-name"`
	BatchSize int32  `yaml:"batch-size"`
}

type ReplicatedMap struct {
	InMemoryFormat    string      `yaml:"in-memory-format"`
	AsyncFillup       bool        `yaml:"async-fillup"`
	StatisticsEnabled bool        `yaml:"statistics-enabled"`
	MergePolicy       MergePolicy `yaml:"merge-policy"`
}

type WanReplicationConfig struct {
	BatchPublisher map[string]BatchPublisherConfig `yaml:"batch-publisher,omitempty"`
}

type BatchPublisherConfig struct {
	ClusterName           string `yaml:"cluster-name,omitempty"`
	BatchSize             int32  `yaml:"batch-size,omitempty"`
	BatchMaxDelayMillis   int32  `yaml:"batch-max-delay-millis,omitempty"`
	ResponseTimeoutMillis int32  `yaml:"response-timeout-millis,omitempty"`
	AcknowledgementType   string `yaml:"acknowledge-type,omitempty"`
	InitialPublisherState string `yaml:"initial-publisher-state,omitempty"`
	QueueFullBehavior     string `yaml:"queue-full-behavior,omitempty"`
	QueueCapacity         int32  `yaml:"queue-capacity,omitempty"`
	TargetEndpoints       string `yaml:"target-endpoints,omitempty"`
}

type NativeMemory struct {
	Enabled                 bool             `yaml:"enabled"`
	AllocatorType           string           `yaml:"allocator-type"`
	Size                    NativeMemorySize `yaml:"size,omitempty"`
	MinBlockSize            int32            `yaml:"min-block-size,omitempty"`
	PageSize                int32            `yaml:"page-size,omitempty"`
	MetadataSpacePercentage int32            `yaml:"metadata-space-percentage,omitempty"`
}

type NativeMemorySize struct {
	Value int64  `yaml:"value"`
	Unit  string `yaml:"unit"`
}

func (hz Hazelcast) HazelcastConfigForcingRestart() Hazelcast {
	return Hazelcast{
		ClusterName: hz.ClusterName,
		Network: Network{
			Join: Join{
				Kubernetes: Kubernetes{
					ServicePerPodLabelName:       hz.Network.Join.Kubernetes.ServicePerPodLabelName,
					ServicePerPodLabelValue:      hz.Network.Join.Kubernetes.ServicePerPodLabelValue,
					UseNodeNameAsExternalAddress: hz.Network.Join.Kubernetes.UseNodeNameAsExternalAddress,
				},
			},
		},
		Jet:                hz.Jet,
		UserCodeDeployment: hz.UserCodeDeployment,
		Properties:         hz.Properties,
	}
}
