package router

import (
	"Driving-school/api/v1"
	"Driving-school/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	SysOperationRecordRouter := Router.Group("sysOperationRecord").Use(middleware.OperationRecord())
	{
		SysOperationRecordRouter.POST("createSysOperationRecord", v1.CreateSysOperationRecord)             // 新建SysOperationRecord
		SysOperationRecordRouter.DELETE("deleteSysOperationRecord", v1.DeleteSysOperationRecord)           // 删除SysOperationRecord
		SysOperationRecordRouter.DELETE("deleteSysOperationRecordByIds", v1.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		SysOperationRecordRouter.GET("findSysOperationRecord", v1.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
		SysOperationRecordRouter.GET("getSysOperationRecordList", v1.GetSysOperationRecordList)            // 获取SysOperationRecord列表

	}
}
