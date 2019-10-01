package domain

import (
	"net/http"

	"github.com/labstack/echo"
	libvirt "github.com/libvirt/libvirt-go"
)

var conn *libvirt.Connect

//Init domain
func Init(connect *libvirt.Connect, e *echo.Echo) {
	conn = connect
	e.GET("/domains", func(c echo.Context) error {
		doms, err := list()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, doms)
	})
	e.POST("/domains", func(c echo.Context) error {
		req := CreateRequest{}
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(400, err.Error())
		}
		resp, err := create(req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, resp)
	})
	e.GET("/domains/:name", func(c echo.Context) error {
		dom, err := get(c.Param("name"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, dom)
	})
	e.GET("/domains/:name/_xml", func(c echo.Context) error {
		content, err := getXML(c.Param("name"))
		if err != nil {
			return err
		}
		return c.XMLBlob(http.StatusOK, []byte(content))
	})
}
