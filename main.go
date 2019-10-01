package main

import (
	"fmt"
	"os"

	_ "github.com/enbiso/libvirt-web/docs"
	"github.com/enbiso/libvirt-web/domain"
	"github.com/enbiso/libvirt-web/network"
	"github.com/labstack/echo"
	libvirt "github.com/libvirt/libvirt-go"
	echoSwagger "github.com/swaggo/echo-swagger"

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

// @title Libvirt Web
// @version 1.0
// @description Libvirt virtualization web API

// @contact.name Faraj Farook

// @BasePath /
func serverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Libvirt Web Server",
		Run: func(cmd *cobra.Command, args []string) {
			addr := cmd.Flags().Lookup("addr").Value.String()
			uri := cmd.Flags().Lookup("uri").Value.String()

			conn, err := libvirt.NewConnect(uri)
			if err != nil {
				panic(err)
			}
			e := echo.New()
			e.GET("/swagger/*", echoSwagger.WrapHandler)
			domain.Init(conn, e)
			network.Init(conn, e)
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
