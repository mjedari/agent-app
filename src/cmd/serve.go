package cmd

import (
	"context"
	"fmt"
	"github.com/mjedari/agent-app/src/pkg"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving service.",
	Long:  `Serving service.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())

		serve(ctx)
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		cancel()
		fmt.Println()

		// Perform any necessary cleanup before exiting
		fmt.Println("\nApplication exited successfully.")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(ctx context.Context) {
	// create agents
	agentPool := pkg.NewAgentPool(pkg.NumberOfAgents)
	agentPool.Initiate()

	// get the target
	newTarget := pkg.NewTarget(9, -1)

	// find the smallest steps to this specific target
	agent := agentPool.SelectAgent(newTarget)

	// set target
	agent.SetTarget(newTarget)

	// move agent
	agent.Move()
}
