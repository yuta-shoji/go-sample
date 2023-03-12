package Http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRequest(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	s := string(body[:])

	if err != nil {
		fmt.Println(err)
		return
	}
	c <- s
}
