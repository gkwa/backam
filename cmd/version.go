package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
	"github.com/gkwa/backam/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of backam",
	Long:  `All software has versions. This is backam's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}