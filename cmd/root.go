package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cmUrl string

func init() {

	cmUrl = os.Getenv("cm_url")

	if len(cmUrl) == 0 {
		cmUrl = "https://api.createsend.com/api/v3.1"
		fmt.Println("using default url:", cmUrl)
	} else {
		fmt.Println("using url:", cmUrl)
	}
}

var rootCmd = &cobra.Command{
	Use:   "cm",
	Short: "cm is a simple cli application that use campaign monitor APIs",
	Long:  `cm is a simple cli application that use campaign monitor APIs details can be found https://github.com/wy3148/cm`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("run './cm help' to see how to use the application")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
