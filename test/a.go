package main

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

func NewElasticsearchClient() (*ElasticsearchClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://127.0.0.1:9200"}, // 替换为您的 Elasticsearch 地址
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ElasticsearchClient{
		Client: esClient,
	}, nil
}

func createIndex(index, indexSettings string) error {
	req := esapi.IndicesCreateRequest{
		Index: index,
		Body:  strings.NewReader(indexSettings),
	}

	res, err := req.Do(context.Background(), esClient.Client)
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

func createDocument(index, id, body string) error {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), esClient.Client)
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

func getDocument(index, id string) (string, error) {
	req := esapi.GetRequest{
		Index:      index,
		DocumentID: id,
	}

	res, err := req.Do(context.Background(), esClient.Client)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.IsError() {
		return "", fmt.Errorf("error getting document: %s", res.String())
	}

	return res.String(), nil
}

func updateDocument(index, id, body string) error {
	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), esClient.Client)
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

func deleteDocument(index, id string) error {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), esClient.Client)
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

func searchDocuments(index, query string) ([]byte, error) {
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(context.Background(), esClient.Client)
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

func main() {
	var err error
	esClient, err = NewElasticsearchClient()
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}
	// 创建索引
	//indexSettings := `{
	//	"settings": {
	//		"index": {
	//			"number_of_shards": 1,
	//			"number_of_replicas": 0
	//		}
	//	}
	//}`
	//err = createIndex("my_index_2", indexSettings)
	//if err != nil {
	//	log.Fatalf("Error creating index: %s", err)
	//}

	// 创建文档
	//err = createDocument("my_index_2", "1", `{"title":"苹果","content":"This is a sample document"}`)
	//if err != nil {
	//	log.Fatalf("Error creating document: %s", err)
	//}
	//err = createDocument("my_index_2", "2", `{"title":"香蕉","content":"This is a sample document"}`)
	//if err != nil {
	//	log.Fatalf("Error creating document: %s", err)
	//}
	//err = createDocument("my_index_2", "3", `{"title":"西瓜","content":"This is a sample document"}`)
	//if err != nil {
	//	log.Fatalf("Error creating document: %s", err)
	//}
	//err = createDocument("my_index_2", "", `{"title":"芒果","content":"This is a sample document"}`)
	//if err != nil {
	//	log.Fatalf("Error creating document: %s", err)
	//}

	// 获取文档
	//doc, err := getDocument("my_index_2", "1")
	//if err != nil {
	//	log.Fatalf("Error getting document: %s", err)
	//}
	//fmt.Println("Document retrieved:", doc)

	// 创建文档
	//err = createDocument("my_index_2", "1", `{"title":"abcde","content":"This is a sample document"}`)
	//if err != nil {
	//	log.Fatalf("Error creating document: %s", err)
	//}
	//
	//if err != nil {
	//	log.Fatalf("Error updating document: %s", err)
	//}
	//
	//// 删除文档
	err = deleteDocument("my_index_2", "1")
	if err != nil {
		log.Fatalf("Error deleting document: %s", err)
	}

	// 搜索文档
	//	query := `{
	//    "query": {
	//        "bool": {
	//            "must": [
	//                {
	//                    "match_phrase": {
	//                        "title": {
	//                            "query": "苹果"
	//                        }
	//                    }
	//                }
	//            ]
	//        }
	//    }
	//}`
	//	result, err := searchDocuments("my_index_2", query)
	//	if err != nil {
	//		log.Fatalf("Error searching documents: %s", err)
	//	}
	//	fmt.Println("Search results (JSON):", string(result))

}
