package main

import (
	"log"

	"gitgub.com/nenych/es_random_indexes/conf"
	"gitgub.com/nenych/es_random_indexes/es"
)

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	var c *conf.Config
	cfg, err := c.Parse("./config.yml")
	check(err)

	esClient, err := es.Connect(cfg.Domain.Host, cfg.Domain.Port)
	check(err)
	log.Println(esClient.Info())

	log.Println(es.CreateIndex(esClient))
}
