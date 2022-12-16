/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	limit       uint64
	outFileName string
	columns     []string

	columnsSupported []string = []string{"name", "phone", "age"}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "us-fake-numbers-gen",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Random Num :  ")
		if len(columns) == 0 {
			columns = columnsSupported
		}
		println("cols ", len(columns))
		data := GenerateData(columns, limit)
		WriteCsv(outFileName, limit, data...)
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.us-fake-numbers-gen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().Uint64VarP(&limit, "limit", "l", 10, "quantity of fake phone numbers. Specify Any number from 0 to 18446744073709551615")
	if err := rootCmd.MarkFlagRequired("limit"); err != nil {
		fmt.Println(err)
	}

	rootCmd.Flags().StringVarP(&outFileName, "out", "o", "fake_nums", "give file name without extension to output data. Give path where file should be created.")
	if err := rootCmd.MarkFlagRequired("out"); err != nil {
		fmt.Println(err)
	}

	rootCmd.Flags().StringArrayVar(&columns, "col", columnsSupported, "give columns to be added in output file. Supported columns are  name,phone,age")

}
