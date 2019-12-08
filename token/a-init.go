package token

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"wx-server/rlog"
	"wx-server/util"
)

var client http.Client

const WxApiServer = "https://api.weixin.qq.com"
const WxCgiApi = WxApiServer + "/cgi-bin"
const WxSnsApi = WxApiServer + "/sns"

func init() {
	tr := &http.Transport{
		//Proxy:           func(r *http.Request) (*url.URL, error) { return url.Parse("http://127.0.0.1:8888") },
		TLSClientConfig: &tls.Config{
			VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
				return nil
			},
			InsecureSkipVerify: true,
		},
	}
	jar, err := util.NewCookieJar(nil)
	rlog.CheckShowError(err)
	client = http.Client{
		Transport: tr,
		Jar:       jar,
	}
}
