package main

import (
	"encoding/json"
	"os"

	_ "github.com/signavio/signa/ext/gmr"
	_ "github.com/signavio/signa/ext/kubernetes/deployment"
	_ "github.com/signavio/signa/ext/kubernetes/get"
	"github.com/signavio/signa/pkg/slack"
)

func loadConfig(file string) map[string]interface{} {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	var c map[string]interface{}
	d := json.NewDecoder(f)
	d.Decode(&c)

	return c
}

func main() {
	c := loadConfig("/etc/signa.conf")
	slack.Run(c["slack-token"].(string))
}