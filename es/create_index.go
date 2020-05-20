package es

import (
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

// CreateIndex function for the creating indexes
func CreateIndex(es *elasticsearch.Client) (string, error) {
	req := esapi.IndexRequest{
		Index:      "test_go",                                // Index name
		Body:       strings.NewReader(`{"title" : "Test4"}`), // Document body
		DocumentID: "4",                                      // Document ID
		Refresh:    "true",                                   // Refresh
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return res.String(), nil
}
