package cmd

import (
	"fmt"
	"os"

	"alfa_agent/internal/constants"
	"alfa_agent/internal/helpers"
	"alfa_agent/internal/services"
	"alfa_agent/internal/log"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Alfa Agent",
	Long: `Start the Alfa Agent and keep it running to perform its configured tasks
and communicate with the master server.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat(fmt.Sprintf("%s/config.yml", constants.CONFIG_PATH))
		if os.IsNotExist(err){
			log.Danger("Config file is not found!, please init before run")
			return
		}

		if helpers.ValidateToken() == false{
			log.Danger("Token is not valid!, Please register token!")
			return
		}

		services.StartingPolling(5)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
