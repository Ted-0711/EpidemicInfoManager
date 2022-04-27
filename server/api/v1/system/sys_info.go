package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// @Summary 获取天气信息
// @Produce  application/json
// @Param data body nil true "城市代码"
// @Success 200 {object} response.Response{msg=string} "返回天气信息"
// @Router /info/getWeather [post]
func (b *BaseApi) GetWeather(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	res, _ := http.Get("http://api.k780.com:88/?app=weather.today&weaid=101020500&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json")
	s, _ := ioutil.ReadAll(res.Body)
	response.OkWithDetailed(response.PageResult{
		List:     string(s),
		Total:    1,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
