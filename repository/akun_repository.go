package repository

import (
	"api-bot-timeline-reminder/model/domain"
	"context"
	"database/sql"
)

type AkunRepository interface {
	LoginAdmin(ctx context.Context, tx *sql.Tx, admin domain.Admin) (domain.Admin, error)
	CreateAdmin(ctx context.Context, tx *sql.Tx, admin domain.Admin) domain.Admin
	RegisterBot(ctx context.Context, tx *sql.Tx, bot domain.RegisterBot) domain.RegisterBot
	LoginUserStudent(ctx context.Context, tx *sql.Tx, student domain.Student) (domain.Student, error)
	LoginUserStaff(ctx context.Context, tx *sql.Tx, staff domain.Staff) (domain.Staff, error)
	InsertDataMahasiswa(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student
	InsertDataStaff(ctx context.Context, tx *sql.Tx, staff domain.Staff) domain.Staff
}
