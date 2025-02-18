package plugin_sdk

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
)

var TaskServiceInstance = &taskService{}
var TaskRepository = repository.GetTaskRepository()

type taskService struct {
}

func (*taskService) GetTaskTypeList() []*model.TaskType {
	return TaskRepository.GetTaskTypeList()
}

func (*taskService) GetTaskList(taskType string) []*model.Task {
	return TaskRepository.GetTaskList(taskType)
}

func (inst *taskService) StartAndFinishAllTask() {
	TaskList := inst.GetTaskList("2")
	for _, Task := range TaskList {
		inst.FinishTask(Task.TaskId, Task.TaskTypeN)
		inst.StartTask(Task.TaskId)
	}
}
func (inst *taskService) StartAndFinishTask(TaskId string, TaskTypeN string) bool {
	inst.StartTask(TaskId)
	return inst.FinishTask(TaskId, TaskTypeN)
}
func (*taskService) StartTask(TaskId string) bool {
	return TaskRepository.StartTask(TaskId)
}
func (*taskService) FinishTask(TaskId string, TaskTypeN string) bool {
	return TaskRepository.FinishTask(TaskId, TaskTypeN)
}
func (*taskService) GetTaskDetail(TaskId string) *model.Task {
	return TaskRepository.GetTaskDetail(TaskId)
}
