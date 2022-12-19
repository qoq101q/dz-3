package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
)
func main() {
	httpClient := &http.Client{}

	n := Note{
		name: "1",
		surname: "2",
		text: "3",
	}
	js, err := json.Marshal(&n)
	if err != nil {
		log.Println(errors.Wrap(err, 'json.Marshal(&n)'))
		os.Exit(1)
	}

	bb := bytes.Buffer{}
	bb.Write(js)

	req, err := http.NewRequest("POST", url "http://127.0.0.1:4000", &bb)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(errors.Wrap(err, "httpClient.Do(req)"))
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Print("Status 0")
}