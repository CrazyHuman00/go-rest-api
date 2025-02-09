package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidator(task model.Task) error
}

type taskValidator struct{}

// 構造体を生成するコンストラクタ
func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

// Taskのバリデーション
func (tv *taskValidator) TaskValidator(task model.Task) error {
	// 第一引数にバリデーション対象の構造体
	// 第二引数以降にタスクのタイトルに対するバリデーションルールを指定
	// Requiredで必須項目、RuneLengthで文字数制限（最小1最大10）を指定
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title, 
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
	)
}