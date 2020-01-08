package oss

import (
	"encoding/base64"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/signers"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"time"
	"wx-server/config"
	"wx-server/rlog"
)

type AuthType string

const (
	Temp = "temp"
)

var regionId = config.Cfg().Oss.RegionId
var accessKeyId = config.Cfg().Oss.AccessKeyId
var accessKeySecret = config.Cfg().Oss.AccessKeySecret
var roleArn = config.Cfg().Oss.RoleArn
var Domain = config.Cfg().Oss.Domain
var bucket = config.Cfg().Oss.Bucket

//var policy = `{"Version":"1","Statement":[{"Action":["oss:PutObject","oss:put"],"Resource":["acs:oss:*:*:*"],"Effect":"Allow"}]}`

func GetOssTempToken() (token StsTokenResult, filePath string) {
	token, filePath = CreateOssOneAuth(Temp, fmt.Sprintf("%v%v-%v.png", time.Now().Unix(), time.Now().Nanosecond(), utils.RandStringBytes(16)))
	return
}

type WxToken struct {
	OssAccessKeyId string `json:"accessKeyId"`
	Key            string `json:"key"`
	Name           string `json:"name"`
	Policy         string `json:"policy"`
	Signature      string `json:"signature"`
}

func CreateWxOssSign() (token WxToken) {
	sourceName := fmt.Sprintf("%v%v-%v.png", time.Now().Unix(), time.Now().Nanosecond(), utils.RandStringBytes(16))
	//s := NewStatement(Allow).
	//	AddResource(Resource(fmt.Sprintf("acs:oss:*:*:%v/%v/%v", bucket, Temp, sourceName))).
	//	AddAction("oss:PutObject")
	//p := NewPolicy().AddStatement(*s)
	//dd, _ := json.Marshal(p)
	policyText := `{
    "expiration": "` + time.Now().Add(time.Hour).Format("2006-01-02T15:04:05.999Z") + `", 
    "conditions": [
    ["content-length-range", 0, 1048576000] 
    ]
}`

	data := base64.StdEncoding.EncodeToString([]byte(policyText))
	//key, _ := base64.StdEncoding.DecodeString(accessKeySecret)
	//sh := hmac.New(crypto.SHA1.New, key)
	//sh.Write([]byte(data))
	//rr := fmt.Sprintf("%+x", sh.Sum(nil))
	//rlog.Warn(rr)

	token.OssAccessKeyId = config.Cfg().Oss.DirectAccessKeyId
	token.Policy = data
	token.Key = "temp/" + sourceName
	token.Name = sourceName
	token.Signature = signers.ShaHmac1(data, config.Cfg().Oss.DirectAccessKeySecret)
	rlog.WarnF("%+v", token)
	return
}

func CreateOssOneAuth(oType AuthType, sourceName string) (StsTokenResult, string) {
	s := NewStatement(Allow).
		AddResource(Resource(fmt.Sprintf("acs:oss:*:*:%v/%v/%v", bucket, oType, sourceName))).
		AddAction("oss:PutObject")
	p := NewPolicy().AddStatement(*s)
	//con := condition.NewCondition().Add(condition.OpeIpAddress, condition.KeySourceIp, "192.168.*.*")
	//p.Condition = *con
	token := getStsToken(bucket, p.String())
	token.FileName = fmt.Sprintf("%v/%v", oType, sourceName)
	return *token, fmt.Sprintf("%v/%v/%v", Domain, oType, sourceName)
}

func getStsToken(bucketName, policy string) *StsTokenResult {
	client, err := sts.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err == nil {
		req := sts.CreateAssumeRoleRequest()
		req.Policy = policy
		req.RoleArn = roleArn
		req.SetDomain("sts.aliyuncs.com")
		req.Method = "POST"
		req.DurationSeconds = requests.NewInteger(60 * 60)
		req.Scheme = "https"
		req.SetVersion("2015-04-01")
		req.RoleSessionName = fmt.Sprintf("%v", time.Now().Unix())
		resp, err := client.AssumeRole(req)
		if err == nil {
			if resp.IsSuccess() {
				return &StsTokenResult{
					AssumedRoleUser: resp.AssumedRoleUser,
					Credentials:     resp.Credentials,
					Region:          regionId,
					Bucket:          bucketName,
					Domain:          Domain,
				}
			}
		} else {
			rlog.Error(err)
		}
	} else {
		rlog.Error(err)
	}
	return &StsTokenResult{}
}

type StsTokenResult struct {
	AssumedRoleUser sts.AssumedRoleUser
	Credentials     sts.Credentials `json:"credentials"`
	FileName        string          `json:"fileName"`
	Region          string          `json:"region"`
	Bucket          string          `json:"bucket"`
	Domain          string          `json:"domain"`
}
