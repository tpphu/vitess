// Code generated by protoc-gen-go.
// source: topo.proto
// DO NOT EDIT!

/*
Package topo is a generated protocol buffer package.

It is generated from these files:
	topo.proto

It has these top-level messages:
	KeyRange
	TabletAlias
	Tablet
	Shard
	Keyspace
	ShardReplication
	EndPoint
	EndPoints
	SrvShard
	ShardReference
	SrvKeyspace
*/
package topo

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// KeyspaceIdType describes the type of the sharding key for a
// range-based sharded keyspace.
type KeyspaceIdType int32

const (
	// UNSET is the default value, when range-based sharding is not used.
	KeyspaceIdType_UNSET KeyspaceIdType = 0
	// UINT64 is when uint64 value is used.
	// This is represented as 'unsigned bigint' in mysql
	KeyspaceIdType_UINT64 KeyspaceIdType = 1
	// BYTES is when an array of bytes is used.
	// This is represented as 'varbinary' in mysql
	KeyspaceIdType_BYTES KeyspaceIdType = 2
)

var KeyspaceIdType_name = map[int32]string{
	0: "UNSET",
	1: "UINT64",
	2: "BYTES",
}
var KeyspaceIdType_value = map[string]int32{
	"UNSET":  0,
	"UINT64": 1,
	"BYTES":  2,
}

func (x KeyspaceIdType) String() string {
	return proto.EnumName(KeyspaceIdType_name, int32(x))
}

// TabletType represents the type of a given tablet.
type TabletType int32

const (
	TabletType_UNKNOWN        TabletType = 0
	TabletType_IDLE           TabletType = 1
	TabletType_MASTER         TabletType = 2
	TabletType_REPLICA        TabletType = 3
	TabletType_RDONLY         TabletType = 4
	TabletType_BATCH          TabletType = 4
	TabletType_SPARE          TabletType = 5
	TabletType_EXPERIMENTAL   TabletType = 6
	TabletType_SCHEMA_UPGRADE TabletType = 7
	TabletType_BACKUP         TabletType = 8
	TabletType_RESTORE        TabletType = 9
	TabletType_WORKER         TabletType = 10
	TabletType_SCRAP          TabletType = 11
)

var TabletType_name = map[int32]string{
	0: "UNKNOWN",
	1: "IDLE",
	2: "MASTER",
	3: "REPLICA",
	4: "RDONLY",
	// Duplicate value: 4: "BATCH",
	5:  "SPARE",
	6:  "EXPERIMENTAL",
	7:  "SCHEMA_UPGRADE",
	8:  "BACKUP",
	9:  "RESTORE",
	10: "WORKER",
	11: "SCRAP",
}
var TabletType_value = map[string]int32{
	"UNKNOWN":        0,
	"IDLE":           1,
	"MASTER":         2,
	"REPLICA":        3,
	"RDONLY":         4,
	"BATCH":          4,
	"SPARE":          5,
	"EXPERIMENTAL":   6,
	"SCHEMA_UPGRADE": 7,
	"BACKUP":         8,
	"RESTORE":        9,
	"WORKER":         10,
	"SCRAP":          11,
}

func (x TabletType) String() string {
	return proto.EnumName(TabletType_name, int32(x))
}

// KeyRange describes a range of sharding keys, when range-based
// sharding is used.
type KeyRange struct {
	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   []byte `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *KeyRange) Reset()         { *m = KeyRange{} }
func (m *KeyRange) String() string { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()    {}

// TabletAlias is a globally unique tablet identifier.
type TabletAlias struct {
	// cell is the cell (or datacenter) the tablet is in
	Cell string `protobuf:"bytes,1,opt,name=cell" json:"cell,omitempty"`
	// uid is a unique id for this tablet within the shard
	// (this is the MySQL server id as well).
	Uid uint32 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *TabletAlias) Reset()         { *m = TabletAlias{} }
func (m *TabletAlias) String() string { return proto.CompactTextString(m) }
func (*TabletAlias) ProtoMessage()    {}

// Tablet represents information about a running instance of vttablet.
type Tablet struct {
	// alias is the unique name of the tablet.
	Alias *TabletAlias `protobuf:"bytes,1,opt,name=alias" json:"alias,omitempty"`
	// Fully qualified domain name of the host.
	Hostname string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	// IP address, stored as a string.
	Ip string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	// Map of named ports. Normally this should include vt, vts, and mysql.
	Portmap map[string]int32 `protobuf:"bytes,4,rep,name=portmap" json:"portmap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Keyspace name.
	Keyspace string `protobuf:"bytes,5,opt,name=keyspace" json:"keyspace,omitempty"`
	// Shard name. If range based sharding is used, it should match
	// key_range.
	Shard string `protobuf:"bytes,6,opt,name=shard" json:"shard,omitempty"`
	// If range based sharding is used, range for the tablet's shard.
	KeyRange *KeyRange `protobuf:"bytes,7,opt,name=key_range" json:"key_range,omitempty"`
	// type is the current type of the tablet.
	Type TabletType `protobuf:"varint,8,opt,name=type,enum=topo.TabletType" json:"type,omitempty"`
	// It this is set, it is used as the database name instead of the
	// normal "vt_" + keyspace.
	DbNameOverride string `protobuf:"bytes,9,opt,name=db_name_override" json:"db_name_override,omitempty"`
	// tablet tags
	Tags map[string]string `protobuf:"bytes,10,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// tablet health information
	HealthMap map[string]string `protobuf:"bytes,11,rep,name=health_map" json:"health_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Tablet) Reset()         { *m = Tablet{} }
func (m *Tablet) String() string { return proto.CompactTextString(m) }
func (*Tablet) ProtoMessage()    {}

func (m *Tablet) GetAlias() *TabletAlias {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Tablet) GetPortmap() map[string]int32 {
	if m != nil {
		return m.Portmap
	}
	return nil
}

func (m *Tablet) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Tablet) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Tablet) GetHealthMap() map[string]string {
	if m != nil {
		return m.HealthMap
	}
	return nil
}

// A Shard contains data about a subset of the data whithin a keyspace.
type Shard struct {
	// There can be only at most one master, but there may be none. (0)
	MasterAlias *TabletAlias `protobuf:"bytes,1,opt,name=master_alias" json:"master_alias,omitempty"`
	// This must match the shard name based on our other conventions, but
	// helpful to have it decomposed here.
	KeyRange *KeyRange `protobuf:"bytes,2,opt,name=key_range" json:"key_range,omitempty"`
	// served_type_map is a list here, but there is at most one entry
	// per TabletType
	ServedTypeMap []*Shard_ShardServedType `protobuf:"bytes,3,rep,name=served_type_map" json:"served_type_map,omitempty"`
	// SourceShards is the list of shards we're replicating from,
	// using filtered replication.
	SourceShards []*Shard_SourceShard `protobuf:"bytes,4,rep,name=source_shards" json:"source_shards,omitempty"`
	// Cells is the list of cells that contain tablets for this shard.
	Cells []string `protobuf:"bytes,5,rep,name=cells" json:"cells,omitempty"`
	// tablet_control_maps is a map in go, but a list in proto.
	TabletControls []*Shard_TabletControl `protobuf:"bytes,6,rep,name=tablet_controls" json:"tablet_controls,omitempty"`
}

func (m *Shard) Reset()         { *m = Shard{} }
func (m *Shard) String() string { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()    {}

func (m *Shard) GetMasterAlias() *TabletAlias {
	if m != nil {
		return m.MasterAlias
	}
	return nil
}

func (m *Shard) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Shard) GetServedTypeMap() []*Shard_ShardServedType {
	if m != nil {
		return m.ServedTypeMap
	}
	return nil
}

func (m *Shard) GetSourceShards() []*Shard_SourceShard {
	if m != nil {
		return m.SourceShards
	}
	return nil
}

func (m *Shard) GetTabletControls() []*Shard_TabletControl {
	if m != nil {
		return m.TabletControls
	}
	return nil
}

// ShardServedType is an entry in the served_type_map
type Shard_ShardServedType struct {
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,enum=topo.TabletType" json:"tablet_type,omitempty"`
	Cells      []string   `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
}

func (m *Shard_ShardServedType) Reset()         { *m = Shard_ShardServedType{} }
func (m *Shard_ShardServedType) String() string { return proto.CompactTextString(m) }
func (*Shard_ShardServedType) ProtoMessage()    {}

// SourceShard represents a data source for filtered replication
// accross shards. When this is used in a destination shard, the master
// of that shard will run filtered replication.
type Shard_SourceShard struct {
	// Uid is the unique ID for this SourceShard object.
	Uid uint32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	// the source keyspace
	Keyspace string `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
	// the source shard
	Shard string `protobuf:"bytes,3,opt,name=shard" json:"shard,omitempty"`
	// the source shard keyrange
	KeyRange *KeyRange `protobuf:"bytes,4,opt,name=key_range" json:"key_range,omitempty"`
	// the source table list to replicate
	Tables []string `protobuf:"bytes,5,rep,name=tables" json:"tables,omitempty"`
}

func (m *Shard_SourceShard) Reset()         { *m = Shard_SourceShard{} }
func (m *Shard_SourceShard) String() string { return proto.CompactTextString(m) }
func (*Shard_SourceShard) ProtoMessage()    {}

func (m *Shard_SourceShard) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

// TabletControl controls tablet's behavior
type Shard_TabletControl struct {
	// which tablet type is affected
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,enum=topo.TabletType" json:"tablet_type,omitempty"`
	Cells      []string   `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
	// what to do
	DisableQueryService bool     `protobuf:"varint,3,opt,name=disable_query_service" json:"disable_query_service,omitempty"`
	BlacklistedTables   []string `protobuf:"bytes,4,rep,name=blacklisted_tables" json:"blacklisted_tables,omitempty"`
}

func (m *Shard_TabletControl) Reset()         { *m = Shard_TabletControl{} }
func (m *Shard_TabletControl) String() string { return proto.CompactTextString(m) }
func (*Shard_TabletControl) ProtoMessage()    {}

// A Keyspace contains data about a keyspace.
type Keyspace struct {
	// name of the column used for sharding
	// empty if the keyspace is not sharded
	ShardingColumnName string `protobuf:"bytes,1,opt,name=sharding_column_name" json:"sharding_column_name,omitempty"`
	// type of the column used for sharding
	// KIT_UNSET if the keyspace is not sharded
	ShardingColumnType KeyspaceIdType `protobuf:"varint,2,opt,name=sharding_column_type,enum=topo.KeyspaceIdType" json:"sharding_column_type,omitempty"`
	// SplitShardCount stores the number of jobs to run to be sure to
	// always have at most one job per shard (used during resharding).
	SplitShardCount int32 `protobuf:"varint,3,opt,name=split_shard_count" json:"split_shard_count,omitempty"`
	// KeyspaceServedFrom will redirect the appropriate traffic to
	// another keyspace.
	ServedFroms []*Keyspace_KeyspaceServedFrom `protobuf:"bytes,4,rep,name=served_froms" json:"served_froms,omitempty"`
}

func (m *Keyspace) Reset()         { *m = Keyspace{} }
func (m *Keyspace) String() string { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()    {}

func (m *Keyspace) GetServedFroms() []*Keyspace_KeyspaceServedFrom {
	if m != nil {
		return m.ServedFroms
	}
	return nil
}

// KeyspaceServedFrom indicates a relationship between a TabletType and the
// keyspace name that's serving it.
type Keyspace_KeyspaceServedFrom struct {
	// the tablet type (key for the map)
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,enum=topo.TabletType" json:"tablet_type,omitempty"`
	// the cells to limit this to
	Cells []string `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
	// the keyspace name that's serving it
	Keyspace string `protobuf:"bytes,3,opt,name=keyspace" json:"keyspace,omitempty"`
}

func (m *Keyspace_KeyspaceServedFrom) Reset()         { *m = Keyspace_KeyspaceServedFrom{} }
func (m *Keyspace_KeyspaceServedFrom) String() string { return proto.CompactTextString(m) }
func (*Keyspace_KeyspaceServedFrom) ProtoMessage()    {}

// ShardReplication describes the MySQL replication relationships
// whithin a cell.
type ShardReplication struct {
	// Note there can be only one ReplicationLink in this array
	// for a given Slave.
	ReplicationLinks []*ShardReplication_ReplicationLink `protobuf:"bytes,1,rep,name=replication_links" json:"replication_links,omitempty"`
}

func (m *ShardReplication) Reset()         { *m = ShardReplication{} }
func (m *ShardReplication) String() string { return proto.CompactTextString(m) }
func (*ShardReplication) ProtoMessage()    {}

func (m *ShardReplication) GetReplicationLinks() []*ShardReplication_ReplicationLink {
	if m != nil {
		return m.ReplicationLinks
	}
	return nil
}

// ReplicationLink describes a tablet instance within the cell
type ShardReplication_ReplicationLink struct {
	TabletAlias *TabletAlias `protobuf:"bytes,1,opt,name=tablet_alias" json:"tablet_alias,omitempty"`
}

func (m *ShardReplication_ReplicationLink) Reset()         { *m = ShardReplication_ReplicationLink{} }
func (m *ShardReplication_ReplicationLink) String() string { return proto.CompactTextString(m) }
func (*ShardReplication_ReplicationLink) ProtoMessage()    {}

func (m *ShardReplication_ReplicationLink) GetTabletAlias() *TabletAlias {
	if m != nil {
		return m.TabletAlias
	}
	return nil
}

// EndPoint corresponds to a single tablet.
type EndPoint struct {
	// The uid of the tablet.
	Uid uint32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	// The host the tablet is running on (FQDN).
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	// The ports opened for service.
	Portmap map[string]int32 `protobuf:"bytes,3,rep,name=portmap" json:"portmap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// The health entries.
	HealthMap map[string]string `protobuf:"bytes,4,rep,name=health_map" json:"health_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EndPoint) Reset()         { *m = EndPoint{} }
func (m *EndPoint) String() string { return proto.CompactTextString(m) }
func (*EndPoint) ProtoMessage()    {}

func (m *EndPoint) GetPortmap() map[string]int32 {
	if m != nil {
		return m.Portmap
	}
	return nil
}

func (m *EndPoint) GetHealthMap() map[string]string {
	if m != nil {
		return m.HealthMap
	}
	return nil
}

// EndPoints corresponds to a list of tablets.
type EndPoints struct {
	Entries []*EndPoint `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *EndPoints) Reset()         { *m = EndPoints{} }
func (m *EndPoints) String() string { return proto.CompactTextString(m) }
func (*EndPoints) ProtoMessage()    {}

func (m *EndPoints) GetEntries() []*EndPoint {
	if m != nil {
		return m.Entries
	}
	return nil
}

// SrvShard is a rollup node for the shard itself.
type SrvShard struct {
	// Copied from Shard.
	Name     string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	KeyRange *KeyRange `protobuf:"bytes,2,opt,name=key_range" json:"key_range,omitempty"`
	// The cell that master tablet resides in.
	MasterCell string `protobuf:"bytes,3,opt,name=master_cell" json:"master_cell,omitempty"`
}

func (m *SrvShard) Reset()         { *m = SrvShard{} }
func (m *SrvShard) String() string { return proto.CompactTextString(m) }
func (*SrvShard) ProtoMessage()    {}

func (m *SrvShard) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

// ShardReference is used as a pointer from a SrvKeyspace to a SrvShard
type ShardReference struct {
	// Copied from Shard.
	Name     string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	KeyRange *KeyRange `protobuf:"bytes,2,opt,name=key_range" json:"key_range,omitempty"`
}

func (m *ShardReference) Reset()         { *m = ShardReference{} }
func (m *ShardReference) String() string { return proto.CompactTextString(m) }
func (*ShardReference) ProtoMessage()    {}

func (m *ShardReference) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

// SrvKeyspace is a rollup node for the keyspace itself.
type SrvKeyspace struct {
	// The partitions this keyspace is serving, per tablet type.
	Partitions []*SrvKeyspace_KeyspacePartition `protobuf:"bytes,1,rep,name=partitions" json:"partitions,omitempty"`
	// copied from Keyspace
	ShardingColumnName string                    `protobuf:"bytes,2,opt,name=sharding_column_name" json:"sharding_column_name,omitempty"`
	ShardingColumnType KeyspaceIdType            `protobuf:"varint,3,opt,name=sharding_column_type,enum=topo.KeyspaceIdType" json:"sharding_column_type,omitempty"`
	ServedFrom         []*SrvKeyspace_ServedFrom `protobuf:"bytes,4,rep,name=served_from" json:"served_from,omitempty"`
	SplitShardCount    int32                     `protobuf:"varint,5,opt,name=split_shard_count" json:"split_shard_count,omitempty"`
}

func (m *SrvKeyspace) Reset()         { *m = SrvKeyspace{} }
func (m *SrvKeyspace) String() string { return proto.CompactTextString(m) }
func (*SrvKeyspace) ProtoMessage()    {}

func (m *SrvKeyspace) GetPartitions() []*SrvKeyspace_KeyspacePartition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *SrvKeyspace) GetServedFrom() []*SrvKeyspace_ServedFrom {
	if m != nil {
		return m.ServedFrom
	}
	return nil
}

type SrvKeyspace_KeyspacePartition struct {
	// The type this partition applies to.
	ServedType TabletType `protobuf:"varint,1,opt,name=served_type,enum=topo.TabletType" json:"served_type,omitempty"`
	// List of non-overlapping continuous shards sorted by range.
	ShardReferences []*ShardReference `protobuf:"bytes,2,rep,name=shard_references" json:"shard_references,omitempty"`
}

func (m *SrvKeyspace_KeyspacePartition) Reset()         { *m = SrvKeyspace_KeyspacePartition{} }
func (m *SrvKeyspace_KeyspacePartition) String() string { return proto.CompactTextString(m) }
func (*SrvKeyspace_KeyspacePartition) ProtoMessage()    {}

func (m *SrvKeyspace_KeyspacePartition) GetShardReferences() []*ShardReference {
	if m != nil {
		return m.ShardReferences
	}
	return nil
}

// ServedFrom indicates a relationship between a TabletType and the
// keyspace name that's serving it.
type SrvKeyspace_ServedFrom struct {
	// the tablet type
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,enum=topo.TabletType" json:"tablet_type,omitempty"`
	// the keyspace name that's serving it
	Keyspace string `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
}

func (m *SrvKeyspace_ServedFrom) Reset()         { *m = SrvKeyspace_ServedFrom{} }
func (m *SrvKeyspace_ServedFrom) String() string { return proto.CompactTextString(m) }
func (*SrvKeyspace_ServedFrom) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("topo.KeyspaceIdType", KeyspaceIdType_name, KeyspaceIdType_value)
	proto.RegisterEnum("topo.TabletType", TabletType_name, TabletType_value)
}