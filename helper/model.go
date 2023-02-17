package helper

import (
	"api-bot-timeline-reminder/model/domain"
	"api-bot-timeline-reminder/model/web"
)

func ToCreateAkunResponse(user domain.Admin) web.CreateAdminResponse {
	return web.CreateAdminResponse{
		IdUser:   user.Id,
		Username: user.User,
		Status:   "berhasil",
	}
}

func ToLoginAdminResponse(user domain.Admin) web.LoginAdminResponse {
	return web.LoginAdminResponse{
		IdUser:   user.Id,
		Username: user.User,
		Status:   "berhasil",
	}
}

func ToRegisterBotResponse(bot domain.RegisterBot, pesanvalidasi string) web.RegisterBotResponse {
	return web.RegisterBotResponse{
		Email:         bot.Email,
		PesanValidasi: pesanvalidasi,
	}
}

func ToLoginStudentResponse(student domain.Student) web.LoginUserResponse {
	return web.LoginUserResponse{
		IdUser: student.Uuid,
		Email:  student.Email,
		Role:   "student",
		Status: "login berhasil",
	}
}

func ToLoginStaffResponse(student domain.Staff) web.LoginUserResponse {
	return web.LoginUserResponse{
		IdUser: student.Id,
		Email:  student.Email,
		Role:   "staff",
		Status: "login berhasil",
	}
}

func ToInserDataMahasiswaResponse(student domain.Student) web.InsertDataMahasiswaResponse {
	return web.InsertDataMahasiswaResponse{
		Nama:   student.Nama,
		Email:  student.Email,
		Status: "insert berhasil",
	}
}

func ToInserDataStaffResponse(staff domain.Staff) web.InsertDataStaffResponse {
	return web.InsertDataStaffResponse{
		Nama:   staff.Nama,
		Email:  staff.Email,
		Status: "insert berhasil",
	}
}

func ToOtpResponse(otp domain.Otp) web.OtpResponse {
	return web.OtpResponse{
		Email:  otp.Email,
		NoHp:   otp.NoHp,
		Status: "OTP Terkirim",
	}
}

func ToOtpValidationResponse(otp domain.Otp) web.OtpValidationResponse {
	return web.OtpValidationResponse{
		Email:  otp.Email,
		Status: otp.Status,
	}
}

func ToGetMessagesResponse(messages domain.ChatbotMessages) web.GetChatbotMessagesResponse {
	return web.GetChatbotMessagesResponse{
		Id:       messages.Id,
		Tag1:     messages.Tag1,
		Tag2:     messages.Tag2,
		Tag3:     messages.Tag3,
		Tag4:     messages.Tag4,
		Tag5:     messages.Tag5,
		Messages: messages.Messages,
	}
}

func ToGetMessagesResponses(messages []domain.ChatbotMessages) []web.GetChatbotMessagesResponse {
	var getMessages []web.GetChatbotMessagesResponse
	for _, allMessages := range messages {
		getMessages = append(getMessages, ToGetMessagesResponse(allMessages))
	}
	return getMessages
}
