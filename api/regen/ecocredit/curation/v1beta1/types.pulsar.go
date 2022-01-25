package curationv1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_TagTarget            protoreflect.MessageDescriptor
	fd_TagTarget_class_id   protoreflect.FieldDescriptor
	fd_TagTarget_project_id protoreflect.FieldDescriptor
	fd_TagTarget_batch_id   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_curation_v1beta1_types_proto_init()
	md_TagTarget = File_regen_ecocredit_curation_v1beta1_types_proto.Messages().ByName("TagTarget")
	fd_TagTarget_class_id = md_TagTarget.Fields().ByName("class_id")
	fd_TagTarget_project_id = md_TagTarget.Fields().ByName("project_id")
	fd_TagTarget_batch_id = md_TagTarget.Fields().ByName("batch_id")
}

var _ protoreflect.Message = (*fastReflection_TagTarget)(nil)

type fastReflection_TagTarget TagTarget

func (x *TagTarget) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TagTarget)(x)
}

func (x *TagTarget) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_curation_v1beta1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TagTarget_messageType fastReflection_TagTarget_messageType
var _ protoreflect.MessageType = fastReflection_TagTarget_messageType{}

type fastReflection_TagTarget_messageType struct{}

func (x fastReflection_TagTarget_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TagTarget)(nil)
}
func (x fastReflection_TagTarget_messageType) New() protoreflect.Message {
	return new(fastReflection_TagTarget)
}
func (x fastReflection_TagTarget_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TagTarget
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TagTarget) Descriptor() protoreflect.MessageDescriptor {
	return md_TagTarget
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TagTarget) Type() protoreflect.MessageType {
	return _fastReflection_TagTarget_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TagTarget) New() protoreflect.Message {
	return new(fastReflection_TagTarget)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TagTarget) Interface() protoreflect.ProtoMessage {
	return (*TagTarget)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TagTarget) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *TagTarget_ClassId:
			v := o.ClassId
			value := protoreflect.ValueOfString(v)
			if !f(fd_TagTarget_class_id, value) {
				return
			}
		case *TagTarget_ProjectId:
			v := o.ProjectId
			value := protoreflect.ValueOfString(v)
			if !f(fd_TagTarget_project_id, value) {
				return
			}
		case *TagTarget_BatchId:
			v := o.BatchId
			value := protoreflect.ValueOfString(v)
			if !f(fd_TagTarget_batch_id, value) {
				return
			}
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
func (x *fastReflection_TagTarget) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagTarget.class_id":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*TagTarget_ClassId); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.curation.v1beta1.TagTarget.project_id":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*TagTarget_ProjectId); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.curation.v1beta1.TagTarget.batch_id":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*TagTarget_BatchId); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagTarget"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagTarget does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TagTarget) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagTarget.class_id":
		x.Sum = nil
	case "regen.ecocredit.curation.v1beta1.TagTarget.project_id":
		x.Sum = nil
	case "regen.ecocredit.curation.v1beta1.TagTarget.batch_id":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagTarget"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagTarget does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TagTarget) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagTarget.class_id":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*TagTarget_ClassId); ok {
			return protoreflect.ValueOfString(v.ClassId)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.curation.v1beta1.TagTarget.project_id":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*TagTarget_ProjectId); ok {
			return protoreflect.ValueOfString(v.ProjectId)
		} else {
			return protoreflect.ValueOfString("")
		}
	case "regen.ecocredit.curation.v1beta1.TagTarget.batch_id":
		if x.Sum == nil {
			return protoreflect.ValueOfString("")
		} else if v, ok := x.Sum.(*TagTarget_BatchId); ok {
			return protoreflect.ValueOfString(v.BatchId)
		} else {
			return protoreflect.ValueOfString("")
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagTarget"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagTarget does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TagTarget) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagTarget.class_id":
		cv := value.Interface().(string)
		x.Sum = &TagTarget_ClassId{ClassId: cv}
	case "regen.ecocredit.curation.v1beta1.TagTarget.project_id":
		cv := value.Interface().(string)
		x.Sum = &TagTarget_ProjectId{ProjectId: cv}
	case "regen.ecocredit.curation.v1beta1.TagTarget.batch_id":
		cv := value.Interface().(string)
		x.Sum = &TagTarget_BatchId{BatchId: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagTarget"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagTarget does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TagTarget) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagTarget.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.curation.v1beta1.TagTarget is not mutable"))
	case "regen.ecocredit.curation.v1beta1.TagTarget.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.curation.v1beta1.TagTarget is not mutable"))
	case "regen.ecocredit.curation.v1beta1.TagTarget.batch_id":
		panic(fmt.Errorf("field batch_id of message regen.ecocredit.curation.v1beta1.TagTarget is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagTarget"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagTarget does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TagTarget) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagTarget.class_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.curation.v1beta1.TagTarget.project_id":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.curation.v1beta1.TagTarget.batch_id":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.curation.v1beta1.TagTarget"))
		}
		panic(fmt.Errorf("message regen.ecocredit.curation.v1beta1.TagTarget does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TagTarget) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.ecocredit.curation.v1beta1.TagTarget.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *TagTarget_ClassId:
			return x.Descriptor().Fields().ByName("class_id")
		case *TagTarget_ProjectId:
			return x.Descriptor().Fields().ByName("project_id")
		case *TagTarget_BatchId:
			return x.Descriptor().Fields().ByName("batch_id")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.curation.v1beta1.TagTarget", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TagTarget) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TagTarget) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TagTarget) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TagTarget) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TagTarget)
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
		switch x := x.Sum.(type) {
		case *TagTarget_ClassId:
			if x == nil {
				break
			}
			l = len(x.ClassId)
			n += 1 + l + runtime.Sov(uint64(l))
		case *TagTarget_ProjectId:
			if x == nil {
				break
			}
			l = len(x.ProjectId)
			n += 1 + l + runtime.Sov(uint64(l))
		case *TagTarget_BatchId:
			if x == nil {
				break
			}
			l = len(x.BatchId)
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
		x := input.Message.Interface().(*TagTarget)
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
		switch x := x.Sum.(type) {
		case *TagTarget_ClassId:
			i -= len(x.ClassId)
			copy(dAtA[i:], x.ClassId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ClassId)))
			i--
			dAtA[i] = 0xa
		case *TagTarget_ProjectId:
			i -= len(x.ProjectId)
			copy(dAtA[i:], x.ProjectId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectId)))
			i--
			dAtA[i] = 0x12
		case *TagTarget_BatchId:
			i -= len(x.BatchId)
			copy(dAtA[i:], x.BatchId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchId)))
			i--
			dAtA[i] = 0x1a
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
		x := input.Message.Interface().(*TagTarget)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TagTarget: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TagTarget: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
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
				x.Sum = &TagTarget_ClassId{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
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
				x.Sum = &TagTarget_ProjectId{string(dAtA[iNdEx:postIndex])}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchId", wireType)
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
				x.Sum = &TagTarget_BatchId{string(dAtA[iNdEx:postIndex])}
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
// source: regen/ecocredit/curation/v1beta1/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TagTarget represents the target of a tag.
type TagTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sum:
	//	*TagTarget_ClassId
	//	*TagTarget_ProjectId
	//	*TagTarget_BatchId
	Sum isTagTarget_Sum `protobuf_oneof:"sum"`
}

func (x *TagTarget) Reset() {
	*x = TagTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_curation_v1beta1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagTarget) ProtoMessage() {}

// Deprecated: Use TagTarget.ProtoReflect.Descriptor instead.
func (*TagTarget) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_curation_v1beta1_types_proto_rawDescGZIP(), []int{0}
}

func (x *TagTarget) GetSum() isTagTarget_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *TagTarget) GetClassId() string {
	if x, ok := x.GetSum().(*TagTarget_ClassId); ok {
		return x.ClassId
	}
	return ""
}

func (x *TagTarget) GetProjectId() string {
	if x, ok := x.GetSum().(*TagTarget_ProjectId); ok {
		return x.ProjectId
	}
	return ""
}

func (x *TagTarget) GetBatchId() string {
	if x, ok := x.GetSum().(*TagTarget_BatchId); ok {
		return x.BatchId
	}
	return ""
}

type isTagTarget_Sum interface {
	isTagTarget_Sum()
}

type TagTarget_ClassId struct {
	// class_id indicates that this tag targets all the credits within a given class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3,oneof"`
}

type TagTarget_ProjectId struct {
	// project_id indicates that this tag targets all the credits within a given project.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3,oneof"`
}

type TagTarget_BatchId struct {
	// batch_id indicates that this tag targets the credits within a batch.
	BatchId string `protobuf:"bytes,3,opt,name=batch_id,json=batchId,proto3,oneof"`
}

func (*TagTarget_ClassId) isTagTarget_Sum() {}

func (*TagTarget_ProjectId) isTagTarget_Sum() {}

func (*TagTarget_BatchId) isTagTarget_Sum() {}

var File_regen_ecocredit_curation_v1beta1_types_proto protoreflect.FileDescriptor

var file_regen_ecocredit_curation_v1beta1_types_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x22, 0x6d, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x0a,
	0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x42,
	0xb1, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x3b, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x45, 0x43, 0xaa, 0x02, 0x20, 0x52, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x43, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02,
	0x2c, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x5c, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23,
	0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x3a, 0x3a, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_curation_v1beta1_types_proto_rawDescOnce sync.Once
	file_regen_ecocredit_curation_v1beta1_types_proto_rawDescData = file_regen_ecocredit_curation_v1beta1_types_proto_rawDesc
)

func file_regen_ecocredit_curation_v1beta1_types_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_curation_v1beta1_types_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_curation_v1beta1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_curation_v1beta1_types_proto_rawDescData)
	})
	return file_regen_ecocredit_curation_v1beta1_types_proto_rawDescData
}

var file_regen_ecocredit_curation_v1beta1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_regen_ecocredit_curation_v1beta1_types_proto_goTypes = []interface{}{
	(*TagTarget)(nil), // 0: regen.ecocredit.curation.v1beta1.TagTarget
}
var file_regen_ecocredit_curation_v1beta1_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_curation_v1beta1_types_proto_init() }
func file_regen_ecocredit_curation_v1beta1_types_proto_init() {
	if File_regen_ecocredit_curation_v1beta1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_curation_v1beta1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagTarget); i {
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
	file_regen_ecocredit_curation_v1beta1_types_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TagTarget_ClassId)(nil),
		(*TagTarget_ProjectId)(nil),
		(*TagTarget_BatchId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_curation_v1beta1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_curation_v1beta1_types_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_curation_v1beta1_types_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_curation_v1beta1_types_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_curation_v1beta1_types_proto = out.File
	file_regen_ecocredit_curation_v1beta1_types_proto_rawDesc = nil
	file_regen_ecocredit_curation_v1beta1_types_proto_goTypes = nil
	file_regen_ecocredit_curation_v1beta1_types_proto_depIdxs = nil
}
