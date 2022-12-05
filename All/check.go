// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {

	// http :=
	res, err := http.Get("https://api.coingecko.com/api/v3/exchange_rates")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(body)

	sb := string(body)
	log.Printf(sb)

	vp := viper.New()

}
