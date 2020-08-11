package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	db "github.com/runpixelrun/deputy_test/internal/data"
	"github.com/spf13/cobra"
)

var pgGetSubCMD = &cobra.Command{
	Use:   "getSubordinates",
	Short: "Returns subordinates or passed userID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(userID) < 1 {
			fmt.Println("You must provide a userID (eg, `--userID 2`)")
			return
		}

		fmt.Printf("\n..fetching subordinates for userID: %v\n", userID)
		ctx := context.Background()

		dbClient := db.NewClient().WithNeoAndPostgres()
		defer dbClient.Pg.Conn().Close(ctx)
		defer dbClient.Neo.Conn().Close()

		result, err := dbClient.GetSubordinates(ctx, userID)
		if err != nil {
			log.Printf("Error fetching subordinates: %v\n", err)
		}
		if result == nil {
			fmt.Println("No subordinates found for userID", userID)
			return
		}

		bytes, err := json.Marshal(result)
		if err != nil {
			log.Println("json.Marshal error:", err)
		}
		fmt.Printf("Subordinates for userID %v are:\n%v\n", userID, string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(pgGetSubCMD)
}