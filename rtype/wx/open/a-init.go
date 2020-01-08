package open

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"wx-server/rlog"
	"wx-server/util"
)

var client http.Client

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

func PostWx(uri string, body []byte, bind interface{}) (data []byte, err error) {
	var req *http.Request
	var resp *http.Response
	req, err = http.NewRequest("POST", uri, bytes.NewReader(body))
	if err == nil {
		resp, err = client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			data, err = ioutil.ReadAll(resp.Body)
			if err == nil && bind != nil {
				err = json.Unmarshal(data, bind)
			}
		}
	}
	return
}
