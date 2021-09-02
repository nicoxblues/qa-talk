package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main(){
	reader := strings.NewReader(`{"body":123}`)
	//START OMIT
	valueToSend := "6"
	//END OMIT
	request, _ := http.NewRequest("GET", "https://api.isevenapi.xyz/api/iseven/" +valueToSend+"/", reader)

	// TODO: check err
	client := &http.Client{}
	resp, _ := client.Do(request)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

} // HL12
