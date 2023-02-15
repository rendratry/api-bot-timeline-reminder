package service

import (
	"api-bot-timeline-reminder/model/web"
	"context"
)

type AkunService interface {
	LoginAdmin(ctx context.Context, request web.LoginAdminRequest) web.LoginAdminResponse
	CreateAdmin(ctx context.Context, request web.CreateAdminRequest) web.CreateAdminResponse
	RegisterBot(ctx context.Context, request web.RegisterBotRequest) web.RegisterBotResponse
	LoginUser(ctx context.Context, request web.LoginUserRequest) web.LoginUserResponse
	InsertDataMahasiswa(ctx context.Context, request web.InsertDataMahasiswaRequest) web.InsertDataMahasiswaResponse
	InsertDataStaff(ctx context.Context, request web.InsertDataStaffRequest) web.InsertDataStaffResponse
}
