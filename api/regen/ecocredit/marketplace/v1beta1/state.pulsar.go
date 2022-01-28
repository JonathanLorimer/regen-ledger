package marketplacev1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_SellOrder                     protoreflect.MessageDescriptor
	fd_SellOrder_id                  protoreflect.FieldDescriptor
	fd_SellOrder_seller              protoreflect.FieldDescriptor
	fd_SellOrder_batch_denom         protoreflect.FieldDescriptor
	fd_SellOrder_quantity            protoreflect.FieldDescriptor
	fd_SellOrder_market_id           protoreflect.FieldDescriptor
	fd_SellOrder_ask_price           protoreflect.FieldDescriptor
	fd_SellOrder_disable_auto_retire protoreflect.FieldDescriptor
	fd_SellOrder_expiration          protoreflect.FieldDescriptor
	fd_SellOrder_maker               protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_state_proto_init()
	md_SellOrder = File_regen_ecocredit_marketplace_v1beta1_state_proto.Messages().ByName("SellOrder")
	fd_SellOrder_id = md_SellOrder.Fields().ByName("id")
	fd_SellOrder_seller = md_SellOrder.Fields().ByName("seller")
	fd_SellOrder_batch_denom = md_SellOrder.Fields().ByName("batch_denom")
	fd_SellOrder_quantity = md_SellOrder.Fields().ByName("quantity")
	fd_SellOrder_market_id = md_SellOrder.Fields().ByName("market_id")
	fd_SellOrder_ask_price = md_SellOrder.Fields().ByName("ask_price")
	fd_SellOrder_disable_auto_retire = md_SellOrder.Fields().ByName("disable_auto_retire")
	fd_SellOrder_expiration = md_SellOrder.Fields().ByName("expiration")
	fd_SellOrder_maker = md_SellOrder.Fields().ByName("maker")
}

var _ protoreflect.Message = (*fastReflection_SellOrder)(nil)

type fastReflection_SellOrder SellOrder

func (x *SellOrder) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SellOrder)(x)
}

func (x *SellOrder) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SellOrder_messageType fastReflection_SellOrder_messageType
var _ protoreflect.MessageType = fastReflection_SellOrder_messageType{}

type fastReflection_SellOrder_messageType struct{}

func (x fastReflection_SellOrder_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SellOrder)(nil)
}
func (x fastReflection_SellOrder_messageType) New() protoreflect.Message {
	return new(fastReflection_SellOrder)
}
func (x fastReflection_SellOrder_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SellOrder
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SellOrder) Descriptor() protoreflect.MessageDescriptor {
	return md_SellOrder
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SellOrder) Type() protoreflect.MessageType {
	return _fastReflection_SellOrder_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SellOrder) New() protoreflect.Message {
	return new(fastReflection_SellOrder)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SellOrder) Interface() protoreflect.ProtoMessage {
	return (*SellOrder)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SellOrder) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_SellOrder_id, value) {
			return
		}
	}
	if len(x.Seller) != 0 {
		value := protoreflect.ValueOfBytes(x.Seller)
		if !f(fd_SellOrder_seller, value) {
			return
		}
	}
	if x.BatchDenom != "" {
		value := protoreflect.ValueOfString(x.BatchDenom)
		if !f(fd_SellOrder_batch_denom, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_SellOrder_quantity, value) {
			return
		}
	}
	if x.MarketId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.MarketId)
		if !f(fd_SellOrder_market_id, value) {
			return
		}
	}
	if x.AskPrice != nil {
		value := protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
		if !f(fd_SellOrder_ask_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_SellOrder_disable_auto_retire, value) {
			return
		}
	}
	if x.Expiration != nil {
		value := protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
		if !f(fd_SellOrder_expiration, value) {
			return
		}
	}
	if x.Maker != false {
		value := protoreflect.ValueOfBool(x.Maker)
		if !f(fd_SellOrder_maker, value) {
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
func (x *fastReflection_SellOrder) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.id":
		return x.Id != uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.seller":
		return len(x.Seller) != 0
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.batch_denom":
		return x.BatchDenom != ""
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.market_id":
		return x.MarketId != uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.ask_price":
		return x.AskPrice != nil
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.expiration":
		return x.Expiration != nil
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.maker":
		return x.Maker != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.SellOrder does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SellOrder) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.id":
		x.Id = uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.seller":
		x.Seller = nil
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.batch_denom":
		x.BatchDenom = ""
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.quantity":
		x.Quantity = ""
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.market_id":
		x.MarketId = uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.ask_price":
		x.AskPrice = nil
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.expiration":
		x.Expiration = nil
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.maker":
		x.Maker = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.SellOrder does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SellOrder) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.seller":
		value := x.Seller
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.batch_denom":
		value := x.BatchDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.market_id":
		value := x.MarketId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.ask_price":
		value := x.AskPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.expiration":
		value := x.Expiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.maker":
		value := x.Maker
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.SellOrder does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SellOrder) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.id":
		x.Id = value.Uint()
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.seller":
		x.Seller = value.Bytes()
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.batch_denom":
		x.BatchDenom = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.market_id":
		x.MarketId = value.Uint()
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.ask_price":
		x.AskPrice = value.Message().Interface().(*v1beta1.Coin)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.expiration":
		x.Expiration = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.maker":
		x.Maker = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.SellOrder does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SellOrder) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.ask_price":
		if x.AskPrice == nil {
			x.AskPrice = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.AskPrice.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.expiration":
		if x.Expiration == nil {
			x.Expiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.id":
		panic(fmt.Errorf("field id of message regen.ecocredit.marketplace.v1beta1.SellOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.seller":
		panic(fmt.Errorf("field seller of message regen.ecocredit.marketplace.v1beta1.SellOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.batch_denom":
		panic(fmt.Errorf("field batch_denom of message regen.ecocredit.marketplace.v1beta1.SellOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.marketplace.v1beta1.SellOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.market_id":
		panic(fmt.Errorf("field market_id of message regen.ecocredit.marketplace.v1beta1.SellOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.marketplace.v1beta1.SellOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.maker":
		panic(fmt.Errorf("field maker of message regen.ecocredit.marketplace.v1beta1.SellOrder is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.SellOrder does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SellOrder) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.seller":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.batch_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.market_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.ask_price":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.SellOrder.maker":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.SellOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.SellOrder does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SellOrder) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.SellOrder", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SellOrder) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SellOrder) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SellOrder) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SellOrder) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SellOrder)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Seller)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BatchDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MarketId != 0 {
			n += 1 + runtime.Sov(uint64(x.MarketId))
		}
		if x.AskPrice != nil {
			l = options.Size(x.AskPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.Expiration != nil {
			l = options.Size(x.Expiration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Maker {
			n += 2
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
		x := input.Message.Interface().(*SellOrder)
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
		if x.Maker {
			i--
			if x.Maker {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x50
		}
		if x.Expiration != nil {
			encoded, err := options.Marshal(x.Expiration)
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
			dAtA[i] = 0x4a
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x38
		}
		if x.AskPrice != nil {
			encoded, err := options.Marshal(x.AskPrice)
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
		if x.MarketId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MarketId))
			i--
			dAtA[i] = 0x28
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.BatchDenom) > 0 {
			i -= len(x.BatchDenom)
			copy(dAtA[i:], x.BatchDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BatchDenom)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Seller) > 0 {
			i -= len(x.Seller)
			copy(dAtA[i:], x.Seller)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Seller)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
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
		x := input.Message.Interface().(*SellOrder)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SellOrder: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SellOrder: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
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
				x.Seller = append(x.Seller[:0], dAtA[iNdEx:postIndex]...)
				if x.Seller == nil {
					x.Seller = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
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
				x.BatchDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
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
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
				}
				x.MarketId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MarketId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AskPrice", wireType)
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
				if x.AskPrice == nil {
					x.AskPrice = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AskPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisableAutoRetire = bool(v != 0)
			case 9:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
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
				if x.Expiration == nil {
					x.Expiration = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Expiration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 10:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Maker", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Maker = bool(v != 0)
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
	md_BuyOrder                      protoreflect.MessageDescriptor
	fd_BuyOrder_id                   protoreflect.FieldDescriptor
	fd_BuyOrder_buyer                protoreflect.FieldDescriptor
	fd_BuyOrder_selection            protoreflect.FieldDescriptor
	fd_BuyOrder_quantity             protoreflect.FieldDescriptor
	fd_BuyOrder_market_id            protoreflect.FieldDescriptor
	fd_BuyOrder_bid_price            protoreflect.FieldDescriptor
	fd_BuyOrder_disable_auto_retire  protoreflect.FieldDescriptor
	fd_BuyOrder_disable_partial_fill protoreflect.FieldDescriptor
	fd_BuyOrder_expiration           protoreflect.FieldDescriptor
	fd_BuyOrder_maker                protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_state_proto_init()
	md_BuyOrder = File_regen_ecocredit_marketplace_v1beta1_state_proto.Messages().ByName("BuyOrder")
	fd_BuyOrder_id = md_BuyOrder.Fields().ByName("id")
	fd_BuyOrder_buyer = md_BuyOrder.Fields().ByName("buyer")
	fd_BuyOrder_selection = md_BuyOrder.Fields().ByName("selection")
	fd_BuyOrder_quantity = md_BuyOrder.Fields().ByName("quantity")
	fd_BuyOrder_market_id = md_BuyOrder.Fields().ByName("market_id")
	fd_BuyOrder_bid_price = md_BuyOrder.Fields().ByName("bid_price")
	fd_BuyOrder_disable_auto_retire = md_BuyOrder.Fields().ByName("disable_auto_retire")
	fd_BuyOrder_disable_partial_fill = md_BuyOrder.Fields().ByName("disable_partial_fill")
	fd_BuyOrder_expiration = md_BuyOrder.Fields().ByName("expiration")
	fd_BuyOrder_maker = md_BuyOrder.Fields().ByName("maker")
}

var _ protoreflect.Message = (*fastReflection_BuyOrder)(nil)

type fastReflection_BuyOrder BuyOrder

func (x *BuyOrder) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BuyOrder)(x)
}

func (x *BuyOrder) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BuyOrder_messageType fastReflection_BuyOrder_messageType
var _ protoreflect.MessageType = fastReflection_BuyOrder_messageType{}

type fastReflection_BuyOrder_messageType struct{}

func (x fastReflection_BuyOrder_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BuyOrder)(nil)
}
func (x fastReflection_BuyOrder_messageType) New() protoreflect.Message {
	return new(fastReflection_BuyOrder)
}
func (x fastReflection_BuyOrder_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BuyOrder) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BuyOrder) Type() protoreflect.MessageType {
	return _fastReflection_BuyOrder_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BuyOrder) New() protoreflect.Message {
	return new(fastReflection_BuyOrder)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BuyOrder) Interface() protoreflect.ProtoMessage {
	return (*BuyOrder)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BuyOrder) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_BuyOrder_id, value) {
			return
		}
	}
	if len(x.Buyer) != 0 {
		value := protoreflect.ValueOfBytes(x.Buyer)
		if !f(fd_BuyOrder_buyer, value) {
			return
		}
	}
	if x.Selection != nil {
		value := protoreflect.ValueOfMessage(x.Selection.ProtoReflect())
		if !f(fd_BuyOrder_selection, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_BuyOrder_quantity, value) {
			return
		}
	}
	if x.MarketId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.MarketId)
		if !f(fd_BuyOrder_market_id, value) {
			return
		}
	}
	if x.BidPrice != "" {
		value := protoreflect.ValueOfString(x.BidPrice)
		if !f(fd_BuyOrder_bid_price, value) {
			return
		}
	}
	if x.DisableAutoRetire != false {
		value := protoreflect.ValueOfBool(x.DisableAutoRetire)
		if !f(fd_BuyOrder_disable_auto_retire, value) {
			return
		}
	}
	if x.DisablePartialFill != false {
		value := protoreflect.ValueOfBool(x.DisablePartialFill)
		if !f(fd_BuyOrder_disable_partial_fill, value) {
			return
		}
	}
	if x.Expiration != nil {
		value := protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
		if !f(fd_BuyOrder_expiration, value) {
			return
		}
	}
	if x.Maker != false {
		value := protoreflect.ValueOfBool(x.Maker)
		if !f(fd_BuyOrder_maker, value) {
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
func (x *fastReflection_BuyOrder) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.id":
		return x.Id != uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.buyer":
		return len(x.Buyer) != 0
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.selection":
		return x.Selection != nil
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.quantity":
		return x.Quantity != ""
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.market_id":
		return x.MarketId != uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.bid_price":
		return x.BidPrice != ""
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_auto_retire":
		return x.DisableAutoRetire != false
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_partial_fill":
		return x.DisablePartialFill != false
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.expiration":
		return x.Expiration != nil
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.maker":
		return x.Maker != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.id":
		x.Id = uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.buyer":
		x.Buyer = nil
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.selection":
		x.Selection = nil
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.quantity":
		x.Quantity = ""
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.market_id":
		x.MarketId = uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.bid_price":
		x.BidPrice = ""
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_auto_retire":
		x.DisableAutoRetire = false
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_partial_fill":
		x.DisablePartialFill = false
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.expiration":
		x.Expiration = nil
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.maker":
		x.Maker = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BuyOrder) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.buyer":
		value := x.Buyer
		return protoreflect.ValueOfBytes(value)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.selection":
		value := x.Selection
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.market_id":
		value := x.MarketId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.bid_price":
		value := x.BidPrice
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_auto_retire":
		value := x.DisableAutoRetire
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_partial_fill":
		value := x.DisablePartialFill
		return protoreflect.ValueOfBool(value)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.expiration":
		value := x.Expiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.maker":
		value := x.Maker
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BuyOrder) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.id":
		x.Id = value.Uint()
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.buyer":
		x.Buyer = value.Bytes()
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.selection":
		x.Selection = value.Message().Interface().(*BuyOrder_Selection)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.quantity":
		x.Quantity = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.market_id":
		x.MarketId = value.Uint()
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.bid_price":
		x.BidPrice = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_auto_retire":
		x.DisableAutoRetire = value.Bool()
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_partial_fill":
		x.DisablePartialFill = value.Bool()
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.expiration":
		x.Expiration = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.maker":
		x.Maker = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BuyOrder) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.selection":
		if x.Selection == nil {
			x.Selection = new(BuyOrder_Selection)
		}
		return protoreflect.ValueOfMessage(x.Selection.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.expiration":
		if x.Expiration == nil {
			x.Expiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.id":
		panic(fmt.Errorf("field id of message regen.ecocredit.marketplace.v1beta1.BuyOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.buyer":
		panic(fmt.Errorf("field buyer of message regen.ecocredit.marketplace.v1beta1.BuyOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.quantity":
		panic(fmt.Errorf("field quantity of message regen.ecocredit.marketplace.v1beta1.BuyOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.market_id":
		panic(fmt.Errorf("field market_id of message regen.ecocredit.marketplace.v1beta1.BuyOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.bid_price":
		panic(fmt.Errorf("field bid_price of message regen.ecocredit.marketplace.v1beta1.BuyOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_auto_retire":
		panic(fmt.Errorf("field disable_auto_retire of message regen.ecocredit.marketplace.v1beta1.BuyOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_partial_fill":
		panic(fmt.Errorf("field disable_partial_fill of message regen.ecocredit.marketplace.v1beta1.BuyOrder is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.maker":
		panic(fmt.Errorf("field maker of message regen.ecocredit.marketplace.v1beta1.BuyOrder is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BuyOrder) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.buyer":
		return protoreflect.ValueOfBytes(nil)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.selection":
		m := new(BuyOrder_Selection)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.quantity":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.market_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.bid_price":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_auto_retire":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.disable_partial_fill":
		return protoreflect.ValueOfBool(false)
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.maker":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BuyOrder) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.BuyOrder", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BuyOrder) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BuyOrder) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BuyOrder) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BuyOrder)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Buyer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Selection != nil {
			l = options.Size(x.Selection)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MarketId != 0 {
			n += 1 + runtime.Sov(uint64(x.MarketId))
		}
		l = len(x.BidPrice)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DisableAutoRetire {
			n += 2
		}
		if x.DisablePartialFill {
			n += 2
		}
		if x.Expiration != nil {
			l = options.Size(x.Expiration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Maker {
			n += 2
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
		x := input.Message.Interface().(*BuyOrder)
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
		if x.Maker {
			i--
			if x.Maker {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x50
		}
		if x.Expiration != nil {
			encoded, err := options.Marshal(x.Expiration)
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
			dAtA[i] = 0x4a
		}
		if x.DisablePartialFill {
			i--
			if x.DisablePartialFill {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x40
		}
		if x.DisableAutoRetire {
			i--
			if x.DisableAutoRetire {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x38
		}
		if len(x.BidPrice) > 0 {
			i -= len(x.BidPrice)
			copy(dAtA[i:], x.BidPrice)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BidPrice)))
			i--
			dAtA[i] = 0x32
		}
		if x.MarketId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MarketId))
			i--
			dAtA[i] = 0x28
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x22
		}
		if x.Selection != nil {
			encoded, err := options.Marshal(x.Selection)
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
			dAtA[i] = 0x1a
		}
		if len(x.Buyer) > 0 {
			i -= len(x.Buyer)
			copy(dAtA[i:], x.Buyer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Buyer)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
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
		x := input.Message.Interface().(*BuyOrder)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
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
				x.Buyer = append(x.Buyer[:0], dAtA[iNdEx:postIndex]...)
				if x.Buyer == nil {
					x.Buyer = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Selection", wireType)
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
				if x.Selection == nil {
					x.Selection = &BuyOrder_Selection{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Selection); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
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
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
				}
				x.MarketId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MarketId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BidPrice", wireType)
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
				x.BidPrice = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisableAutoRetire = bool(v != 0)
			case 8:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisablePartialFill", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.DisablePartialFill = bool(v != 0)
			case 9:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
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
				if x.Expiration == nil {
					x.Expiration = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Expiration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 10:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Maker", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Maker = bool(v != 0)
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
	md_BuyOrder_Selection               protoreflect.MessageDescriptor
	fd_BuyOrder_Selection_sell_order_id protoreflect.FieldDescriptor
	fd_BuyOrder_Selection_filter        protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_state_proto_init()
	md_BuyOrder_Selection = File_regen_ecocredit_marketplace_v1beta1_state_proto.Messages().ByName("BuyOrder").Messages().ByName("Selection")
	fd_BuyOrder_Selection_sell_order_id = md_BuyOrder_Selection.Fields().ByName("sell_order_id")
	fd_BuyOrder_Selection_filter = md_BuyOrder_Selection.Fields().ByName("filter")
}

var _ protoreflect.Message = (*fastReflection_BuyOrder_Selection)(nil)

type fastReflection_BuyOrder_Selection BuyOrder_Selection

func (x *BuyOrder_Selection) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BuyOrder_Selection)(x)
}

func (x *BuyOrder_Selection) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BuyOrder_Selection_messageType fastReflection_BuyOrder_Selection_messageType
var _ protoreflect.MessageType = fastReflection_BuyOrder_Selection_messageType{}

type fastReflection_BuyOrder_Selection_messageType struct{}

func (x fastReflection_BuyOrder_Selection_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BuyOrder_Selection)(nil)
}
func (x fastReflection_BuyOrder_Selection_messageType) New() protoreflect.Message {
	return new(fastReflection_BuyOrder_Selection)
}
func (x fastReflection_BuyOrder_Selection_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder_Selection
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BuyOrder_Selection) Descriptor() protoreflect.MessageDescriptor {
	return md_BuyOrder_Selection
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BuyOrder_Selection) Type() protoreflect.MessageType {
	return _fastReflection_BuyOrder_Selection_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BuyOrder_Selection) New() protoreflect.Message {
	return new(fastReflection_BuyOrder_Selection)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BuyOrder_Selection) Interface() protoreflect.ProtoMessage {
	return (*BuyOrder_Selection)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BuyOrder_Selection) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			v := o.SellOrderId
			value := protoreflect.ValueOfUint64(v)
			if !f(fd_BuyOrder_Selection_sell_order_id, value) {
				return
			}
		case *BuyOrder_Selection_Filter:
			v := o.Filter
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_BuyOrder_Selection_filter, value) {
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
func (x *fastReflection_BuyOrder_Selection) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.sell_order_id":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*BuyOrder_Selection_SellOrderId); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.filter":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*BuyOrder_Selection_Filter); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder_Selection) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.sell_order_id":
		x.Sum = nil
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.filter":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BuyOrder_Selection) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.sell_order_id":
		if x.Sum == nil {
			return protoreflect.ValueOfUint64(uint64(0))
		} else if v, ok := x.Sum.(*BuyOrder_Selection_SellOrderId); ok {
			return protoreflect.ValueOfUint64(v.SellOrderId)
		} else {
			return protoreflect.ValueOfUint64(uint64(0))
		}
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.filter":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*Filter)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*BuyOrder_Selection_Filter); ok {
			return protoreflect.ValueOfMessage(v.Filter.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Filter)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BuyOrder_Selection) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.sell_order_id":
		cv := value.Uint()
		x.Sum = &BuyOrder_Selection_SellOrderId{SellOrderId: cv}
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.filter":
		cv := value.Message().Interface().(*Filter)
		x.Sum = &BuyOrder_Selection_Filter{Filter: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BuyOrder_Selection) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.filter":
		if x.Sum == nil {
			value := &Filter{}
			oneofValue := &BuyOrder_Selection_Filter{Filter: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *BuyOrder_Selection_Filter:
			return protoreflect.ValueOfMessage(m.Filter.ProtoReflect())
		default:
			value := &Filter{}
			oneofValue := &BuyOrder_Selection_Filter{Filter: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.sell_order_id":
		panic(fmt.Errorf("field sell_order_id of message regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BuyOrder_Selection) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.sell_order_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.filter":
		value := &Filter{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BuyOrder_Selection) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *BuyOrder_Selection_SellOrderId:
			return x.Descriptor().Fields().ByName("sell_order_id")
		case *BuyOrder_Selection_Filter:
			return x.Descriptor().Fields().ByName("filter")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BuyOrder_Selection) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BuyOrder_Selection) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BuyOrder_Selection) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BuyOrder_Selection) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BuyOrder_Selection)
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
		case *BuyOrder_Selection_SellOrderId:
			if x == nil {
				break
			}
			n += 1 + runtime.Sov(uint64(x.SellOrderId))
		case *BuyOrder_Selection_Filter:
			if x == nil {
				break
			}
			l = options.Size(x.Filter)
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
		x := input.Message.Interface().(*BuyOrder_Selection)
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
		case *BuyOrder_Selection_SellOrderId:
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SellOrderId))
			i--
			dAtA[i] = 0x8
		case *BuyOrder_Selection_Filter:
			encoded, err := options.Marshal(x.Filter)
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
		x := input.Message.Interface().(*BuyOrder_Selection)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder_Selection: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BuyOrder_Selection: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SellOrderId", wireType)
				}
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Sum = &BuyOrder_Selection_SellOrderId{v}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
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
				v := &Filter{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &BuyOrder_Selection_Filter{v}
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
	md_AllowedDenom               protoreflect.MessageDescriptor
	fd_AllowedDenom_bank_denom    protoreflect.FieldDescriptor
	fd_AllowedDenom_display_denom protoreflect.FieldDescriptor
	fd_AllowedDenom_exponent      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_state_proto_init()
	md_AllowedDenom = File_regen_ecocredit_marketplace_v1beta1_state_proto.Messages().ByName("AllowedDenom")
	fd_AllowedDenom_bank_denom = md_AllowedDenom.Fields().ByName("bank_denom")
	fd_AllowedDenom_display_denom = md_AllowedDenom.Fields().ByName("display_denom")
	fd_AllowedDenom_exponent = md_AllowedDenom.Fields().ByName("exponent")
}

var _ protoreflect.Message = (*fastReflection_AllowedDenom)(nil)

type fastReflection_AllowedDenom AllowedDenom

func (x *AllowedDenom) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AllowedDenom)(x)
}

func (x *AllowedDenom) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AllowedDenom_messageType fastReflection_AllowedDenom_messageType
var _ protoreflect.MessageType = fastReflection_AllowedDenom_messageType{}

type fastReflection_AllowedDenom_messageType struct{}

func (x fastReflection_AllowedDenom_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AllowedDenom)(nil)
}
func (x fastReflection_AllowedDenom_messageType) New() protoreflect.Message {
	return new(fastReflection_AllowedDenom)
}
func (x fastReflection_AllowedDenom_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AllowedDenom
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AllowedDenom) Descriptor() protoreflect.MessageDescriptor {
	return md_AllowedDenom
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AllowedDenom) Type() protoreflect.MessageType {
	return _fastReflection_AllowedDenom_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AllowedDenom) New() protoreflect.Message {
	return new(fastReflection_AllowedDenom)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AllowedDenom) Interface() protoreflect.ProtoMessage {
	return (*AllowedDenom)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AllowedDenom) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BankDenom != "" {
		value := protoreflect.ValueOfString(x.BankDenom)
		if !f(fd_AllowedDenom_bank_denom, value) {
			return
		}
	}
	if x.DisplayDenom != "" {
		value := protoreflect.ValueOfString(x.DisplayDenom)
		if !f(fd_AllowedDenom_display_denom, value) {
			return
		}
	}
	if x.Exponent != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Exponent)
		if !f(fd_AllowedDenom_exponent, value) {
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
func (x *fastReflection_AllowedDenom) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.bank_denom":
		return x.BankDenom != ""
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.display_denom":
		return x.DisplayDenom != ""
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.exponent":
		return x.Exponent != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.AllowedDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.AllowedDenom does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AllowedDenom) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.bank_denom":
		x.BankDenom = ""
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.display_denom":
		x.DisplayDenom = ""
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.exponent":
		x.Exponent = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.AllowedDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.AllowedDenom does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AllowedDenom) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.bank_denom":
		value := x.BankDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.display_denom":
		value := x.DisplayDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.exponent":
		value := x.Exponent
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.AllowedDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.AllowedDenom does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_AllowedDenom) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.bank_denom":
		x.BankDenom = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.display_denom":
		x.DisplayDenom = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.exponent":
		x.Exponent = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.AllowedDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.AllowedDenom does not contain field %s", fd.FullName()))
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
func (x *fastReflection_AllowedDenom) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.bank_denom":
		panic(fmt.Errorf("field bank_denom of message regen.ecocredit.marketplace.v1beta1.AllowedDenom is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.display_denom":
		panic(fmt.Errorf("field display_denom of message regen.ecocredit.marketplace.v1beta1.AllowedDenom is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.exponent":
		panic(fmt.Errorf("field exponent of message regen.ecocredit.marketplace.v1beta1.AllowedDenom is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.AllowedDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.AllowedDenom does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AllowedDenom) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.bank_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.display_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1beta1.AllowedDenom.exponent":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.AllowedDenom"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.AllowedDenom does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AllowedDenom) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.AllowedDenom", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AllowedDenom) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AllowedDenom) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_AllowedDenom) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AllowedDenom) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AllowedDenom)
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
		l = len(x.BankDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DisplayDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Exponent != 0 {
			n += 1 + runtime.Sov(uint64(x.Exponent))
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
		x := input.Message.Interface().(*AllowedDenom)
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
		if x.Exponent != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Exponent))
			i--
			dAtA[i] = 0x18
		}
		if len(x.DisplayDenom) > 0 {
			i -= len(x.DisplayDenom)
			copy(dAtA[i:], x.DisplayDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DisplayDenom)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BankDenom) > 0 {
			i -= len(x.BankDenom)
			copy(dAtA[i:], x.BankDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BankDenom)))
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
		x := input.Message.Interface().(*AllowedDenom)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AllowedDenom: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AllowedDenom: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BankDenom", wireType)
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
				x.BankDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DisplayDenom", wireType)
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
				x.DisplayDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
				}
				x.Exponent = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Exponent |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
	md_Market                    protoreflect.MessageDescriptor
	fd_Market_id                 protoreflect.FieldDescriptor
	fd_Market_credit_type        protoreflect.FieldDescriptor
	fd_Market_bank_denom         protoreflect.FieldDescriptor
	fd_Market_precision_modifier protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_state_proto_init()
	md_Market = File_regen_ecocredit_marketplace_v1beta1_state_proto.Messages().ByName("Market")
	fd_Market_id = md_Market.Fields().ByName("id")
	fd_Market_credit_type = md_Market.Fields().ByName("credit_type")
	fd_Market_bank_denom = md_Market.Fields().ByName("bank_denom")
	fd_Market_precision_modifier = md_Market.Fields().ByName("precision_modifier")
}

var _ protoreflect.Message = (*fastReflection_Market)(nil)

type fastReflection_Market Market

func (x *Market) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Market)(x)
}

func (x *Market) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Market_messageType fastReflection_Market_messageType
var _ protoreflect.MessageType = fastReflection_Market_messageType{}

type fastReflection_Market_messageType struct{}

func (x fastReflection_Market_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Market)(nil)
}
func (x fastReflection_Market_messageType) New() protoreflect.Message {
	return new(fastReflection_Market)
}
func (x fastReflection_Market_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Market
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Market) Descriptor() protoreflect.MessageDescriptor {
	return md_Market
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Market) Type() protoreflect.MessageType {
	return _fastReflection_Market_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Market) New() protoreflect.Message {
	return new(fastReflection_Market)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Market) Interface() protoreflect.ProtoMessage {
	return (*Market)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Market) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Market_id, value) {
			return
		}
	}
	if x.CreditType != "" {
		value := protoreflect.ValueOfString(x.CreditType)
		if !f(fd_Market_credit_type, value) {
			return
		}
	}
	if x.BankDenom != "" {
		value := protoreflect.ValueOfString(x.BankDenom)
		if !f(fd_Market_bank_denom, value) {
			return
		}
	}
	if x.PrecisionModifier != uint32(0) {
		value := protoreflect.ValueOfUint32(x.PrecisionModifier)
		if !f(fd_Market_precision_modifier, value) {
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
func (x *fastReflection_Market) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Market.id":
		return x.Id != uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.Market.credit_type":
		return x.CreditType != ""
	case "regen.ecocredit.marketplace.v1beta1.Market.bank_denom":
		return x.BankDenom != ""
	case "regen.ecocredit.marketplace.v1beta1.Market.precision_modifier":
		return x.PrecisionModifier != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Market"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Market does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Market) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Market.id":
		x.Id = uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.Market.credit_type":
		x.CreditType = ""
	case "regen.ecocredit.marketplace.v1beta1.Market.bank_denom":
		x.BankDenom = ""
	case "regen.ecocredit.marketplace.v1beta1.Market.precision_modifier":
		x.PrecisionModifier = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Market"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Market does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Market) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Market.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.marketplace.v1beta1.Market.credit_type":
		value := x.CreditType
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1beta1.Market.bank_denom":
		value := x.BankDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1beta1.Market.precision_modifier":
		value := x.PrecisionModifier
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Market"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Market does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Market) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Market.id":
		x.Id = value.Uint()
	case "regen.ecocredit.marketplace.v1beta1.Market.credit_type":
		x.CreditType = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1beta1.Market.bank_denom":
		x.BankDenom = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1beta1.Market.precision_modifier":
		x.PrecisionModifier = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Market"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Market does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Market) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Market.id":
		panic(fmt.Errorf("field id of message regen.ecocredit.marketplace.v1beta1.Market is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.Market.credit_type":
		panic(fmt.Errorf("field credit_type of message regen.ecocredit.marketplace.v1beta1.Market is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.Market.bank_denom":
		panic(fmt.Errorf("field bank_denom of message regen.ecocredit.marketplace.v1beta1.Market is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.Market.precision_modifier":
		panic(fmt.Errorf("field precision_modifier of message regen.ecocredit.marketplace.v1beta1.Market is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Market"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Market does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Market) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Market.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.marketplace.v1beta1.Market.credit_type":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1beta1.Market.bank_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1beta1.Market.precision_modifier":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Market"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Market does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Market) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.Market", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Market) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Market) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Market) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Market) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Market)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.CreditType)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BankDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.PrecisionModifier != 0 {
			n += 1 + runtime.Sov(uint64(x.PrecisionModifier))
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
		x := input.Message.Interface().(*Market)
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
		if x.PrecisionModifier != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.PrecisionModifier))
			i--
			dAtA[i] = 0x20
		}
		if len(x.BankDenom) > 0 {
			i -= len(x.BankDenom)
			copy(dAtA[i:], x.BankDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BankDenom)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.CreditType) > 0 {
			i -= len(x.CreditType)
			copy(dAtA[i:], x.CreditType)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CreditType)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
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
		x := input.Message.Interface().(*Market)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Market: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreditType", wireType)
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
				x.CreditType = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BankDenom", wireType)
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
				x.BankDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PrecisionModifier", wireType)
				}
				x.PrecisionModifier = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.PrecisionModifier |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: regen/ecocredit/marketplace/v1beta1/state.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SellOrder represents the information for a sell order.
type SellOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the unique ID of sell order.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// seller is the address of the owner of the credits being sold.
	Seller []byte `protobuf:"bytes,2,opt,name=seller,proto3" json:"seller,omitempty"`
	// batch_denom is the credit batch being sold.
	BatchDenom string `protobuf:"bytes,3,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// quantity is the quantity of credits being sold.
	Quantity string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// market_id is the market in which this sell order exists and specifies
	// the back denom ask_price corresponds to.
	MarketId uint64 `protobuf:"varint,5,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// ask_price is the integer price the seller is asking for each unit of the
	// batch_denom. Each credit unit of the batch will be sold for at least the
	// ask_price or more.
	AskPrice *v1beta1.Coin `protobuf:"bytes,6,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	// disable_auto_retire disables auto-retirement of credits which allows a
	// buyer to disable auto-retirement in their buy order enabling them to
	// resell the credits to another buyer.
	DisableAutoRetire bool `protobuf:"varint,7,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// expiration is an optional timestamp when the sell order expires. When the
	// expiration time is reached, the sell order is removed from state.
	Expiration *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// maker indicates that this is a maker order, meaning that when it hit
	// the order book, there were no matching buy orders.
	Maker bool `protobuf:"varint,10,opt,name=maker,proto3" json:"maker,omitempty"`
}

func (x *SellOrder) Reset() {
	*x = SellOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellOrder) ProtoMessage() {}

// Deprecated: Use SellOrder.ProtoReflect.Descriptor instead.
func (*SellOrder) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescGZIP(), []int{0}
}

func (x *SellOrder) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SellOrder) GetSeller() []byte {
	if x != nil {
		return x.Seller
	}
	return nil
}

func (x *SellOrder) GetBatchDenom() string {
	if x != nil {
		return x.BatchDenom
	}
	return ""
}

func (x *SellOrder) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *SellOrder) GetMarketId() uint64 {
	if x != nil {
		return x.MarketId
	}
	return 0
}

func (x *SellOrder) GetAskPrice() *v1beta1.Coin {
	if x != nil {
		return x.AskPrice
	}
	return nil
}

func (x *SellOrder) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *SellOrder) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *SellOrder) GetMaker() bool {
	if x != nil {
		return x.Maker
	}
	return false
}

// BuyOrder represents the information for a buy order.
type BuyOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the unique ID of buy order.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// buyer is the address that created the buy order
	Buyer []byte `protobuf:"bytes,2,opt,name=buyer,proto3" json:"buyer,omitempty"`
	// selection is the buy order selection.
	Selection *BuyOrder_Selection `protobuf:"bytes,3,opt,name=selection,proto3" json:"selection,omitempty"`
	// quantity is the quantity of credits to buy. If the quantity of credits
	// available is less than this amount the order will be partially filled
	// unless disable_partial_fill is true.
	Quantity string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// market_id is the market in which this sell order exists and specifies
	// the back denom ask_price corresponds to.
	MarketId uint64 `protobuf:"varint,5,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// bid price is the integer bid price for this buy order. A credit unit will be
	// settled at a purchase price that is no more than the bid price. The
	// buy order will fail if the buyer does not have enough funds available
	// to complete the purchase.
	BidPrice string `protobuf:"bytes,6,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	// disable_auto_retire allows auto-retirement to be disabled. If it is set to true
	// the credits will not auto-retire and can be resold assuming that the
	// corresponding sell order has auto-retirement disabled. If the sell order
	// hasn't disabled auto-retirement and the buy order tries to disable it,
	// that buy order will fail.
	DisableAutoRetire bool `protobuf:"varint,7,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// disable_partial_fill disables the default behavior of partially filling
	// buy orders if the requested quantity is not available.
	DisablePartialFill bool `protobuf:"varint,8,opt,name=disable_partial_fill,json=disablePartialFill,proto3" json:"disable_partial_fill,omitempty"`
	// expiration is the optional timestamp when the buy order expires. When the
	// expiration time is reached, the buy order is removed from state.
	Expiration *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// maker indicates that this is a maker order, meaning that when it hit
	// the order book, there were no matching sell orders.
	Maker bool `protobuf:"varint,10,opt,name=maker,proto3" json:"maker,omitempty"`
}

func (x *BuyOrder) Reset() {
	*x = BuyOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrder) ProtoMessage() {}

// Deprecated: Use BuyOrder.ProtoReflect.Descriptor instead.
func (*BuyOrder) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescGZIP(), []int{1}
}

func (x *BuyOrder) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BuyOrder) GetBuyer() []byte {
	if x != nil {
		return x.Buyer
	}
	return nil
}

func (x *BuyOrder) GetSelection() *BuyOrder_Selection {
	if x != nil {
		return x.Selection
	}
	return nil
}

func (x *BuyOrder) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *BuyOrder) GetMarketId() uint64 {
	if x != nil {
		return x.MarketId
	}
	return 0
}

func (x *BuyOrder) GetBidPrice() string {
	if x != nil {
		return x.BidPrice
	}
	return ""
}

func (x *BuyOrder) GetDisableAutoRetire() bool {
	if x != nil {
		return x.DisableAutoRetire
	}
	return false
}

func (x *BuyOrder) GetDisablePartialFill() bool {
	if x != nil {
		return x.DisablePartialFill
	}
	return false
}

func (x *BuyOrder) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *BuyOrder) GetMaker() bool {
	if x != nil {
		return x.Maker
	}
	return false
}

// AllowedDenom represents the information for an allowed ask/bid denom.
type AllowedDenom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// denom is the bank denom to allow (ex. ibc/GLKHDSG423SGS)
	BankDenom string `protobuf:"bytes,1,opt,name=bank_denom,json=bankDenom,proto3" json:"bank_denom,omitempty"`
	// display_denom is the denom to display to the user and is informational.
	// Because the denom is likely an IBC denom, this should be chosen by
	// governance to represent the consensus trusted name of the denom.
	DisplayDenom string `protobuf:"bytes,2,opt,name=display_denom,json=displayDenom,proto3" json:"display_denom,omitempty"`
	// exponent is the exponent that relates the denom to the display_denom and is
	// informational
	Exponent uint32 `protobuf:"varint,3,opt,name=exponent,proto3" json:"exponent,omitempty"`
}

func (x *AllowedDenom) Reset() {
	*x = AllowedDenom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedDenom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedDenom) ProtoMessage() {}

// Deprecated: Use AllowedDenom.ProtoReflect.Descriptor instead.
func (*AllowedDenom) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescGZIP(), []int{2}
}

func (x *AllowedDenom) GetBankDenom() string {
	if x != nil {
		return x.BankDenom
	}
	return ""
}

func (x *AllowedDenom) GetDisplayDenom() string {
	if x != nil {
		return x.DisplayDenom
	}
	return ""
}

func (x *AllowedDenom) GetExponent() uint32 {
	if x != nil {
		return x.Exponent
	}
	return 0
}

// Market describes a distinctly processed market between a credit type and
// allowed bank denom. Each market has its own precision in the order book
// and is processed independently of other markets. Governance must enable
// markets one by one. Every additional enabled market potentially adds more
// processing overhead to the blockchain and potentially weakens liquidity in
// competing markets. For instance, enabling side by side USD/Carbon and
// EUR/Carbon markets may have the end result that each market individually has
// less liquidity and longer settlement times. Such decisions should be taken
// with care.
type Market struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the unique ID of the market.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// credit_type is the abbreviation of the credit type.
	CreditType string `protobuf:"bytes,2,opt,name=credit_type,json=creditType,proto3" json:"credit_type,omitempty"`
	// bank_denom is an allowed bank denom.
	BankDenom string `protobuf:"bytes,3,opt,name=bank_denom,json=bankDenom,proto3" json:"bank_denom,omitempty"`
	// precision_modifier is an optional modifier used to convert arbitrary
	// precision integer bank amounts to uint64 values used for sorting in the
	// order book. Given an arbitrary precision integer x, its uint64 conversion
	// will be x / 10^precision_modifier using round half away from zero
	// rounding.
	//
	// uint64 values range from 0 to 1,8446,744,073,709,551,615.
	// This allows for a full 18 digits of precision. In most real world cases,
	// a precision modifier of 0 (meaning no conversion) is probably sufficient.
	//
	// Consider a USD stable coin with 6 decimal digits of precision. A credit
	// would need to be worth over 1 trillion USD for a non-zero precision_modifier
	// to be needed.
	//
	// precision_modifier is provided to accomodate cases where the bank denom
	// has either a huge precision (such as 18 decimal digits on its own) or
	// extreme hyper-inflationary scenarios.
	//
	// In cases where there is a non-zero precision_modifier, bids and asks
	// which are rounded to the same number will be ordered equivalently and
	// differentiated only by time priority. The precision integer amount in the
	// order will still be used in settlement.
	//
	// In cases where an arbitrary precision integer overflows its conversion to
	// uint64, the order fill fail and the market will eventually become stuck
	// if prices are in this region. Governance should intervene at this
	// point to change the precision_modifier to a reasonable value to re-enable
	// the market. When precision_modifier changes, all active orders in the order
	// book will be updated to reflect this.
	PrecisionModifier uint32 `protobuf:"varint,4,opt,name=precision_modifier,json=precisionModifier,proto3" json:"precision_modifier,omitempty"`
}

func (x *Market) Reset() {
	*x = Market{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Market) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Market) ProtoMessage() {}

// Deprecated: Use Market.ProtoReflect.Descriptor instead.
func (*Market) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescGZIP(), []int{3}
}

func (x *Market) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Market) GetCreditType() string {
	if x != nil {
		return x.CreditType
	}
	return ""
}

func (x *Market) GetBankDenom() string {
	if x != nil {
		return x.BankDenom
	}
	return ""
}

func (x *Market) GetPrecisionModifier() uint32 {
	if x != nil {
		return x.PrecisionModifier
	}
	return 0
}

// Selection defines a buy order selection.
type BuyOrder_Selection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sum defines the type of selection.
	//
	// Types that are assignable to Sum:
	//	*BuyOrder_Selection_SellOrderId
	//	*BuyOrder_Selection_Filter
	Sum isBuyOrder_Selection_Sum `protobuf_oneof:"sum"`
}

func (x *BuyOrder_Selection) Reset() {
	*x = BuyOrder_Selection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrder_Selection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrder_Selection) ProtoMessage() {}

// Deprecated: Use BuyOrder_Selection.ProtoReflect.Descriptor instead.
func (*BuyOrder_Selection) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescGZIP(), []int{1, 0}
}

func (x *BuyOrder_Selection) GetSum() isBuyOrder_Selection_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *BuyOrder_Selection) GetSellOrderId() uint64 {
	if x, ok := x.GetSum().(*BuyOrder_Selection_SellOrderId); ok {
		return x.SellOrderId
	}
	return 0
}

func (x *BuyOrder_Selection) GetFilter() *Filter {
	if x, ok := x.GetSum().(*BuyOrder_Selection_Filter); ok {
		return x.Filter
	}
	return nil
}

type isBuyOrder_Selection_Sum interface {
	isBuyOrder_Selection_Sum()
}

type BuyOrder_Selection_SellOrderId struct {
	// sell_order_id is the sell order ID against which the buyer is trying to buy.
	// When sell_order_id is set, this is known as a direct buy order because it
	// is placed directly against a specific sell order.
	SellOrderId uint64 `protobuf:"varint,1,opt,name=sell_order_id,json=sellOrderId,proto3,oneof"`
}

type BuyOrder_Selection_Filter struct {
	// filter selects credits to buy based upon the specified filter criteria.
	Filter *Filter `protobuf:"bytes,2,opt,name=filter,proto3,oneof"`
}

func (*BuyOrder_Selection_SellOrderId) isBuyOrder_Selection_Sum() {}

func (*BuyOrder_Selection_Filter) isBuyOrder_Selection_Sum() {}

var File_regen_ecocredit_marketplace_v1beta1_state_proto protoreflect.FileDescriptor

var file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x23, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f,
	0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12,
	0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x61, 0x6b, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x61, 0x6b, 0x65,
	0x72, 0x3a, 0x3b, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x35, 0x0a, 0x04, 0x0a, 0x02, 0x69, 0x64, 0x12,
	0x0f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0x18, 0x01, 0x22, 0xbd,
	0x04, 0x0a, 0x08, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x75, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x75, 0x79, 0x65,
	0x72, 0x12, 0x55, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x75, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x74, 0x69, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x12, 0x30,
	0x0a, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x6c,
	0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x61, 0x6b,
	0x65, 0x72, 0x1a, 0x7f, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x6c, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x05, 0x0a, 0x03,
	0x73, 0x75, 0x6d, 0x3a, 0x29, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x23, 0x0a, 0x04, 0x0a, 0x02, 0x69,
	0x64, 0x12, 0x09, 0x0a, 0x05, 0x62, 0x75, 0x79, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x18, 0x02, 0x22, 0x9b,
	0x01, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x65,
	0x6e, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x3a,
	0x2b, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x25, 0x0a, 0x0c, 0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x5f,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x13, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x10, 0x01, 0x18, 0x01, 0x18, 0x03, 0x22, 0xb5, 0x01, 0x0a,
	0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b,
	0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61,
	0x6e, 0x6b, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x3a, 0x2c, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x26, 0x0a, 0x04,
	0x0a, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x16, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2c, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x10, 0x01,
	0x18, 0x01, 0x18, 0x04, 0x42, 0xc6, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f,
	0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x52, 0x45, 0x4d, 0xaa, 0x02, 0x23, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x45,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x23, 0x52,
	0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xe2, 0x02, 0x2f, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x5c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x26, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescOnce sync.Once
	file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescData = file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDesc
)

func file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescData)
	})
	return file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDescData
}

var file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_regen_ecocredit_marketplace_v1beta1_state_proto_goTypes = []interface{}{
	(*SellOrder)(nil),             // 0: regen.ecocredit.marketplace.v1beta1.SellOrder
	(*BuyOrder)(nil),              // 1: regen.ecocredit.marketplace.v1beta1.BuyOrder
	(*AllowedDenom)(nil),          // 2: regen.ecocredit.marketplace.v1beta1.AllowedDenom
	(*Market)(nil),                // 3: regen.ecocredit.marketplace.v1beta1.Market
	(*BuyOrder_Selection)(nil),    // 4: regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection
	(*v1beta1.Coin)(nil),          // 5: cosmos.base.v1beta1.Coin
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*Filter)(nil),                // 7: regen.ecocredit.marketplace.v1beta1.Filter
}
var file_regen_ecocredit_marketplace_v1beta1_state_proto_depIdxs = []int32{
	5, // 0: regen.ecocredit.marketplace.v1beta1.SellOrder.ask_price:type_name -> cosmos.base.v1beta1.Coin
	6, // 1: regen.ecocredit.marketplace.v1beta1.SellOrder.expiration:type_name -> google.protobuf.Timestamp
	4, // 2: regen.ecocredit.marketplace.v1beta1.BuyOrder.selection:type_name -> regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection
	6, // 3: regen.ecocredit.marketplace.v1beta1.BuyOrder.expiration:type_name -> google.protobuf.Timestamp
	7, // 4: regen.ecocredit.marketplace.v1beta1.BuyOrder.Selection.filter:type_name -> regen.ecocredit.marketplace.v1beta1.Filter
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_marketplace_v1beta1_state_proto_init() }
func file_regen_ecocredit_marketplace_v1beta1_state_proto_init() {
	if File_regen_ecocredit_marketplace_v1beta1_state_proto != nil {
		return
	}
	file_regen_ecocredit_marketplace_v1beta1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellOrder); i {
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
		file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrder); i {
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
		file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedDenom); i {
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
		file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Market); i {
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
		file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrder_Selection); i {
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
	file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*BuyOrder_Selection_SellOrderId)(nil),
		(*BuyOrder_Selection_Filter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_marketplace_v1beta1_state_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_marketplace_v1beta1_state_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_marketplace_v1beta1_state_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_marketplace_v1beta1_state_proto = out.File
	file_regen_ecocredit_marketplace_v1beta1_state_proto_rawDesc = nil
	file_regen_ecocredit_marketplace_v1beta1_state_proto_goTypes = nil
	file_regen_ecocredit_marketplace_v1beta1_state_proto_depIdxs = nil
}
