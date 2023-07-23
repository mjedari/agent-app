package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "agent-app",
		Short: "short description",
		Long:  `long description`,
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Mahdi Jedari", "i.jedari@gmail.com")
}
