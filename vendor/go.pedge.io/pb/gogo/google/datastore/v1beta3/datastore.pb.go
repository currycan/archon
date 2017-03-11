// Code generated by protoc-gen-gogo.
// source: google/datastore/v1beta3/datastore.proto
// DO NOT EDIT!

package google_datastore_v1beta3

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// Commit modes.
type CommitRequest_Mode int32

const (
	// Unspecified.
	CommitRequest_MODE_UNSPECIFIED CommitRequest_Mode = 0
	// Transactional.
	CommitRequest_TRANSACTIONAL CommitRequest_Mode = 1
	// Non-transactional.
	CommitRequest_NON_TRANSACTIONAL CommitRequest_Mode = 2
)

var CommitRequest_Mode_name = map[int32]string{
	0: "MODE_UNSPECIFIED",
	1: "TRANSACTIONAL",
	2: "NON_TRANSACTIONAL",
}
var CommitRequest_Mode_value = map[string]int32{
	"MODE_UNSPECIFIED":  0,
	"TRANSACTIONAL":     1,
	"NON_TRANSACTIONAL": 2,
}

func (x CommitRequest_Mode) String() string {
	return proto.EnumName(CommitRequest_Mode_name, int32(x))
}
func (CommitRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorDatastore, []int{8, 0}
}

// Read consistencies.
type ReadOptions_ReadConsistency int32

const (
	// Unspecified.
	ReadOptions_READ_CONSISTENCY_UNSPECIFIED ReadOptions_ReadConsistency = 0
	// Strong consistency.
	ReadOptions_STRONG ReadOptions_ReadConsistency = 1
	// Eventual consistency.
	ReadOptions_EVENTUAL ReadOptions_ReadConsistency = 2
)

var ReadOptions_ReadConsistency_name = map[int32]string{
	0: "READ_CONSISTENCY_UNSPECIFIED",
	1: "STRONG",
	2: "EVENTUAL",
}
var ReadOptions_ReadConsistency_value = map[string]int32{
	"READ_CONSISTENCY_UNSPECIFIED": 0,
	"STRONG":                       1,
	"EVENTUAL":                     2,
}

func (x ReadOptions_ReadConsistency) String() string {
	return proto.EnumName(ReadOptions_ReadConsistency_name, int32(x))
}
func (ReadOptions_ReadConsistency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorDatastore, []int{14, 0}
}

// The request for [google.datastore.v1beta3.Datastore.Lookup][google.datastore.v1beta3.Datastore.Lookup].
type LookupRequest struct {
	// Project ID against which to make the request.
	ProjectId string `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Options for this lookup request.
	ReadOptions *ReadOptions `protobuf:"bytes,1,opt,name=read_options,json=readOptions" json:"read_options,omitempty"`
	// Keys of entities to look up.
	Keys []*Key `protobuf:"bytes,3,rep,name=keys" json:"keys,omitempty"`
}

func (m *LookupRequest) Reset()                    { *m = LookupRequest{} }
func (m *LookupRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()               {}
func (*LookupRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{0} }

func (m *LookupRequest) GetReadOptions() *ReadOptions {
	if m != nil {
		return m.ReadOptions
	}
	return nil
}

func (m *LookupRequest) GetKeys() []*Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

// The response for [google.datastore.v1beta3.Datastore.Lookup][google.datastore.v1beta3.Datastore.Lookup].
type LookupResponse struct {
	// Entities found as `ResultType.FULL` entities. The order of results in this
	// field is undefined and has no relation to the order of the keys in the
	// input.
	Found []*EntityResult `protobuf:"bytes,1,rep,name=found" json:"found,omitempty"`
	// Entities not found as `ResultType.KEY_ONLY` entities. The order of results
	// in this field is undefined and has no relation to the order of the keys
	// in the input.
	Missing []*EntityResult `protobuf:"bytes,2,rep,name=missing" json:"missing,omitempty"`
	// A list of keys that were not looked up due to resource constraints. The
	// order of results in this field is undefined and has no relation to the
	// order of the keys in the input.
	Deferred []*Key `protobuf:"bytes,3,rep,name=deferred" json:"deferred,omitempty"`
}

func (m *LookupResponse) Reset()                    { *m = LookupResponse{} }
func (m *LookupResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()               {}
func (*LookupResponse) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{1} }

func (m *LookupResponse) GetFound() []*EntityResult {
	if m != nil {
		return m.Found
	}
	return nil
}

func (m *LookupResponse) GetMissing() []*EntityResult {
	if m != nil {
		return m.Missing
	}
	return nil
}

func (m *LookupResponse) GetDeferred() []*Key {
	if m != nil {
		return m.Deferred
	}
	return nil
}

// The request for [google.datastore.v1beta3.Datastore.RunQuery][google.datastore.v1beta3.Datastore.RunQuery].
type RunQueryRequest struct {
	// Project ID against which to make the request.
	ProjectId string `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Entities are partitioned into subsets, identified by a partition ID.
	// Queries are scoped to a single partition.
	// This partition ID is normalized with the standard default context
	// partition ID.
	PartitionId *PartitionId `protobuf:"bytes,2,opt,name=partition_id,json=partitionId" json:"partition_id,omitempty"`
	// The options for this query.
	ReadOptions *ReadOptions `protobuf:"bytes,1,opt,name=read_options,json=readOptions" json:"read_options,omitempty"`
	// The type of query.
	//
	// Types that are valid to be assigned to QueryType:
	//	*RunQueryRequest_Query
	//	*RunQueryRequest_GqlQuery
	QueryType isRunQueryRequest_QueryType `protobuf_oneof:"query_type"`
}

func (m *RunQueryRequest) Reset()                    { *m = RunQueryRequest{} }
func (m *RunQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*RunQueryRequest) ProtoMessage()               {}
func (*RunQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{2} }

type isRunQueryRequest_QueryType interface {
	isRunQueryRequest_QueryType()
}

type RunQueryRequest_Query struct {
	Query *Query `protobuf:"bytes,3,opt,name=query,oneof"`
}
type RunQueryRequest_GqlQuery struct {
	GqlQuery *GqlQuery `protobuf:"bytes,7,opt,name=gql_query,json=gqlQuery,oneof"`
}

func (*RunQueryRequest_Query) isRunQueryRequest_QueryType()    {}
func (*RunQueryRequest_GqlQuery) isRunQueryRequest_QueryType() {}

func (m *RunQueryRequest) GetQueryType() isRunQueryRequest_QueryType {
	if m != nil {
		return m.QueryType
	}
	return nil
}

func (m *RunQueryRequest) GetPartitionId() *PartitionId {
	if m != nil {
		return m.PartitionId
	}
	return nil
}

func (m *RunQueryRequest) GetReadOptions() *ReadOptions {
	if m != nil {
		return m.ReadOptions
	}
	return nil
}

func (m *RunQueryRequest) GetQuery() *Query {
	if x, ok := m.GetQueryType().(*RunQueryRequest_Query); ok {
		return x.Query
	}
	return nil
}

func (m *RunQueryRequest) GetGqlQuery() *GqlQuery {
	if x, ok := m.GetQueryType().(*RunQueryRequest_GqlQuery); ok {
		return x.GqlQuery
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RunQueryRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RunQueryRequest_OneofMarshaler, _RunQueryRequest_OneofUnmarshaler, _RunQueryRequest_OneofSizer, []interface{}{
		(*RunQueryRequest_Query)(nil),
		(*RunQueryRequest_GqlQuery)(nil),
	}
}

func _RunQueryRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RunQueryRequest)
	// query_type
	switch x := m.QueryType.(type) {
	case *RunQueryRequest_Query:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Query); err != nil {
			return err
		}
	case *RunQueryRequest_GqlQuery:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GqlQuery); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RunQueryRequest.QueryType has unexpected type %T", x)
	}
	return nil
}

func _RunQueryRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RunQueryRequest)
	switch tag {
	case 3: // query_type.query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Query)
		err := b.DecodeMessage(msg)
		m.QueryType = &RunQueryRequest_Query{msg}
		return true, err
	case 7: // query_type.gql_query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GqlQuery)
		err := b.DecodeMessage(msg)
		m.QueryType = &RunQueryRequest_GqlQuery{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RunQueryRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RunQueryRequest)
	// query_type
	switch x := m.QueryType.(type) {
	case *RunQueryRequest_Query:
		s := proto.Size(x.Query)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RunQueryRequest_GqlQuery:
		s := proto.Size(x.GqlQuery)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The response for [google.datastore.v1beta3.Datastore.RunQuery][google.datastore.v1beta3.Datastore.RunQuery].
type RunQueryResponse struct {
	// A batch of query results (always present).
	Batch *QueryResultBatch `protobuf:"bytes,1,opt,name=batch" json:"batch,omitempty"`
	// The parsed form of the `GqlQuery` from the request, if it was set.
	Query *Query `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
}

func (m *RunQueryResponse) Reset()                    { *m = RunQueryResponse{} }
func (m *RunQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*RunQueryResponse) ProtoMessage()               {}
func (*RunQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{3} }

func (m *RunQueryResponse) GetBatch() *QueryResultBatch {
	if m != nil {
		return m.Batch
	}
	return nil
}

func (m *RunQueryResponse) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

// The request for [google.datastore.v1beta3.Datastore.BeginTransaction][google.datastore.v1beta3.Datastore.BeginTransaction].
type BeginTransactionRequest struct {
	// Project ID against which to make the request.
	ProjectId string `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (m *BeginTransactionRequest) Reset()                    { *m = BeginTransactionRequest{} }
func (m *BeginTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*BeginTransactionRequest) ProtoMessage()               {}
func (*BeginTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{4} }

// The response for [google.datastore.v1beta3.Datastore.BeginTransaction][google.datastore.v1beta3.Datastore.BeginTransaction].
type BeginTransactionResponse struct {
	// The transaction identifier (always present).
	Transaction []byte `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (m *BeginTransactionResponse) Reset()         { *m = BeginTransactionResponse{} }
func (m *BeginTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*BeginTransactionResponse) ProtoMessage()    {}
func (*BeginTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorDatastore, []int{5}
}

// The request for [google.datastore.v1beta3.Datastore.Rollback][google.datastore.v1beta3.Datastore.Rollback].
type RollbackRequest struct {
	// Project ID against which to make the request.
	ProjectId string `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The transaction identifier, returned by a call to
	// [google.datastore.v1beta3.Datastore.BeginTransaction][google.datastore.v1beta3.Datastore.BeginTransaction].
	Transaction []byte `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (m *RollbackRequest) Reset()                    { *m = RollbackRequest{} }
func (m *RollbackRequest) String() string            { return proto.CompactTextString(m) }
func (*RollbackRequest) ProtoMessage()               {}
func (*RollbackRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{6} }

// The response for [google.datastore.v1beta3.Datastore.Rollback][google.datastore.v1beta3.Datastore.Rollback]
// (an empty message).
type RollbackResponse struct {
}

func (m *RollbackResponse) Reset()                    { *m = RollbackResponse{} }
func (m *RollbackResponse) String() string            { return proto.CompactTextString(m) }
func (*RollbackResponse) ProtoMessage()               {}
func (*RollbackResponse) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{7} }

// The request for [google.datastore.v1beta3.Datastore.Commit][google.datastore.v1beta3.Datastore.Commit].
type CommitRequest struct {
	// Project ID against which to make the request.
	ProjectId string `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The type of commit to perform. Defaults to `TRANSACTIONAL`.
	Mode CommitRequest_Mode `protobuf:"varint,5,opt,name=mode,proto3,enum=google.datastore.v1beta3.CommitRequest_Mode" json:"mode,omitempty"`
	// Must be set when mode is `TRANSACTIONAL`.
	//
	// Types that are valid to be assigned to TransactionSelector:
	//	*CommitRequest_Transaction
	TransactionSelector isCommitRequest_TransactionSelector `protobuf_oneof:"transaction_selector"`
	// The mutations to perform.
	//
	// When mode is `TRANSACTIONAL`, mutations affecting a single entity are
	// applied in order. The following sequences of mutations affecting a single
	// entity are not permitted in a single `Commit` request:
	// - `insert` followed by `insert`
	// - `update` followed by `insert`
	// - `upsert` followed by `insert`
	// - `delete` followed by `update`
	//
	// When mode is `NON_TRANSACTIONAL`, no two mutations may affect a single
	// entity.
	Mutations []*Mutation `protobuf:"bytes,6,rep,name=mutations" json:"mutations,omitempty"`
}

func (m *CommitRequest) Reset()                    { *m = CommitRequest{} }
func (m *CommitRequest) String() string            { return proto.CompactTextString(m) }
func (*CommitRequest) ProtoMessage()               {}
func (*CommitRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{8} }

type isCommitRequest_TransactionSelector interface {
	isCommitRequest_TransactionSelector()
}

type CommitRequest_Transaction struct {
	Transaction []byte `protobuf:"bytes,1,opt,name=transaction,proto3,oneof"`
}

func (*CommitRequest_Transaction) isCommitRequest_TransactionSelector() {}

func (m *CommitRequest) GetTransactionSelector() isCommitRequest_TransactionSelector {
	if m != nil {
		return m.TransactionSelector
	}
	return nil
}

func (m *CommitRequest) GetTransaction() []byte {
	if x, ok := m.GetTransactionSelector().(*CommitRequest_Transaction); ok {
		return x.Transaction
	}
	return nil
}

func (m *CommitRequest) GetMutations() []*Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CommitRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CommitRequest_OneofMarshaler, _CommitRequest_OneofUnmarshaler, _CommitRequest_OneofSizer, []interface{}{
		(*CommitRequest_Transaction)(nil),
	}
}

func _CommitRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CommitRequest)
	// transaction_selector
	switch x := m.TransactionSelector.(type) {
	case *CommitRequest_Transaction:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.Transaction)
	case nil:
	default:
		return fmt.Errorf("CommitRequest.TransactionSelector has unexpected type %T", x)
	}
	return nil
}

func _CommitRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CommitRequest)
	switch tag {
	case 1: // transaction_selector.transaction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.TransactionSelector = &CommitRequest_Transaction{x}
		return true, err
	default:
		return false, nil
	}
}

func _CommitRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CommitRequest)
	// transaction_selector
	switch x := m.TransactionSelector.(type) {
	case *CommitRequest_Transaction:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Transaction)))
		n += len(x.Transaction)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The response for [google.datastore.v1beta3.Datastore.Commit][google.datastore.v1beta3.Datastore.Commit].
type CommitResponse struct {
	// The result of performing the mutations.
	// The i-th mutation result corresponds to the i-th mutation in the request.
	MutationResults []*MutationResult `protobuf:"bytes,3,rep,name=mutation_results,json=mutationResults" json:"mutation_results,omitempty"`
	// The number of index entries updated during the commit.
	IndexUpdates int32 `protobuf:"varint,4,opt,name=index_updates,json=indexUpdates,proto3" json:"index_updates,omitempty"`
}

func (m *CommitResponse) Reset()                    { *m = CommitResponse{} }
func (m *CommitResponse) String() string            { return proto.CompactTextString(m) }
func (*CommitResponse) ProtoMessage()               {}
func (*CommitResponse) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{9} }

func (m *CommitResponse) GetMutationResults() []*MutationResult {
	if m != nil {
		return m.MutationResults
	}
	return nil
}

// The request for [google.datastore.v1beta3.Datastore.AllocateIds][google.datastore.v1beta3.Datastore.AllocateIds].
type AllocateIdsRequest struct {
	// Project ID against which to make the request.
	ProjectId string `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// A list of keys with incomplete key paths for which to allocate IDs.
	// No key may be reserved/read-only.
	Keys []*Key `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
}

func (m *AllocateIdsRequest) Reset()                    { *m = AllocateIdsRequest{} }
func (m *AllocateIdsRequest) String() string            { return proto.CompactTextString(m) }
func (*AllocateIdsRequest) ProtoMessage()               {}
func (*AllocateIdsRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{10} }

func (m *AllocateIdsRequest) GetKeys() []*Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

// The response for [google.datastore.v1beta3.Datastore.AllocateIds][google.datastore.v1beta3.Datastore.AllocateIds].
type AllocateIdsResponse struct {
	// The keys specified in the request (in the same order), each with
	// its key path completed with a newly allocated ID.
	Keys []*Key `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
}

func (m *AllocateIdsResponse) Reset()                    { *m = AllocateIdsResponse{} }
func (m *AllocateIdsResponse) String() string            { return proto.CompactTextString(m) }
func (*AllocateIdsResponse) ProtoMessage()               {}
func (*AllocateIdsResponse) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{11} }

func (m *AllocateIdsResponse) GetKeys() []*Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

// A mutation to apply to an entity.
type Mutation struct {
	// The mutation operation.
	//
	// For `insert`, `update`, and `upsert`:
	// - The entity's key must not be reserved/read-only.
	// - No property in the entity may have a reserved name,
	//   not even a property in an entity in a value.
	// - No value in the entity may have meaning 18,
	//   not even a value in an entity in another value.
	//
	// Types that are valid to be assigned to Operation:
	//	*Mutation_Insert
	//	*Mutation_Update
	//	*Mutation_Upsert
	//	*Mutation_Delete
	Operation isMutation_Operation `protobuf_oneof:"operation"`
}

func (m *Mutation) Reset()                    { *m = Mutation{} }
func (m *Mutation) String() string            { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()               {}
func (*Mutation) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{12} }

type isMutation_Operation interface {
	isMutation_Operation()
}

type Mutation_Insert struct {
	Insert *Entity `protobuf:"bytes,4,opt,name=insert,oneof"`
}
type Mutation_Update struct {
	Update *Entity `protobuf:"bytes,5,opt,name=update,oneof"`
}
type Mutation_Upsert struct {
	Upsert *Entity `protobuf:"bytes,6,opt,name=upsert,oneof"`
}
type Mutation_Delete struct {
	Delete *Key `protobuf:"bytes,7,opt,name=delete,oneof"`
}

func (*Mutation_Insert) isMutation_Operation() {}
func (*Mutation_Update) isMutation_Operation() {}
func (*Mutation_Upsert) isMutation_Operation() {}
func (*Mutation_Delete) isMutation_Operation() {}

func (m *Mutation) GetOperation() isMutation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *Mutation) GetInsert() *Entity {
	if x, ok := m.GetOperation().(*Mutation_Insert); ok {
		return x.Insert
	}
	return nil
}

func (m *Mutation) GetUpdate() *Entity {
	if x, ok := m.GetOperation().(*Mutation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *Mutation) GetUpsert() *Entity {
	if x, ok := m.GetOperation().(*Mutation_Upsert); ok {
		return x.Upsert
	}
	return nil
}

func (m *Mutation) GetDelete() *Key {
	if x, ok := m.GetOperation().(*Mutation_Delete); ok {
		return x.Delete
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Mutation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Mutation_OneofMarshaler, _Mutation_OneofUnmarshaler, _Mutation_OneofSizer, []interface{}{
		(*Mutation_Insert)(nil),
		(*Mutation_Update)(nil),
		(*Mutation_Upsert)(nil),
		(*Mutation_Delete)(nil),
	}
}

func _Mutation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Mutation)
	// operation
	switch x := m.Operation.(type) {
	case *Mutation_Insert:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Insert); err != nil {
			return err
		}
	case *Mutation_Update:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *Mutation_Upsert:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Upsert); err != nil {
			return err
		}
	case *Mutation_Delete:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Delete); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Mutation.Operation has unexpected type %T", x)
	}
	return nil
}

func _Mutation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Mutation)
	switch tag {
	case 4: // operation.insert
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Entity)
		err := b.DecodeMessage(msg)
		m.Operation = &Mutation_Insert{msg}
		return true, err
	case 5: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Entity)
		err := b.DecodeMessage(msg)
		m.Operation = &Mutation_Update{msg}
		return true, err
	case 6: // operation.upsert
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Entity)
		err := b.DecodeMessage(msg)
		m.Operation = &Mutation_Upsert{msg}
		return true, err
	case 7: // operation.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Key)
		err := b.DecodeMessage(msg)
		m.Operation = &Mutation_Delete{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Mutation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Mutation)
	// operation
	switch x := m.Operation.(type) {
	case *Mutation_Insert:
		s := proto.Size(x.Insert)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mutation_Update:
		s := proto.Size(x.Update)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mutation_Upsert:
		s := proto.Size(x.Upsert)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mutation_Delete:
		s := proto.Size(x.Delete)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The result of applying a mutation.
type MutationResult struct {
	// The automatically allocated key.
	// Set only when the mutation allocated a key.
	Key *Key `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
}

func (m *MutationResult) Reset()                    { *m = MutationResult{} }
func (m *MutationResult) String() string            { return proto.CompactTextString(m) }
func (*MutationResult) ProtoMessage()               {}
func (*MutationResult) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{13} }

func (m *MutationResult) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

// Options shared by read requests.
type ReadOptions struct {
	// If not specified, lookups and ancestor queries default to
	// `read_consistency`=`STRONG`, global queries default to
	// `read_consistency`=`EVENTUAL`.
	//
	// Types that are valid to be assigned to ConsistencyType:
	//	*ReadOptions_ReadConsistency_
	//	*ReadOptions_Transaction
	ConsistencyType isReadOptions_ConsistencyType `protobuf_oneof:"consistency_type"`
}

func (m *ReadOptions) Reset()                    { *m = ReadOptions{} }
func (m *ReadOptions) String() string            { return proto.CompactTextString(m) }
func (*ReadOptions) ProtoMessage()               {}
func (*ReadOptions) Descriptor() ([]byte, []int) { return fileDescriptorDatastore, []int{14} }

type isReadOptions_ConsistencyType interface {
	isReadOptions_ConsistencyType()
}

type ReadOptions_ReadConsistency_ struct {
	ReadConsistency ReadOptions_ReadConsistency `protobuf:"varint,1,opt,name=read_consistency,json=readConsistency,proto3,enum=google.datastore.v1beta3.ReadOptions_ReadConsistency,oneof"`
}
type ReadOptions_Transaction struct {
	Transaction []byte `protobuf:"bytes,2,opt,name=transaction,proto3,oneof"`
}

func (*ReadOptions_ReadConsistency_) isReadOptions_ConsistencyType() {}
func (*ReadOptions_Transaction) isReadOptions_ConsistencyType()      {}

func (m *ReadOptions) GetConsistencyType() isReadOptions_ConsistencyType {
	if m != nil {
		return m.ConsistencyType
	}
	return nil
}

func (m *ReadOptions) GetReadConsistency() ReadOptions_ReadConsistency {
	if x, ok := m.GetConsistencyType().(*ReadOptions_ReadConsistency_); ok {
		return x.ReadConsistency
	}
	return ReadOptions_READ_CONSISTENCY_UNSPECIFIED
}

func (m *ReadOptions) GetTransaction() []byte {
	if x, ok := m.GetConsistencyType().(*ReadOptions_Transaction); ok {
		return x.Transaction
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReadOptions) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReadOptions_OneofMarshaler, _ReadOptions_OneofUnmarshaler, _ReadOptions_OneofSizer, []interface{}{
		(*ReadOptions_ReadConsistency_)(nil),
		(*ReadOptions_Transaction)(nil),
	}
}

func _ReadOptions_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReadOptions)
	// consistency_type
	switch x := m.ConsistencyType.(type) {
	case *ReadOptions_ReadConsistency_:
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.ReadConsistency))
	case *ReadOptions_Transaction:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.Transaction)
	case nil:
	default:
		return fmt.Errorf("ReadOptions.ConsistencyType has unexpected type %T", x)
	}
	return nil
}

func _ReadOptions_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReadOptions)
	switch tag {
	case 1: // consistency_type.read_consistency
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConsistencyType = &ReadOptions_ReadConsistency_{ReadOptions_ReadConsistency(x)}
		return true, err
	case 2: // consistency_type.transaction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ConsistencyType = &ReadOptions_Transaction{x}
		return true, err
	default:
		return false, nil
	}
}

func _ReadOptions_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReadOptions)
	// consistency_type
	switch x := m.ConsistencyType.(type) {
	case *ReadOptions_ReadConsistency_:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.ReadConsistency))
	case *ReadOptions_Transaction:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Transaction)))
		n += len(x.Transaction)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*LookupRequest)(nil), "google.datastore.v1beta3.LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "google.datastore.v1beta3.LookupResponse")
	proto.RegisterType((*RunQueryRequest)(nil), "google.datastore.v1beta3.RunQueryRequest")
	proto.RegisterType((*RunQueryResponse)(nil), "google.datastore.v1beta3.RunQueryResponse")
	proto.RegisterType((*BeginTransactionRequest)(nil), "google.datastore.v1beta3.BeginTransactionRequest")
	proto.RegisterType((*BeginTransactionResponse)(nil), "google.datastore.v1beta3.BeginTransactionResponse")
	proto.RegisterType((*RollbackRequest)(nil), "google.datastore.v1beta3.RollbackRequest")
	proto.RegisterType((*RollbackResponse)(nil), "google.datastore.v1beta3.RollbackResponse")
	proto.RegisterType((*CommitRequest)(nil), "google.datastore.v1beta3.CommitRequest")
	proto.RegisterType((*CommitResponse)(nil), "google.datastore.v1beta3.CommitResponse")
	proto.RegisterType((*AllocateIdsRequest)(nil), "google.datastore.v1beta3.AllocateIdsRequest")
	proto.RegisterType((*AllocateIdsResponse)(nil), "google.datastore.v1beta3.AllocateIdsResponse")
	proto.RegisterType((*Mutation)(nil), "google.datastore.v1beta3.Mutation")
	proto.RegisterType((*MutationResult)(nil), "google.datastore.v1beta3.MutationResult")
	proto.RegisterType((*ReadOptions)(nil), "google.datastore.v1beta3.ReadOptions")
	proto.RegisterEnum("google.datastore.v1beta3.CommitRequest_Mode", CommitRequest_Mode_name, CommitRequest_Mode_value)
	proto.RegisterEnum("google.datastore.v1beta3.ReadOptions_ReadConsistency", ReadOptions_ReadConsistency_name, ReadOptions_ReadConsistency_value)
}

var fileDescriptorDatastore = []byte{
	// 1086 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x57, 0x5f, 0x73, 0xe2, 0x54,
	0x14, 0xdf, 0xd0, 0x42, 0xe1, 0x40, 0x29, 0x7b, 0x5d, 0x95, 0x61, 0xba, 0x23, 0x73, 0xb5, 0x6e,
	0x65, 0x56, 0xd8, 0xb2, 0x76, 0x56, 0x3b, 0xfb, 0x50, 0xa0, 0xec, 0x96, 0x71, 0x0b, 0x35, 0x50,
	0x67, 0x7c, 0xca, 0x04, 0x72, 0x8b, 0xb1, 0x21, 0x49, 0x93, 0xe0, 0xc8, 0x38, 0xbe, 0x38, 0x3e,
	0xe9, 0xf8, 0xe4, 0xf8, 0x01, 0x7c, 0xf6, 0x59, 0xfd, 0x10, 0xce, 0xf8, 0xe2, 0x57, 0xf0, 0x43,
	0xf8, 0xe8, 0xcd, 0xcd, 0x0d, 0x21, 0x6c, 0x03, 0xa9, 0xb3, 0x6f, 0xe4, 0x70, 0x7e, 0xe7, 0xfc,
	0xce, 0xff, 0x04, 0xf6, 0xc7, 0x86, 0x31, 0xd6, 0x48, 0x4d, 0x91, 0x1d, 0xd9, 0x76, 0x0c, 0x8b,
	0xd4, 0xbe, 0x3c, 0x18, 0x12, 0x47, 0x7e, 0x1c, 0x48, 0xaa, 0xa6, 0x65, 0x38, 0x06, 0x2a, 0x7a,
	0x9a, 0xd5, 0x40, 0xce, 0x35, 0x4b, 0xbb, 0xdc, 0x86, 0x6c, 0xaa, 0x35, 0x59, 0xd7, 0x0d, 0x47,
	0x76, 0x54, 0x43, 0xb7, 0x3d, 0x5c, 0x69, 0x2f, 0xd2, 0x03, 0xd1, 0x1d, 0xd5, 0x99, 0x71, 0xb5,
	0x77, 0x22, 0xd5, 0xae, 0xa7, 0xc4, 0xe2, 0x5a, 0xf8, 0x57, 0x01, 0xb6, 0x5f, 0x18, 0xc6, 0xd5,
	0xd4, 0x14, 0x09, 0x95, 0xdb, 0x0e, 0xba, 0x0f, 0x40, 0xff, 0xfa, 0x82, 0x8c, 0x1c, 0x49, 0x55,
	0x8a, 0xe9, 0xb2, 0xb0, 0x9f, 0x11, 0x33, 0x5c, 0xd2, 0x51, 0xd0, 0x29, 0xe4, 0x2c, 0x22, 0x2b,
	0x92, 0x61, 0x32, 0x4e, 0x45, 0x81, 0x2a, 0x64, 0xeb, 0x7b, 0xd5, 0xa8, 0x60, 0xaa, 0x22, 0xd5,
	0xee, 0x79, 0xca, 0x62, 0xd6, 0x0a, 0x1e, 0xd0, 0x01, 0x6c, 0x5e, 0x91, 0x99, 0x5d, 0xdc, 0x28,
	0x6f, 0x50, 0x0b, 0xf7, 0xa3, 0x2d, 0x7c, 0x4c, 0x66, 0x22, 0x53, 0xc5, 0x7f, 0x0a, 0x90, 0xf7,
	0xd9, 0xda, 0x26, 0x35, 0x42, 0xd0, 0x53, 0x48, 0x5e, 0x1a, 0x53, 0x5d, 0xa1, 0x44, 0x5c, 0x33,
	0xef, 0x46, 0x9b, 0x69, 0xb3, 0xec, 0x50, 0xe0, 0x54, 0x73, 0x44, 0x0f, 0x84, 0x8e, 0x61, 0x6b,
	0xa2, 0xda, 0xb6, 0xaa, 0x8f, 0x8b, 0x89, 0x5b, 0xe1, 0x7d, 0x18, 0xfa, 0x08, 0xd2, 0x0a, 0xb9,
	0x24, 0x96, 0x45, 0x94, 0x78, 0x91, 0xcc, 0xd5, 0xf1, 0x5f, 0x09, 0xd8, 0x11, 0xa7, 0xfa, 0x27,
	0x6e, 0x39, 0xe2, 0x67, 0xdf, 0x94, 0x2d, 0xca, 0x83, 0x66, 0xd0, 0x55, 0x48, 0xac, 0xcb, 0xfe,
	0xb9, 0xaf, 0xdd, 0x51, 0xc4, 0xac, 0x19, 0x3c, 0xbc, 0xc2, 0x3a, 0x3e, 0x81, 0x24, 0xeb, 0x28,
	0x1a, 0xbe, 0x6b, 0xe2, 0xad, 0x68, 0x13, 0x2c, 0xd2, 0xd3, 0x3b, 0xa2, 0xa7, 0x8f, 0x1a, 0x90,
	0x19, 0x5f, 0x6b, 0x92, 0x07, 0xde, 0x62, 0x60, 0x1c, 0x0d, 0x7e, 0x7e, 0xad, 0xf9, 0xf8, 0xf4,
	0x98, 0xff, 0x6e, 0xe6, 0x00, 0x18, 0x5c, 0x72, 0x66, 0x26, 0xc1, 0x3f, 0x08, 0x50, 0x08, 0x12,
	0xca, 0x1b, 0xe4, 0x18, 0x92, 0x43, 0xd9, 0x19, 0x7d, 0xce, 0x23, 0xac, 0xac, 0xa1, 0xe7, 0xd5,
	0xb7, 0xe9, 0x22, 0x44, 0x0f, 0x88, 0x0e, 0xfd, 0x00, 0x13, 0xb1, 0x02, 0xe4, 0xe1, 0xe1, 0x0f,
	0xe1, 0xcd, 0x26, 0x19, 0xab, 0xfa, 0xc0, 0x92, 0x75, 0x5b, 0x1e, 0xb9, 0xc9, 0x8a, 0x57, 0x65,
	0xfc, 0x14, 0x8a, 0x2f, 0x23, 0x79, 0x38, 0x65, 0xc8, 0x3a, 0x81, 0x98, 0x05, 0x95, 0x13, 0x17,
	0x45, 0x58, 0xa4, 0x5d, 0x65, 0x68, 0xda, 0x50, 0x1e, 0x5d, 0xc5, 0xec, 0xaa, 0xf5, 0x36, 0x11,
	0x4d, 0xec, 0xdc, 0xa6, 0xc7, 0x04, 0xff, 0x9e, 0x80, 0xed, 0x96, 0x31, 0x99, 0xa8, 0x4e, 0x4c,
	0x37, 0xc7, 0xb0, 0x39, 0x31, 0x14, 0x52, 0x4c, 0xd2, 0x3f, 0xf2, 0xf5, 0x87, 0xd1, 0x69, 0x0c,
	0x59, 0xad, 0x9e, 0x51, 0x8c, 0xc8, 0x90, 0x08, 0xdf, 0x40, 0x94, 0xf6, 0xc3, 0xa2, 0x90, 0x7a,
	0xc9, 0x4c, 0xa6, 0x7c, 0x63, 0x16, 0x53, 0x6c, 0x22, 0x57, 0x74, 0xd5, 0x19, 0x57, 0x15, 0x03,
	0x10, 0x7e, 0x06, 0x9b, 0xae, 0x4f, 0x74, 0x0f, 0x0a, 0x67, 0xbd, 0x93, 0xb6, 0x74, 0xd1, 0xed,
	0x9f, 0xb7, 0x5b, 0x9d, 0x67, 0x9d, 0xf6, 0x49, 0xe1, 0x0e, 0xba, 0x0b, 0xdb, 0x03, 0xb1, 0xd1,
	0xed, 0x37, 0x5a, 0x83, 0x4e, 0xaf, 0xdb, 0x78, 0x51, 0x10, 0xd0, 0xeb, 0x70, 0xb7, 0xdb, 0xeb,
	0x4a, 0x61, 0x71, 0xa2, 0xf9, 0x06, 0xdc, 0x5b, 0x20, 0x26, 0xd9, 0x44, 0xa3, 0x79, 0x30, 0x2c,
	0xfc, 0x3d, 0xdd, 0x62, 0x7e, 0x88, 0xbc, 0xaa, 0x7d, 0x28, 0xf8, 0xfe, 0x25, 0x8b, 0x75, 0xa0,
	0xbf, 0x17, 0xf7, 0x63, 0x70, 0xf7, 0x56, 0xd2, 0xce, 0x24, 0xf4, 0x6c, 0xa3, 0xb7, 0x61, 0x5b,
	0xd5, 0x15, 0xf2, 0x95, 0x34, 0x35, 0x29, 0x98, 0xd8, 0xc5, 0x4d, 0x9a, 0xaf, 0xa4, 0x98, 0x63,
	0xc2, 0x0b, 0x4f, 0x86, 0x2f, 0x01, 0x35, 0x34, 0xcd, 0x18, 0xd1, 0x87, 0x8e, 0x62, 0xc7, 0xac,
	0xa4, 0xbf, 0xba, 0x85, 0xf8, 0xab, 0xfb, 0x14, 0x5e, 0x0b, 0xf9, 0xe1, 0x81, 0xff, 0x0f, 0x4b,
	0xdf, 0x25, 0x20, 0xed, 0x87, 0x8e, 0x8e, 0x20, 0xa5, 0x52, 0x3b, 0x96, 0xc3, 0x82, 0xcb, 0xd6,
	0xcb, 0xeb, 0xf6, 0x37, 0x6d, 0x17, 0x8e, 0x70, 0xb1, 0x5e, 0x66, 0x58, 0x47, 0xc6, 0xc4, 0x7a,
	0x08, 0x0f, 0xcb, 0xfc, 0xa6, 0x6e, 0x83, 0x65, 0x7e, 0x9f, 0x40, 0x4a, 0xa1, 0xbd, 0x40, 0xfd,
	0x7a, 0x4b, 0x6f, 0x75, 0xd4, 0x2e, 0xd0, 0x53, 0x6f, 0x66, 0x21, 0x63, 0x98, 0xc4, 0x62, 0x91,
	0xe3, 0x06, 0xe4, 0xc3, 0x0d, 0x80, 0x6a, 0xb0, 0x41, 0x13, 0xc4, 0xd7, 0xf0, 0x9a, 0x54, 0xba,
	0x9a, 0xf8, 0x5f, 0x01, 0xb2, 0x0b, 0x6b, 0x1d, 0x0d, 0xa1, 0xc0, 0x6e, 0xc2, 0x88, 0x3e, 0xa8,
	0xb6, 0x43, 0xf4, 0xd1, 0x8c, 0xcd, 0x58, 0xbe, 0x7e, 0x18, 0xeb, 0x2e, 0xb0, 0xdf, 0xad, 0x00,
	0x4c, 0xa9, 0xef, 0x58, 0x61, 0xd1, 0xf2, 0x08, 0x27, 0x6e, 0x18, 0x61, 0x7c, 0x46, 0x37, 0xd8,
	0x12, 0xac, 0x0c, 0xbb, 0x62, 0xbb, 0x71, 0x22, 0xb5, 0x7a, 0xdd, 0x7e, 0xa7, 0x3f, 0x68, 0x77,
	0x5b, 0x9f, 0x2d, 0xcd, 0x25, 0x40, 0xaa, 0x3f, 0x10, 0x7b, 0xdd, 0xe7, 0x74, 0x20, 0x73, 0x90,
	0x6e, 0x7f, 0xda, 0xee, 0x0e, 0x2e, 0xd8, 0x1c, 0xd2, 0xe5, 0xb5, 0x10, 0x11, 0x3b, 0x15, 0xf5,
	0x3f, 0xb6, 0x20, 0x73, 0xe2, 0xc7, 0x82, 0x7e, 0x14, 0x20, 0xe5, 0xbd, 0x57, 0xa0, 0x07, 0xd1,
	0x91, 0x86, 0xde, 0x93, 0x4a, 0xfb, 0xeb, 0x15, 0xf9, 0xa2, 0x7c, 0xf4, 0xed, 0xdf, 0xff, 0xfc,
	0x94, 0xa8, 0xe0, 0xbd, 0xf9, 0x1b, 0x18, 0x9f, 0x24, 0xbb, 0xf6, 0x75, 0x30, 0x65, 0xdf, 0x1c,
	0x69, 0x0c, 0x76, 0x24, 0x54, 0xd0, 0xcf, 0x02, 0xa4, 0xfd, 0x43, 0x86, 0xde, 0x5b, 0x91, 0xfb,
	0xf0, 0xdb, 0x43, 0xa9, 0x12, 0x47, 0x95, 0xb3, 0xaa, 0x33, 0x56, 0x0f, 0xf1, 0x83, 0x35, 0xac,
	0x2c, 0x0e, 0x74, 0x79, 0xfd, 0x46, 0x0f, 0xec, 0xf2, 0x65, 0x42, 0x07, 0xd1, 0x4e, 0x23, 0xee,
	0x5f, 0xa9, 0x7e, 0x1b, 0x08, 0xe7, 0x7b, 0xc4, 0xf8, 0x7e, 0x80, 0x6b, 0x6b, 0xf8, 0x0e, 0x97,
	0x0c, 0xb8, 0xbc, 0xdd, 0xfa, 0x7a, 0x1b, 0x77, 0x55, 0x7d, 0x43, 0x67, 0x67, 0x55, 0x7d, 0xc3,
	0xcb, 0x3b, 0x76, 0x7d, 0x47, 0x0c, 0x36, 0xaf, 0x2f, 0xbf, 0xa7, 0x2b, 0xeb, 0x1b, 0xbe, 0xe3,
	0x2b, 0xeb, 0xbb, 0x7c, 0x9e, 0x63, 0xd7, 0x97, 0x03, 0x5d, 0x5e, 0xbf, 0xd0, 0x85, 0xb0, 0xb0,
	0xa5, 0xd1, 0x8a, 0x1b, 0xfd, 0xf2, 0xd1, 0x28, 0xbd, 0x1f, 0x53, 0x9b, 0x13, 0x3c, 0x64, 0x04,
	0x6b, 0xb8, 0xb2, 0x86, 0xa0, 0x1c, 0x60, 0x29, 0xc7, 0xe6, 0x23, 0xd8, 0xa5, 0x89, 0x8c, 0x74,
	0xd5, 0xcc, 0xcf, 0xc7, 0xfa, 0xdc, 0xfd, 0xc2, 0x39, 0x17, 0x86, 0x29, 0xf6, 0xa9, 0xf3, 0xf8,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xa9, 0xa3, 0x78, 0x9b, 0x0d, 0x00, 0x00,
}