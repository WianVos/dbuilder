package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wianvos/dbuilder/pkg/config"
	"github.com/wianvos/dbuilder/pkg/repo"
)

var dbuilder = &cobra.Command{

	Short: "dbuilder is a build tool for docker images that will let you setup and automate seperate docker image builds based on version info, configuration and templated dockerfiles",
}

var build = &cobra.Command{
	Use:   "build",
	Short: "build a directory",
	Long:  "build docker images based on a git repo containing dbuilder info", //TODO document the long output,
	Run:   dbuild,
}

func init() {
	dbuilder.AddCommand(build)
}

func main() {

	dbuilder.Execute()

}

func dbuild(cmd *cobra.Command, args []string) {
	//get the git repo locally

	//run on a straight directory
	dir := args[0]
	r, err := repo.NewRepo(dir)
	e, err := config.NewConfigFromFile(dir + "/dbuilder.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//spew.Dump(project)
	fmt.Println(r)

	r.SetConfig(e)

	r.Execute()
	//verify our requirements
	// is the config in the repo
	// is there a templated docker build file in /templates

	//build the image

}
