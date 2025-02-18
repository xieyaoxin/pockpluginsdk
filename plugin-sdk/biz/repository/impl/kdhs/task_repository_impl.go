package kdhs

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"strings"
)

var TaskRepositoryKdhsImplInstance = &taskRepositoryKdhsImpl{}

type taskRepositoryKdhsImpl struct{}

func (*taskRepositoryKdhsImpl) GetTaskTypeList() []*model.TaskType {
	//panic("implement me")
	//http://43.248.129.148:567/?title_vary=1
	Params := util.InitParam()
	Params["title_vary"] = "1"
	response := CallServerGetInterface("function/taskshow.php", Params)
	// 根据格式获取任务列表
	TaskList := []*model.TaskType{}
	lines := strings.Split(response, "<ul")
	for _, line := range lines {
		if strings.Contains(line, "getTaskDetail") {
			array := strings.Split(util.ReplaceAllString(line, "<", "'", ">"), "<")
			TaskList = append(TaskList, &model.TaskType{
				TaskTypeId:   array[7],
				TaskTypeName: array[11],
			})

		}
	}
	return TaskList
}

func (*taskRepositoryKdhsImpl) GetTaskList(taskType string) []*model.Task {
	Params := util.InitParam()
	if taskType == "2" {
		Params["title_vary"] = "3"
	} else {
		Params["title_vary"] = "2"
	}
	Params["bid"] = taskType
	Params["rd"] = "0.9080655236537869"
	response := CallServerGetInterface("function/taskshow.php", Params)
	taskList := []*model.Task{}
	for _, line := range strings.Split(response, "<li") {
		if strings.Contains(line, "taskASwap") {
			array := strings.Split(util.ReplaceAllString(line, ",", "'", ">", "<"), ",")
			taskList = append(
				taskList, &model.Task{
					TaskId:    array[5],
					TaskName:  array[8],
					TaskTypeN: array[6],
				},
			)
		}
	}
	return taskList
}

func (*taskRepositoryKdhsImpl) StartTask(TaskId string) bool {
	//	 getTask.php?taskid=571&type=get
	Params := util.InitParam()
	Params["taskid"] = TaskId
	Params["type"] = "get"
	response := CallServerGetInterface("function/getTask.php", Params)
	plugin_log.Info(response)
	return "恭喜您，成功接受此任务！" == response
}

func (*taskRepositoryKdhsImpl) FinishTask(TaskId, TaskTypeN string) bool {
	Params := util.InitParam()
	Params["taskid"] = TaskId
	Params["type"] = "complate"
	Params["n"] = TaskTypeN
	response := CallServerGetInterface("function/getTask.php", Params)
	plugin_log.Info(response)
	return strings.Contains(response, "任务完成")
}

func (*taskRepositoryKdhsImpl) GetTaskDetail(TaskId string) *model.Task {
	//TODO implement me
	panic("implement me")
}
