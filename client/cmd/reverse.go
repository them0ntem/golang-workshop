/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

// reverseCmd represents the reverse command
var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server_url, _ := cmd.Flags().GetString("server_url")
		text, _ := cmd.Flags().GetString("text")
		postBody, _ := json.Marshal(map[string]string{
			"text": text,
		})
		responseBody := bytes.NewBuffer(postBody)
		out, err := http.Post(server_url+"/reverse", "application/json", responseBody)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	reverseCmd.PersistentFlags().String("server_url", "http://localhost:8080", "server url to connect to")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	reverseCmd.Flags().String("text", "t", "text to reverse")
}
