/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
	"github.com/spf13/cobra"
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:   "kbot",
	Aliases: []string{"start"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s\n", appVersion)
		pref := tele.Settings{
			URL: "",
			Token:  os.Getenv("TOKEN"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}
	
		b, err := tele.NewBot(pref)
		if err != nil {
			log.Fatal(err)
			return
		}
	

		b.Handle(tele.OnText, func(m tele.Context) error {
			log.Print(m.Message().Payload, m.Text())
			log.Printf("hello, i'm kbot %s!\n", appVersion)
			payload:=m.Text()

			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("hello, i'm kbot %s!\n", appVersion))
			}
			return err
		})


		b.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
