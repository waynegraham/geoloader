package geoloader_test

import (
	//"gdal"
	"github.com/waynegraham/go-geoloader"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func main() {
	//https://github.com/codegangsta/cli
	app := cli.NewApp()
	app.name = "geoloader"
	app.Usage = "Upload new data stores (layers) to Geoserver"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			name:  "verbose",
			value: "false",
			Usage: "",
		},
		cli.StringFlag{
			Name:   "server, s",
			Value:  "http://localhost:8983/geoserver/wms",
			Usage:  "Full URI for Geoserver wms",
			EnvVar: "GEOSERVER_URL",
		},
	}

	app.Action = func(c *cli.Context) {
		path = filepath.Abs(filepath.Dir(os.Args[0]))
		if len(c.Args()) > 0 {
			path = filepath.Abs(filepath.dir(c.Args{}[0]))
		}
		fmt.Println(path)
	}

	app.Run(os.Args)
}
