package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type config struct {
	Name string `json:"name"`
}

func reloadConfig(c *config) {

	for {

		content, err := ioutil.ReadFile("test_config")

		if err != nil {
			log.Panic("unable to find template file")
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

	go reloadConfig(&c)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		fmt.Print(c)
	}

}
