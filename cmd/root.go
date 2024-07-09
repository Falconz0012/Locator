/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var address string 

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goLocator",
	Short: "CLI Tool made to convert Addresses into geo coordenates or vice-versa",
	Long: `Welcome to goLocator a CLI tool made to convert addresses into geo coordinates or vice Versa
  This tool implements Google's geocoding API to transcribe the data and will be used as a root for further projects `,

	 Run: convertAddress,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goLocator.yaml)")
  rootCmd.Flags().StringVarP(&address, "address","a","","Use -a [STREET] [NUMBER] [STATE] [COUNTRY] ")
  rootCmd.MarkFlagRequired("addresss")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
