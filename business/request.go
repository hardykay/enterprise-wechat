package business

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func requestGet(url string, argv ...interface{}) ([]byte, error) {
	url = fmt.Sprintf(url, argv...)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
