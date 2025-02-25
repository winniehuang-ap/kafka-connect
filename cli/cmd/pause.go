// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/winniehuang-ap/kafka-connect/v3/lib/connectors"
)

// pauseCmd represents the pause command
var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause a connector",
	RunE:  RunEPause,
}

// RunEPause ...
func RunEPause(cmd *cobra.Command, args []string) error {
	req := connectors.ConnectorRequest{
		Name: connector,
	}
	resp, err := getClient().PauseConnector(req, sync)
	if err != nil {
		return err
	}
	return printResponse(resp)
}

func init() {
	RootCmd.AddCommand(pauseCmd)

	pauseCmd.PersistentFlags().BoolVarP(&sync, "sync", "y", false, "execute synchronously")
	pauseCmd.PersistentFlags().StringVarP(&connector, "connector", "n", "", "name of the target connector")
}
