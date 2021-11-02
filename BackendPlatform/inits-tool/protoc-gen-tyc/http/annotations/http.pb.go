// Code generated by protoc-gen-go. DO NOT EDIT.
// source: http.proto

package annotations

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// `HttpRule` defines the mapping of an RPC method to one or more HTTP
// REST API methods. The mapping specifies how different portions of the RPC
// request message are mapped to URL path, URL query parameters, and
// HTTP request body. The mapping is typically specified as an
// `google.api.http` annotation on the RPC method,
// see "google/api/annotations.proto" for details.
//
// The mapping consists of a field specifying the path template and
// method kind.  The path template can refer to fields in the request
// message, as in the example below which describes a REST GET
// operation on a resource collection of messages:
//
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http).get = "/v1/messages/{message_id}/{sub.subfield}";
//       }
//     }
//     message GetMessageRequest {
//       message SubMessage {
//         string subfield = 1;
//       }
//       string message_id = 1; // mapped to the URL
//       SubMessage sub = 2;    // `sub.subfield` is url-mapped
//     }
//     message Message {
//       string text = 1; // content of the resource
//     }
//
// The same http annotation can alternatively be expressed inside the
// `GRPC API Configuration` YAML file.
//
//     http:
//       rules:
//         - selector: <proto_package_name>.Messaging.GetMessage
//           get: /v1/messages/{message_id}/{sub.subfield}
//
// This definition enables an automatic, bidrectional mapping of HTTP
// JSON to RPC. Example:
//
// HTTP | RPC
// -----|-----
// `GET /v1/messages/123456/foo`  | `GetMessage(message_id: "123456" sub: SubMessage(subfield: "foo"))`
//
// In general, not only fields but also field paths can be referenced
// from a path pattern. Fields mapped to the path pattern cannot be
// repeated and must have a primitive (non-message) type.
//
// Any fields in the request message which are not bound by the path
// pattern automatically become (optional) HTTP query
// parameters. Assume the following definition of the request message:
//
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http).get = "/v1/messages/{message_id}";
//       }
//     }
//     message GetMessageRequest {
//       message SubMessage {
//         string subfield = 1;
//       }
//       string message_id = 1; // mapped to the URL
//       int64 revision = 2;    // becomes a parameter
//       SubMessage sub = 3;    // `sub.subfield` becomes a parameter
//     }
//
//
// This enables a HTTP JSON to RPC mapping as below:
//
// HTTP | RPC
// -----|-----
// `GET /v1/messages/123456?revision=2&sub.subfield=foo` | `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield: "foo"))`
//
// Note that fields which are mapped to HTTP parameters must have a
// primitive type or a repeated primitive type. Message types are not
// allowed. In the case of a repeated type, the parameter can be
// repeated in the URL, as in `...?param=A&param=B`.
//
// For HTTP method kinds which allow a request body, the `body` field
// specifies the mapping. Consider a REST update method on the
// message resource collection:
//
//
//     service Messaging {
//       rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
//         option (google.api.http) = {
//           put: "/v1/messages/{message_id}"
//           body: "message"
//         };
//       }
//     }
//     message UpdateMessageRequest {
//       string message_id = 1; // mapped to the URL
//       Message message = 2;   // mapped to the body
//     }
//
//
// The following HTTP JSON to RPC mapping is enabled, where the
// representation of the JSON in the request body is determined by
// protos JSON encoding:
//
// HTTP | RPC
// -----|-----
// `PUT /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id: "123456" message { text: "Hi!" })`
//
// The special name `*` can be used in the body mapping to define that
// every field not bound by the path template should be mapped to the
// request body.  This enables the following alternative definition of
// the update method:
//
//     service Messaging {
//       rpc UpdateMessage(Message) returns (Message) {
//         option (google.api.http) = {
//           put: "/v1/messages/{message_id}"
//           body: "*"
//         };
//       }
//     }
//     message Message {
//       string message_id = 1;
//       string text = 2;
//     }
//
//
// The following HTTP JSON to RPC mapping is enabled:
//
// HTTP | RPC
// -----|-----
// `PUT /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id: "123456" text: "Hi!")`
//
// Note that when using `*` in the body mapping, it is not possible to
// have HTTP parameters, as all fields not bound by the path end in
// the body. This makes this option more rarely used in practice of
// defining REST APIs. The common usage of `*` is in custom methods
// which don't use the URL at all for transferring data.
//
// It is possible to define multiple HTTP methods for one RPC by using
// the `additional_bindings` option. Example:
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http) = {
//           get: "/v1/messages/{message_id}"
//           additional_bindings {
//             get: "/v1/users/{user_id}/messages/{message_id}"
//           }
//         };
//       }
//     }
//     message GetMessageRequest {
//       string message_id = 1;
//       string user_id = 2;
//     }
//
//
// This enables the following two alternative HTTP JSON to RPC
// mappings:
//
// HTTP | RPC
// -----|-----
// `GET /v1/messages/123456` | `GetMessage(message_id: "123456")`
// `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id: "123456")`
//
// # Rules for HTTP mapping
//
// The rules for mapping HTTP path, query parameters, and body fields
// to the request message are as follows:
//
// 1. The `body` field specifies either `*` or a field path, or is
//    omitted. If omitted, it indicates there is no HTTP request body.
// 2. Leaf fields (recursive expansion of nested messages in the
//    request) can be classified into three types:
//     (a) Matched in the URL template.
//     (b) Covered by body (if body is `*`, everything except (a) fields;
//         else everything under the body field)
//     (c) All other fields.
// 3. URL query parameters found in the HTTP request are mapped to (c) fields.
// 4. Any body sent with an HTTP request can contain only (b) fields.
//
// The syntax of the path template is as follows:
//
//     Template = "/" Segments [ Verb ] ;
//     Segments = Segment { "/" Segment } ;
//     Segment  = "*" | "**" | LITERAL | Variable ;
//     Variable = "{" FieldPath [ "=" Segments ] "}" ;
//     FieldPath = IDENT { "." IDENT } ;
//     Verb     = ":" LITERAL ;
//
// The syntax `*` matches a single path segment. The syntax `**` matches zero
// or more path segments, which must be the last part of the path except the
// `Verb`. The syntax `LITERAL` matches literal text in the path.
//
// The syntax `Variable` matches part of the URL path as specified by its
// template. A variable template must not contain other variables. If a variable
// matches a single path segment, its template may be omitted, e.g. `{var}`
// is equivalent to `{var=*}`.
//
// If a variable contains exactly one path segment, such as `"{var}"` or
// `"{var=*}"`, when such a variable is expanded into a URL path, all characters
// except `[-_.~0-9a-zA-Z]` are percent-encoded. Such variables show up in the
// Discovery Document as `{var}`.
//
// If a variable contains one or more path segments, such as `"{var=foo/*}"`
// or `"{var=**}"`, when such a variable is expanded into a URL path, all
// characters except `[-_.~/0-9a-zA-Z]` are percent-encoded. Such variables
// show up in the Discovery Document as `{+var}`.
//
// NOTE: While the single segment variable matches the semantics of
// [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2
// Simple String Expansion, the multi segment variable **does not** match
// RFC 6570 Reserved Expansion. The reason is that the Reserved Expansion
// does not expand special characters like `?` and `#`, which would lead
// to invalid URLs.
//
// NOTE: the field paths in variables and in the `body` must not refer to
// repeated fields or map fields.
type HttpRule struct {
	// Selects methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector *string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// Determines the URL pattern is matched by this rules. This pattern can be
	// used with any of the {get|put|post|delete|patch} methods. A custom method
	// can be defined using the 'custom' field.
	//oneof pattern {
	//  // Used for listing and getting information about resources.
	Get *string `protobuf:"bytes,2,opt,name=get" json:"get,omitempty"`
	// Used for updating a resource.
	Put *string `protobuf:"bytes,3,opt,name=put" json:"put,omitempty"`
	// Used for creating a resource.
	Post *string `protobuf:"bytes,4,opt,name=post" json:"post,omitempty"`
	// Used for deleting a resource.
	Delete *string `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	// Used for updating a resource.
	Patch *string `protobuf:"bytes,6,opt,name=patch" json:"patch,omitempty"`
	//  // The custom pattern is used for specifying an HTTP method that is not
	//  // included in the `pattern` field, such as HEAD, or "*" to leave the
	//  // HTTP method unspecified for this rule. The wild-card rule is useful
	//  // for services that provide content to Web (HTML) clients.
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom" json:"custom,omitempty"`
	// The name of the request field whose value is mapped to the HTTP body, or
	// `*` for mapping all fields not captured by the path pattern to the HTTP
	// body. NOTE: the referred field must not be a repeated field and must be
	// present at the top-level of request message type.
	Body *string `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	// Optional. The name of the response field whose value is mapped to the HTTP
	// body of response. Other response fields are ignored. When
	// not set, the response message will be used as HTTP body of response.
	ResponseBody *string `protobuf:"bytes,12,opt,name=response_body,json=responseBody" json:"response_body,omitempty"`
	// Additional HTTP bindings for the selector. Nested bindings must
	// not contain an `additional_bindings` field themselves (that is,
	// the nesting may only be one level deep).
	AdditionalBindings   []*HttpRule `protobuf:"bytes,11,rep,name=additional_bindings,json=additionalBindings" json:"additional_bindings,omitempty"`
	Method               *string     `protobuf:"bytes,13,opt,name=method" json:"method,omitempty"`
	Pattern              *string     `protobuf:"bytes,14,opt,name=pattern" json:"pattern,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HttpRule) Reset()         { *m = HttpRule{} }
func (m *HttpRule) String() string { return proto.CompactTextString(m) }
func (*HttpRule) ProtoMessage()    {}
func (*HttpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_11b04836674e6f94, []int{0}
}

func (m *HttpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRule.Unmarshal(m, b)
}
func (m *HttpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRule.Marshal(b, m, deterministic)
}
func (m *HttpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRule.Merge(m, src)
}
func (m *HttpRule) XXX_Size() int {
	return xxx_messageInfo_HttpRule.Size(m)
}
func (m *HttpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRule proto.InternalMessageInfo

func (m *HttpRule) GetSelector() string {
	if m != nil && m.Selector != nil {
		return *m.Selector
	}
	return ""
}

func (m *HttpRule) GetGet() string {
	if m != nil && m.Get != nil {
		return *m.Get
	}
	return ""
}

func (m *HttpRule) GetPut() string {
	if m != nil && m.Put != nil {
		return *m.Put
	}
	return ""
}

func (m *HttpRule) GetPost() string {
	if m != nil && m.Post != nil {
		return *m.Post
	}
	return ""
}

func (m *HttpRule) GetDelete() string {
	if m != nil && m.Delete != nil {
		return *m.Delete
	}
	return ""
}

func (m *HttpRule) GetPatch() string {
	if m != nil && m.Patch != nil {
		return *m.Patch
	}
	return ""
}

func (m *HttpRule) GetCustom() *CustomHttpPattern {
	if m != nil {
		return m.Custom
	}
	return nil
}

func (m *HttpRule) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *HttpRule) GetResponseBody() string {
	if m != nil && m.ResponseBody != nil {
		return *m.ResponseBody
	}
	return ""
}

func (m *HttpRule) GetAdditionalBindings() []*HttpRule {
	if m != nil {
		return m.AdditionalBindings
	}
	return nil
}

func (m *HttpRule) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *HttpRule) GetPattern() string {
	if m != nil && m.Pattern != nil {
		return *m.Pattern
	}
	return ""
}

// A custom pattern is used for defining custom HTTP verb.
type CustomHttpPattern struct {
	// The name of this custom HTTP verb.
	Kind *string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	// The path matched by this custom verb.
	Path                 *string  `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomHttpPattern) Reset()         { *m = CustomHttpPattern{} }
func (m *CustomHttpPattern) String() string { return proto.CompactTextString(m) }
func (*CustomHttpPattern) ProtoMessage()    {}
func (*CustomHttpPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_11b04836674e6f94, []int{1}
}

func (m *CustomHttpPattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomHttpPattern.Unmarshal(m, b)
}
func (m *CustomHttpPattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomHttpPattern.Marshal(b, m, deterministic)
}
func (m *CustomHttpPattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomHttpPattern.Merge(m, src)
}
func (m *CustomHttpPattern) XXX_Size() int {
	return xxx_messageInfo_CustomHttpPattern.Size(m)
}
func (m *CustomHttpPattern) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomHttpPattern.DiscardUnknown(m)
}

var xxx_messageInfo_CustomHttpPattern proto.InternalMessageInfo

func (m *CustomHttpPattern) GetKind() string {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return ""
}

func (m *CustomHttpPattern) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         51234,
	Name:          "inits.api.http",
	Tag:           "bytes,51234,opt,name=http",
	Filename:      "http.proto",
}

func init() {
	proto.RegisterType((*HttpRule)(nil), "inits.api.HttpRule")
	proto.RegisterType((*CustomHttpPattern)(nil), "inits.api.CustomHttpPattern")
	proto.RegisterExtension(E_Http)
}

func init() { proto.RegisterFile("http.proto", fileDescriptor_11b04836674e6f94) }

var fileDescriptor_11b04836674e6f94 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x55, 0x9a, 0x34, 0x0d, 0x4e, 0x8a, 0xc0, 0xa0, 0xca, 0xea, 0x01, 0xa2, 0x72, 0xc9, 0xa5,
	0x5e, 0xa9, 0x17, 0x24, 0x7a, 0x0b, 0x48, 0x70, 0x00, 0x81, 0x82, 0xb8, 0x70, 0xa9, 0x1c, 0x7b,
	0x70, 0xac, 0x6e, 0x3c, 0xd6, 0x7a, 0xf6, 0xd0, 0xbf, 0xe0, 0x1b, 0xf8, 0x4b, 0x6e, 0xc8, 0xb3,
	0xbb, 0x25, 0x08, 0xf5, 0xb4, 0x6f, 0xde, 0x78, 0xfc, 0xe6, 0xbd, 0xb5, 0x10, 0x3b, 0xa2, 0xa4,
	0x53, 0x83, 0x84, 0x72, 0xe1, 0x0c, 0x44, 0x68, 0xee, 0xb2, 0x36, 0x29, 0x9c, 0xbf, 0xf3, 0x81,
	0x76, 0xed, 0x56, 0x5b, 0xdc, 0x57, 0x1e, 0x6b, 0x13, 0x7d, 0xc5, 0xc7, 0xb6, 0xed, 0x8f, 0x0e,
	0xd8, 0x4b, 0x0f, 0xf1, 0xd2, 0x63, 0xe5, 0x20, 0xdb, 0x26, 0x24, 0xc2, 0xe6, 0x00, 0x76, 0x77,
	0x5e, 0xfc, 0x3e, 0x12, 0xb3, 0x0f, 0x44, 0x69, 0xd3, 0xd6, 0x20, 0xcf, 0xc5, 0x2c, 0x43, 0x0d,
	0x96, 0xb0, 0x51, 0xa3, 0xe5, 0x68, 0xf5, 0x68, 0x73, 0x5f, 0xcb, 0x27, 0x62, 0xec, 0x81, 0xd4,
	0x11, 0xd3, 0x05, 0x16, 0x26, 0xb5, 0xa4, 0xc6, 0x1d, 0x93, 0x5a, 0x92, 0x52, 0x4c, 0x12, 0x66,
	0x52, 0x13, 0xa6, 0x18, 0xcb, 0x33, 0x31, 0x75, 0x50, 0x03, 0x81, 0x3a, 0x66, 0xb6, 0xaf, 0xe4,
	0x73, 0x71, 0x9c, 0x0c, 0xd9, 0x9d, 0x9a, 0x32, 0xdd, 0x15, 0xf2, 0xb5, 0x98, 0xda, 0x36, 0x13,
	0xee, 0xd5, 0x6c, 0x39, 0x5a, 0xcd, 0xaf, 0x5e, 0xea, 0x43, 0xcf, 0xfa, 0x2d, 0xf7, 0xca, 0xbe,
	0x5f, 0x0c, 0x11, 0x34, 0x71, 0xd3, 0x1f, 0x2f, 0xd2, 0x5b, 0x74, 0x77, 0xea, 0xa4, 0x93, 0x2e,
	0x58, 0xbe, 0x12, 0xa7, 0x0d, 0xe4, 0x84, 0x31, 0xc3, 0x0d, 0x37, 0x17, 0xdc, 0x5c, 0x0c, 0xe4,
	0xba, 0x1c, 0x7a, 0x2f, 0x9e, 0x19, 0xe7, 0x02, 0x05, 0x8c, 0xa6, 0xbe, 0xd9, 0x86, 0xe8, 0x42,
	0xf4, 0x59, 0xcd, 0x97, 0xe3, 0xd5, 0xfc, 0xea, 0xec, 0x5f, 0xf9, 0x21, 0xa8, 0x8d, 0xfc, 0x3b,
	0xb2, 0xee, 0x27, 0x8a, 0xd1, 0x3d, 0xd0, 0x0e, 0x9d, 0x3a, 0xed, 0x8c, 0x76, 0x95, 0x54, 0xe2,
	0x24, 0x75, 0xcb, 0xaa, 0xc7, 0xdc, 0x18, 0xca, 0x8b, 0x6b, 0xf1, 0xf4, 0x3f, 0x43, 0xc5, 0xc8,
	0x6d, 0x88, 0xae, 0xcf, 0x9f, 0x31, 0xe7, 0x6a, 0x68, 0xd7, 0x87, 0xcf, 0xf8, 0xcd, 0x47, 0x31,
	0x29, 0x4f, 0x43, 0xbe, 0xd0, 0x1e, 0xd1, 0xd7, 0xa0, 0x87, 0x9f, 0xaf, 0x3f, 0xb1, 0xee, 0xe7,
	0x54, 0xf6, 0xcb, 0xea, 0xd7, 0xcf, 0x31, 0x27, 0xf9, 0x90, 0x15, 0xbe, 0x65, 0xfd, 0xed, 0xfb,
	0x57, 0x1f, 0x48, 0x87, 0x78, 0x0b, 0xda, 0xc6, 0xaa, 0x7c, 0x6b, 0xf4, 0xc1, 0x56, 0xc3, 0x50,
	0x65, 0xf7, 0xee, 0xf0, 0x59, 0xdd, 0xf3, 0x65, 0xbc, 0x32, 0x31, 0x22, 0x19, 0x16, 0xbc, 0x3e,
	0xc0, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x51, 0xf0, 0x4d, 0xbe, 0x02, 0x00, 0x00,
}