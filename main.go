package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	serverCmd := serverCmd()
	serverCmd.AddCommand(versionCmd())

	if err := serverCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func serverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Lib Virt Web Server",
		Run: func(cmd *cobra.Command, args []string) {
			addr := cmd.Flags().Lookup("addr").Value.String()
			uri := cmd.Flags().Lookup("uri").Value.String()

			conn, err := newConnection(uri)
			if err != nil {
				panic(err)
			}

			e := echo.New()
			e.GET("/domains", func(c echo.Context) error {
				doms, err := conn.ListDomains()
				if err != nil {
					return err
				}
				return c.JSON(http.StatusOK, doms)
			})
			e.GET("/domains/:name", func(c echo.Context) error {
				dom, err := conn.GetDomain(c.Param("name"))
				if err != nil {
					return err
				}
				return c.JSON(http.StatusOK, dom)
			})
			e.GET("/domains/:name/_xml", func(c echo.Context) error {
				content, err := conn.GetDomainXML(c.Param("name"))
				if err != nil {
					return err
				}
				return c.XMLBlob(http.StatusOK, []byte(content))
			})
			e.Start(addr)
		},
	}
	cmd.Flags().StringP("uri", "u", "qemu:///system", "Virtd URI")
	cmd.Flags().StringP("addr", "a", ":8080", "API listening address")
	viper.BindPFlag("uri", cmd.Flags().Lookup("uri"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	return cmd
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of LibVirt Web",
		Long:  `All software has versions. This is LibVirt Web's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("LibVirt Web v0.1.0")
		},
	}
}
