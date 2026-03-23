package kafkamanager

import (
	// "fmt"
	// "net"
	// "strconv"
	//
	"github.com/segmentio/kafka-go"
)

type KafkaConf struct {
	Host string
	Port int
}

func (kc *KafkaConf) ConnectKafka() *kafka.Conn {
	/* 	conn, err := kafka.Dial("tcp", fmt.Sprintf("%s:%s", kc.Host, kc.Port)) */
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer conn.Close()
	//
	// controller, err := conn.Controller()
	// var controllerConn *kafka.Conn
	// controllerConn, err = kafka.Dial(
	// 	"tcp",
	// 	net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)),
	// )
	// fmt.Println(controllerConn)
	// return controllerConn
	return nil
}
