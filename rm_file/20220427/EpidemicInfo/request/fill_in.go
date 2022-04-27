package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type Fill_inSearch struct{
    EpidemicInfo.Fill_in
    request.PageInfo
}
