// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/sql/jobs/jobs.proto
// DO NOT EDIT!

/*
	Package jobs is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/sql/jobs/jobs.proto

	It has these top-level messages:
		BackupDetails
		RestoreDetails
		ResumeSpanList
		SchemaChangeJobDetails
		Payload
*/
package jobs

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/pkg/roachpb"

import github_com_cockroachdb_cockroach_pkg_sql_sqlbase "github.com/cockroachdb/cockroach/pkg/sql/sqlbase"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BackupDetails struct {
}

func (m *BackupDetails) Reset()                    { *m = BackupDetails{} }
func (m *BackupDetails) String() string            { return proto.CompactTextString(m) }
func (*BackupDetails) ProtoMessage()               {}
func (*BackupDetails) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{0} }

type RestoreDetails struct {
	LowWaterMark []byte `protobuf:"bytes,1,opt,name=low_water_mark,json=lowWaterMark,proto3" json:"low_water_mark,omitempty"`
}

func (m *RestoreDetails) Reset()                    { *m = RestoreDetails{} }
func (m *RestoreDetails) String() string            { return proto.CompactTextString(m) }
func (*RestoreDetails) ProtoMessage()               {}
func (*RestoreDetails) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{1} }

type ResumeSpanList struct {
	ResumeSpans []cockroach_roachpb1.Span `protobuf:"bytes,1,rep,name=resume_spans,json=resumeSpans" json:"resume_spans"`
}

func (m *ResumeSpanList) Reset()                    { *m = ResumeSpanList{} }
func (m *ResumeSpanList) String() string            { return proto.CompactTextString(m) }
func (*ResumeSpanList) ProtoMessage()               {}
func (*ResumeSpanList) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{2} }

type SchemaChangeJobDetails struct {
	ReadAsOf int64 `protobuf:"varint,1,opt,name=readAsOf,proto3" json:"readAsOf,omitempty"`
	// A schema change can involve running multiple processors backfilling
	// or deleting data. They occasionally checkpoint Spans so that the
	// processing can resume in the event of a node failure. The spans are
	// non-overlapping contiguous areas of the KV space that still need to
	// be processed.
	ResumeSpanList []ResumeSpanList `protobuf:"bytes,2,rep,name=resume_span_list,json=resumeSpanList" json:"resume_span_list"`
}

func (m *SchemaChangeJobDetails) Reset()                    { *m = SchemaChangeJobDetails{} }
func (m *SchemaChangeJobDetails) String() string            { return proto.CompactTextString(m) }
func (*SchemaChangeJobDetails) ProtoMessage()               {}
func (*SchemaChangeJobDetails) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{3} }

type Payload struct {
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// For consistency with the SQL timestamp type, which has microsecond
	// precision, we avoid the timestamp.Timestamp WKT, which has nanosecond
	// precision, and use microsecond integers directly.
	StartedMicros     int64                                                 `protobuf:"varint,3,opt,name=started_micros,json=startedMicros,proto3" json:"started_micros,omitempty"`
	FinishedMicros    int64                                                 `protobuf:"varint,4,opt,name=finished_micros,json=finishedMicros,proto3" json:"finished_micros,omitempty"`
	ModifiedMicros    int64                                                 `protobuf:"varint,5,opt,name=modified_micros,json=modifiedMicros,proto3" json:"modified_micros,omitempty"`
	DescriptorIDs     []github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID `protobuf:"varint,6,rep,packed,name=descriptor_ids,json=descriptorIds,casttype=github.com/cockroachdb/cockroach/pkg/sql/sqlbase.ID" json:"descriptor_ids,omitempty"`
	FractionCompleted float32                                               `protobuf:"fixed32,7,opt,name=fraction_completed,json=fractionCompleted,proto3" json:"fraction_completed,omitempty"`
	Error             string                                                `protobuf:"bytes,8,opt,name=error,proto3" json:"error,omitempty"`
	// Types that are valid to be assigned to Details:
	//	*Payload_Backup
	//	*Payload_Restore
	//	*Payload_SchemaChange
	Details isPayload_Details `protobuf_oneof:"details"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{4} }

type isPayload_Details interface {
	isPayload_Details()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Payload_Backup struct {
	Backup *BackupDetails `protobuf:"bytes,10,opt,name=backup,oneof"`
}
type Payload_Restore struct {
	Restore *RestoreDetails `protobuf:"bytes,11,opt,name=restore,oneof"`
}
type Payload_SchemaChange struct {
	SchemaChange *SchemaChangeJobDetails `protobuf:"bytes,12,opt,name=schemaChange,oneof"`
}

func (*Payload_Backup) isPayload_Details()       {}
func (*Payload_Restore) isPayload_Details()      {}
func (*Payload_SchemaChange) isPayload_Details() {}

func (m *Payload) GetDetails() isPayload_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Payload) GetBackup() *BackupDetails {
	if x, ok := m.GetDetails().(*Payload_Backup); ok {
		return x.Backup
	}
	return nil
}

func (m *Payload) GetRestore() *RestoreDetails {
	if x, ok := m.GetDetails().(*Payload_Restore); ok {
		return x.Restore
	}
	return nil
}

func (m *Payload) GetSchemaChange() *SchemaChangeJobDetails {
	if x, ok := m.GetDetails().(*Payload_SchemaChange); ok {
		return x.SchemaChange
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Payload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Payload_OneofMarshaler, _Payload_OneofUnmarshaler, _Payload_OneofSizer, []interface{}{
		(*Payload_Backup)(nil),
		(*Payload_Restore)(nil),
		(*Payload_SchemaChange)(nil),
	}
}

func _Payload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Payload)
	// details
	switch x := m.Details.(type) {
	case *Payload_Backup:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Backup); err != nil {
			return err
		}
	case *Payload_Restore:
		_ = b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Restore); err != nil {
			return err
		}
	case *Payload_SchemaChange:
		_ = b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SchemaChange); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Payload.Details has unexpected type %T", x)
	}
	return nil
}

func _Payload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Payload)
	switch tag {
	case 10: // details.backup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BackupDetails)
		err := b.DecodeMessage(msg)
		m.Details = &Payload_Backup{msg}
		return true, err
	case 11: // details.restore
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RestoreDetails)
		err := b.DecodeMessage(msg)
		m.Details = &Payload_Restore{msg}
		return true, err
	case 12: // details.schemaChange
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SchemaChangeJobDetails)
		err := b.DecodeMessage(msg)
		m.Details = &Payload_SchemaChange{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Payload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Payload)
	// details
	switch x := m.Details.(type) {
	case *Payload_Backup:
		s := proto.Size(x.Backup)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Payload_Restore:
		s := proto.Size(x.Restore)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Payload_SchemaChange:
		s := proto.Size(x.SchemaChange)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BackupDetails)(nil), "cockroach.sql.jobs.BackupDetails")
	proto.RegisterType((*RestoreDetails)(nil), "cockroach.sql.jobs.RestoreDetails")
	proto.RegisterType((*ResumeSpanList)(nil), "cockroach.sql.jobs.ResumeSpanList")
	proto.RegisterType((*SchemaChangeJobDetails)(nil), "cockroach.sql.jobs.SchemaChangeJobDetails")
	proto.RegisterType((*Payload)(nil), "cockroach.sql.jobs.Payload")
}
func (m *BackupDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *RestoreDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestoreDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LowWaterMark) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.LowWaterMark)))
		i += copy(dAtA[i:], m.LowWaterMark)
	}
	return i, nil
}

func (m *ResumeSpanList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResumeSpanList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ResumeSpans) > 0 {
		for _, msg := range m.ResumeSpans {
			dAtA[i] = 0xa
			i++
			i = encodeVarintJobs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SchemaChangeJobDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaChangeJobDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReadAsOf != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.ReadAsOf))
	}
	if len(m.ResumeSpanList) > 0 {
		for _, msg := range m.ResumeSpanList {
			dAtA[i] = 0x12
			i++
			i = encodeVarintJobs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if m.StartedMicros != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.StartedMicros))
	}
	if m.FinishedMicros != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.FinishedMicros))
	}
	if m.ModifiedMicros != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.ModifiedMicros))
	}
	if len(m.DescriptorIDs) > 0 {
		dAtA2 := make([]byte, len(m.DescriptorIDs)*10)
		var j1 int
		for _, num := range m.DescriptorIDs {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x32
		i++
		i = encodeVarintJobs(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.FractionCompleted != 0 {
		dAtA[i] = 0x3d
		i++
		i = encodeFixed32Jobs(dAtA, i, uint32(math.Float32bits(float32(m.FractionCompleted))))
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if m.Details != nil {
		nn3, err := m.Details.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	return i, nil
}

func (m *Payload_Backup) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Backup != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.Backup.Size()))
		n4, err := m.Backup.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *Payload_Restore) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Restore != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.Restore.Size()))
		n5, err := m.Restore.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *Payload_SchemaChange) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.SchemaChange != nil {
		dAtA[i] = 0x62
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.SchemaChange.Size()))
		n6, err := m.SchemaChange.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func encodeFixed64Jobs(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Jobs(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintJobs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BackupDetails) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *RestoreDetails) Size() (n int) {
	var l int
	_ = l
	l = len(m.LowWaterMark)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}

func (m *ResumeSpanList) Size() (n int) {
	var l int
	_ = l
	if len(m.ResumeSpans) > 0 {
		for _, e := range m.ResumeSpans {
			l = e.Size()
			n += 1 + l + sovJobs(uint64(l))
		}
	}
	return n
}

func (m *SchemaChangeJobDetails) Size() (n int) {
	var l int
	_ = l
	if m.ReadAsOf != 0 {
		n += 1 + sovJobs(uint64(m.ReadAsOf))
	}
	if len(m.ResumeSpanList) > 0 {
		for _, e := range m.ResumeSpanList {
			l = e.Size()
			n += 1 + l + sovJobs(uint64(l))
		}
	}
	return n
}

func (m *Payload) Size() (n int) {
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	if m.StartedMicros != 0 {
		n += 1 + sovJobs(uint64(m.StartedMicros))
	}
	if m.FinishedMicros != 0 {
		n += 1 + sovJobs(uint64(m.FinishedMicros))
	}
	if m.ModifiedMicros != 0 {
		n += 1 + sovJobs(uint64(m.ModifiedMicros))
	}
	if len(m.DescriptorIDs) > 0 {
		l = 0
		for _, e := range m.DescriptorIDs {
			l += sovJobs(uint64(e))
		}
		n += 1 + sovJobs(uint64(l)) + l
	}
	if m.FractionCompleted != 0 {
		n += 5
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	if m.Details != nil {
		n += m.Details.Size()
	}
	return n
}

func (m *Payload_Backup) Size() (n int) {
	var l int
	_ = l
	if m.Backup != nil {
		l = m.Backup.Size()
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}
func (m *Payload_Restore) Size() (n int) {
	var l int
	_ = l
	if m.Restore != nil {
		l = m.Restore.Size()
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}
func (m *Payload_SchemaChange) Size() (n int) {
	var l int
	_ = l
	if m.SchemaChange != nil {
		l = m.SchemaChange.Size()
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}

func sovJobs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozJobs(x uint64) (n int) {
	return sovJobs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BackupDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RestoreDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RestoreDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestoreDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowWaterMark", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LowWaterMark = append(m.LowWaterMark[:0], dAtA[iNdEx:postIndex]...)
			if m.LowWaterMark == nil {
				m.LowWaterMark = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResumeSpanList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResumeSpanList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResumeSpanList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResumeSpans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResumeSpans = append(m.ResumeSpans, cockroach_roachpb1.Span{})
			if err := m.ResumeSpans[len(m.ResumeSpans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SchemaChangeJobDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaChangeJobDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaChangeJobDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadAsOf", wireType)
			}
			m.ReadAsOf = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReadAsOf |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResumeSpanList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResumeSpanList = append(m.ResumeSpanList, ResumeSpanList{})
			if err := m.ResumeSpanList[len(m.ResumeSpanList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedMicros", wireType)
			}
			m.StartedMicros = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartedMicros |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinishedMicros", wireType)
			}
			m.FinishedMicros = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinishedMicros |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModifiedMicros", wireType)
			}
			m.ModifiedMicros = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModifiedMicros |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType == 0 {
				var v github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowJobs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.DescriptorIDs = append(m.DescriptorIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowJobs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthJobs
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowJobs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.DescriptorIDs = append(m.DescriptorIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DescriptorIDs", wireType)
			}
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FractionCompleted", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.FractionCompleted = float32(math.Float32frombits(v))
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backup", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BackupDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &Payload_Backup{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Restore", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RestoreDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &Payload_Restore{v}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaChange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SchemaChangeJobDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &Payload_SchemaChange{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipJobs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthJobs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJobs
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipJobs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthJobs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJobs   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/sql/jobs/jobs.proto", fileDescriptorJobs) }

var fileDescriptorJobs = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x4d, 0xd6, 0x6e, 0xdd, 0xdc, 0x36, 0x1b, 0xd6, 0x04, 0xd1, 0x0e, 0x59, 0x16, 0x81, 0xa8,
	0x90, 0x48, 0xa4, 0x4d, 0xe2, 0x82, 0x84, 0xa0, 0xeb, 0xa1, 0x9b, 0x98, 0x98, 0xbc, 0x03, 0x12,
	0x97, 0xc8, 0x49, 0xdc, 0xd6, 0x34, 0x89, 0x33, 0xdb, 0xd5, 0xc4, 0x37, 0xe0, 0xc8, 0xc7, 0xda,
	0x91, 0x13, 0xe2, 0x54, 0x41, 0xf9, 0x16, 0x9c, 0x50, 0x9c, 0x3f, 0x6d, 0xb5, 0x72, 0xa9, 0xec,
	0xf7, 0x7b, 0xef, 0xf9, 0xe9, 0xd7, 0x17, 0x70, 0x12, 0xb2, 0x70, 0xca, 0x19, 0x0e, 0x27, 0x5e,
	0x36, 0x1d, 0x7b, 0xe2, 0x36, 0xf6, 0x3e, 0xb3, 0x40, 0xa8, 0x1f, 0x37, 0xe3, 0x4c, 0x32, 0x08,
	0x6b, 0x8a, 0x2b, 0x6e, 0x63, 0x37, 0x9f, 0x1c, 0xd9, 0xeb, 0x32, 0x75, 0xca, 0x02, 0x2f, 0xc2,
	0x12, 0x17, 0xaa, 0xa3, 0xc3, 0x31, 0x1b, 0x33, 0x75, 0xf4, 0xf2, 0x53, 0x81, 0x3a, 0xfb, 0xa0,
	0xdb, 0xc7, 0xe1, 0x74, 0x96, 0x0d, 0x88, 0xc4, 0x34, 0x16, 0xce, 0x2b, 0x60, 0x20, 0x22, 0x24,
	0xe3, 0xa4, 0x44, 0xe0, 0x53, 0x60, 0xc4, 0xec, 0xce, 0xbf, 0xc3, 0x92, 0x70, 0x3f, 0xc1, 0x7c,
	0x6a, 0xea, 0xb6, 0xde, 0xeb, 0xa0, 0x4e, 0xcc, 0xee, 0x3e, 0xe6, 0xe0, 0x15, 0xe6, 0x53, 0x07,
	0x29, 0xdd, 0x2c, 0x21, 0x37, 0x19, 0x4e, 0xdf, 0x53, 0x21, 0xe1, 0x5b, 0xd0, 0xe1, 0x0a, 0xf1,
	0x45, 0x86, 0x53, 0x61, 0xea, 0x76, 0xa3, 0xd7, 0x3e, 0x7d, 0xe2, 0x2e, 0xd3, 0x97, 0x29, 0xdd,
	0x5c, 0xd2, 0x6f, 0xde, 0xcf, 0x8f, 0x35, 0xd4, 0xe6, 0xb5, 0x89, 0x70, 0xbe, 0xea, 0xe0, 0xf1,
	0x4d, 0x38, 0x21, 0x09, 0x3e, 0x9f, 0xe0, 0x74, 0x4c, 0x2e, 0x59, 0x50, 0x85, 0x3a, 0x02, 0xbb,
	0x9c, 0xe0, 0xe8, 0x9d, 0xf8, 0x30, 0x52, 0x71, 0x1a, 0xa8, 0xbe, 0x43, 0x04, 0x0e, 0x56, 0x1e,
	0xf6, 0x63, 0x2a, 0xa4, 0xb9, 0xa5, 0x1e, 0x77, 0xdc, 0x87, 0xab, 0x73, 0xd7, 0x63, 0x97, 0x39,
	0x0c, 0xbe, 0x86, 0x3a, 0x3f, 0x9a, 0xa0, 0x75, 0x8d, 0xbf, 0xc4, 0x0c, 0x47, 0xd0, 0x06, 0xed,
	0x88, 0x88, 0x90, 0xd3, 0x4c, 0x52, 0x96, 0xaa, 0xe7, 0xf7, 0xd0, 0x2a, 0x94, 0xa7, 0x9b, 0x09,
	0xc2, 0x53, 0x9c, 0x10, 0x73, 0x4b, 0x8d, 0xeb, 0x3b, 0x7c, 0x06, 0x0c, 0x21, 0x31, 0x97, 0x24,
	0xf2, 0x13, 0x1a, 0x72, 0x26, 0xcc, 0x86, 0xca, 0xdf, 0x2d, 0xd1, 0x2b, 0x05, 0xc2, 0xe7, 0x60,
	0x7f, 0x44, 0x53, 0x2a, 0x26, 0x4b, 0x5e, 0x53, 0xf1, 0x8c, 0x0a, 0x5e, 0x12, 0x13, 0x16, 0xd1,
	0x11, 0x5d, 0x12, 0xb7, 0x0b, 0x62, 0x05, 0x97, 0x44, 0x06, 0x8c, 0x2a, 0x23, 0xe3, 0x3e, 0x8d,
	0x84, 0xb9, 0x63, 0x37, 0x7a, 0xdd, 0xfe, 0x70, 0x31, 0x3f, 0xee, 0x0e, 0xea, 0xc9, 0xc5, 0x40,
	0xfc, 0x9d, 0x1f, 0x9f, 0x8d, 0xa9, 0x9c, 0xcc, 0x02, 0x37, 0x64, 0x89, 0x57, 0xef, 0x2c, 0x0a,
	0xbc, 0x87, 0xed, 0x14, 0xb7, 0x71, 0x80, 0x05, 0x71, 0x2f, 0x06, 0xa8, 0xbb, 0xf4, 0xbf, 0x88,
	0x04, 0x7c, 0x09, 0xe0, 0x88, 0xe3, 0x30, 0xdf, 0x88, 0x1f, 0xb2, 0x24, 0x8b, 0x89, 0x24, 0x91,
	0xd9, 0xb2, 0xf5, 0xde, 0x16, 0x7a, 0x54, 0x4d, 0xce, 0xab, 0x01, 0x3c, 0x04, 0xdb, 0x84, 0x73,
	0xc6, 0xcd, 0x5d, 0xb5, 0xb1, 0xe2, 0x02, 0x5f, 0x83, 0x9d, 0x40, 0x15, 0xd4, 0x04, 0xb6, 0xde,
	0x6b, 0x9f, 0x9e, 0x6c, 0xfa, 0x0b, 0xd7, 0x2a, 0x3c, 0xd4, 0x50, 0x29, 0x81, 0x6f, 0x40, 0x8b,
	0x17, 0x65, 0x36, 0xdb, 0x4a, 0xfd, 0xbf, 0x02, 0xac, 0xf4, 0x7d, 0xa8, 0xa1, 0x4a, 0x04, 0xaf,
	0x41, 0x47, 0xac, 0xf4, 0xcf, 0xec, 0x28, 0x93, 0x17, 0x9b, 0x4c, 0x36, 0xf7, 0x74, 0xa8, 0xa1,
	0x35, 0x87, 0xfe, 0x1e, 0x68, 0x45, 0xc5, 0xe8, 0xb2, 0xb9, 0xbb, 0x77, 0x00, 0xfa, 0xd6, 0xfd,
	0x6f, 0x4b, 0xbb, 0x5f, 0x58, 0xfa, 0xf7, 0x85, 0xa5, 0xff, 0x5c, 0x58, 0xfa, 0xaf, 0x85, 0xa5,
	0x7f, 0xfb, 0x63, 0x69, 0x9f, 0x9a, 0xb9, 0x6f, 0xb0, 0xa3, 0xbe, 0xd3, 0xb3, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x86, 0xff, 0xfe, 0xd3, 0x18, 0x04, 0x00, 0x00,
}
