package network

import (
	"net/http"

	"github.com/labstack/echo"
	libvirt "github.com/libvirt/libvirt-go"
)

var conn *libvirt.Connect

//Init network
func Init(connect *libvirt.Connect, e *echo.Echo) {
	conn = connect
	e.GET("/networks", func(c echo.Context) error {
		doms, err := list()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, doms)
	})
	e.GET("/networks/:name", func(c echo.Context) error {
		dom, err := get(c.Param("name"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, dom)
	})
	e.GET("/networks/:name/_xml", func(c echo.Context) error {
		content, err := getXML(c.Param("name"))
		if err != nil {
			return err
		}
		return c.XMLBlob(http.StatusOK, []byte(content))
	})
}
