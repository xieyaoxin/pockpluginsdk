package test

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"reflect"
	"testing"
)

func Test_taskService_GetTaskTypeList(t *testing.T) {
	tests := []struct {
		name string
		want []*model.TaskType
	}{
		// TODO: Add test cases.
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := plugin_sdk.TaskServiceInstance.GetTaskTypeList(); got != nil {
				for _, taskType := range got {
					plugin_log.Info("%s:%s", taskType.TaskTypeId, taskType.TaskTypeName)
				}
			}
		})
	}
}

func Test_taskService_StartAndFinishTask(t *testing.T) {
	type args struct {
		TaskId string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{TaskId: "45"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := plugin_sdk.TaskServiceInstance
			if got := ta.StartAndFinishTask(tt.args.TaskId, "6"); got != tt.want {
				t.Errorf("FinishTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskService_GetTaskDetail(t *testing.T) {
	type args struct {
		TaskId string
	}
	tests := []struct {
		name string
		args args
		want *model.Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := plugin_sdk.TaskServiceInstance
			if got := ta.GetTaskDetail(tt.args.TaskId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTaskDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskService_GetTaskList(t *testing.T) {
	type args struct {
		taskType string
	}
	type test struct {
		name string
		args args
		want []*model.Task
	}
	tests := []test{}
	taskTypeList := plugin_sdk.TaskServiceInstance.GetTaskTypeList()
	for _, taskType := range taskTypeList {
		tests = append(tests, test{
			name: taskType.TaskTypeName,
			args: args{taskType: taskType.TaskTypeId},
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := plugin_sdk.TaskServiceInstance
			plugin_log.Info("%s", tt.name)
			if got := ta.GetTaskList(tt.args.taskType); got != nil {
				for _, task := range got {
					plugin_log.Info("\t%s:%s:%s", task.TaskId, task.TaskName, task.TaskTypeN)
				}
			}
		})
	}
}

func Test_taskService_StartTask(t *testing.T) {
	type args struct {
		TaskId string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := plugin_sdk.TaskServiceInstance
			ta.StartAndFinishAllTask()
		})
	}
}
