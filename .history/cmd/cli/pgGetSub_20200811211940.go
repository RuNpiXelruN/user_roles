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
	Use:   "pgGetSub",
	Short: "Returns subordinates from the Postgres database given a userID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(userID) < 1 {
			fmt.Println("You must provide a userID (eg, `--userID 2`)")
			return
		}

		ctx := context.Background()

		dbClient := db.NewClient().WithPostgres()
		defer dbClient.Pg.Conn().Close(ctx)

		users, err := dbClient.GetSubordinates(ctx, userID)
		if err != nil {
			log.Printf("Error fetching subordinates: %v\n", err)
			return
		}

		if len(users) < 1 {
			fmt.Printf("No subordinates found for userID: %v\n", userID)
			return
		}

		bytes, err := json.Marshal(users)
		if err != nil {
			log.Println("json.Marshal error:", err)
			return
		}

		fmt.Printf("Subordinates for userID %v are:\n%v\n", userID, string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(pgGetSubCMD)
}
