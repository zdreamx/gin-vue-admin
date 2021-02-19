// 自动生成模板Seedling
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Seedling struct {
	global.GVA_MODEL
	Enterprise string `json:"enterprise " form:"enterprise "`
	Mobile     string `json:"mobile" form:"mobile"`
	Technology string `json:"technology " form:"technology "`
	Process    string `json:"process" form:"process"`
	SeedDate   string `json:"seedDate" form:"seedDate"`
	Passwd     string `json:"passwd" form:"passwd"`
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type SeedlingWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Seedling   `json:"business"`
// }

// func (Seedling) TableName() string {
// 	return ""
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["seed"] = func() model.GVA_Workflow {
//   return new(model.SeedlingWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["seed"] = func() interface{} {
// 	return new(model.Seedling)
// }
