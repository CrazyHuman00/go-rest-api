package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

// コンストラクタ（repositoryとのDIとvalidatorのDI）
func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	// 構造体の実体のポインタを返す
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	// TaskResponseの新しい構造体を作成
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	// 空のtaskオブジェクトを作成
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTasks := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTasks, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	// タスクのバリデーションを行う
	if err := tu.tv.TaskValidator(task); err != nil {
		return model.TaskResponse{}, err
	}
	// タスクの作成
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	resTasks := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTasks, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	// タスクのバリデーションを行う
	if err := tu.tv.TaskValidator(task); err != nil {
		return model.TaskResponse{}, err
	}
	// タスクの更新
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTasks := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTasks, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}