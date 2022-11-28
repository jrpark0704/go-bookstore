package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// byte slice나 문자열을 자료구조로 변환
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
