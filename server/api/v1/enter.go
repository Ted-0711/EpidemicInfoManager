package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup       system.ApiGroup
	ExampleApiGroup      example.ApiGroup
	EpidemicInfoApiGroup EpidemicInfo.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
