package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CounselorSearch struct{
    EpidemicInfo.Counselor
    request.PageInfo
}
