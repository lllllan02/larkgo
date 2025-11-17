package core

import "encoding/json"

type Serializable interface {
	// Serialize 序列化
	Serialize(v interface{}) ([]byte, error)

	// Deserialize 反序列化
	Deserialize(data []byte, v interface{}) error
}

type defaultSerialization struct{}

func (d *defaultSerialization) Serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (d *defaultSerialization) Deserialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
