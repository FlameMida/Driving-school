package upload

import (
	"Driving-school/global"
	"mime/multipart"
)

//@interface_name: OSS
//@description: OSS接口

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

//@function: NewOss
//@description: OSS接口
//@description: OSS的实例化方法
//@return: OSS

func NewOss() OSS {
	switch global.GVA_CONFIG.System.OssType {
	case "local":
		return &Local{}
	default:
		return &Local{}
	}
}
