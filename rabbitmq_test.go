package main

import (
	"github.com/streadway/amqp"
	"log"
	"testing"
	"time"
)

func TestPublish1(t *testing.T) {
	//rabbitMq := app.RabbitMqConn()
	conn, err := amqp.Dial("amqp://amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ: ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error opening channel: ", err)
	}
	defer ch.Close()
	_, err = ch.QueueDeclare("schedule_email_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Error declaring queue: ", err)
	}

	//Mendapatkan waktu saat ini
	now := time.Now()
	// Menentukan jadwal untuk mengirim pesan
	scheduledTime := time.Date(2023, 2, 17, 24, 5, 0, 0, time.Local)
	//Menghitung delay time
	delay := scheduledTime.Sub(now)

	// Mengirim pesan pada jadwal yang ditentukan
	go func() {
		time.Sleep(delay)
		msg := amqp.Publishing{
			Body: []byte("Hello, this is scheduled message"),
		}
		err := ch.Publish("", "schedule_email_queue", false, false, msg)
		if err != nil {
			log.Printf("Error publishing message: %s", err)
		}
	}()
	log.Printf("Scheduled message to be sent at %s", scheduledTime)

	select {}
}

func TestPublish2(t *testing.T) {
	//rabbitMq := app.RabbitMqConn()
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ: ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error opening channel: ", err)
	}
	defer ch.Close()
	_, err = ch.QueueDeclare("schedule_email_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Error declaring queue: ", err)
	}

	//Mendapatkan waktu saat ini
	now := time.Now()
	// Menentukan jadwal untuk mengirim pesan
	scheduledTime := time.Date(2023, 2, 17, 26, 5, 0, 0, time.Local)
	//Menghitung delay time
	delay := scheduledTime.Sub(now)

	// Mengirim pesan pada jadwal yang ditentukan
	go func() {
		time.Sleep(delay)
		msg := amqp.Publishing{
			Body: []byte("Hello, this is scheduled message"),
		}
		err := ch.Publish("", "schedule_email_queue", false, false, msg)
		if err != nil {
			log.Printf("Error publishing message: %s", err)
		}
	}()
	log.Printf("Scheduled message to be sent at %s", scheduledTime)

	select {}
}

func TestPublishDelay(t *testing.T) {
	// Connect ke RabbitMQ
	conn, err := amqp.Dial("amqp://admin:Adminmyfin123@api.myfin.id:5672/")
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ: ", err)
	}
	defer conn.Close()

	// Buka channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error opening channel: ", err)
	}
	defer ch.Close()

	// Membuat exchange
	err = ch.ExchangeDeclare("delayed_exchange", "x-delayed-message", true, false, false, false, amqp.Table{
		"x-delayed-type": "direct",
	})
	if err != nil {
		log.Fatal("Error declaring exchange: ", err)
	}

	// Membuat antrian dengan binding ke exchange
	q, err := ch.QueueDeclare("delayed_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Error declaring queue: ", err)
	}
	err = ch.QueueBind(q.Name, "", "delayed_exchange", false, nil)
	if err != nil {
		log.Fatal("Error binding queue to exchange: ", err)
	}

	// Mendapatkan waktu saat ini
	now := time.Now()
	// Menentukan jadwal untuk mengirim pesan
	scheduledTime := time.Date(2023, 2, 28, 19, 14, 0, 0, time.Local)
	// Menghitung delay time
	delay := scheduledTime.Sub(now)

	// Mengirim pesan pada jadwal yang ditentukan
	go func() {
		time.Sleep(delay)
		msg := amqp.Publishing{
			Body:      []byte("Hello, this is scheduled message with x-delay"),
			Headers:   amqp.Table{"x-delay": int(delay / time.Millisecond)},
			Timestamp: time.Now(),
		}
		err := ch.Publish("delayed_exchange", "", false, false, msg)
		if err != nil {
			log.Printf("Error publishing message: %s", err)
		}
	}()
	log.Printf("Scheduled message to be sent at %s", scheduledTime)

	select {}
}
