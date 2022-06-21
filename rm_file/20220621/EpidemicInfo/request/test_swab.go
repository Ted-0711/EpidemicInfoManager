package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type Test_swabSearch struct{
    EpidemicInfo.Test_swab
    request.PageInfo
}
