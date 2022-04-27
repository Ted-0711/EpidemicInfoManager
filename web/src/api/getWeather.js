import service from '@/utils/request'
// @Summary 获取天气信息
// @Produce  application/json
// @Param data body {}
// @Router /base/captcha [post]
export const getWeather = (data) => {
  return service({
    url: '/info/getWeather',
    method: 'post',
    data
  })
}