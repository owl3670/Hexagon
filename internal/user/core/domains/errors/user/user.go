package user

type UserError string

func (e UserError) Error() string {
	return string(e)
}

const (
	EmailInvalid                   = UserError("Email id invalid")
	NicknameIsEmpty                = UserError("Nickname is empty")
	PasswordIsEmpty                = UserError("Password is empty")
	NameIsEmpty                    = UserError("Name is empty")
	PhoneNumberIsEmpty             = UserError("Phone number is empty")
	ConfirmPasswordIsInvalid       = UserError("Confirm password is invalid")
	UserNotFound                   = UserError("User not found")
	UserEmailAlreadyExists         = UserError("User email already exists")
	UserNicknameAlreadyExists      = UserError("User nickname already exists")
	UserPhoneNumberAlreadyExists   = UserError("User phone number already exists")
	UserIdentificationParamIsEmpty = UserError("User identifiacation param is empty")
)
