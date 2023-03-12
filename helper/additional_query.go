package helper

import (
	"api-bot-timeline-reminder/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

func GetConnection2() *sql.DB {
	db, err := sql.Open("mysql", "myfin:Admin@myfin123@tcp(103.189.234.90:3306)/d3ti_psdku")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

type Presensi struct {
	JamMasuk  string
	JamPulang string
}

func GetOffice() Presensi {
	tx := GetConnection2()
	ctx := context.Background()
	script := "select jam_masuk, jam_pulang from office where id = ?"
	rows, err := tx.QueryContext(ctx, script, "1")
	PanicIfError(err)
	defer rows.Close()

	office := Presensi{}

	if rows.Next() {
		rows.Scan(&office.JamMasuk, &office.JamPulang)
		return office
	}
	return office
}

func GetIssuer(Iss interface{}) bool {
	tx := GetConnection2()
	ctx := context.Background()
	script := "select id from admin where id = ?"
	rows, err := tx.QueryContext(ctx, script, Iss)
	PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}
}

func GetReceiverMahasiswa(id string) (domain.NotificationReceiver, error) {
	tx := GetConnection2()
	ctx := context.Background()
	script := "select bot_user_mahasiswa.no_wa, id_telegram from bot_user_mahasiswa join mahasiswa on bot_user_mahasiswa.id_mahasiswa = mahasiswa.id_mhs where bot_user_mahasiswa.id_mahasiswa = ?"
	stmt, err := tx.PrepareContext(ctx, script)
	PanicIfError(err)
	defer stmt.Close()

	receiver := domain.NotificationReceiver{}

	err = stmt.QueryRow(1).Scan(&receiver.Whatsapp, &receiver.Telegram, &receiver.Email)
	if err != nil {
		return receiver, err
	}

	return receiver, nil
}

func GetReceiverStaff(id string) (domain.NotificationReceiver, error) {
	tx := GetConnection2()
	ctx := context.Background()
	script := "select bot_user_staff.no_wa, id_telegram from bot_user_staff join staf on bot_user_staff.id_mahasiswa = staf.id_staf where bot_user_staff.id_staff = ?"
	stmt, err := tx.PrepareContext(ctx, script)
	PanicIfError(err)
	defer stmt.Close()

	receiver := domain.NotificationReceiver{}

	err = stmt.QueryRow(1).Scan(&receiver.Whatsapp, &receiver.Telegram, &receiver.Email)
	if err != nil {
		return receiver, err
	}

	return receiver, nil
}

func UpdateLastAccessLoginAdmin(time string) {
	tx := GetConnection2()
	ctx := context.Background()
	script := "update admin set last_time_access = ?"
	_, err := tx.ExecContext(ctx, script, time)
	PanicIfError(err)
}

func GetUUIDMahasiswa(email string) (string, error) {
	tx := GetConnection2()
	ctx := context.Background()
	script := "select id_mhs from mahasiswa where email = ?"
	rows, err := tx.QueryContext(ctx, script, email)
	PanicIfError(err)
	defer rows.Close()

	var uuid string
	if rows.Next() {
		err := rows.Scan(&uuid)
		PanicIfError(err)
		return uuid, nil
	} else {
		return uuid, errors.New("mahasiswa belum terdaftar")
	}
}

func GetRegisteredApp(idregister string) bool {
	tx := GetConnection2()
	ctx := context.Background()
	script := "select id, registered_id, app_detail from bot_registered_app where registered_id = ?"
	rows, err := tx.QueryContext(ctx, script, idregister)
	PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}
}

func LogPublishDelay(recipient string, scheduleTime int64, passTime string, wa bool, email bool, telegram bool, statusNotification bool) int {
	tx := GetConnection2()
	ctx := context.Background()
	script := "insert into bot_notification_log(recipient, create_time, schedule_time, pass_time, wa, email, telegram, status_notification) values (?,?,?,?,?,?,?)"
	createTime := time.Now().UnixNano() / int64(time.Millisecond)
	id, err := tx.ExecContext(ctx, script, recipient, createTime, scheduleTime, passTime, wa, email, telegram, statusNotification)
	PanicIfError(err)

	lastId, err := id.LastInsertId()
	PanicIfError(err)

	return int(lastId)
}
