package requester

import (
	"log"
	"testing"
)

func TestParseCurlFile(t *testing.T) {
	var path = "../examples/test.chrome.curl.txt"
	curl, _ := ParseCurlFile(path)
	log.Print(curl.Data)
	url := curl.GetURL()
	log.Println(url)
}
