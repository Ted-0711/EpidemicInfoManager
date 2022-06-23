package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/EpidemicInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type Quarantine_siteSearch struct{
    EpidemicInfo.Quarantine_site
    request.PageInfo
}
