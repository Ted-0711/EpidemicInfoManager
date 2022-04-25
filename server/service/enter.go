package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	EpidemicInfoServiceGroup EpidemicInfo.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
