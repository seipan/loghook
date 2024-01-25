package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "e2e test",
	Long:  `e2e test of loghook`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		discordURL, err := cmd.Flags().GetString("discord")
		if err != nil {
			log.Println(err)
			return err
		}
		slackURL, err := cmd.Flags().GetString("slack")
		if err != nil {
			log.Println(err)
			return err
		}
		err = eToeTest(discordURL, slackURL)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}

func init() {
	rootCmd.Flags().StringP("discord", "d", "", "discord webhook url")
	rootCmd.Flags().StringP("slack", "s", "", "slack webhook url")
}
