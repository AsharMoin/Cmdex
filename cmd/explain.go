package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain what a command or keyword does",
	Long:  "Will do this later",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(explainCommand(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}

func explainCommand(command string) string {
	data, err := os.Open("./resources/commands.json")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()
	cmdMap := make(map[string]interface{})

	if err := json.NewDecoder(data).Decode(&cmdMap); err != nil {
		fmt.Println(err)
	}

	commands, ok := cmdMap["commands"].([]interface{})
	if !ok {
		fmt.Println("commands is not a slice")
	}

	for _, item := range commands {
		cmdInfo, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("item is not a map")
			continue
		}

		// Extract the "command" and "description" fields as strings
		cmdName, _ := cmdInfo["command"].(string)
		desc, _ := cmdInfo["description"].(string)

		if cmdName == command {
			return desc
		}
	}

	return ""
}
