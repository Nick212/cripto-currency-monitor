package cmd

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/EnubeRepos/CriptoCurrencyMonitor/app"
	"github.com/EnubeRepos/CriptoCurrencyMonitor/models/crm"

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
		ctx := a.NewContext(a.Config.Version, "en-us", "", "")

		for _, host := range a.Config.HOST_API {
			a.Config.HOST_CRM = host
			ctx.Logger.Println("HOST Session : ", a.Config.HOST_CRM)

			ctx.Logger.Debug(apiExternal.Total)

		}

		return nil
	},
}


func init() {
	rootCmd.AddCommand(trickerCmd)
}
