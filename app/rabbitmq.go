package app

import (
	"github.com/streadway/amqp"
	"log"
)

func RabbitMqConn() *amqp.Connection {
	conn, err := amqp.Dial("amqp://admin:Adminmyfin123@api.myfin.id:5672/")
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ: ", err)
	}
	defer conn.Close()
	
	return conn
}
