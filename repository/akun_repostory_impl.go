package repository

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/domain"
	"context"
	"database/sql"
	"errors"
)

type AkunRepositoryImpl struct {
}

func NewAkunRepositoryImpl() *AkunRepositoryImpl {
	return &AkunRepositoryImpl{}
}

func (repository *AkunRepositoryImpl) LoginAdmin(ctx context.Context, tx *sql.Tx, admin domain.Admin) (domain.Admin, error) {
	script := "select id, `username`, password from admin where `username` = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, admin.User)
	helper.PanicIfError(err)
	user := domain.Admin{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.User, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("email atau password salah")
	}
}

func (repository *AkunRepositoryImpl) CreateAdmin(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin {
	script := "insert into admin(id, username, password, last_time_access) values (?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, admin.Id, admin.User, admin.Password, admin.LastTimeAccess)
	helper.PanicIfError(err)

	return admin
}

func (repository *AkunRepositoryImpl) RegisterBot(ctx context.Context, tx *sql.Tx, bot domain.RegisterBot) domain.RegisterBot {
	script := "insert into bot_registrasi(id_registrasi, email, nama_lengkap, no_wa, create_at, regis_status) values (?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, bot.IdRegister, bot.Email, bot.NamaLengkap, bot.NoHp, bot.CreateAt, bot.StatusRegis)
	helper.PanicIfError(err)

	return bot
}

func (repository *AkunRepositoryImpl) LoginUserStudent(ctx context.Context, tx *sql.Tx, student domain.Student) (domain.Student, error) {
	script := "select uuid, nama, email, nim, prodi, password from user where `email` = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, student.Email)
	helper.PanicIfError(err)
	userStudent := domain.Student{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&userStudent.Uuid, &userStudent.Nama, &userStudent.Email, &userStudent.Nim, &userStudent.Prodi, &userStudent.Password)
		helper.PanicIfError(err)
		return userStudent, nil
	} else {
		return userStudent, errors.New("email atau password salah")
	}
}

func (repository *AkunRepositoryImpl) LoginUserStaff(ctx context.Context, tx *sql.Tx, staff domain.Staff) (domain.Staff, error) {
	script := "select id, nama, email, password from user where `email` = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, staff.Email)
	helper.PanicIfError(err)
	userStaff := domain.Staff{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&userStaff.Id, &userStaff.Nama, &userStaff.Email, &userStaff.Password)
		helper.PanicIfError(err)
		return userStaff, nil
	} else {
		return userStaff, errors.New("email atau password salah")
	}
}

func (repository *AkunRepositoryImpl) InsertDataMahasiswa(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	script := "insert into user(uuid, nama, email, tahun_angkatan, nim, prodi) values(?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, student.Uuid, student.Nama, student.Email, student.TahunAngkatan, student.Nim, student.Prodi)
	helper.PanicIfError(err)

	return student
}

func (repository *AkunRepositoryImpl) InsertDataStaff(ctx context.Context, tx *sql.Tx, staff domain.Staff) domain.Staff {
	script := "insert into dosen(id, nama, nip, tgl_lahir, alamat, status, email, notif_status) values (?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, staff.Id, staff.Nama, staff.Nip, staff.TglLahir, staff.Alamat, staff.Status, staff.Email, staff.NotifStatus)
	helper.PanicIfError(err)

	return staff
}
