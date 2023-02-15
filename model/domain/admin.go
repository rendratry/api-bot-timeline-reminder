package domain

type Admin struct {
	Id             string
	User           string
	Password       string
	LastTimeAccess int64
	IdUser         interface{}
}
