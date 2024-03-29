package main

import (
	"flag"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"github.com/mft-labs/runb2bqueries/utils"
)

var (
	config *ini.File

)
func main() {
	var conf string
	flag.StringVar(&conf,"conf","b2bqueries.conf","Config file")
	flag.Parse()
	log.Printf("Running runb2bqueries\n")
	util := &utils.Util{}
	err := util.LoadConfig(conf)
	if err!=nil {
		log.Printf("Failed to load config:%v",err)
		os.Exit(1)
	}
	proc := &Process{Util:util}
	proc.RunQueries()
}
