package helper

import (
	"api-bot-timeline-reminder/model/domain"
	"errors"
	"net/http"
	"net/url"
)

func GetReceiver(id string) (domain.NotificationReceiver, error) {
	mahasiswa, err := GetReceiverMahasiswa(id)
	if err != nil {
		staff, err := GetReceiverStaff(id)
		if err != nil {
			return staff, errors.New("receiver not found")
		}
		return staff, nil
	}
	return mahasiswa, nil
}

func SendMessageWhatsapp(receiver string, message string) {

}

func SendMessageTelegram(receiver string, message string) {
	_, err := http.Get("https://api.telegram.org/bot5836520934:AAGQ_iQvY7Hbm5goVRPbwH-k57p25dA-Gns/SendMessage?chat_id=" + receiver + "&text=" + url.QueryEscape(string(message)))
	if err != nil {
		panic(err)
	}
}
