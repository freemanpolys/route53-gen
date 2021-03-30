/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var Name string
var Ip string
var Wildcard bool

type ResourceRecord struct {
	Value string `json:"Value"`
}

type Change struct {
	Action            string `json:"Action"`
	ResourceRecordSet struct {
		Name            string           `json:"Name"`
		Type            string           `json:"Type"`
		TTL             int              `json:"TTL"`
		ResourceRecords []ResourceRecord `json:"ResourceRecords"`
	} `json:"ResourceRecordSet"`
}
type Route53Record struct {
	Comment string   `json:"Comment"`
	Changes []Change `json:"Changes"`
}

// recordsetsCmd represents the recordsets command
var recordsetsCmd = &cobra.Command{
	Use:   "recordsets",
	Short: "Generate AWS Route53 ResourceRecordSet type A json",
	Long:  `Generate AWS Route53 ResourceRecordSet type A json.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("recordsets called")
		createRecord(args)
	},
}

func init() {
	rootCmd.AddCommand(recordsetsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordsetsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordsetsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	recordsetsCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Record name (required)")
	recordsetsCmd.PersistentFlags().StringVarP(&Ip, "ip", "i", "", "Record IP (required)")
	recordsetsCmd.PersistentFlags().BoolVarP(&Wildcard, "wildcard", "w", false, "Is Record name wirldcard")
	recordsetsCmd.MarkPersistentFlagRequired("name")
	recordsetsCmd.MarkPersistentFlagRequired("ip")

}

func createRecord(args []string) {
	record := Route53Record{}
	record.Comment = fmt.Sprintf("Creating %s entry", Name)
	change := make([]Change, 1)
	change[0].Action = "CREATE"
	change[0].ResourceRecordSet.Type = "A"
	change[0].ResourceRecordSet.TTL = 300
	if Wildcard {
		Name = "*." + Name
	}
	change[0].ResourceRecordSet.Name = Name

	// set ip address
	recordsets := make([]ResourceRecord, 1)
	recordsets[0].Value = Ip
	change[0].ResourceRecordSet.ResourceRecords = recordsets
	record.Changes = change

	if res, err := json.Marshal(record); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(res))
	}

}
