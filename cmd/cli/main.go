package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ccfos/nightingale/v6/cli"
	"github.com/ccfos/nightingale/v6/pkg/version"
)

var (
	upgrade     = flag.Bool("upgrade", false, "Upgrade the database.")
	showVersion = flag.Bool("version", false, "Show version.")
	configDir   = flag.String("config", "", "Specify configuration directory.(env:N9E_CONFIGS)")
)

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	if *upgrade {
		if *configDir == "" {
			fmt.Println("Please specify the configuration directory.")
			os.Exit(1)
		}

		err := cli.Upgrade(*configDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print("Upgrade successfully.")
		os.Exit(0)
	}
}
