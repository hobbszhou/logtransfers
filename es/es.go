package es

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type ESClient struct {
	client      *elastic.Client
	index       string
	logDataChan chan interface{}
}

var (
	esClient *ESClient
)

func Init(addr, index string, goroutineNum, maxSize int) (err error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + addr))
	if err != nil {

		panic(err)
	}
	fmt.Println("11111-", client)
	esClient = &ESClient{
		client:      client,
		index:       index,
		logDataChan: make(chan interface{}, maxSize),
	}

	fmt.Println("connect to es success")

	for i := 0; i < goroutineNum; i++ {
		go sendToES()
	}

	return
}
func PutLogData(msg interface{}) {
	esClient.logDataChan <- msg
}
func sendToES() {
	for m1 := range esClient.logDataChan {

		fmt.Println("sendToES2222----", m1)

		esClient.client.Index().Index(esClient.index).BodyJson(m1).Do(context.Background())
		_, err := esClient.client.Index().Index(esClient.index).BodyJson(m1).Do(context.Background())
		if err != nil {

			panic(err)
		}

	}
}
