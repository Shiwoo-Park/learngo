package modules

import (
	"log"
	"net/http"
	"strings"
)

func HandlelError(err error) {
	if err != nil {
		log.Fatalln("Error occured: ", err)
	}
}
func HandlelErrorForResponse(res *http.Response, err error) {
	HandlelError(err)
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
}
func CleanString(s string) string {
	txtArray := strings.Fields(strings.TrimSpace(s))
	return strings.Join(txtArray, " ")
}
