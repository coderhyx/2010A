package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticsearchClient struct {
	Client *elasticsearch.Client
}

var esClient *ElasticsearchClient

func NewElasticsearchClient(dsn string) (*ElasticsearchClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{dsn},
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &ElasticsearchClient{
		Client: esClient,
	}, nil
}

func (es *ElasticsearchClient) CreateIndex(index, indexSettings string) error {
	req := esapi.IndicesCreateRequest{
		Index: index,
		Body:  strings.NewReader(indexSettings),
	}

	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index: %s", res.String())
	}

	fmt.Println("Index created:", res.String())
	return nil
}

func (es *ElasticsearchClient) CreateDocument(index, id, body string) error {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.Client)
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

func (es *ElasticsearchClient) GetDocument(index, id string) (string, error) {
	req := esapi.GetRequest{
		Index:      index,
		DocumentID: id,
	}

	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.IsError() {
		return "", fmt.Errorf("error getting document: %s", res.String())
	}

	return res.String(), nil
}

func (es *ElasticsearchClient) UpdateDocument(index, id, body string) error {
	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.Client)
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

func (es *ElasticsearchClient) DeleteDocument(index, id string) error {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.Client)
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

func (es *ElasticsearchClient) SearchDocuments(index, query string) ([]byte, error) {
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  strings.NewReader(query),
	}
	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf("error searching documents: %s", res.String())
	}
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return jsonResult, nil
}

func (es *ElasticsearchClient) CountDocuments(index, query string) (int, error) {
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  strings.NewReader(query),
	}
	// 执行请求
	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		log.Fatalf("Error executing the request: %s", err)
	}
	defer res.Body.Close()

	// 检查请求是否成功
	if res.IsError() {
		log.Fatalf("Error response: %s", res.String())
	}

	// 解析响应的 JSON 数据
	var countResult map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&countResult); err != nil {
		log.Fatalf("Error parsing the response: %s", err)
	}

	// 从解析后的数据中获取统计结果
	count := int(countResult["count"].(float64))
	fmt.Printf("Count: %d\n", count)
	return count, nil
}
