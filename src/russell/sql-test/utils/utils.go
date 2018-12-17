package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const SUCCESSCODE int = 10000

const DBERRORCODE int = 99

func HttpPost(urlStr string, params url.Values) ([]byte) {
	resp, err := http.Post(urlStr,
		"application/x-www-form-urlencoded",
		strings.NewReader(params.Encode()))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return body
}
