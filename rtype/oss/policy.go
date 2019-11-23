package oss

import (
	"encoding/json"
	"fmt"
	"time"
	"wx-server/rtype/oss/condition"
)

type Policy struct {
	Version   string
	Statement []Statement
	Condition condition.Condition
}

type PolicyWxm struct {
	Expiration time.Time             `json:"expiration"`
	Conditions []condition.Condition `json:"conditions"`
}

type Statement struct {
	Action   []Action
	Resource []Resource
	Effect   Effect
}

func (p *Policy) String() string {
	data, _ := json.Marshal(p)
	return string(data)
}

type Effect string

const (
	Allow Effect = "Allow"
	Deny  Effect = "Deny"
)

// https://help.aliyun.com/document_detail/100680.html?spm=a2c4g.11186623.2.8.770d41f0vpNmjY#concept-y5r-5rm-2gb
type Action string

/**
Resource 指代的是 OSS 上面的某个具体的资源或者某些资源（支持*通配），
resource的规则是acs:oss:{region}:{bucket_owner}:{bucket_name}/{object_name}。

对于所有 Bucket 级别的操作来说不需要最后的斜杠和{object_name}，
即acs:oss:{region}:{bucket_owner}:{bucket_name}。Resource 也是一个列表，
可以有多个 Resource。其中的 region 字段暂时不做支持，设置为*。
*/
type Resource string

func NewPolicy() *Policy {
	return &Policy{
		Version:   "1",
		Statement: make([]Statement, 0),
	}
}

func (p *Policy) AddStatement(statement Statement) *Policy {
	p.Statement = append(p.Statement, statement)
	return p
}

func NewStatement(effect Effect) *Statement {
	return &Statement{
		Effect:   effect,
		Action:   make([]Action, 0),
		Resource: make([]Resource, 0),
	}
}

func (statement *Statement) AddAction(action Action) *Statement {
	statement.Action = append(statement.Action, action)
	return statement
}
func (statement *Statement) AddResource(resource Resource) *Statement {
	statement.Resource = append(statement.Resource, resource)
	return statement
}
func (statement *Statement) AddOssResource(bucket, baseUrl string) *Statement {
	statement.Resource = append(statement.Resource, Resource(fmt.Sprintf("acs:oss:*:*:%v/%v/*", bucket, baseUrl)))
	return statement
}
