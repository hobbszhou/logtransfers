package kafka

import (
	"encoding/json"
	"fmt"
	"mylogtransfer/es"

	"github.com/Shopify/sarama"
)

func Init(addr []string, topic string) (err error) {
	fmt.Println("1111111111--", addr, "  2222-", topic)
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println("分区列表:", partitionList)

	for partition := range partitionList {

		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			panic(err)
		}

		fmt.Println("start to consume......")
		go func(sarama.PartitionConsumer) {
			fmt.Println("in sranmam ......")
			for msg := range pc.Messages() {

				fmt.Println("22222--msg.Topic=", msg.Topic, "  msg.Value=", msg.Value)
				var m1 map[string]interface{}
				err = json.Unmarshal(msg.Value, &m1)

				if err != nil {
					fmt.Println("aaaaaaaaaaaaaaaaaaaerror\n")
					continue
				}
				fmt.Println("map====", m1)
				es.PutLogData(m1)
			}
		}(pc)
	}

	return

}
