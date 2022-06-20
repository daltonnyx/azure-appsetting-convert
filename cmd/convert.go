/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"muzzammil.xyz/jsonc"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A tool for converting appsettings.json to Azure App Configuration",
	Long:  `A tool for converting appsettings.json to Azure App Configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("convert called")
		inputFile := cmd.Flag("input").Value.String()
		outputFile := cmd.Flag("output").Value.String()
		inputJson, err := os.Open(inputFile)
		if err != nil {
			fmt.Println(err)
		}
		defer inputJson.Close()

		byteValues, _ := ioutil.ReadAll(inputJson)
		var result map[string]any
		err = json.Unmarshal(jsonc.ToJSON(byteValues), &result)
		if err != nil {
			fmt.Println(err)
		}

		settings := flattenJson(result, "")
		outputByte, err := json.Marshal(settings)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(outputFile, outputByte, 0644)
		fmt.Println("convert done")
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	convertCmd.Flags().StringP("input", "i", "appsettings.json", "path to appsettings.json file")
	convertCmd.Flags().StringP("output", "o", "output.json", "path to output file")
}

func flattenJson(jsonObject map[string]any, prefix string) []Setting {
	var results []Setting
	for key, value := range jsonObject {
		switch child := value.(type) {
		case map[string]any:
			parseSetting(child, prefix+key+":", &results)
		default:
			results = append(results, Setting{Name: prefix + key, Value: value, SlotSetting: true})
		}
	}
	return results
}

func parseSetting(jsonObject map[string]any, prefix string, results *[]Setting) {
	for key, value := range jsonObject {
		switch child := value.(type) {
		case map[string]any:
			parseSetting(child, prefix+key+":", results)
		default:
			*results = append(*results, Setting{Name: prefix + key, Value: value, SlotSetting: true})
		}
	}
}
