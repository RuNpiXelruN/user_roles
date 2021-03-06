package cli

import (
	"fmt"
	"unicode"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var setRolesCMD = &cobra.Command{
	Use:   "setRoles",
	Short: "Sets roles to the DB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n%v%v %q\n", rootCmd.Use, "Setting roles...", unicode.Avestan.R16[3])
	},
}

func init() {
	rootCmd.AddCommand(setRolesCMD)
}
