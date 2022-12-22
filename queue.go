package core

import (
	"encoding/base64"
	"strings"

	"github.com/golang/protobuf/proto"
)

// SerializeEvent uses protobuf to marshal given event
func SerializeEvent(event proto.Message) (string, error) {

	protoContent, err := proto.Marshal(event)
	return base64.StdEncoding.EncodeToString(protoContent), err
}

// DeserializeEvent wil try to unmarshal given event data to passed event.
func DeserializeEvent(data string, event proto.Message) error {

	stringData := strings.TrimSuffix(strings.TrimPrefix(data, "\""), "\"")
	protoData, err := base64.StdEncoding.DecodeString(stringData)
	if err != nil {
		return err
	}
	return proto.Unmarshal(protoData, event)
}
