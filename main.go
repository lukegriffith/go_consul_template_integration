package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type config struct {
	Servers []string `json:"memcache_servers"`
}



func reloadConfig(c *config) {

	for {
		content, err := ioutil.ReadFile("/app/test_config")

		if err != nil {
			log.Panic("unable to find config file")
		}

		var data config

		if err := json.Unmarshal(content, &data); err != nil {
			log.Panic(err)
		}

		*c = data

		time.Sleep(3 * time.Second)
	}
}

func main() {

	c := config{}

	fmt.Println("started.")

	go reloadConfig(&c)

	fmt.Println("loaded")

	for {
		time.Sleep(5 * time.Second)

		fmt.Println(c)
	}

}
