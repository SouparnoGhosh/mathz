package mathz

import (
	"fmt"

	"github.com/souparno/mathz/pkg/mathz"
	"github.com/spf13/cobra"
)

var mulCmd = &cobra.Command{
	Use: "mul",
	Aliases: []string{"mul"},
	Short: "Multiplies 2 numbers",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res := mathz.Multiply(args[0], args[1])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(mulCmd)
}

