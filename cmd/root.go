/*
Copyright © 2022 Mohammad Ali Ashraf ranaaliashraf12@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	limit       uint64
	outFileName string
	columns     []string
	unique      bool

	columnsSupported []string = []string{"name", "age", "phone"}
)

var rootCmd = &cobra.Command{
	Use:   "fake-gen",
	Short: "",
	Long: `Description:
	Single command cli tool for generating US fake numbers. To generate 10,000 fake numbers use following command:
			
			fake-gen --limit 10000 --unique --out huge_data --column name --column name --column age --column phone --column phone --column phone
									
								OR

			fake-gen -l 10000 -u -o huge_data -c name -c name -c age -c phone -c phone -c phone

			The output csv file will have 6 columns. We will mark them as firstName,LastName,Age,Mobile,Home,Work
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Random Num :  ")
		if len(columns) == 0 {
			columns = columnsSupported
		}
		data := GenerateData(columns, limit, unique)
		WriteCsv(outFileName, limit, data...)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Uint64VarP(&limit, "limit", "l", 10, "Rows of fake data to generate.")
	rootCmd.Flags().StringVarP(&outFileName, "out", "o", "fake_nums", "give file name without extension to output data. Give path where file should be created.")
	rootCmd.Flags().BoolVarP(&unique, "unique", "u", false, "set unique flag if you want to generate only unique phone numbers")
	rootCmd.Flags().StringArrayVarP(&columns, "column", "c", columnsSupported, "give columns in sequence to be added in output csv file.")
}
