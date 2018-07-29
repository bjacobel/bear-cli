package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new note",
	Long: `Create a new Bear note from stdin.

Examples:
• http --json get https://example.com/json | bear new
• bear new --title crashdump < error.log
• echo !! | bear new --tags=snippets,oneliners

`,
	Run: func(cmd *cobra.Command, args []string) {
		consoleScanner := bufio.NewScanner(os.Stdin)
		var buffer bytes.Buffer

		for consoleScanner.Scan() {
			buffer.WriteString(consoleScanner.Text())
			buffer.WriteString("\n")
		}

		input := buffer.String()
		fmt.Println(input)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
