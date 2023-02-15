package service

import (
	"api-bot-timeline-reminder/exception"
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/domain"
	"api-bot-timeline-reminder/model/web"
	"api-bot-timeline-reminder/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"time"
)

type AkunServiceImpl struct {
	AkunRepository repository.AkunRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(akunRepository repository.AkunRepository, DB *sql.DB, validate *validator.Validate) *AkunServiceImpl {
	return &AkunServiceImpl{
		AkunRepository: akunRepository,
		DB:             DB,
		Validate:       validate}
}

func (service *AkunServiceImpl) LoginAdmin(ctx context.Context, request web.LoginAdminRequest) web.LoginAdminResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	admin := domain.Admin{
		User: request.Username,
	}

	login, err := service.AkunRepository.LoginAdmin(ctx, tx, admin)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	_, err = helper.CheckPasswordHash(request.Password, login.Password, "email atau password salah")
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	login.User = request.Username

	helper.UpdateLastAccessLoginAdmin(strconv.Itoa(int(time.Now().UnixNano() / 1000000)))

	return helper.ToLoginAdminResponse(login)
}

func (service *AkunServiceImpl) CreateAdmin(ctx context.Context, request web.CreateAdminRequest) web.CreateAdminResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pwhash, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	admin := domain.Admin{
		Id:             uuid.NewV4().String(),
		User:           request.Username,
		Password:       pwhash,
		LastTimeAccess: time.Now().UnixNano() / 1000000,
	}

	admin = service.AkunRepository.CreateAdmin(ctx, tx, admin)

	return helper.ToCreateAkunResponse(admin)
}

func (service *AkunServiceImpl) RegisterBot(ctx context.Context, request web.RegisterBotRequest) web.RegisterBotResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	str1 := helper.EncodeToString(4)
	str2 := helper.EncodeToString(5)

	uuidstr, err := helper.GetUUIDMahasiswa(request.Email)
	helper.PanicIfError(err)
	bot := domain.RegisterBot{
		IdRegister:  str1 + "-" + str2,
		Email:       request.Email,
		NamaLengkap: request.NamaLengkap,
		NoHp:        request.NoHp,
		CreateAt:    time.Now().UnixNano() / 1000000,
		StatusRegis: 0,
	}

	bot = service.AkunRepository.RegisterBot(ctx, tx, bot)

	pesanvalidasi := "Jangan merubah pesan ini! silahkan dikirim ke chatbot, validasi:"

	return helper.ToRegisterBotResponse(bot, pesanvalidasi+str1+"-"+str2+"."+uuidstr)
}

func (service *AkunServiceImpl) LoginUser(ctx context.Context, request web.LoginUserRequest) web.LoginUserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	if strings.Contains(request.Email, "staff.uns.ac.id") {
		staff := domain.Staff{
			Email: request.Email,
		}

		login, err := service.AkunRepository.LoginUserStaff(ctx, tx, staff)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		_, err = helper.CheckPasswordHash(request.Password, login.Password, "email atau password salah")
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		return helper.ToLoginStaffResponse(login)

	}

	student := domain.Student{
		Email: request.Email,
	}

	login, err := service.AkunRepository.LoginUserStudent(ctx, tx, student)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	_, err = helper.CheckPasswordHash(request.Password, login.Password, "email atau password salah")
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToLoginStudentResponse(login)

}

func (service *AkunServiceImpl) InsertDataMahasiswa(ctx context.Context, request web.InsertDataMahasiswaRequest) web.InsertDataMahasiswaResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	insert := domain.Student{
		Uuid:          uuid.NewV4().String(),
		Nama:          request.Nama,
		Email:         request.Email,
		TahunAngkatan: request.TahunAngkatan,
		Nim:           request.Nim,
		Prodi:         request.Prodi,
	}

	insert = service.AkunRepository.InsertDataMahasiswa(ctx, tx, insert)
	return helper.ToInserDataMahasiswaResponse(insert)
}

func (service *AkunServiceImpl) InsertDataDosen(ctx context.Context, request web.InsertDataStaffRequest) web.InsertDataStaffResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	insert := domain.Staff{
		Id:          uuid.NewV4().String(),
		Nama:        request.Nama,
		Nip:         request.Nip,
		TglLahir:    "",
		Alamat:      request.Alamat,
		Status:      request.Status,
		Email:       request.Email,
		NotifStatus: true,
	}

	insert = service.AkunRepository.InsertDataStaff(ctx, tx, insert)
	return helper.ToInserDataStaffResponse(insert)
}
