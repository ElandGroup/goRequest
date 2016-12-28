package api

import (
	"goRequest/logHelper"

	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func AcceptJson(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return nil
	}
	u.Name = u.Name + ",Good"
	logHelper.Debug.Printf("%+v", u)
	return c.JSON(http.StatusOK, u)
}
