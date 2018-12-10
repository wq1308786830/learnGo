package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const SUCCESSCODE int = 10000

const DBERRORCODE int = 99

func HttpPost(url string, params []byte) ([]byte, error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		bytes.NewBuffer(params))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return body, err
}
