package tools

import (
	"github.com/chanxuehong/wechat.v2/mch/core"
	"github.com/chanxuehong/wechat.v2/util"
)

// ShortURL 转换短链接.
func ShortURL(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/tools/shorturl", req)
}

type ShortURLRequest struct {
	LongURL  string `xml:"long_url"`  // URL链接
	NonceStr string `xml:"nonce_str"` // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"` // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
}

type ShortURLResponse struct {
	ShortURL string `xml:"short_url"` // 转换后的URL
}

// ShortURL2 转换短链接.
func ShortURL2(clt *core.Client, req *ShortURLRequest) (resp *ShortURLResponse, err error) {
	m1 := make(map[string]string, 8)
	m1["appid"] = clt.AppId()
	m1["mch_id"] = clt.MchId()
	m1["long_url"] = req.LongURL
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = util.NonceStr()
	}
	if req.SignType != "" {
		m1["sign_type"] = req.SignType
	}

	m2, err := ShortURL(clt, m1)
	if err != nil {
		return nil, err
	}

	resp = &ShortURLResponse{
		ShortURL: m2["short_url"],
	}
	return resp, nil
}
