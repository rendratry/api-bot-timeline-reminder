package helper

import (
	"api-bot-timeline-reminder/model/web"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func PublishRabbitmqDelay(structPublish web.PublishDelayRequest) time.Time {
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
	scheduledTime := time.Date(structPublish.Time.Year, time.Month(structPublish.Time.Month), structPublish.Time.Day, structPublish.Time.Hour, structPublish.Time.Minute, 0, 0, time.Local)
	// Menghitung delay time
	delay := scheduledTime.Sub(now)

	// Mengirim pesan pada jadwal yang ditentukan
	marshal, err := json.Marshal(structPublish)
	if err != nil {
		return time.Time{}
	}

	go func() {
		time.Sleep(delay)
		msg := amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshal,
			Headers:     amqp.Table{"x-delay": int(delay / time.Millisecond)},
			Timestamp:   time.Now(),
		}
		err := ch.Publish("delayed_exchange", "", false, false, msg)
		if err != nil {
			log.Printf("Error publishing message: %s", err)
		}
	}()
	log.Printf("Scheduled message to be sent at %s", scheduledTime)
	select {}
	return scheduledTime
}

func PublishRabbitMQSendEmail(message web.SendEmailRequest) {
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

	q, err := ch.QueueDeclare(
		"SendEmail",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("gagal declare", err)
	}

	type Person struct {
		name  string
		email string
		nohp  string
	}

	DataJson, _ := json.Marshal(message)

	if err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        DataJson,
		},
	); err != nil {
		fmt.Println("error publishing", err)
	}
}
