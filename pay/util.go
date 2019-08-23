package pay

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
	"sandpay/pay/reponse"
	"strconv"
	"strings"
	"time"
)

/**
  生成随机字符串
*/
func randomString(lens int) string {
	now := time.Now()
	return MakeMd5(strconv.FormatInt(now.UnixNano(), 10))[:lens]
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
	http := TimeoutClient()
	resp, err := http.Post(requrl, "application/x-www-form-urlencoded", strings.NewReader(HttpBuildQuery(request)))

	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		return response, fmt.Errorf("http request response StatusCode:%v", resp.StatusCode)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
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
			vals.Set(key, val)
		}
	}
	result, err := PublicSha1Verify(vals)
	if err != nil {
		return response, err
	}
	mapInfo := result.(map[string]string)
	for key, value := range mapInfo {
		//log.Println("mapinfo result ", key, value)
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

func GenerateSignString(body interface{}, header interface{}) (str string) {
	signData := make(map[string]interface{}, 2)
	signData["body"] = body
	signData["head"] = header
	signDataJson, err := json.Marshal(signData)
	if err != nil {
		return
	}
	return string(signDataJson[:])
}

func GeneratePostData(signDataJsonString string, sign string) map[string]string {
	postData := make(map[string]string, 4)
	postData["chart"] = `utf-8`
	postData["signType"] = `01`
	postData["data"] = signDataJsonString
	postData["sign"] = sign
	return postData
}
