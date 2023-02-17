package main

import (
	"api-bot-timeline-reminder/helper"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestEpoch(t *testing.T) {
	waktu := time.Now().UnixNano() / 1000000
	fmt.Println(waktu)
	fmt.Println(reflect.TypeOf(waktu))
}

func TestConvertEpoch(t *testing.T) {
	time := time.Unix(-2211750432000, 0)
	fmt.Println(time)
}

func TestCreateUUID(t *testing.T) {
	uuidstr := uuid.NewV4()
	fmt.Println(uuidstr)
}

func TestEkstrakString(t *testing.T) {
	s := "berikut ini adalah string nya:1839-94893.93d0cff1-4768-49a2-ad05-679168d218a9"
	parts := strings.Split(s, ":")
	if len(parts) > 1 {
		result := parts[1]
		fmt.Println(result)
		parts := strings.Split(result, ".")
		if len(parts) > 1 {
			result := parts[0]
			result1 := parts[1]
			fmt.Println(result)
			fmt.Println(result1)
		} else {
			fmt.Println("Tidak ditemukan karakter '.' dalam string")
		}
	} else {
		fmt.Println("Tidak ditemukan karakter ':' dalam string")
	}
}

//func TestCron(t *testing.T) {
//	epochMillis := int64(1674482340000)
//
//	// Convert milliseconds to time
//	scheduledTime := time.Unix(0, epochMillis*int64(time.Millisecond))
//	fmt.Println(scheduledTime)
//	fmt.Println(scheduledTime.Format("15:04"))
//
//	c := cron.New()
//	c.AddFunc("0 0 "+scheduledTime.Format("15:04")+" 23 1 * 2023", func() {
//		_, err := http.Get("https://api.telegram.org/bot5836520934:AAGQ_iQvY7Hbm5goVRPbwH-k57p25dA-Gns/SendMessage?chat_id=1359729975&text=Halorendrainitescron")
//		if err != nil {
//			panic(err)
//		}
//	})
//	c.Start()
//	select {}
//}

func TestDataType(t *testing.T) {
	OtpTime := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println("tipe data :", reflect.TypeOf(OtpTime))
}

func TestParsing(t *testing.T) {
	jam := "7.45"
	parts := strings.Split(jam, ".")
	hour := parts[0]
	minute := parts[1]

	fmt.Println("hour:", hour)
	fmt.Println("minute:", minute)
}

func TestJwtCreate(t *testing.T) {
	claims := jwt.MapClaims{
		"iss": "Rendra",
		"sub": "rendratrykusuma@gmail.com",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 30).Unix(),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the JWT token with a secret key
	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(signedToken)
}

func TestJwtValidate(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY3NjE4OTIsImlhdCI6MTY3NjQ3Mzg5MiwiaXNzIjoiYm90LnRpbWVsaW5lLnJlbWluZGVyIiwic3ViIjoiYTNlMjI0ZjEtMTI5My00ODcyLTg1N2ItYmNmM2QyZGMzNzMwIn0.x1gphh_lz4SFUAbOKxW7cFLah6eboFbrsF59HfooKI8"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("admin-bot-timeline"), nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(reflect.TypeOf(claims))
		fmt.Println("Token is valid. Claims:")
		fmt.Println("Issuer:", claims["iss"])
		fmt.Println("Subject:", claims["sub"])
		fmt.Println("Iat:", claims["iat"])
		fmt.Println("Expiration:", claims["exp"])
	} else {
		fmt.Println("Token is invalid")
	}
}

func TestGetUUID(t *testing.T) {
	tx := helper.GetConnection2()
	ctx := context.Background()
	script := "select uuid from user where email = ?"
	rows, err := tx.QueryContext(ctx, script, "rendratrikusuma@student.uns.ac.id")
	helper.PanicIfError(err)
	defer rows.Close()

	var uuidstr string
	if rows.Next() {
		err := rows.Scan(&uuidstr)
		helper.PanicIfError(err)
		fmt.Println(uuidstr)
	} else {
		fmt.Println("failed")
	}
}

func TestValidateJWT(t *testing.T) {
	helper.ValidateJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU5MzA5MDgsImlhdCI6MTY3NTkzMDg3OCwiaXNzIjoiUmVuZHJhIiwic3ViIjoiUmVuZHJhIn0.pmyVijF9GrEHIPZk6__MeCQMTFTwJ81DOi1dNIUyCvw", "secret")
}

func TestUploadPdf(t *testing.T) {

}
