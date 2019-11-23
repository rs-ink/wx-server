package oss

import (
	"log"
	"testing"
	"wx-server/rtype/oss/condition"
)

//var policy = `{"Version":"1","Statement":[{"Action":["oss:PutObject","oss:put"],"Resource":["acs:oss:*:*:*"],"Effect":"Allow"}]}`

func TestNewPolicy(t *testing.T) {
	p := NewPolicy()

	con := condition.NewCondition()
	con.Add(condition.OpeNumericLessThanEquals, condition.KeySecureTransport, "https")

	statement := NewStatement(Allow).AddAction(Action("oss:PutObject")).AddResource(Resource("acs:oss:*:*:duduimage/adsfasdf"))
	p.AddStatement(*statement)
	p.Condition = *con

	log.Println(p.String())
}
