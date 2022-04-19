/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"golang_workshop/facade"
	"net/http"
)

type ReverseAPIInput struct {
	Text string
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		runServer(port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().String("port", "8000", "Port to run the server on")
}

func runServer(port string) {
	fmt.Printf("runServer at port %v", port)
	router := gin.Default()
	router.POST("/reverse", func(c *gin.Context) {
		var apiInput ReverseAPIInput
		if err := c.BindJSON(&apiInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"reverse": "",
				"message": "Invalid input",
			})
		}
		rev := facade.Reverse(apiInput.Text)
		c.JSON(http.StatusOK, gin.H{
			"reverse": rev,
			"message": "",
		})
	})
	err := router.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		return
	}
}
