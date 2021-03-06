// Copyright © 2019 IBM Corporation and others.
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
	"errors"
	"github.com/spf13/cobra"
)

func newTestCmd(rootConfig *RootCommandConfig) *cobra.Command {
	config := &devCommonConfig{RootCommandConfig: rootConfig}
	// testCmd represents the test command
	var testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test your project in the local Appsody environment.",
		Long: `Run the local Appsody environment, starting a container-based, continuous build environment for your project, and running the test suite each time a file changes.
		
Run this command from the root directory of your Appsody project.`,
		Example: `  appsody test
  Runs the tests for your Appsody project.
		
  appsody test --no-watcher
  Runs the tests for your Appsody project without monitoring your project files for changes. The command completes after the tests are run once.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return errors.New("Unexpected argument. Use 'appsody [command] --help' for more information about a command")
			}

			rootConfig.Info.log("Running test environment")
			return commonCmd(config, "test")
		},
	}

	addDevCommonFlags(testCmd, config)
	return testCmd
}
