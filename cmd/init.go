package cmd

import (
	"alfa_agent/internal/config"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the agent configuration and database",
	Long: `Prepare Alfa Agent for first use by generating the required configuration
files and initializing its local database. Existing files are preserved.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Initializing configuration...")
		if err := createDir(configPath, "Config"); err != nil {
			log.Fatalln("Failed to inititalize config directory: ", err)
		}

		if err := writeFile(filepath.Join(configPath, "config.yml"), "config.yml", config.ConfigFile); err != nil {
			log.Fatalln("Failed to initialize config file: %w", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createDir(path, name string) error {
	if info, err := os.Stat(path); err == nil {
		if !info.IsDir() {
			return fmt.Errorf("%s path exists but is not a directory: %s", name, path)
		}
		fmt.Printf("%s directory already exists", name)
		return nil
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create %s directory %s: %w", name, path, err)
	}

	fmt.Printf("%s directory created", name)

	return nil
}

func writeFile(path, name, content string) error {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("%s already exists", name)

		return nil
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create %s %s: %w", name, path, err)
	}

	fmt.Printf("%s created", name)

	return nil
}
