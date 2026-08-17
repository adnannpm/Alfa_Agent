package cmd

import (
	"alfa_agent/internal/helpers"
	"alfa_agent/internal/version"
	"bytes"
	"encoding/json"
	"fmt"
	"alfa_agent/internal/log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	token string
	host  string
	name  string
)

var enrollCmd = &cobra.Command{
	Use:   "enroll",
	Short: "Enroll the agent with a master server",
	Long: `Register this Alfa Agent with a master server using an enrollment token.
The command saves the token and server address to the local configuration, then
sends the agent identity and version information to the master server.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.InsertToken(token, host)
		if err != nil {
			log.Error(err)
		}

		log.Info("Susscessfuly insert to config!")

		endpoint := fmt.Sprintf("http://%s/api/token", host)
		data := map[string]interface{}{
			"token_agent": token,
			"name":        name,
			"version":     version.VERSION,
			"ip_address":  host,
		}

		jsonData, _ := json.Marshal(data)

		if _, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData)); err != nil {
			log.Error(err)
		}

		log.Info("Susscessfuly add new server to master")
	},
}

func init() {
	rootCmd.AddCommand(enrollCmd)

	enrollCmd.Flags().StringVar(&token, "token", "", "Need token please...")
	enrollCmd.Flags().StringVar(&host, "host", "", "Need host please...")
	enrollCmd.Flags().StringVar(&name, "name", "", "Need name please...")
}
