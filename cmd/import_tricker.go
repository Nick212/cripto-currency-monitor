package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	// "github.com/EnubeRepos/CriptoCurrencyMonitor/app"
	"github.com/Nick212/cripto-currency-monitor/app"
	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/spf13/cobra"
)

var trickerCmd = &cobra.Command{
	Use: "tricker",
	RunE: func(cmd *cobra.Command, args []string) error {

		a, err := app.New()

		if err != nil {
			fmt.Println(err)
		}
		defer a.Close()

		ctx := a.NewContext(a.Config.Version, a.Config.Culture, "", "")
		ctx.Logger.Println("HOST Session : ", a.Config.HOST_API)

		result, err := a.GetTricker()

		if err != nil {
			ctx.Logger.Error("Error Get Tricker: ", err)
		}

		client, errEs := elasticsearch6.NewDefaultClient()

		if errEs != nil {
			ctx.Logger.Error("Error Get errEs: ", err)
		}

		document, err := json.Marshal(result)

		if err != nil {
			ctx.Logger.Error("Error Marshal Struct: ", err)
		}

		req := esapi.IndexRequest{
			Index:      a.Config.IndexName,
			DocumentID: "1",
			Body:       strings.NewReader(string(document)),
		}
		req.Do(context.Background(), client)

		// ctx.Logger.Println(client.Info())

		ctx.Logger.Println(result.Last)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(trickerCmd)
}
