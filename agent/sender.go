package helperAgent

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// SendToFalconAgent POST metrics to falcon agent PUSH API.
func SendToFalconAgent(urlFalcon, dataPost string) (respCode int, respBody string) {
	req, err := http.NewRequest("POST", urlFalcon, strings.NewReader(dataPost))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response

	connectTimeout := time.Second * time.Duration(2)
	clinet := http.Client{
		Timeout: connectTimeout,
	}
	resp, err = clinet.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	respCode = resp.StatusCode

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	respBody = string(b)
	return
}
