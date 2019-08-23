package reponse

type header struct {
	RespTime string `json:"respTime"`
	RespMsg  string `json:"respMsg"`
	Version  string `json:"version"`
	RespCode string `json:"respCode"`
}

type Response struct {
	Charset  string `json:"charset"`
	Data     string `json:"data"`
	SignType string `json:"signType"`
	Body     string `json:"body"`
}

func (s *Response) SetBody(body string) *Response {
	s.Body = body
	return s
}
func (s *Response) SetData(data string) *Response {
	s.Data = data
	return s
}
func (s *Response) SetSignType(signType string) *Response {
	s.SignType = signType
	return s
}
func (s *Response) SetCharset(charset string) *Response {
	s.Charset = charset
	return s
}

func (s *Response) SetKeyValue(key string, value string) *Response {
	if key == "charset" {
		s.Charset = value
	} else if key == "signType" {
		s.SignType = value
	} else if key == "data" {
		s.Data = value
	}
	return s
}
