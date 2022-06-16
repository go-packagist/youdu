package youdu

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type Response struct {
	restyResponse *resty.Response
	encryptor     *encryptor

	decryptResult *DecryptResult
}

func NewResponse(restyResponse *resty.Response) *Response {
	return &Response{
		restyResponse: restyResponse,
	}
}

func (r *Response) Body() []byte {
	return r.restyResponse.Body()
}

func (r *Response) String() string {
	return r.restyResponse.String()
}

func (r *Response) Json() (map[string]interface{}, error) {
	var v map[string]interface{}
	if err := json.Unmarshal(r.restyResponse.Body(), &v); err != nil {
		return nil, err
	}

	return v, nil
}

// func (r *Response) Decrypt() (*Response, error) {
// 	v, err := r.Json()
// 	if err != nil {
// 		return r, err
// 	}
//
// 	r.decryptResult, err = r.encryptor.Decrypt(v["encrypt"].(string))
// 	if err != nil {
// 		return r, err
// 	}
//
// 	return r, nil
// }

// func (r *Response) DecryptResult() *DecryptResult {
// 	return r.decryptResult
// }

func (r *Response) StatusCode() int {
	return r.restyResponse.StatusCode()
}

func (r *Response) Header() map[string][]string {
	return r.restyResponse.Header()
}

func (r *Response) IsSuccess() bool {
	return r.StatusCode() == 200
}

// func (r *Response) Cookies() []*resty.Cookie {
// 	return r.restyResponse.Cookies()
// }
//
// func (r *Response) ContentLength() int64 {
// 	return r.restyResponse.ContentLength()
// }
//
// func (r *Response) ContentType() string {
// 	return r.restyResponse.ContentType()
// }
