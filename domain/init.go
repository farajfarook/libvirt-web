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
		doms, err := List()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, doms)
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
