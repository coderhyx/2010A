package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticsearchClient struct {
	Client *elasticsearch.Client
}

func NewElasticsearchClient() (*ElasticsearchClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ElasticsearchClient{
		Client: esClient,
	}, nil
}

func (ec *ElasticsearchClient) CreateDocument(index, id string, body string) error {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(body),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), ec.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating document: %s", res.String())
	}

	fmt.Println("Document created:", res.String())
	return nil
}

func (ec *ElasticsearchClient) GetDocument(index, id string) (string, error) {
	req := esapi.GetRequest{
		Index:      index,
		DocumentID: id,
	}

	res, err := req.Do(context.Background(), ec.Client)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.IsError() {
		return "", fmt.Errorf("error getting document: %s", res.String())
	}

	return res.String(), nil
}

func (ec *ElasticsearchClient) UpdateDocument(index, id, body string) error {
	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), ec.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error updating document: %s", res.String())
	}

	fmt.Println("Document updated:", res.String())
	return nil
}

func (ec *ElasticsearchClient) DeleteDocument(index, id string) error {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), ec.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error deleting document: %s", res.String())
	}

	fmt.Println("Document deleted:", res.String())
	return nil
}
