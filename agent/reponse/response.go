package reponse

type Response struct {
	AccessType string `json:"charset"`
	Data       string `json:"data"`
	TransCode  string `json:"transCode"`
	MerId      string `json:"merId"`
}

func (s *Response) SetData(data string) *Response {
	s.Data = data
	return s
}

func (s *Response) SetKeyValue(key string, value string) *Response {
	if key == "accessType" {
		s.AccessType = value
	} else if key == "transCode" {
		s.TransCode = value
	} else if key == "data" {
		s.Data = value
	} else if key == "merId" {
		s.MerId = value
	}
	return s
}
