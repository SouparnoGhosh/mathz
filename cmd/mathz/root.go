package mathz

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mathz",
	Short: "mathz - basic calculations",
	Long: "mathz is a CLI for basic calculations.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}