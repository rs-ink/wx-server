package condition

import "encoding/json"

type Operator string

const (
	OpeStringEquals              Operator = "StringEquals"
	OpeStringNotEquals           Operator = "StringNotEquals"
	OpeStringEqualsIgnoreCase    Operator = "StringEqualsIgnoreCase"
	OpeStringNotEqualsIgnoreCase Operator = "StringNotEqualsIgnoreCase"
	OpeStringLike                Operator = "StringLike"
	OpeStringNotLike             Operator = "StringNotLike"
	OpeNumericEquals             Operator = "NumericEquals"
	OpeNumericNotEquals          Operator = "NumericNotEquals"
	OpeNumericLessThan           Operator = "NumericLessThan"
	OpeNumericLessThanEquals     Operator = "NumericLessThanEquals"
	OpeNumericGreaterThan        Operator = "NumericGreaterThan"
	OpeNumericGreaterThanEquals  Operator = "NumericGreaterThanEquals"
	OpeDateEquals                Operator = "DateEquals"
	OpeDateNotEquals             Operator = "DateNotEquals"
	OpeDateLessThan              Operator = "DateLessThan"
	OpeDateGreaterThan           Operator = "DateGreaterThan"
	OpeDateGreaterThanEquals     Operator = "DateGreaterThanEquals"
	OpeBool                      Operator = "Bool"
	OpeIpAddress                 Operator = "IpAddress"
	OpeNotIpAddress              Operator = "NotIpAddress"
)

type Key string

const (
	KeySourceIp        Key = "acs:SourceIp"
	KeyUserAgent       Key = "acs:UserAgent"
	KeyCurrentTime     Key = "acs:CurrentTime"
	KeySecureTransport Key = "acs:SecureTransport"
	KeyPrefix          Key = "acs:Prefix"
	KeyDelimiter       Key = "acs:Delimiter"
)

func (kc *KeyCondition) Put(key Key, value string) *KeyCondition {
	if kl, ok := (*kc)[key]; ok {
		for _, v := range kl {
			if v == value {
				return kc
			}
		}
		(*kc)[key] = append(kl, value)
	} else {
		(*kc)[key] = []string{value}
	}
	return kc
}

type KeyCondition map[Key][]string

func NewKeyCondition() KeyCondition {
	return KeyCondition(make(map[Key][]string, 0))
}

type Condition map[Operator]KeyCondition

func NewCondition() *Condition {
	con := Condition(make(map[Operator]KeyCondition, 0))
	return &con
}

func (con *Condition) String() string {
	data, _ := json.Marshal(con)
	return string(data)
}

func (con *Condition) toMap() *map[Operator]KeyCondition {
	return (*map[Operator]KeyCondition)(con)
}

func (con *Condition) Add(operator Operator, key Key, value string) *Condition {
	mm := *con.toMap()
	if op, ok := mm[operator]; ok {
		//mm[operator] = op.Put(secret, value)
		op.Put(key, value)
	} else {
		kc := NewKeyCondition()
		mm[operator] = *kc.Put(key, value)
	}
	return con
}
