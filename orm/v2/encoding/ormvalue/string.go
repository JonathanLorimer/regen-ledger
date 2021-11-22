package ormvalue

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type StringCodec struct{}

func (s StringCodec) FixedSize() int {
	return -1
}

func (s StringCodec) Size(value protoreflect.Value) (int, error) {
	return len(value.String()), nil
}

func (s StringCodec) IsOrdered() bool {
	return true
}

func (s StringCodec) Compare(v1, v2 protoreflect.Value) int {
	return strings.Compare(v1.String(), v2.String())
}

func (s StringCodec) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	bz, err := io.ReadAll(r)
	return protoreflect.ValueOfString(string(bz)), err
}

func (s StringCodec) Encode(value protoreflect.Value, w io.Writer) error {
	_, err := w.Write([]byte(value.String()))
	return err
}

func (s StringCodec) IsEmpty(value protoreflect.Value) bool {
	return value.String() == ""
}

type NonTerminalStringCodec struct{}

func (s NonTerminalStringCodec) FixedSize() int {
	return -1
}

func (s NonTerminalStringCodec) Size(value protoreflect.Value) (int, error) {
	return len(value.String()) + 1, nil
}

func (s NonTerminalStringCodec) IsOrdered() bool {
	return true
}

func (s NonTerminalStringCodec) IsEmpty(value protoreflect.Value) bool {
	return value.String() == ""
}

func (s NonTerminalStringCodec) Compare(v1, v2 protoreflect.Value) int {
	return strings.Compare(v1.String(), v2.String())
}

func (s NonTerminalStringCodec) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	var bz []byte
	for {
		b, err := r.ReadByte()
		if b == 0 || err == io.EOF {
			return protoreflect.ValueOfString(string(bz)), err
		}
		bz = append(bz, b)
	}
}

func (s NonTerminalStringCodec) Encode(value protoreflect.Value, w io.Writer) error {
	str := value.String()
	bz := []byte(str)
	for _, b := range bz {
		if b == 0 {
			return fmt.Errorf("illegal null terminator found in index string: %s", str)
		}
	}
	_, err := w.Write([]byte(str))
	if err != nil {
		return err
	}
	_, err = w.Write(nullTerminator)
	return err
}

var nullTerminator = []byte{0}
