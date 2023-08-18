package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return nil
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	if e != nil {
		utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
		return ioutil.ReadAll(utf8Reader)
	} else {
		return ioutil.ReadAll(bodyReader)
	}
}
