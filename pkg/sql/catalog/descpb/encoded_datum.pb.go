// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/catalog/descpb/encoded_datum.proto

package descpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// DatumEncoding identifies the encoding used for an EncDatum.
type DatumEncoding int32

const (
	// Indicates that the datum is encoded using the order-preserving encoding
	// used for keys (ascending order).
	DatumEncoding_ASCENDING_KEY DatumEncoding = 0
	// Indicates that the datum is encoded using the order-preserving encoding
	// used for keys (descending order).
	DatumEncoding_DESCENDING_KEY DatumEncoding = 1
	// Indicates that the datum is encoded using the encoding used for values.
	DatumEncoding_VALUE DatumEncoding = 2
)

var DatumEncoding_name = map[int32]string{
	0: "ASCENDING_KEY",
	1: "DESCENDING_KEY",
	2: "VALUE",
}

var DatumEncoding_value = map[string]int32{
	"ASCENDING_KEY":  0,
	"DESCENDING_KEY": 1,
	"VALUE":          2,
}

func (x DatumEncoding) Enum() *DatumEncoding {
	p := new(DatumEncoding)
	*p = x
	return p
}

func (x DatumEncoding) String() string {
	return proto.EnumName(DatumEncoding_name, int32(x))
}

func (x *DatumEncoding) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DatumEncoding_value, data, "DatumEncoding")
	if err != nil {
		return err
	}
	*x = DatumEncoding(value)
	return nil
}

func (DatumEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_efe5a7af8eada883, []int{0}
}

func init() {
	proto.RegisterEnum("cockroach.sql.sqlbase.DatumEncoding", DatumEncoding_name, DatumEncoding_value)
}

func init() {
	proto.RegisterFile("sql/catalog/descpb/encoded_datum.proto", fileDescriptor_efe5a7af8eada883)
}

var fileDescriptor_efe5a7af8eada883 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x2e, 0xcc, 0xd1,
	0x4f, 0x4e, 0x2c, 0x49, 0xcc, 0xc9, 0x4f, 0xd7, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0x48, 0xd2, 0x4f,
	0xcd, 0x4b, 0xce, 0x4f, 0x49, 0x4d, 0x89, 0x4f, 0x49, 0x2c, 0x29, 0xcd, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x4d, 0xce, 0x4f, 0xce, 0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0, 0x2b, 0x2e,
	0xcc, 0x01, 0xe1, 0xa4, 0xc4, 0xe2, 0x54, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x0a, 0x7d,
	0x10, 0x0b, 0xa2, 0x58, 0xcb, 0x91, 0x8b, 0xd7, 0x05, 0xa4, 0xd7, 0x15, 0x64, 0x50, 0x66, 0x5e,
	0xba, 0x90, 0x20, 0x17, 0xaf, 0x63, 0xb0, 0xb3, 0xab, 0x9f, 0x8b, 0xa7, 0x9f, 0x7b, 0xbc, 0xb7,
	0x6b, 0xa4, 0x00, 0x83, 0x90, 0x10, 0x17, 0x9f, 0x8b, 0x2b, 0x8a, 0x18, 0xa3, 0x10, 0x27, 0x17,
	0x6b, 0x98, 0xa3, 0x4f, 0xa8, 0xab, 0x00, 0x93, 0x93, 0xc6, 0x89, 0x87, 0x72, 0x0c, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0x78, 0xe3, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0x1b, 0xc4, 0xb5,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xdf, 0x29, 0x66, 0xc2, 0x00, 0x00, 0x00,
}
