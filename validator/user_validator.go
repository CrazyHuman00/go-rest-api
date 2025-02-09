package validator

import (
	"go-rest-api/model"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IUserValidator interface {
	UserValidator(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidator(user model.User) error {
	// 第一引数にバリデーション対象の構造体
	// 第二引数以降にユーザのemailとpasswordに対するバリデーションルールを指定
	// Requiredで必須項目、RuneLengthで文字数制限（最小1最大30）を指定
	// is.Emailでemailの形式チェック
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited max 30 char"),
		),
	)
}