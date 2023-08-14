package mathz

import (
	"fmt"
	// "os"

	"github.com/souparno/mathz/pkg/mathz"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Aliases: []string{"addition"},
	Short: "Adds 2 numbers",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res := mathz.Add(args[0], args[1])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

