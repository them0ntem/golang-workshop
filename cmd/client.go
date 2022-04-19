/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverUrl, _ := cmd.Flags().GetString("server_url")
		text, _ := cmd.Flags().GetString("text")
		runClient(serverUrl, text)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	clientCmd.Flags().String("server_url", "http://localhost:8000", "Server url to connect to")
	clientCmd.Flags().String("text", "", "Text to reverse")
}

func runClient(serverUrl string, text string) {
	postBody, _ := json.Marshal(map[string]string{
		"text": text,
	})
	responseBody := bytes.NewBuffer(postBody)
	out, err := http.Post(serverUrl+"/reverse", "application/json", responseBody)
	if err != nil {
		log.Fatalln(err)
	}
	bytesArr, _ := io.ReadAll(out.Body)
	fmt.Println(string(bytesArr))
}
