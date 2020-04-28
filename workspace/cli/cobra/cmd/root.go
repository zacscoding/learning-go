package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// https://github.com/spf13/cobra

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Short message",
	Long:  "Long Long Long message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running :)")
	},
}

func init() {
	fmt.Println("cmd/root.go::init() is called")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
