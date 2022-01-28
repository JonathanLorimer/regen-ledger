package groupv1alpha1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Member          protoreflect.MessageDescriptor
	fd_Member_address  protoreflect.FieldDescriptor
	fd_Member_weight   protoreflect.FieldDescriptor
	fd_Member_metadata protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_types_proto_init()
	md_Member = File_regen_group_v1alpha1_types_proto.Messages().ByName("Member")
	fd_Member_address = md_Member.Fields().ByName("address")
	fd_Member_weight = md_Member.Fields().ByName("weight")
	fd_Member_metadata = md_Member.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_Member)(nil)

type fastReflection_Member Member

func (x *Member) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Member)(x)
}

func (x *Member) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Member_messageType fastReflection_Member_messageType
var _ protoreflect.MessageType = fastReflection_Member_messageType{}

type fastReflection_Member_messageType struct{}

func (x fastReflection_Member_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Member)(nil)
}
func (x fastReflection_Member_messageType) New() protoreflect.Message {
	return new(fastReflection_Member)
}
func (x fastReflection_Member_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Member
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Member) Descriptor() protoreflect.MessageDescriptor {
	return md_Member
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Member) Type() protoreflect.MessageType {
	return _fastReflection_Member_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Member) New() protoreflect.Message {
	return new(fastReflection_Member)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Member) Interface() protoreflect.ProtoMessage {
	return (*Member)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Member) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_Member_address, value) {
			return
		}
	}
	if x.Weight != "" {
		value := protoreflect.ValueOfString(x.Weight)
		if !f(fd_Member_weight, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_Member_metadata, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Member) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Member.address":
		return x.Address != ""
	case "regen.group.v1alpha1.Member.weight":
		return x.Weight != ""
	case "regen.group.v1alpha1.Member.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Member"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Member does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Member) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Member.address":
		x.Address = ""
	case "regen.group.v1alpha1.Member.weight":
		x.Weight = ""
	case "regen.group.v1alpha1.Member.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Member"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Member does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Member) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.Member.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.Member.weight":
		value := x.Weight
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.Member.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Member"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Member does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Member) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Member.address":
		x.Address = value.Interface().(string)
	case "regen.group.v1alpha1.Member.weight":
		x.Weight = value.Interface().(string)
	case "regen.group.v1alpha1.Member.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Member"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Member does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Member) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Member.address":
		panic(fmt.Errorf("field address of message regen.group.v1alpha1.Member is not mutable"))
	case "regen.group.v1alpha1.Member.weight":
		panic(fmt.Errorf("field weight of message regen.group.v1alpha1.Member is not mutable"))
	case "regen.group.v1alpha1.Member.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.Member is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Member"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Member does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Member) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Member.address":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.Member.weight":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.Member.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Member"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Member does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Member) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.Member", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Member) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Member) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Member) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Member) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Member)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Weight)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Member)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Weight) > 0 {
			i -= len(x.Weight)
			copy(dAtA[i:], x.Weight)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Weight)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Member)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Member: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Weight = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_Members_1_list)(nil)

type _Members_1_list struct {
	list *[]*Member
}

func (x *_Members_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Members_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Members_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	(*x.list)[i] = concreteValue
}

func (x *_Members_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Members_1_list) AppendMutable() protoreflect.Value {
	v := new(Member)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Members_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Members_1_list) NewElement() protoreflect.Value {
	v := new(Member)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Members_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Members         protoreflect.MessageDescriptor
	fd_Members_members protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_types_proto_init()
	md_Members = File_regen_group_v1alpha1_types_proto.Messages().ByName("Members")
	fd_Members_members = md_Members.Fields().ByName("members")
}

var _ protoreflect.Message = (*fastReflection_Members)(nil)

type fastReflection_Members Members

func (x *Members) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Members)(x)
}

func (x *Members) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Members_messageType fastReflection_Members_messageType
var _ protoreflect.MessageType = fastReflection_Members_messageType{}

type fastReflection_Members_messageType struct{}

func (x fastReflection_Members_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Members)(nil)
}
func (x fastReflection_Members_messageType) New() protoreflect.Message {
	return new(fastReflection_Members)
}
func (x fastReflection_Members_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Members
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Members) Descriptor() protoreflect.MessageDescriptor {
	return md_Members
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Members) Type() protoreflect.MessageType {
	return _fastReflection_Members_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Members) New() protoreflect.Message {
	return new(fastReflection_Members)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Members) Interface() protoreflect.ProtoMessage {
	return (*Members)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Members) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Members) != 0 {
		value := protoreflect.ValueOfList(&_Members_1_list{list: &x.Members})
		if !f(fd_Members_members, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Members) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Members.members":
		return len(x.Members) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Members"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Members does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Members) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Members.members":
		x.Members = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Members"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Members does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Members) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.Members.members":
		if len(x.Members) == 0 {
			return protoreflect.ValueOfList(&_Members_1_list{})
		}
		listValue := &_Members_1_list{list: &x.Members}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Members"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Members does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Members) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Members.members":
		lv := value.List()
		clv := lv.(*_Members_1_list)
		x.Members = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Members"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Members does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Members) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Members.members":
		if x.Members == nil {
			x.Members = []*Member{}
		}
		value := &_Members_1_list{list: &x.Members}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Members"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Members does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Members) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Members.members":
		list := []*Member{}
		return protoreflect.ValueOfList(&_Members_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Members"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Members does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Members) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.Members", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Members) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Members) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Members) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Members) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Members)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Members) > 0 {
			for _, e := range x.Members {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Members)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Members) > 0 {
			for iNdEx := len(x.Members) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Members[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Members)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Members: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Members: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Members = append(x.Members, &Member{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Members[len(x.Members)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_ThresholdDecisionPolicy           protoreflect.MessageDescriptor
	fd_ThresholdDecisionPolicy_threshold protoreflect.FieldDescriptor
	fd_ThresholdDecisionPolicy_timeout   protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_types_proto_init()
	md_ThresholdDecisionPolicy = File_regen_group_v1alpha1_types_proto.Messages().ByName("ThresholdDecisionPolicy")
	fd_ThresholdDecisionPolicy_threshold = md_ThresholdDecisionPolicy.Fields().ByName("threshold")
	fd_ThresholdDecisionPolicy_timeout = md_ThresholdDecisionPolicy.Fields().ByName("timeout")
}

var _ protoreflect.Message = (*fastReflection_ThresholdDecisionPolicy)(nil)

type fastReflection_ThresholdDecisionPolicy ThresholdDecisionPolicy

func (x *ThresholdDecisionPolicy) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ThresholdDecisionPolicy)(x)
}

func (x *ThresholdDecisionPolicy) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ThresholdDecisionPolicy_messageType fastReflection_ThresholdDecisionPolicy_messageType
var _ protoreflect.MessageType = fastReflection_ThresholdDecisionPolicy_messageType{}

type fastReflection_ThresholdDecisionPolicy_messageType struct{}

func (x fastReflection_ThresholdDecisionPolicy_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ThresholdDecisionPolicy)(nil)
}
func (x fastReflection_ThresholdDecisionPolicy_messageType) New() protoreflect.Message {
	return new(fastReflection_ThresholdDecisionPolicy)
}
func (x fastReflection_ThresholdDecisionPolicy_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ThresholdDecisionPolicy
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ThresholdDecisionPolicy) Descriptor() protoreflect.MessageDescriptor {
	return md_ThresholdDecisionPolicy
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ThresholdDecisionPolicy) Type() protoreflect.MessageType {
	return _fastReflection_ThresholdDecisionPolicy_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ThresholdDecisionPolicy) New() protoreflect.Message {
	return new(fastReflection_ThresholdDecisionPolicy)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ThresholdDecisionPolicy) Interface() protoreflect.ProtoMessage {
	return (*ThresholdDecisionPolicy)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ThresholdDecisionPolicy) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Threshold != "" {
		value := protoreflect.ValueOfString(x.Threshold)
		if !f(fd_ThresholdDecisionPolicy_threshold, value) {
			return
		}
	}
	if x.Timeout != nil {
		value := protoreflect.ValueOfMessage(x.Timeout.ProtoReflect())
		if !f(fd_ThresholdDecisionPolicy_timeout, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ThresholdDecisionPolicy) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.threshold":
		return x.Threshold != ""
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.timeout":
		return x.Timeout != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.ThresholdDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.ThresholdDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ThresholdDecisionPolicy) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.threshold":
		x.Threshold = ""
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.timeout":
		x.Timeout = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.ThresholdDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.ThresholdDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ThresholdDecisionPolicy) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.threshold":
		value := x.Threshold
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.timeout":
		value := x.Timeout
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.ThresholdDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.ThresholdDecisionPolicy does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ThresholdDecisionPolicy) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.threshold":
		x.Threshold = value.Interface().(string)
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.timeout":
		x.Timeout = value.Message().Interface().(*durationpb.Duration)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.ThresholdDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.ThresholdDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ThresholdDecisionPolicy) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.timeout":
		if x.Timeout == nil {
			x.Timeout = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.Timeout.ProtoReflect())
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.threshold":
		panic(fmt.Errorf("field threshold of message regen.group.v1alpha1.ThresholdDecisionPolicy is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.ThresholdDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.ThresholdDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ThresholdDecisionPolicy) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.threshold":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.ThresholdDecisionPolicy.timeout":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.ThresholdDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.ThresholdDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ThresholdDecisionPolicy) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.ThresholdDecisionPolicy", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ThresholdDecisionPolicy) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ThresholdDecisionPolicy) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ThresholdDecisionPolicy) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ThresholdDecisionPolicy) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ThresholdDecisionPolicy)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Threshold)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Timeout != nil {
			l = options.Size(x.Timeout)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ThresholdDecisionPolicy)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Timeout != nil {
			encoded, err := options.Marshal(x.Timeout)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Threshold) > 0 {
			i -= len(x.Threshold)
			copy(dAtA[i:], x.Threshold)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Threshold)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ThresholdDecisionPolicy)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ThresholdDecisionPolicy: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ThresholdDecisionPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Threshold = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Timeout == nil {
					x.Timeout = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Timeout); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GroupInfo              protoreflect.MessageDescriptor
	fd_GroupInfo_group_id     protoreflect.FieldDescriptor
	fd_GroupInfo_admin        protoreflect.FieldDescriptor
	fd_GroupInfo_metadata     protoreflect.FieldDescriptor
	fd_GroupInfo_version      protoreflect.FieldDescriptor
	fd_GroupInfo_total_weight protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_types_proto_init()
	md_GroupInfo = File_regen_group_v1alpha1_types_proto.Messages().ByName("GroupInfo")
	fd_GroupInfo_group_id = md_GroupInfo.Fields().ByName("group_id")
	fd_GroupInfo_admin = md_GroupInfo.Fields().ByName("admin")
	fd_GroupInfo_metadata = md_GroupInfo.Fields().ByName("metadata")
	fd_GroupInfo_version = md_GroupInfo.Fields().ByName("version")
	fd_GroupInfo_total_weight = md_GroupInfo.Fields().ByName("total_weight")
}

var _ protoreflect.Message = (*fastReflection_GroupInfo)(nil)

type fastReflection_GroupInfo GroupInfo

func (x *GroupInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GroupInfo)(x)
}

func (x *GroupInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GroupInfo_messageType fastReflection_GroupInfo_messageType
var _ protoreflect.MessageType = fastReflection_GroupInfo_messageType{}

type fastReflection_GroupInfo_messageType struct{}

func (x fastReflection_GroupInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GroupInfo)(nil)
}
func (x fastReflection_GroupInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_GroupInfo)
}
func (x fastReflection_GroupInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GroupInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GroupInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_GroupInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GroupInfo) Type() protoreflect.MessageType {
	return _fastReflection_GroupInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GroupInfo) New() protoreflect.Message {
	return new(fastReflection_GroupInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GroupInfo) Interface() protoreflect.ProtoMessage {
	return (*GroupInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GroupInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_GroupInfo_group_id, value) {
			return
		}
	}
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_GroupInfo_admin, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_GroupInfo_metadata, value) {
			return
		}
	}
	if x.Version != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Version)
		if !f(fd_GroupInfo_version, value) {
			return
		}
	}
	if x.TotalWeight != "" {
		value := protoreflect.ValueOfString(x.TotalWeight)
		if !f(fd_GroupInfo_total_weight, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GroupInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupInfo.group_id":
		return x.GroupId != uint64(0)
	case "regen.group.v1alpha1.GroupInfo.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.GroupInfo.metadata":
		return len(x.Metadata) != 0
	case "regen.group.v1alpha1.GroupInfo.version":
		return x.Version != uint64(0)
	case "regen.group.v1alpha1.GroupInfo.total_weight":
		return x.TotalWeight != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupInfo.group_id":
		x.GroupId = uint64(0)
	case "regen.group.v1alpha1.GroupInfo.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.GroupInfo.metadata":
		x.Metadata = nil
	case "regen.group.v1alpha1.GroupInfo.version":
		x.Version = uint64(0)
	case "regen.group.v1alpha1.GroupInfo.total_weight":
		x.TotalWeight = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GroupInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.GroupInfo.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.GroupInfo.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.GroupInfo.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.group.v1alpha1.GroupInfo.version":
		value := x.Version
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.GroupInfo.total_weight":
		value := x.TotalWeight
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupInfo.group_id":
		x.GroupId = value.Uint()
	case "regen.group.v1alpha1.GroupInfo.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.GroupInfo.metadata":
		x.Metadata = value.Bytes()
	case "regen.group.v1alpha1.GroupInfo.version":
		x.Version = value.Uint()
	case "regen.group.v1alpha1.GroupInfo.total_weight":
		x.TotalWeight = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupInfo.group_id":
		panic(fmt.Errorf("field group_id of message regen.group.v1alpha1.GroupInfo is not mutable"))
	case "regen.group.v1alpha1.GroupInfo.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.GroupInfo is not mutable"))
	case "regen.group.v1alpha1.GroupInfo.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.GroupInfo is not mutable"))
	case "regen.group.v1alpha1.GroupInfo.version":
		panic(fmt.Errorf("field version of message regen.group.v1alpha1.GroupInfo is not mutable"))
	case "regen.group.v1alpha1.GroupInfo.total_weight":
		panic(fmt.Errorf("field total_weight of message regen.group.v1alpha1.GroupInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GroupInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupInfo.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.GroupInfo.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.GroupInfo.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.group.v1alpha1.GroupInfo.version":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.GroupInfo.total_weight":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GroupInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.GroupInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GroupInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GroupInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GroupInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GroupInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Version != 0 {
			n += 1 + runtime.Sov(uint64(x.Version))
		}
		l = len(x.TotalWeight)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GroupInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.TotalWeight) > 0 {
			i -= len(x.TotalWeight)
			copy(dAtA[i:], x.TotalWeight)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TotalWeight)))
			i--
			dAtA[i] = 0x2a
		}
		if x.Version != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Version))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0x12
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GroupInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GroupInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GroupInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
				}
				x.Version = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Version |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.TotalWeight = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GroupMember          protoreflect.MessageDescriptor
	fd_GroupMember_group_id protoreflect.FieldDescriptor
	fd_GroupMember_member   protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_types_proto_init()
	md_GroupMember = File_regen_group_v1alpha1_types_proto.Messages().ByName("GroupMember")
	fd_GroupMember_group_id = md_GroupMember.Fields().ByName("group_id")
	fd_GroupMember_member = md_GroupMember.Fields().ByName("member")
}

var _ protoreflect.Message = (*fastReflection_GroupMember)(nil)

type fastReflection_GroupMember GroupMember

func (x *GroupMember) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GroupMember)(x)
}

func (x *GroupMember) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GroupMember_messageType fastReflection_GroupMember_messageType
var _ protoreflect.MessageType = fastReflection_GroupMember_messageType{}

type fastReflection_GroupMember_messageType struct{}

func (x fastReflection_GroupMember_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GroupMember)(nil)
}
func (x fastReflection_GroupMember_messageType) New() protoreflect.Message {
	return new(fastReflection_GroupMember)
}
func (x fastReflection_GroupMember_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GroupMember
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GroupMember) Descriptor() protoreflect.MessageDescriptor {
	return md_GroupMember
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GroupMember) Type() protoreflect.MessageType {
	return _fastReflection_GroupMember_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GroupMember) New() protoreflect.Message {
	return new(fastReflection_GroupMember)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GroupMember) Interface() protoreflect.ProtoMessage {
	return (*GroupMember)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GroupMember) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_GroupMember_group_id, value) {
			return
		}
	}
	if x.Member != nil {
		value := protoreflect.ValueOfMessage(x.Member.ProtoReflect())
		if !f(fd_GroupMember_member, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GroupMember) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupMember.group_id":
		return x.GroupId != uint64(0)
	case "regen.group.v1alpha1.GroupMember.member":
		return x.Member != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupMember"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupMember does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupMember) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupMember.group_id":
		x.GroupId = uint64(0)
	case "regen.group.v1alpha1.GroupMember.member":
		x.Member = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupMember"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupMember does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GroupMember) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.GroupMember.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.GroupMember.member":
		value := x.Member
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupMember"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupMember does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupMember) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupMember.group_id":
		x.GroupId = value.Uint()
	case "regen.group.v1alpha1.GroupMember.member":
		x.Member = value.Message().Interface().(*Member)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupMember"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupMember does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupMember) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupMember.member":
		if x.Member == nil {
			x.Member = new(Member)
		}
		return protoreflect.ValueOfMessage(x.Member.ProtoReflect())
	case "regen.group.v1alpha1.GroupMember.group_id":
		panic(fmt.Errorf("field group_id of message regen.group.v1alpha1.GroupMember is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupMember"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupMember does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GroupMember) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupMember.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.GroupMember.member":
		m := new(Member)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupMember"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupMember does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GroupMember) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.GroupMember", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GroupMember) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupMember) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GroupMember) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GroupMember) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GroupMember)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		if x.Member != nil {
			l = options.Size(x.Member)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GroupMember)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Member != nil {
			encoded, err := options.Marshal(x.Member)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GroupMember)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GroupMember: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GroupMember: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Member == nil {
					x.Member = &Member{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Member); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_GroupAccountInfo                 protoreflect.MessageDescriptor
	fd_GroupAccountInfo_address         protoreflect.FieldDescriptor
	fd_GroupAccountInfo_group_id        protoreflect.FieldDescriptor
	fd_GroupAccountInfo_admin           protoreflect.FieldDescriptor
	fd_GroupAccountInfo_metadata        protoreflect.FieldDescriptor
	fd_GroupAccountInfo_version         protoreflect.FieldDescriptor
	fd_GroupAccountInfo_decision_policy protoreflect.FieldDescriptor
	fd_GroupAccountInfo_derivation_key  protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_types_proto_init()
	md_GroupAccountInfo = File_regen_group_v1alpha1_types_proto.Messages().ByName("GroupAccountInfo")
	fd_GroupAccountInfo_address = md_GroupAccountInfo.Fields().ByName("address")
	fd_GroupAccountInfo_group_id = md_GroupAccountInfo.Fields().ByName("group_id")
	fd_GroupAccountInfo_admin = md_GroupAccountInfo.Fields().ByName("admin")
	fd_GroupAccountInfo_metadata = md_GroupAccountInfo.Fields().ByName("metadata")
	fd_GroupAccountInfo_version = md_GroupAccountInfo.Fields().ByName("version")
	fd_GroupAccountInfo_decision_policy = md_GroupAccountInfo.Fields().ByName("decision_policy")
	fd_GroupAccountInfo_derivation_key = md_GroupAccountInfo.Fields().ByName("derivation_key")
}

var _ protoreflect.Message = (*fastReflection_GroupAccountInfo)(nil)

type fastReflection_GroupAccountInfo GroupAccountInfo

func (x *GroupAccountInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GroupAccountInfo)(x)
}

func (x *GroupAccountInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GroupAccountInfo_messageType fastReflection_GroupAccountInfo_messageType
var _ protoreflect.MessageType = fastReflection_GroupAccountInfo_messageType{}

type fastReflection_GroupAccountInfo_messageType struct{}

func (x fastReflection_GroupAccountInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GroupAccountInfo)(nil)
}
func (x fastReflection_GroupAccountInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_GroupAccountInfo)
}
func (x fastReflection_GroupAccountInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GroupAccountInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GroupAccountInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_GroupAccountInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GroupAccountInfo) Type() protoreflect.MessageType {
	return _fastReflection_GroupAccountInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GroupAccountInfo) New() protoreflect.Message {
	return new(fastReflection_GroupAccountInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GroupAccountInfo) Interface() protoreflect.ProtoMessage {
	return (*GroupAccountInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GroupAccountInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_GroupAccountInfo_address, value) {
			return
		}
	}
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_GroupAccountInfo_group_id, value) {
			return
		}
	}
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_GroupAccountInfo_admin, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_GroupAccountInfo_metadata, value) {
			return
		}
	}
	if x.Version != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Version)
		if !f(fd_GroupAccountInfo_version, value) {
			return
		}
	}
	if x.DecisionPolicy != nil {
		value := protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
		if !f(fd_GroupAccountInfo_decision_policy, value) {
			return
		}
	}
	if len(x.DerivationKey) != 0 {
		value := protoreflect.ValueOfBytes(x.DerivationKey)
		if !f(fd_GroupAccountInfo_derivation_key, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GroupAccountInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupAccountInfo.address":
		return x.Address != ""
	case "regen.group.v1alpha1.GroupAccountInfo.group_id":
		return x.GroupId != uint64(0)
	case "regen.group.v1alpha1.GroupAccountInfo.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.GroupAccountInfo.metadata":
		return len(x.Metadata) != 0
	case "regen.group.v1alpha1.GroupAccountInfo.version":
		return x.Version != uint64(0)
	case "regen.group.v1alpha1.GroupAccountInfo.decision_policy":
		return x.DecisionPolicy != nil
	case "regen.group.v1alpha1.GroupAccountInfo.derivation_key":
		return len(x.DerivationKey) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupAccountInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupAccountInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupAccountInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupAccountInfo.address":
		x.Address = ""
	case "regen.group.v1alpha1.GroupAccountInfo.group_id":
		x.GroupId = uint64(0)
	case "regen.group.v1alpha1.GroupAccountInfo.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.GroupAccountInfo.metadata":
		x.Metadata = nil
	case "regen.group.v1alpha1.GroupAccountInfo.version":
		x.Version = uint64(0)
	case "regen.group.v1alpha1.GroupAccountInfo.decision_policy":
		x.DecisionPolicy = nil
	case "regen.group.v1alpha1.GroupAccountInfo.derivation_key":
		x.DerivationKey = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupAccountInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupAccountInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GroupAccountInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.GroupAccountInfo.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.GroupAccountInfo.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.GroupAccountInfo.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.GroupAccountInfo.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.group.v1alpha1.GroupAccountInfo.version":
		value := x.Version
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.GroupAccountInfo.decision_policy":
		value := x.DecisionPolicy
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.group.v1alpha1.GroupAccountInfo.derivation_key":
		value := x.DerivationKey
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupAccountInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupAccountInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupAccountInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupAccountInfo.address":
		x.Address = value.Interface().(string)
	case "regen.group.v1alpha1.GroupAccountInfo.group_id":
		x.GroupId = value.Uint()
	case "regen.group.v1alpha1.GroupAccountInfo.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.GroupAccountInfo.metadata":
		x.Metadata = value.Bytes()
	case "regen.group.v1alpha1.GroupAccountInfo.version":
		x.Version = value.Uint()
	case "regen.group.v1alpha1.GroupAccountInfo.decision_policy":
		x.DecisionPolicy = value.Message().Interface().(*anypb.Any)
	case "regen.group.v1alpha1.GroupAccountInfo.derivation_key":
		x.DerivationKey = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupAccountInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupAccountInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupAccountInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupAccountInfo.decision_policy":
		if x.DecisionPolicy == nil {
			x.DecisionPolicy = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
	case "regen.group.v1alpha1.GroupAccountInfo.address":
		panic(fmt.Errorf("field address of message regen.group.v1alpha1.GroupAccountInfo is not mutable"))
	case "regen.group.v1alpha1.GroupAccountInfo.group_id":
		panic(fmt.Errorf("field group_id of message regen.group.v1alpha1.GroupAccountInfo is not mutable"))
	case "regen.group.v1alpha1.GroupAccountInfo.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.GroupAccountInfo is not mutable"))
	case "regen.group.v1alpha1.GroupAccountInfo.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.GroupAccountInfo is not mutable"))
	case "regen.group.v1alpha1.GroupAccountInfo.version":
		panic(fmt.Errorf("field version of message regen.group.v1alpha1.GroupAccountInfo is not mutable"))
	case "regen.group.v1alpha1.GroupAccountInfo.derivation_key":
		panic(fmt.Errorf("field derivation_key of message regen.group.v1alpha1.GroupAccountInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupAccountInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupAccountInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GroupAccountInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GroupAccountInfo.address":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.GroupAccountInfo.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.GroupAccountInfo.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.GroupAccountInfo.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.group.v1alpha1.GroupAccountInfo.version":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.GroupAccountInfo.decision_policy":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.group.v1alpha1.GroupAccountInfo.derivation_key":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GroupAccountInfo"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GroupAccountInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GroupAccountInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.GroupAccountInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GroupAccountInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GroupAccountInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GroupAccountInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GroupAccountInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GroupAccountInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Version != 0 {
			n += 1 + runtime.Sov(uint64(x.Version))
		}
		if x.DecisionPolicy != nil {
			l = options.Size(x.DecisionPolicy)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DerivationKey)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GroupAccountInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.DerivationKey) > 0 {
			i -= len(x.DerivationKey)
			copy(dAtA[i:], x.DerivationKey)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DerivationKey)))
			i--
			dAtA[i] = 0x3a
		}
		if x.DecisionPolicy != nil {
			encoded, err := options.Marshal(x.DecisionPolicy)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x32
		}
		if x.Version != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Version))
			i--
			dAtA[i] = 0x28
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0x1a
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GroupAccountInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GroupAccountInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GroupAccountInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
				}
				x.Version = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Version |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DecisionPolicy", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.DecisionPolicy == nil {
					x.DecisionPolicy = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DecisionPolicy); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DerivationKey", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DerivationKey = append(x.DerivationKey[:0], dAtA[iNdEx:postIndex]...)
				if x.DerivationKey == nil {
					x.DerivationKey = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_Proposal_4_list)(nil)

type _Proposal_4_list struct {
	list *[]string
}

func (x *_Proposal_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Proposal_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_Proposal_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_Proposal_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_Proposal_4_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Proposal at list field Proposers as it is not of Message kind"))
}

func (x *_Proposal_4_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_Proposal_4_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_Proposal_4_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_Proposal_13_list)(nil)

type _Proposal_13_list struct {
	list *[]*anypb.Any
}

func (x *_Proposal_13_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Proposal_13_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Proposal_13_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	(*x.list)[i] = concreteValue
}

func (x *_Proposal_13_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Proposal_13_list) AppendMutable() protoreflect.Value {
	v := new(anypb.Any)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Proposal_13_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Proposal_13_list) NewElement() protoreflect.Value {
	v := new(anypb.Any)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Proposal_13_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Proposal                       protoreflect.MessageDescriptor
	fd_Proposal_proposal_id           protoreflect.FieldDescriptor
	fd_Proposal_address               protoreflect.FieldDescriptor
	fd_Proposal_metadata              protoreflect.FieldDescriptor
	fd_Proposal_proposers             protoreflect.FieldDescriptor
	fd_Proposal_submitted_at          protoreflect.FieldDescriptor
	fd_Proposal_group_version         protoreflect.FieldDescriptor
	fd_Proposal_group_account_version protoreflect.FieldDescriptor
	fd_Proposal_status                protoreflect.FieldDescriptor
	fd_Proposal_result                protoreflect.FieldDescriptor
	fd_Proposal_vote_state            protoreflect.FieldDescriptor
	fd_Proposal_timeout               protoreflect.FieldDescriptor
	fd_Proposal_executor_result       protoreflect.FieldDescriptor
	fd_Proposal_msgs                  protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_types_proto_init()
	md_Proposal = File_regen_group_v1alpha1_types_proto.Messages().ByName("Proposal")
	fd_Proposal_proposal_id = md_Proposal.Fields().ByName("proposal_id")
	fd_Proposal_address = md_Proposal.Fields().ByName("address")
	fd_Proposal_metadata = md_Proposal.Fields().ByName("metadata")
	fd_Proposal_proposers = md_Proposal.Fields().ByName("proposers")
	fd_Proposal_submitted_at = md_Proposal.Fields().ByName("submitted_at")
	fd_Proposal_group_version = md_Proposal.Fields().ByName("group_version")
	fd_Proposal_group_account_version = md_Proposal.Fields().ByName("group_account_version")
	fd_Proposal_status = md_Proposal.Fields().ByName("status")
	fd_Proposal_result = md_Proposal.Fields().ByName("result")
	fd_Proposal_vote_state = md_Proposal.Fields().ByName("vote_state")
	fd_Proposal_timeout = md_Proposal.Fields().ByName("timeout")
	fd_Proposal_executor_result = md_Proposal.Fields().ByName("executor_result")
	fd_Proposal_msgs = md_Proposal.Fields().ByName("msgs")
}

var _ protoreflect.Message = (*fastReflection_Proposal)(nil)

type fastReflection_Proposal Proposal

func (x *Proposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Proposal)(x)
}

func (x *Proposal) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Proposal_messageType fastReflection_Proposal_messageType
var _ protoreflect.MessageType = fastReflection_Proposal_messageType{}

type fastReflection_Proposal_messageType struct{}

func (x fastReflection_Proposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Proposal)(nil)
}
func (x fastReflection_Proposal_messageType) New() protoreflect.Message {
	return new(fastReflection_Proposal)
}
func (x fastReflection_Proposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Proposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Proposal) Descriptor() protoreflect.MessageDescriptor {
	return md_Proposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Proposal) Type() protoreflect.MessageType {
	return _fastReflection_Proposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Proposal) New() protoreflect.Message {
	return new(fastReflection_Proposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Proposal) Interface() protoreflect.ProtoMessage {
	return (*Proposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Proposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_Proposal_proposal_id, value) {
			return
		}
	}
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_Proposal_address, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_Proposal_metadata, value) {
			return
		}
	}
	if len(x.Proposers) != 0 {
		value := protoreflect.ValueOfList(&_Proposal_4_list{list: &x.Proposers})
		if !f(fd_Proposal_proposers, value) {
			return
		}
	}
	if x.SubmittedAt != nil {
		value := protoreflect.ValueOfMessage(x.SubmittedAt.ProtoReflect())
		if !f(fd_Proposal_submitted_at, value) {
			return
		}
	}
	if x.GroupVersion != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupVersion)
		if !f(fd_Proposal_group_version, value) {
			return
		}
	}
	if x.GroupAccountVersion != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupAccountVersion)
		if !f(fd_Proposal_group_account_version, value) {
			return
		}
	}
	if x.Status != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Status))
		if !f(fd_Proposal_status, value) {
			return
		}
	}
	if x.Result != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Result))
		if !f(fd_Proposal_result, value) {
			return
		}
	}
	if x.VoteState != nil {
		value := protoreflect.ValueOfMessage(x.VoteState.ProtoReflect())
		if !f(fd_Proposal_vote_state, value) {
			return
		}
	}
	if x.Timeout != nil {
		value := protoreflect.ValueOfMessage(x.Timeout.ProtoReflect())
		if !f(fd_Proposal_timeout, value) {
			return
		}
	}
	if x.ExecutorResult != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.ExecutorResult))
		if !f(fd_Proposal_executor_result, value) {
			return
		}
	}
	if len(x.Msgs) != 0 {
		value := protoreflect.ValueOfList(&_Proposal_13_list{list: &x.Msgs})
		if !f(fd_Proposal_msgs, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Proposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Proposal.proposal_id":
		return x.ProposalId != uint64(0)
	case "regen.group.v1alpha1.Proposal.address":
		return x.Address != ""
	case "regen.group.v1alpha1.Proposal.metadata":
		return len(x.Metadata) != 0
	case "regen.group.v1alpha1.Proposal.proposers":
		return len(x.Proposers) != 0
	case "regen.group.v1alpha1.Proposal.submitted_at":
		return x.SubmittedAt != nil
	case "regen.group.v1alpha1.Proposal.group_version":
		return x.GroupVersion != uint64(0)
	case "regen.group.v1alpha1.Proposal.group_account_version":
		return x.GroupAccountVersion != uint64(0)
	case "regen.group.v1alpha1.Proposal.status":
		return x.Status != 0
	case "regen.group.v1alpha1.Proposal.result":
		return x.Result != 0
	case "regen.group.v1alpha1.Proposal.vote_state":
		return x.VoteState != nil
	case "regen.group.v1alpha1.Proposal.timeout":
		return x.Timeout != nil
	case "regen.group.v1alpha1.Proposal.executor_result":
		return x.ExecutorResult != 0
	case "regen.group.v1alpha1.Proposal.msgs":
		return len(x.Msgs) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Proposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Proposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Proposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Proposal.proposal_id":
		x.ProposalId = uint64(0)
	case "regen.group.v1alpha1.Proposal.address":
		x.Address = ""
	case "regen.group.v1alpha1.Proposal.metadata":
		x.Metadata = nil
	case "regen.group.v1alpha1.Proposal.proposers":
		x.Proposers = nil
	case "regen.group.v1alpha1.Proposal.submitted_at":
		x.SubmittedAt = nil
	case "regen.group.v1alpha1.Proposal.group_version":
		x.GroupVersion = uint64(0)
	case "regen.group.v1alpha1.Proposal.group_account_version":
		x.GroupAccountVersion = uint64(0)
	case "regen.group.v1alpha1.Proposal.status":
		x.Status = 0
	case "regen.group.v1alpha1.Proposal.result":
		x.Result = 0
	case "regen.group.v1alpha1.Proposal.vote_state":
		x.VoteState = nil
	case "regen.group.v1alpha1.Proposal.timeout":
		x.Timeout = nil
	case "regen.group.v1alpha1.Proposal.executor_result":
		x.ExecutorResult = 0
	case "regen.group.v1alpha1.Proposal.msgs":
		x.Msgs = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Proposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Proposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Proposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.Proposal.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.Proposal.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.Proposal.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.group.v1alpha1.Proposal.proposers":
		if len(x.Proposers) == 0 {
			return protoreflect.ValueOfList(&_Proposal_4_list{})
		}
		listValue := &_Proposal_4_list{list: &x.Proposers}
		return protoreflect.ValueOfList(listValue)
	case "regen.group.v1alpha1.Proposal.submitted_at":
		value := x.SubmittedAt
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.group.v1alpha1.Proposal.group_version":
		value := x.GroupVersion
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.Proposal.group_account_version":
		value := x.GroupAccountVersion
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.Proposal.status":
		value := x.Status
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "regen.group.v1alpha1.Proposal.result":
		value := x.Result
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "regen.group.v1alpha1.Proposal.vote_state":
		value := x.VoteState
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.group.v1alpha1.Proposal.timeout":
		value := x.Timeout
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.group.v1alpha1.Proposal.executor_result":
		value := x.ExecutorResult
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "regen.group.v1alpha1.Proposal.msgs":
		if len(x.Msgs) == 0 {
			return protoreflect.ValueOfList(&_Proposal_13_list{})
		}
		listValue := &_Proposal_13_list{list: &x.Msgs}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Proposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Proposal does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Proposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Proposal.proposal_id":
		x.ProposalId = value.Uint()
	case "regen.group.v1alpha1.Proposal.address":
		x.Address = value.Interface().(string)
	case "regen.group.v1alpha1.Proposal.metadata":
		x.Metadata = value.Bytes()
	case "regen.group.v1alpha1.Proposal.proposers":
		lv := value.List()
		clv := lv.(*_Proposal_4_list)
		x.Proposers = *clv.list
	case "regen.group.v1alpha1.Proposal.submitted_at":
		x.SubmittedAt = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.group.v1alpha1.Proposal.group_version":
		x.GroupVersion = value.Uint()
	case "regen.group.v1alpha1.Proposal.group_account_version":
		x.GroupAccountVersion = value.Uint()
	case "regen.group.v1alpha1.Proposal.status":
		x.Status = (Proposal_Status)(value.Enum())
	case "regen.group.v1alpha1.Proposal.result":
		x.Result = (Proposal_Result)(value.Enum())
	case "regen.group.v1alpha1.Proposal.vote_state":
		x.VoteState = value.Message().Interface().(*Tally)
	case "regen.group.v1alpha1.Proposal.timeout":
		x.Timeout = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.group.v1alpha1.Proposal.executor_result":
		x.ExecutorResult = (Proposal_ExecutorResult)(value.Enum())
	case "regen.group.v1alpha1.Proposal.msgs":
		lv := value.List()
		clv := lv.(*_Proposal_13_list)
		x.Msgs = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Proposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Proposal does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Proposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Proposal.proposers":
		if x.Proposers == nil {
			x.Proposers = []string{}
		}
		value := &_Proposal_4_list{list: &x.Proposers}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.Proposal.submitted_at":
		if x.SubmittedAt == nil {
			x.SubmittedAt = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.SubmittedAt.ProtoReflect())
	case "regen.group.v1alpha1.Proposal.vote_state":
		if x.VoteState == nil {
			x.VoteState = new(Tally)
		}
		return protoreflect.ValueOfMessage(x.VoteState.ProtoReflect())
	case "regen.group.v1alpha1.Proposal.timeout":
		if x.Timeout == nil {
			x.Timeout = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Timeout.ProtoReflect())
	case "regen.group.v1alpha1.Proposal.msgs":
		if x.Msgs == nil {
			x.Msgs = []*anypb.Any{}
		}
		value := &_Proposal_13_list{list: &x.Msgs}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.Proposal.proposal_id":
		panic(fmt.Errorf("field proposal_id of message regen.group.v1alpha1.Proposal is not mutable"))
	case "regen.group.v1alpha1.Proposal.address":
		panic(fmt.Errorf("field address of message regen.group.v1alpha1.Proposal is not mutable"))
	case "regen.group.v1alpha1.Proposal.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.Proposal is not mutable"))
	case "regen.group.v1alpha1.Proposal.group_version":
		panic(fmt.Errorf("field group_version of message regen.group.v1alpha1.Proposal is not mutable"))
	case "regen.group.v1alpha1.Proposal.group_account_version":
		panic(fmt.Errorf("field group_account_version of message regen.group.v1alpha1.Proposal is not mutable"))
	case "regen.group.v1alpha1.Proposal.status":
		panic(fmt.Errorf("field status of message regen.group.v1alpha1.Proposal is not mutable"))
	case "regen.group.v1alpha1.Proposal.result":
		panic(fmt.Errorf("field result of message regen.group.v1alpha1.Proposal is not mutable"))
	case "regen.group.v1alpha1.Proposal.executor_result":
		panic(fmt.Errorf("field executor_result of message regen.group.v1alpha1.Proposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Proposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Proposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Proposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Proposal.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.Proposal.address":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.Proposal.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.group.v1alpha1.Proposal.proposers":
		list := []string{}
		return protoreflect.ValueOfList(&_Proposal_4_list{list: &list})
	case "regen.group.v1alpha1.Proposal.submitted_at":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.group.v1alpha1.Proposal.group_version":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.Proposal.group_account_version":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.Proposal.status":
		return protoreflect.ValueOfEnum(0)
	case "regen.group.v1alpha1.Proposal.result":
		return protoreflect.ValueOfEnum(0)
	case "regen.group.v1alpha1.Proposal.vote_state":
		m := new(Tally)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.group.v1alpha1.Proposal.timeout":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.group.v1alpha1.Proposal.executor_result":
		return protoreflect.ValueOfEnum(0)
	case "regen.group.v1alpha1.Proposal.msgs":
		list := []*anypb.Any{}
		return protoreflect.ValueOfList(&_Proposal_13_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Proposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Proposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Proposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.Proposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Proposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Proposal) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Proposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Proposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Proposal)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Proposers) > 0 {
			for _, s := range x.Proposers {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.SubmittedAt != nil {
			l = options.Size(x.SubmittedAt)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupVersion != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupVersion))
		}
		if x.GroupAccountVersion != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupAccountVersion))
		}
		if x.Status != 0 {
			n += 1 + runtime.Sov(uint64(x.Status))
		}
		if x.Result != 0 {
			n += 1 + runtime.Sov(uint64(x.Result))
		}
		if x.VoteState != nil {
			l = options.Size(x.VoteState)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Timeout != nil {
			l = options.Size(x.Timeout)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ExecutorResult != 0 {
			n += 1 + runtime.Sov(uint64(x.ExecutorResult))
		}
		if len(x.Msgs) > 0 {
			for _, e := range x.Msgs {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Proposal)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Msgs) > 0 {
			for iNdEx := len(x.Msgs) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Msgs[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x6a
			}
		}
		if x.ExecutorResult != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ExecutorResult))
			i--
			dAtA[i] = 0x60
		}
		if x.Timeout != nil {
			encoded, err := options.Marshal(x.Timeout)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x5a
		}
		if x.VoteState != nil {
			encoded, err := options.Marshal(x.VoteState)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x52
		}
		if x.Result != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Result))
			i--
			dAtA[i] = 0x48
		}
		if x.Status != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Status))
			i--
			dAtA[i] = 0x40
		}
		if x.GroupAccountVersion != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupAccountVersion))
			i--
			dAtA[i] = 0x38
		}
		if x.GroupVersion != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupVersion))
			i--
			dAtA[i] = 0x30
		}
		if x.SubmittedAt != nil {
			encoded, err := options.Marshal(x.SubmittedAt)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Proposers) > 0 {
			for iNdEx := len(x.Proposers) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Proposers[iNdEx])
				copy(dAtA[i:], x.Proposers[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Proposers[iNdEx])))
				i--
				dAtA[i] = 0x22
			}
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Proposal)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Proposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Proposers", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Proposers = append(x.Proposers, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SubmittedAt", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.SubmittedAt == nil {
					x.SubmittedAt = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SubmittedAt); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupVersion", wireType)
				}
				x.GroupVersion = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupVersion |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupAccountVersion", wireType)
				}
				x.GroupAccountVersion = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupAccountVersion |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 8:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
				}
				x.Status = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Status |= Proposal_Status(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 9:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
				}
				x.Result = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Result |= Proposal_Result(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 10:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VoteState", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.VoteState == nil {
					x.VoteState = &Tally{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.VoteState); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 11:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Timeout == nil {
					x.Timeout = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Timeout); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 12:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ExecutorResult", wireType)
				}
				x.ExecutorResult = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ExecutorResult |= Proposal_ExecutorResult(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 13:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Msgs = append(x.Msgs, &anypb.Any{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Msgs[len(x.Msgs)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Tally               protoreflect.MessageDescriptor
	fd_Tally_yes_count     protoreflect.FieldDescriptor
	fd_Tally_no_count      protoreflect.FieldDescriptor
	fd_Tally_abstain_count protoreflect.FieldDescriptor
	fd_Tally_veto_count    protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_types_proto_init()
	md_Tally = File_regen_group_v1alpha1_types_proto.Messages().ByName("Tally")
	fd_Tally_yes_count = md_Tally.Fields().ByName("yes_count")
	fd_Tally_no_count = md_Tally.Fields().ByName("no_count")
	fd_Tally_abstain_count = md_Tally.Fields().ByName("abstain_count")
	fd_Tally_veto_count = md_Tally.Fields().ByName("veto_count")
}

var _ protoreflect.Message = (*fastReflection_Tally)(nil)

type fastReflection_Tally Tally

func (x *Tally) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Tally)(x)
}

func (x *Tally) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Tally_messageType fastReflection_Tally_messageType
var _ protoreflect.MessageType = fastReflection_Tally_messageType{}

type fastReflection_Tally_messageType struct{}

func (x fastReflection_Tally_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Tally)(nil)
}
func (x fastReflection_Tally_messageType) New() protoreflect.Message {
	return new(fastReflection_Tally)
}
func (x fastReflection_Tally_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Tally
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Tally) Descriptor() protoreflect.MessageDescriptor {
	return md_Tally
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Tally) Type() protoreflect.MessageType {
	return _fastReflection_Tally_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Tally) New() protoreflect.Message {
	return new(fastReflection_Tally)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Tally) Interface() protoreflect.ProtoMessage {
	return (*Tally)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Tally) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.YesCount != "" {
		value := protoreflect.ValueOfString(x.YesCount)
		if !f(fd_Tally_yes_count, value) {
			return
		}
	}
	if x.NoCount != "" {
		value := protoreflect.ValueOfString(x.NoCount)
		if !f(fd_Tally_no_count, value) {
			return
		}
	}
	if x.AbstainCount != "" {
		value := protoreflect.ValueOfString(x.AbstainCount)
		if !f(fd_Tally_abstain_count, value) {
			return
		}
	}
	if x.VetoCount != "" {
		value := protoreflect.ValueOfString(x.VetoCount)
		if !f(fd_Tally_veto_count, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Tally) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Tally.yes_count":
		return x.YesCount != ""
	case "regen.group.v1alpha1.Tally.no_count":
		return x.NoCount != ""
	case "regen.group.v1alpha1.Tally.abstain_count":
		return x.AbstainCount != ""
	case "regen.group.v1alpha1.Tally.veto_count":
		return x.VetoCount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Tally"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Tally does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tally) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Tally.yes_count":
		x.YesCount = ""
	case "regen.group.v1alpha1.Tally.no_count":
		x.NoCount = ""
	case "regen.group.v1alpha1.Tally.abstain_count":
		x.AbstainCount = ""
	case "regen.group.v1alpha1.Tally.veto_count":
		x.VetoCount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Tally"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Tally does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Tally) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.Tally.yes_count":
		value := x.YesCount
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.Tally.no_count":
		value := x.NoCount
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.Tally.abstain_count":
		value := x.AbstainCount
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.Tally.veto_count":
		value := x.VetoCount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Tally"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Tally does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tally) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Tally.yes_count":
		x.YesCount = value.Interface().(string)
	case "regen.group.v1alpha1.Tally.no_count":
		x.NoCount = value.Interface().(string)
	case "regen.group.v1alpha1.Tally.abstain_count":
		x.AbstainCount = value.Interface().(string)
	case "regen.group.v1alpha1.Tally.veto_count":
		x.VetoCount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Tally"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Tally does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tally) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Tally.yes_count":
		panic(fmt.Errorf("field yes_count of message regen.group.v1alpha1.Tally is not mutable"))
	case "regen.group.v1alpha1.Tally.no_count":
		panic(fmt.Errorf("field no_count of message regen.group.v1alpha1.Tally is not mutable"))
	case "regen.group.v1alpha1.Tally.abstain_count":
		panic(fmt.Errorf("field abstain_count of message regen.group.v1alpha1.Tally is not mutable"))
	case "regen.group.v1alpha1.Tally.veto_count":
		panic(fmt.Errorf("field veto_count of message regen.group.v1alpha1.Tally is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Tally"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Tally does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Tally) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Tally.yes_count":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.Tally.no_count":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.Tally.abstain_count":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.Tally.veto_count":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Tally"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Tally does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Tally) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.Tally", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Tally) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tally) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Tally) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Tally) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Tally)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.YesCount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NoCount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AbstainCount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.VetoCount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Tally)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.VetoCount) > 0 {
			i -= len(x.VetoCount)
			copy(dAtA[i:], x.VetoCount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.VetoCount)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.AbstainCount) > 0 {
			i -= len(x.AbstainCount)
			copy(dAtA[i:], x.AbstainCount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AbstainCount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.NoCount) > 0 {
			i -= len(x.NoCount)
			copy(dAtA[i:], x.NoCount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NoCount)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.YesCount) > 0 {
			i -= len(x.YesCount)
			copy(dAtA[i:], x.YesCount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.YesCount)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Tally)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Tally: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Tally: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field YesCount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.YesCount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NoCount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NoCount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AbstainCount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AbstainCount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VetoCount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.VetoCount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Vote              protoreflect.MessageDescriptor
	fd_Vote_proposal_id  protoreflect.FieldDescriptor
	fd_Vote_voter        protoreflect.FieldDescriptor
	fd_Vote_choice       protoreflect.FieldDescriptor
	fd_Vote_metadata     protoreflect.FieldDescriptor
	fd_Vote_submitted_at protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_types_proto_init()
	md_Vote = File_regen_group_v1alpha1_types_proto.Messages().ByName("Vote")
	fd_Vote_proposal_id = md_Vote.Fields().ByName("proposal_id")
	fd_Vote_voter = md_Vote.Fields().ByName("voter")
	fd_Vote_choice = md_Vote.Fields().ByName("choice")
	fd_Vote_metadata = md_Vote.Fields().ByName("metadata")
	fd_Vote_submitted_at = md_Vote.Fields().ByName("submitted_at")
}

var _ protoreflect.Message = (*fastReflection_Vote)(nil)

type fastReflection_Vote Vote

func (x *Vote) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Vote)(x)
}

func (x *Vote) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Vote_messageType fastReflection_Vote_messageType
var _ protoreflect.MessageType = fastReflection_Vote_messageType{}

type fastReflection_Vote_messageType struct{}

func (x fastReflection_Vote_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Vote)(nil)
}
func (x fastReflection_Vote_messageType) New() protoreflect.Message {
	return new(fastReflection_Vote)
}
func (x fastReflection_Vote_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Vote
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Vote) Descriptor() protoreflect.MessageDescriptor {
	return md_Vote
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Vote) Type() protoreflect.MessageType {
	return _fastReflection_Vote_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Vote) New() protoreflect.Message {
	return new(fastReflection_Vote)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Vote) Interface() protoreflect.ProtoMessage {
	return (*Vote)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Vote) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_Vote_proposal_id, value) {
			return
		}
	}
	if x.Voter != "" {
		value := protoreflect.ValueOfString(x.Voter)
		if !f(fd_Vote_voter, value) {
			return
		}
	}
	if x.Choice != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Choice))
		if !f(fd_Vote_choice, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_Vote_metadata, value) {
			return
		}
	}
	if x.SubmittedAt != nil {
		value := protoreflect.ValueOfMessage(x.SubmittedAt.ProtoReflect())
		if !f(fd_Vote_submitted_at, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Vote) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Vote.proposal_id":
		return x.ProposalId != uint64(0)
	case "regen.group.v1alpha1.Vote.voter":
		return x.Voter != ""
	case "regen.group.v1alpha1.Vote.choice":
		return x.Choice != 0
	case "regen.group.v1alpha1.Vote.metadata":
		return len(x.Metadata) != 0
	case "regen.group.v1alpha1.Vote.submitted_at":
		return x.SubmittedAt != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Vote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Vote does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Vote) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Vote.proposal_id":
		x.ProposalId = uint64(0)
	case "regen.group.v1alpha1.Vote.voter":
		x.Voter = ""
	case "regen.group.v1alpha1.Vote.choice":
		x.Choice = 0
	case "regen.group.v1alpha1.Vote.metadata":
		x.Metadata = nil
	case "regen.group.v1alpha1.Vote.submitted_at":
		x.SubmittedAt = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Vote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Vote does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Vote) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.Vote.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.Vote.voter":
		value := x.Voter
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.Vote.choice":
		value := x.Choice
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "regen.group.v1alpha1.Vote.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.group.v1alpha1.Vote.submitted_at":
		value := x.SubmittedAt
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Vote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Vote does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Vote) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Vote.proposal_id":
		x.ProposalId = value.Uint()
	case "regen.group.v1alpha1.Vote.voter":
		x.Voter = value.Interface().(string)
	case "regen.group.v1alpha1.Vote.choice":
		x.Choice = (Choice)(value.Enum())
	case "regen.group.v1alpha1.Vote.metadata":
		x.Metadata = value.Bytes()
	case "regen.group.v1alpha1.Vote.submitted_at":
		x.SubmittedAt = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Vote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Vote does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Vote) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Vote.submitted_at":
		if x.SubmittedAt == nil {
			x.SubmittedAt = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.SubmittedAt.ProtoReflect())
	case "regen.group.v1alpha1.Vote.proposal_id":
		panic(fmt.Errorf("field proposal_id of message regen.group.v1alpha1.Vote is not mutable"))
	case "regen.group.v1alpha1.Vote.voter":
		panic(fmt.Errorf("field voter of message regen.group.v1alpha1.Vote is not mutable"))
	case "regen.group.v1alpha1.Vote.choice":
		panic(fmt.Errorf("field choice of message regen.group.v1alpha1.Vote is not mutable"))
	case "regen.group.v1alpha1.Vote.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.Vote is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Vote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Vote does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Vote) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.Vote.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.Vote.voter":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.Vote.choice":
		return protoreflect.ValueOfEnum(0)
	case "regen.group.v1alpha1.Vote.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.group.v1alpha1.Vote.submitted_at":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.Vote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.Vote does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Vote) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.Vote", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Vote) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Vote) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Vote) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Vote) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Vote)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		l = len(x.Voter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Choice != 0 {
			n += 1 + runtime.Sov(uint64(x.Choice))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.SubmittedAt != nil {
			l = options.Size(x.SubmittedAt)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Vote)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.SubmittedAt != nil {
			encoded, err := options.Marshal(x.SubmittedAt)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x22
		}
		if x.Choice != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Choice))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Voter) > 0 {
			i -= len(x.Voter)
			copy(dAtA[i:], x.Voter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Voter)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Vote)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Vote: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Voter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Choice", wireType)
				}
				x.Choice = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Choice |= Choice(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SubmittedAt", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.SubmittedAt == nil {
					x.SubmittedAt = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SubmittedAt); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: regen/group/v1alpha1/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Choice defines available types of choices for voting.
type Choice int32

const (
	// CHOICE_UNSPECIFIED defines a no-op voting choice.
	Choice_CHOICE_UNSPECIFIED Choice = 0
	// CHOICE_NO defines a no voting choice.
	Choice_CHOICE_NO Choice = 1
	// CHOICE_YES defines a yes voting choice.
	Choice_CHOICE_YES Choice = 2
	// CHOICE_ABSTAIN defines an abstaining voting choice.
	Choice_CHOICE_ABSTAIN Choice = 3
	// CHOICE_VETO defines a voting choice with veto.
	Choice_CHOICE_VETO Choice = 4
)

// Enum value maps for Choice.
var (
	Choice_name = map[int32]string{
		0: "CHOICE_UNSPECIFIED",
		1: "CHOICE_NO",
		2: "CHOICE_YES",
		3: "CHOICE_ABSTAIN",
		4: "CHOICE_VETO",
	}
	Choice_value = map[string]int32{
		"CHOICE_UNSPECIFIED": 0,
		"CHOICE_NO":          1,
		"CHOICE_YES":         2,
		"CHOICE_ABSTAIN":     3,
		"CHOICE_VETO":        4,
	}
)

func (x Choice) Enum() *Choice {
	p := new(Choice)
	*p = x
	return p
}

func (x Choice) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Choice) Descriptor() protoreflect.EnumDescriptor {
	return file_regen_group_v1alpha1_types_proto_enumTypes[0].Descriptor()
}

func (Choice) Type() protoreflect.EnumType {
	return &file_regen_group_v1alpha1_types_proto_enumTypes[0]
}

func (x Choice) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Choice.Descriptor instead.
func (Choice) EnumDescriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{0}
}

// Status defines proposal statuses.
type Proposal_Status int32

const (
	// An empty value is invalid and not allowed.
	Proposal_STATUS_UNSPECIFIED Proposal_Status = 0
	// Initial status of a proposal when persisted.
	Proposal_STATUS_SUBMITTED Proposal_Status = 1
	// Final status of a proposal when the final tally was executed.
	Proposal_STATUS_CLOSED Proposal_Status = 2
	// Final status of a proposal when the group was modified before the final tally.
	Proposal_STATUS_ABORTED Proposal_Status = 3
)

// Enum value maps for Proposal_Status.
var (
	Proposal_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUBMITTED",
		2: "STATUS_CLOSED",
		3: "STATUS_ABORTED",
	}
	Proposal_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUBMITTED":   1,
		"STATUS_CLOSED":      2,
		"STATUS_ABORTED":     3,
	}
)

func (x Proposal_Status) Enum() *Proposal_Status {
	p := new(Proposal_Status)
	*p = x
	return p
}

func (x Proposal_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Proposal_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_regen_group_v1alpha1_types_proto_enumTypes[1].Descriptor()
}

func (Proposal_Status) Type() protoreflect.EnumType {
	return &file_regen_group_v1alpha1_types_proto_enumTypes[1]
}

func (x Proposal_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Proposal_Status.Descriptor instead.
func (Proposal_Status) EnumDescriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{6, 0}
}

// Result defines types of proposal results.
type Proposal_Result int32

const (
	// An empty value is invalid and not allowed
	Proposal_RESULT_UNSPECIFIED Proposal_Result = 0
	// Until a final tally has happened the status is unfinalized
	Proposal_RESULT_UNFINALIZED Proposal_Result = 1
	// Final result of the tally
	Proposal_RESULT_ACCEPTED Proposal_Result = 2
	// Final result of the tally
	Proposal_RESULT_REJECTED Proposal_Result = 3
)

// Enum value maps for Proposal_Result.
var (
	Proposal_Result_name = map[int32]string{
		0: "RESULT_UNSPECIFIED",
		1: "RESULT_UNFINALIZED",
		2: "RESULT_ACCEPTED",
		3: "RESULT_REJECTED",
	}
	Proposal_Result_value = map[string]int32{
		"RESULT_UNSPECIFIED": 0,
		"RESULT_UNFINALIZED": 1,
		"RESULT_ACCEPTED":    2,
		"RESULT_REJECTED":    3,
	}
)

func (x Proposal_Result) Enum() *Proposal_Result {
	p := new(Proposal_Result)
	*p = x
	return p
}

func (x Proposal_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Proposal_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_regen_group_v1alpha1_types_proto_enumTypes[2].Descriptor()
}

func (Proposal_Result) Type() protoreflect.EnumType {
	return &file_regen_group_v1alpha1_types_proto_enumTypes[2]
}

func (x Proposal_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Proposal_Result.Descriptor instead.
func (Proposal_Result) EnumDescriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{6, 1}
}

// ExecutorResult defines types of proposal executor results.
type Proposal_ExecutorResult int32

const (
	// An empty value is not allowed.
	Proposal_EXECUTOR_RESULT_UNSPECIFIED Proposal_ExecutorResult = 0
	// We have not yet run the executor.
	Proposal_EXECUTOR_RESULT_NOT_RUN Proposal_ExecutorResult = 1
	// The executor was successful and proposed action updated state.
	Proposal_EXECUTOR_RESULT_SUCCESS Proposal_ExecutorResult = 2
	// The executor returned an error and proposed action didn't update state.
	Proposal_EXECUTOR_RESULT_FAILURE Proposal_ExecutorResult = 3
)

// Enum value maps for Proposal_ExecutorResult.
var (
	Proposal_ExecutorResult_name = map[int32]string{
		0: "EXECUTOR_RESULT_UNSPECIFIED",
		1: "EXECUTOR_RESULT_NOT_RUN",
		2: "EXECUTOR_RESULT_SUCCESS",
		3: "EXECUTOR_RESULT_FAILURE",
	}
	Proposal_ExecutorResult_value = map[string]int32{
		"EXECUTOR_RESULT_UNSPECIFIED": 0,
		"EXECUTOR_RESULT_NOT_RUN":     1,
		"EXECUTOR_RESULT_SUCCESS":     2,
		"EXECUTOR_RESULT_FAILURE":     3,
	}
)

func (x Proposal_ExecutorResult) Enum() *Proposal_ExecutorResult {
	p := new(Proposal_ExecutorResult)
	*p = x
	return p
}

func (x Proposal_ExecutorResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Proposal_ExecutorResult) Descriptor() protoreflect.EnumDescriptor {
	return file_regen_group_v1alpha1_types_proto_enumTypes[3].Descriptor()
}

func (Proposal_ExecutorResult) Type() protoreflect.EnumType {
	return &file_regen_group_v1alpha1_types_proto_enumTypes[3]
}

func (x Proposal_ExecutorResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Proposal_ExecutorResult.Descriptor instead.
func (Proposal_ExecutorResult) EnumDescriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{6, 2}
}

// Member represents a group member with an account address,
// non-zero weight and metadata.
type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the member's account address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// weight is the member's voting weight that should be greater than 0.
	Weight string `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	// metadata is any arbitrary metadata to attached to the member.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Member) GetWeight() string {
	if x != nil {
		return x.Weight
	}
	return ""
}

func (x *Member) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Members defines a repeated slice of Member objects.
type Members struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// members is the list of members.
	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *Members) Reset() {
	*x = Members{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Members) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Members) ProtoMessage() {}

// Deprecated: Use Members.ProtoReflect.Descriptor instead.
func (*Members) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{1}
}

func (x *Members) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

// ThresholdDecisionPolicy implements the DecisionPolicy interface
type ThresholdDecisionPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// threshold is the minimum weighted sum of yes votes that must be met or exceeded for a proposal to succeed.
	Threshold string `protobuf:"bytes,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	// timeout is the duration from submission of a proposal to the end of voting period
	// Within this times votes and exec messages can be submitted.
	Timeout *durationpb.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *ThresholdDecisionPolicy) Reset() {
	*x = ThresholdDecisionPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThresholdDecisionPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThresholdDecisionPolicy) ProtoMessage() {}

// Deprecated: Use ThresholdDecisionPolicy.ProtoReflect.Descriptor instead.
func (*ThresholdDecisionPolicy) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{2}
}

func (x *ThresholdDecisionPolicy) GetThreshold() string {
	if x != nil {
		return x.Threshold
	}
	return ""
}

func (x *ThresholdDecisionPolicy) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// GroupInfo represents the high-level on-chain information for a group.
type GroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// admin is the account address of the group's admin.
	Admin string `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
	// metadata is any arbitrary metadata to attached to the group.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// version is used to track changes to a group's membership structure that
	// would break existing proposals. Whenever any members weight is changed,
	// or any member is added or removed this version is incremented and will
	// cause proposals based on older versions of this group to fail
	Version uint64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// total_weight is the sum of the group members' weights.
	TotalWeight string `protobuf:"bytes,5,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
}

func (x *GroupInfo) Reset() {
	*x = GroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupInfo) ProtoMessage() {}

// Deprecated: Use GroupInfo.ProtoReflect.Descriptor instead.
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{3}
}

func (x *GroupInfo) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupInfo) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *GroupInfo) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GroupInfo) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GroupInfo) GetTotalWeight() string {
	if x != nil {
		return x.TotalWeight
	}
	return ""
}

// GroupMember represents the relationship between a group and a member.
type GroupMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// member is the member data.
	Member *Member `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *GroupMember) Reset() {
	*x = GroupMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMember) ProtoMessage() {}

// Deprecated: Use GroupMember.ProtoReflect.Descriptor instead.
func (*GroupMember) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{4}
}

func (x *GroupMember) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupMember) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

// GroupAccountInfo represents the high-level on-chain information for a group account.
type GroupAccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the group account address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,3,opt,name=admin,proto3" json:"admin,omitempty"`
	// metadata is any arbitrary metadata to attached to the group account.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// version is used to track changes to a group's GroupAccountInfo structure that
	// would create a different result on a running proposal.
	Version uint64 `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	// decision_policy specifies the group account's decision policy.
	DecisionPolicy *anypb.Any `protobuf:"bytes,6,opt,name=decision_policy,json=decisionPolicy,proto3" json:"decision_policy,omitempty"`
	// derivation_key is the "derivation" key of the group account,
	// which is needed to derive the group root module key and execute proposals.
	DerivationKey []byte `protobuf:"bytes,7,opt,name=derivation_key,json=derivationKey,proto3" json:"derivation_key,omitempty"`
}

func (x *GroupAccountInfo) Reset() {
	*x = GroupAccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupAccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupAccountInfo) ProtoMessage() {}

// Deprecated: Use GroupAccountInfo.ProtoReflect.Descriptor instead.
func (*GroupAccountInfo) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{5}
}

func (x *GroupAccountInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GroupAccountInfo) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupAccountInfo) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *GroupAccountInfo) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GroupAccountInfo) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GroupAccountInfo) GetDecisionPolicy() *anypb.Any {
	if x != nil {
		return x.DecisionPolicy
	}
	return nil
}

func (x *GroupAccountInfo) GetDerivationKey() []byte {
	if x != nil {
		return x.DerivationKey
	}
	return nil
}

// Proposal defines a group proposal. Any member of a group can submit a proposal
// for a group account to decide upon.
// A proposal consists of a set of `sdk.Msg`s that will be executed if the proposal
// passes as well as some optional metadata associated with the proposal.
type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proposal_id is the unique id of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// address is the group account address.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// metadata is any arbitrary metadata to attached to the proposal.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// proposers are the account addresses of the proposers.
	Proposers []string `protobuf:"bytes,4,rep,name=proposers,proto3" json:"proposers,omitempty"`
	// submitted_at is a timestamp specifying when a proposal was submitted.
	SubmittedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	// group_version tracks the version of the group that this proposal corresponds to.
	// When group membership is changed, existing proposals from previous group versions will become invalid.
	GroupVersion uint64 `protobuf:"varint,6,opt,name=group_version,json=groupVersion,proto3" json:"group_version,omitempty"`
	// group_account_version tracks the version of the group account that this proposal corresponds to.
	// When a decision policy is changed, existing proposals from previous policy versions will become invalid.
	GroupAccountVersion uint64 `protobuf:"varint,7,opt,name=group_account_version,json=groupAccountVersion,proto3" json:"group_account_version,omitempty"`
	// Status represents the high level position in the life cycle of the proposal. Initial value is Submitted.
	Status Proposal_Status `protobuf:"varint,8,opt,name=status,proto3,enum=regen.group.v1alpha1.Proposal_Status" json:"status,omitempty"`
	// result is the final result based on the votes and election rule. Initial value is unfinalized.
	// The result is persisted so that clients can always rely on this state and not have to replicate the logic.
	Result Proposal_Result `protobuf:"varint,9,opt,name=result,proto3,enum=regen.group.v1alpha1.Proposal_Result" json:"result,omitempty"`
	// vote_state contains the sums of all weighted votes for this proposal.
	VoteState *Tally `protobuf:"bytes,10,opt,name=vote_state,json=voteState,proto3" json:"vote_state,omitempty"`
	// timeout is the timestamp of the block where the proposal execution times out. Header times of the votes and execution messages
	// must be before this end time to be included in the election. After the timeout timestamp the proposal can not be
	// executed anymore and should be considered pending delete.
	Timeout *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// executor_result is the final result based on the votes and election rule. Initial value is NotRun.
	ExecutorResult Proposal_ExecutorResult `protobuf:"varint,12,opt,name=executor_result,json=executorResult,proto3,enum=regen.group.v1alpha1.Proposal_ExecutorResult" json:"executor_result,omitempty"`
	// msgs is a list of Msgs that will be executed if the proposal passes.
	Msgs []*anypb.Any `protobuf:"bytes,13,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{6}
}

func (x *Proposal) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *Proposal) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Proposal) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Proposal) GetProposers() []string {
	if x != nil {
		return x.Proposers
	}
	return nil
}

func (x *Proposal) GetSubmittedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SubmittedAt
	}
	return nil
}

func (x *Proposal) GetGroupVersion() uint64 {
	if x != nil {
		return x.GroupVersion
	}
	return 0
}

func (x *Proposal) GetGroupAccountVersion() uint64 {
	if x != nil {
		return x.GroupAccountVersion
	}
	return 0
}

func (x *Proposal) GetStatus() Proposal_Status {
	if x != nil {
		return x.Status
	}
	return Proposal_STATUS_UNSPECIFIED
}

func (x *Proposal) GetResult() Proposal_Result {
	if x != nil {
		return x.Result
	}
	return Proposal_RESULT_UNSPECIFIED
}

func (x *Proposal) GetVoteState() *Tally {
	if x != nil {
		return x.VoteState
	}
	return nil
}

func (x *Proposal) GetTimeout() *timestamppb.Timestamp {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Proposal) GetExecutorResult() Proposal_ExecutorResult {
	if x != nil {
		return x.ExecutorResult
	}
	return Proposal_EXECUTOR_RESULT_UNSPECIFIED
}

func (x *Proposal) GetMsgs() []*anypb.Any {
	if x != nil {
		return x.Msgs
	}
	return nil
}

// Tally represents the sum of weighted votes.
type Tally struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// yes_count is the weighted sum of yes votes.
	YesCount string `protobuf:"bytes,1,opt,name=yes_count,json=yesCount,proto3" json:"yes_count,omitempty"`
	// no_count is the weighted sum of no votes.
	NoCount string `protobuf:"bytes,2,opt,name=no_count,json=noCount,proto3" json:"no_count,omitempty"`
	// abstain_count is the weighted sum of abstainers
	AbstainCount string `protobuf:"bytes,3,opt,name=abstain_count,json=abstainCount,proto3" json:"abstain_count,omitempty"`
	// veto_count is the weighted sum of vetoes.
	VetoCount string `protobuf:"bytes,4,opt,name=veto_count,json=vetoCount,proto3" json:"veto_count,omitempty"`
}

func (x *Tally) Reset() {
	*x = Tally{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tally) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tally) ProtoMessage() {}

// Deprecated: Use Tally.ProtoReflect.Descriptor instead.
func (*Tally) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{7}
}

func (x *Tally) GetYesCount() string {
	if x != nil {
		return x.YesCount
	}
	return ""
}

func (x *Tally) GetNoCount() string {
	if x != nil {
		return x.NoCount
	}
	return ""
}

func (x *Tally) GetAbstainCount() string {
	if x != nil {
		return x.AbstainCount
	}
	return ""
}

func (x *Tally) GetVetoCount() string {
	if x != nil {
		return x.VetoCount
	}
	return ""
}

// Vote represents a vote for a proposal.
type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proposal is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// voter is the account address of the voter.
	Voter string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	// choice is the voter's choice on the proposal.
	Choice Choice `protobuf:"varint,3,opt,name=choice,proto3,enum=regen.group.v1alpha1.Choice" json:"choice,omitempty"`
	// metadata is any arbitrary metadata to attached to the vote.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// submitted_at is the timestamp when the vote was submitted.
	SubmittedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_types_proto_rawDescGZIP(), []int{8}
}

func (x *Vote) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *Vote) GetVoter() string {
	if x != nil {
		return x.Voter
	}
	return ""
}

func (x *Vote) GetChoice() Choice {
	if x != nil {
		return x.Choice
	}
	return Choice_CHOICE_UNSPECIFIED
}

func (x *Vote) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Vote) GetSubmittedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SubmittedAt
	}
	return nil
}

var File_regen_group_v1alpha1_types_proto protoreflect.FileDescriptor

var file_regen_group_v1alpha1_types_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a,
	0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x17, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x12, 0x39, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x3a, 0x12, 0xca, 0xb4, 0x2d,
	0x0e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22,
	0x95, 0x01, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x5e, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x97, 0x02, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a,
	0x0f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x12, 0xca, 0xb4,
	0x2d, 0x0e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x0e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f,
	0x01, 0x22, 0xed, 0x0a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0b, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a,
	0x15, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x25, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x40, 0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x6c, 0x6c, 0x79,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x3a, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04,
	0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x56, 0x0a,
	0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x22,
	0xd0, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x12, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x1a, 0x19, 0x8a, 0x9d, 0x20, 0x15, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x31, 0x0a,
	0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x1a, 0x1b, 0x8a, 0x9d, 0x20, 0x17, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x12, 0x2b, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x44, 0x10, 0x02, 0x1a, 0x18, 0x8a, 0x9d, 0x20, 0x14, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x2d, 0x0a,
	0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x03, 0x1a, 0x19, 0x8a, 0x9d, 0x20, 0x15, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x1a, 0x04, 0x88, 0xa3,
	0x1e, 0x00, 0x22, 0xda, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a,
	0x12, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x19, 0x8a, 0x9d, 0x20, 0x15, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x12, 0x35, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x46, 0x49, 0x4e,
	0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x1d, 0x8a, 0x9d, 0x20, 0x19, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x55, 0x6e, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x0f, 0x52, 0x45, 0x53, 0x55, 0x4c,
	0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x1a, 0x1a, 0x8a, 0x9d,
	0x20, 0x16, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x0f, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x1a, 0x8a,
	0x9d, 0x20, 0x16, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x1a, 0x04, 0x88, 0xa3, 0x1e, 0x00, 0x22,
	0x99, 0x02, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x42, 0x0a, 0x1b, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x1a, 0x21, 0x8a, 0x9d, 0x20, 0x1d, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x17, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54,
	0x4f, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x55,
	0x4e, 0x10, 0x01, 0x1a, 0x20, 0x8a, 0x9d, 0x20, 0x1c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4e,
	0x6f, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x3e, 0x0a, 0x17, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x4f,
	0x52, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x02, 0x1a, 0x21, 0x8a, 0x9d, 0x20, 0x1d, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x17, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x4f,
	0x52, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45,
	0x10, 0x03, 0x1a, 0x21, 0x8a, 0x9d, 0x20, 0x1d, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x1a, 0x04, 0x88, 0xa3, 0x1e, 0x00, 0x3a, 0x04, 0x88, 0xa0, 0x1f,
	0x00, 0x22, 0x89, 0x01, 0x0a, 0x05, 0x54, 0x61, 0x6c, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x79,
	0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x79, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x62, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x62, 0x73, 0x74,
	0x61, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x74, 0x6f,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65,
	0x74, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x22, 0xd4, 0x01,
	0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a,
	0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x06, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x43, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x2a, 0x64, 0x0a, 0x06, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x12, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45,
	0x5f, 0x4e, 0x4f, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x5f,
	0x59, 0x45, 0x53, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x5f,
	0x41, 0x42, 0x53, 0x54, 0x41, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x4f,
	0x49, 0x43, 0x45, 0x5f, 0x56, 0x45, 0x54, 0x4f, 0x10, 0x04, 0x42, 0xe6, 0x01, 0x0a, 0x18, 0x63,
	0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x47, 0x58, 0xaa, 0x02, 0x14, 0x52, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x14, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x20, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x52, 0x65, 0x67,
	0x65, 0x6e, 0x3a, 0x3a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_group_v1alpha1_types_proto_rawDescOnce sync.Once
	file_regen_group_v1alpha1_types_proto_rawDescData = file_regen_group_v1alpha1_types_proto_rawDesc
)

func file_regen_group_v1alpha1_types_proto_rawDescGZIP() []byte {
	file_regen_group_v1alpha1_types_proto_rawDescOnce.Do(func() {
		file_regen_group_v1alpha1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_group_v1alpha1_types_proto_rawDescData)
	})
	return file_regen_group_v1alpha1_types_proto_rawDescData
}

var file_regen_group_v1alpha1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_regen_group_v1alpha1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_regen_group_v1alpha1_types_proto_goTypes = []interface{}{
	(Choice)(0),                     // 0: regen.group.v1alpha1.Choice
	(Proposal_Status)(0),            // 1: regen.group.v1alpha1.Proposal.Status
	(Proposal_Result)(0),            // 2: regen.group.v1alpha1.Proposal.Result
	(Proposal_ExecutorResult)(0),    // 3: regen.group.v1alpha1.Proposal.ExecutorResult
	(*Member)(nil),                  // 4: regen.group.v1alpha1.Member
	(*Members)(nil),                 // 5: regen.group.v1alpha1.Members
	(*ThresholdDecisionPolicy)(nil), // 6: regen.group.v1alpha1.ThresholdDecisionPolicy
	(*GroupInfo)(nil),               // 7: regen.group.v1alpha1.GroupInfo
	(*GroupMember)(nil),             // 8: regen.group.v1alpha1.GroupMember
	(*GroupAccountInfo)(nil),        // 9: regen.group.v1alpha1.GroupAccountInfo
	(*Proposal)(nil),                // 10: regen.group.v1alpha1.Proposal
	(*Tally)(nil),                   // 11: regen.group.v1alpha1.Tally
	(*Vote)(nil),                    // 12: regen.group.v1alpha1.Vote
	(*durationpb.Duration)(nil),     // 13: google.protobuf.Duration
	(*anypb.Any)(nil),               // 14: google.protobuf.Any
	(*timestamppb.Timestamp)(nil),   // 15: google.protobuf.Timestamp
}
var file_regen_group_v1alpha1_types_proto_depIdxs = []int32{
	4,  // 0: regen.group.v1alpha1.Members.members:type_name -> regen.group.v1alpha1.Member
	13, // 1: regen.group.v1alpha1.ThresholdDecisionPolicy.timeout:type_name -> google.protobuf.Duration
	4,  // 2: regen.group.v1alpha1.GroupMember.member:type_name -> regen.group.v1alpha1.Member
	14, // 3: regen.group.v1alpha1.GroupAccountInfo.decision_policy:type_name -> google.protobuf.Any
	15, // 4: regen.group.v1alpha1.Proposal.submitted_at:type_name -> google.protobuf.Timestamp
	1,  // 5: regen.group.v1alpha1.Proposal.status:type_name -> regen.group.v1alpha1.Proposal.Status
	2,  // 6: regen.group.v1alpha1.Proposal.result:type_name -> regen.group.v1alpha1.Proposal.Result
	11, // 7: regen.group.v1alpha1.Proposal.vote_state:type_name -> regen.group.v1alpha1.Tally
	15, // 8: regen.group.v1alpha1.Proposal.timeout:type_name -> google.protobuf.Timestamp
	3,  // 9: regen.group.v1alpha1.Proposal.executor_result:type_name -> regen.group.v1alpha1.Proposal.ExecutorResult
	14, // 10: regen.group.v1alpha1.Proposal.msgs:type_name -> google.protobuf.Any
	0,  // 11: regen.group.v1alpha1.Vote.choice:type_name -> regen.group.v1alpha1.Choice
	15, // 12: regen.group.v1alpha1.Vote.submitted_at:type_name -> google.protobuf.Timestamp
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_regen_group_v1alpha1_types_proto_init() }
func file_regen_group_v1alpha1_types_proto_init() {
	if File_regen_group_v1alpha1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_group_v1alpha1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Members); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThresholdDecisionPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMember); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupAccountInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tally); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_group_v1alpha1_types_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_group_v1alpha1_types_proto_goTypes,
		DependencyIndexes: file_regen_group_v1alpha1_types_proto_depIdxs,
		EnumInfos:         file_regen_group_v1alpha1_types_proto_enumTypes,
		MessageInfos:      file_regen_group_v1alpha1_types_proto_msgTypes,
	}.Build()
	File_regen_group_v1alpha1_types_proto = out.File
	file_regen_group_v1alpha1_types_proto_rawDesc = nil
	file_regen_group_v1alpha1_types_proto_goTypes = nil
	file_regen_group_v1alpha1_types_proto_depIdxs = nil
}