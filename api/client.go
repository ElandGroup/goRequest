package api

import (
	"goRequest/logHelper"
	"net/http"

	"github.com/labstack/echo"
	"github.com/smallnest/goreq"
)

type (
	APIResult struct {
		Result  interface{} `json:"result"`
		Success bool        `json:"success"`
		Error   APIError    `json:"error"`
	}

	APIError struct {
		Code    int         `json:"code"`
		Details interface{} `json:"details"`
		Message string      `json:"message"`
	}
)

func ReqXml(c echo.Context) error {

	xmldata := `
	<xml><appid>xiaomiao</appid></xml>
	`
	req, body, err := goreq.New().Post("https://api.mch.weixin.qq.com/sandbox/pay/micropay").ContentType("xml").SendRawString(xmldata).End()
	if req.StatusCode == http.StatusOK {
		logHelper.Debug.Println("\n body:", body)
	}
	logHelper.Debug.Println("\n err:", err)
	logHelper.Debug.Println("\n req:", req)

	return c.String(http.StatusOK, body)

}
