package core

import (
	"encoding/json"
	"github.com/yametech/echoer/pkg/utils"
	"time"
)

type Kind string

type Metadata struct {
	Name    string `json:"name" bson:"name"`
	Kind    Kind   `json:"kind"  bson:"kind"`
	Version int64  `json:"version" bson:"version"`
	UUID    string `json:"uuid" bson:"uuid"`

	Labels map[string]interface{} `json:"labels" bson:"labels"`
}

func (m *Metadata) Clone() IObject {
	panic("implement me")
}

func (m *Metadata) SetLabel(key string, value interface{}) {
	if m.Labels == nil {
		m.Labels = make(map[string]interface{})
	}
	m.Labels[key] = value
}

func (m *Metadata) GetUUID() string {
	return m.UUID
}

func (m *Metadata) GetResourceVersion() int64 {
	return m.Version
}

func (m *Metadata) GetName() string {
	return m.Name
}

func (m *Metadata) GetKind() Kind {
	return m.Kind
}

func (m *Metadata) GenerateVersion() IObject {
	m.Version = time.Now().Unix()
	if m.UUID == "" {
		m.UUID = utils.NewSUID().String()
	}
	return m
}

func Clone(src, tag interface{}) {
	b, _ := json.Marshal(src)
	_ = json.Unmarshal(b, tag)
}

func ToMap(i interface{}) (map[string]interface{}, error) {
	var result = make(map[string]interface{})
	bs, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bs, &result); err != nil {
		return nil, err
	}
	return result, err
}

func EncodeFromMap(i interface{}, m map[string]interface{}) error {
	bs, err := json.Marshal(&m)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bs, i); err != nil {
		return err
	}
	return nil
}

func UnmarshalInterfaceToResource(src interface{}, dest IObject) error {
	bs, err := json.Marshal(src)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bs, dest); err != nil {
		return err
	}
	return nil
}

type IObject interface {
	GetKind() Kind
	GetName() string
	Clone() IObject
	GenerateVersion() IObject
	GetResourceVersion() int64
	GetUUID() string
}
