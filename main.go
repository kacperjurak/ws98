package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	hostname := flag.String("h", "", "host number")
	portNumber := flag.String("p", "8080", "port number")
	dirName := flag.String("d", "public", "directory name")

	ok, err := exists(*dirName)
	if err != nil {
		panic(err)
	}

	if !ok {
		fmt.Println(fmt.Sprintf("directory '%s' does not exist", *dirName))
		os.Exit(1)
	}

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       *dirName,
		Browse:     true,
		IgnoreBase: true,
	}))

	e.Static("/", *dirName)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", *hostname, *portNumber)))
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
