package helper

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func GetConnection2() *sql.DB {
	db, err := sql.Open("mysql", "myfin:Admin@myfin123@tcp(103.189.234.90:3306)/bot")
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
	script := "select uuid from user where email = ?"
	rows, err := tx.QueryContext(ctx, script, email)
	PanicIfError(err)
	defer rows.Close()

	var uuid string
	if rows.Next() {
		err := rows.Scan(&uuid)
		PanicIfError(err)
		return uuid, nil
	} else {
		return uuid, errors.New("failed")
	}
}
