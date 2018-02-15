package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/srbry/cf-manifest-updater/manifest"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "cf-manifest-updater"
	app.Usage = "updates golang manifests"
	app.Action = updateManifest
	app.Run(os.Args)
}

func updateManifest(c *cli.Context) error {
	manifestFile := c.Args().First()
	if manifestFile == "" {
		return fmt.Errorf("must provide a source manifest file")
	}

	oldManifest, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		return err
	}

	newManifest, err := manifest.Update(oldManifest)
	if err != nil {
		return err
	}

	fmt.Printf("---\n%s\n", newManifest)
	return nil
}
