package sms

type SMSTokenError string

func (e SMSTokenError) Error() string {
	return string(e)
}

const (
	SMSCodeIsInvalid             = SMSTokenError("SMS code is invalid")
	SMSIsNotConfirmed            = SMSTokenError("SMS is not confirmed")
	SMSIsAlreadyConfirmed        = SMSTokenError("SMS is already confirmed")
	SMSConfirmationTypeIsInvalid = SMSTokenError("SMS confirmation type is invalid")
	SMSTokenNotFound             = SMSTokenError("SMS token not found")
	PhoneNumberIsMismatch        = SMSTokenError("Phone number is mismatch")
)
