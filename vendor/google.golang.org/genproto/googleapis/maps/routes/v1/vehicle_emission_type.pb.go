// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/maps/routes/v1/vehicle_emission_type.proto

package routes

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A set of values describing the vehicle's emission type.
// Applies only to the DRIVE travel mode.
type VehicleEmissionType int32

const (
	// No emission type specified. Default to GASOLINE.
	VehicleEmissionType_VEHICLE_EMISSION_TYPE_UNSPECIFIED VehicleEmissionType = 0
	// Gasoline/petrol fueled vehicle.
	VehicleEmissionType_GASOLINE VehicleEmissionType = 1
	// Electricity powered vehicle.
	VehicleEmissionType_ELECTRIC VehicleEmissionType = 2
	// Hybrid fuel (such as gasoline + electric) vehicle.
	VehicleEmissionType_HYBRID VehicleEmissionType = 3
)

var VehicleEmissionType_name = map[int32]string{
	0: "VEHICLE_EMISSION_TYPE_UNSPECIFIED",
	1: "GASOLINE",
	2: "ELECTRIC",
	3: "HYBRID",
}

var VehicleEmissionType_value = map[string]int32{
	"VEHICLE_EMISSION_TYPE_UNSPECIFIED": 0,
	"GASOLINE":                          1,
	"ELECTRIC":                          2,
	"HYBRID":                            3,
}

func (x VehicleEmissionType) String() string {
	return proto.EnumName(VehicleEmissionType_name, int32(x))
}

func (VehicleEmissionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3eaf884fc84e5c7, []int{0}
}

func init() {
	proto.RegisterEnum("google.maps.routes.v1.VehicleEmissionType", VehicleEmissionType_name, VehicleEmissionType_value)
}

func init() {
	proto.RegisterFile("google/maps/routes/v1/vehicle_emission_type.proto", fileDescriptor_e3eaf884fc84e5c7)
}

var fileDescriptor_e3eaf884fc84e5c7 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0x4d, 0x2c, 0x28, 0xd6, 0x2f, 0xca, 0x2f, 0x2d, 0x49, 0x2d, 0xd6, 0x2f,
	0x33, 0xd4, 0x2f, 0x4b, 0xcd, 0xc8, 0x4c, 0xce, 0x49, 0x8d, 0x4f, 0xcd, 0xcd, 0x2c, 0x2e, 0xce,
	0xcc, 0xcf, 0x8b, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x85,
	0x68, 0xd1, 0x03, 0x69, 0xd1, 0x83, 0x68, 0xd1, 0x2b, 0x33, 0xd4, 0x4a, 0xe1, 0x12, 0x0e, 0x83,
	0xe8, 0x72, 0x85, 0x6a, 0x0a, 0xa9, 0x2c, 0x48, 0x15, 0x52, 0xe5, 0x52, 0x0c, 0x73, 0xf5, 0xf0,
	0x74, 0xf6, 0x71, 0x8d, 0x77, 0xf5, 0xf5, 0x0c, 0x0e, 0xf6, 0xf4, 0xf7, 0x8b, 0x0f, 0x89, 0x0c,
	0x70, 0x8d, 0x0f, 0xf5, 0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x10,
	0xe2, 0xe1, 0xe2, 0x70, 0x77, 0x0c, 0xf6, 0xf7, 0xf1, 0xf4, 0x73, 0x15, 0x60, 0x04, 0xf1, 0x5c,
	0x7d, 0x5c, 0x9d, 0x43, 0x82, 0x3c, 0x9d, 0x05, 0x98, 0x84, 0xb8, 0xb8, 0xd8, 0x3c, 0x22, 0x9d,
	0x82, 0x3c, 0x5d, 0x04, 0x98, 0x9d, 0xd6, 0x31, 0x72, 0x49, 0x26, 0xe7, 0xe7, 0xea, 0x61, 0x75,
	0x83, 0x93, 0x04, 0x16, 0x17, 0x04, 0x80, 0x1c, 0x1d, 0xc0, 0x18, 0x65, 0x0d, 0xd5, 0x92, 0x9e,
	0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f, 0x9e, 0x9a, 0x07, 0xf6, 0x92, 0x3e,
	0x44, 0x2a, 0xb1, 0x20, 0xb3, 0x18, 0x2d, 0x20, 0xac, 0x21, 0xac, 0x1f, 0x8c, 0x8c, 0x8b, 0x98,
	0x58, 0xdc, 0x7d, 0x83, 0x82, 0x57, 0x31, 0x89, 0xba, 0x43, 0xcc, 0xf1, 0x05, 0x59, 0x1d, 0x04,
	0xb1, 0x3a, 0xcc, 0xf0, 0x14, 0x4c, 0x3c, 0x06, 0x24, 0x1e, 0x03, 0x11, 0x8f, 0x09, 0x33, 0x4c,
	0x62, 0x03, 0xdb, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x97, 0xfc, 0x9e, 0x13, 0x69, 0x01,
	0x00, 0x00,
}
