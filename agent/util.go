package agent

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sandpay/agent/reponse"
	"strconv"
	"strings"
	"time"
)

/**
  生成随机字符串
*/
func randomString(lens int) string {
	now := time.Now()
	return MakeMd5(strconv.FormatInt(now.UnixNano(), 16))[:lens]
}

/**
  字符串md5
*/
func MakeMd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	s := fmt.Sprintf("%x", h.Sum(nil))
	return s
}

/**
  获取当前时间戳
*/
func getNowSec() int64 {
	return time.Now().Unix()
}

/**
  获取当前时间戳
*/
func Str2Sec(layout, str string) int64 {
	tm2, _ := time.ParseInLocation(layout, str, time.Local)
	return tm2.Unix()
}

/**
  获取当前时间
*/
func Sec2Str(layout string, sec int64) string {
	t := time.Unix(sec, 0)
	nt := t.Format(layout)
	return nt
}

// base64 加密
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// base64 解密
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func TimeoutClient() *http.Client {
	connectTimeout := time.Duration(20 * time.Second)
	readWriteTimeout := time.Duration(30 * time.Second)
	return &http.Client{
		Transport: &http.Transport{
			Dial:                TimeoutDialer(connectTimeout, readWriteTimeout),
			MaxIdleConnsPerHost: 200,
			DisableKeepAlives:   true,
		},
	}
}
func TimeoutDialer(cTimeout time.Duration,
	rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

// TODO
// 发送post请求
func PayPost(requrl string, request map[string]string) (response reponse.Response, err error) {
	client := TimeoutClient()
	resp, err := client.Post(requrl, "application/x-www-form-urlencoded;charset=UTF-8", strings.NewReader(HttpBuildQuery(request)))

	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		return response, fmt.Errorf("http request response StatusCode:%v", resp.StatusCode)
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	data, err := ioutil.ReadAll(resp.Body)
	if len(data) == 0 {
		return
	}
	dataString, _ := url.QueryUnescape(string(data[:]))

	if err != nil {
		return response, err
	}

	var fields []string
	fields = strings.Split(dataString, "&")

	vals := url.Values{}
	for _, field := range fields {
		f := strings.SplitN(field, "=", 2)
		if len(f) >= 2 {
			key, val := f[0], f[1]
			// 有特殊数据需要处理 不知道为啥会有两种格式一种是带“” 一种是不带 “ 统一过滤
			vals.Set(strings.TrimLeft(key, `"`), strings.TrimRight(val, `"`))
		}
	}
	result, err := PublicSha1Verify(vals)
	if err != nil {
		return response, err
	}
	mapInfo := result.(map[string]string)
	//log.Println(mapInfo)
	for key, value := range mapInfo {
		response.SetKeyValue(key, value)
	}
	return response, err
}

// urlencode
func HttpBuildQuery(params map[string]string) string {
	qs := url.Values{}
	for k, v := range params {
		qs.Add(k, v)
	}
	return qs.Encode()
}

func GenerateSignString(signData interface{}) (str string) {

	signDataJson, err := json.Marshal(signData)
	if err != nil {
		return
	}
	return string(signDataJson[:])
}

func GeneratePostData(encryptData string, encryptKey string, transCode string, merId string, sign string) map[string]string {
	postData := make(map[string]string, 10)
	postData["transCode"] = transCode
	postData["accessType"] = `0`
	postData["merId"] = merId
	postData["plId"] = ""
	postData["encryptData"] = encryptData
	postData["encryptKey"] = encryptKey
	postData["sign"] = sign
	postData["extend"] = ""
	return postData
}
