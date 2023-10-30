package app

func ConvertToStruct(request EmailRequest, email *Email) {
	email.To = request.To
	email.Subject = request.Subject
	email.Message = request.Message
	email.Type = request.Type
}
