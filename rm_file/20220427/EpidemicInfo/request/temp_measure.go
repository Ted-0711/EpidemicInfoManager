package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type Temp_measureSearch struct{
    EpidemicInfo.Temp_measure
    request.PageInfo
}
