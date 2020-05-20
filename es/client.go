package es

import (
	"fmt"

	"github.com/elastic/go-elasticsearch"
)

func esUrl(host, port string) string {
	return fmt.Sprintf("http://%s:%s", host, port)
}

// Connect - creates the file structure for the connection
func Connect(host, port string) (*elasticsearch.Client, error) {
	esCfg := elasticsearch.Config{
		Addresses: []string{
			esUrl(host, port),
		},
	}

	es, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		return nil, err
	}

	return es, nil
}
