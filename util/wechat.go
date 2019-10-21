package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	appId     = "wxc756bd2ddf6fb5f5"
	appSecret = "991baca4143e346a50876cc8a16b148e"
	grantType = "authorization_code"
)

type WechatKey struct {
	Session_key string
	Openid      string
}

//
type WechatUserInfo struct {
	Openid    string
	NickName  string
	Gender    int8
	Language  string
	City      string
	Province  string
	Country   string
	AvatarUrl string
	WechatWaterMark
}

type WechatWaterMark struct {
	TimeStamp time.Time
	AppId     string
}

// 发送 http get 请求 返回 session_key 和 openid
func Code2Session(code string) []byte {
	url := "https://api.weixin.qq.com/sns/jscode2session?" +
		"appid=" + appId + "&secret=" + appSecret + "&js_code=" + code + "&grant_type=" + grantType
	resp, err := http.Get(url)
	if err != nil {
		log.Panicln("http get 请求失败", err)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取内容失败", err)
	}
	return res
}

// 解密微信用户数据
func DecryptionWechat(sessionKey, encryptedData, iv string) ([]byte, error) {
	if len(sessionKey) != 24 {
		return nil, errors.New("微信 session_key 的长度不为 24")
	}
	if len(iv) != 24 {
		return nil, errors.New("微信 iv 的长度不为 24")
	}

	// 解码
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, errors.New("base64解码 微信 session_key 失败")
	}
	aesIv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, errors.New("base64解码 微信 iv 失败")
	}
	aesCrypt, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, errors.New("base64解码 微信 encryptedData 失败")
	}

	// 解密
	cb, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, errors.New("微信解密失败")
	}
	bm := cipher.NewCBCDecrypter(cb, aesIv)
	dst := make([]byte, len(aesCrypt))
	bm.CryptBlocks(dst, aesCrypt)

	// 去掉末尾的占位符
	for k, v := range dst {
		if v == '\x02' {
			dst[k] = ' '
		}
	}

	return dst, nil
}
