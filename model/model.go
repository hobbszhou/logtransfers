package model

type Config struct {
	KafkaConf `ini:"kafka"`
	ESConf    `ini:"es"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}
type ESConf struct {
	Address string `ini:"address"`
	Index   string `ini:"index"`
	MaxSize int    `ini:"max_chan_size"`
	GoNum   int    `ini:"goroutine_num"`
}
