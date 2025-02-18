package repository

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"

type TaskRepository interface {
	GetTaskTypeList() []*model.TaskType
	GetTaskList(taskType string) []*model.Task
	StartTask(TaskId string) bool
	FinishTask(TaskId string, TaskTypeN string) bool
	GetTaskDetail(TaskId string) *model.Task
}
