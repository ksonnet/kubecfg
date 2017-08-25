// Copyright 2017 The kubecfg authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(prototypeCmd)
}

var prototypeCmd = &cobra.Command{
	Use:   "prototype <command>",
	Short: `Instantiate, inspect, and get examples for ksonnet prototypes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Invalid number of arguments to command 'prototype'")
		}

		switch args[0] {
		case "use":
			return fmt.Errorf("Command 'prototype use' not implemented yet")
		case "describe":
			return fmt.Errorf("Command 'prototype describe' not implemented yet")
		case "search":
			return fmt.Errorf("Command 'prototype search' not implemented yet")
		default:
			return fmt.Errorf("Unrecognized command 'prototype %s'", args[0])
		}
	},
	Long: `Manage, inspect, instantiate, and get examples for ksonnet prototypes.

Prototypes are Kubernetes app configuration templates with "holes" that can be
filled in by (e.g.) the ksonnet CLI tool or a language server. For example, a
prototype for a 'apps.v1beta1.Deployment' might require a name and image, and
the ksonnet CLI could expand this to a fully-formed 'Deployment' object.

Commands:
  use      Instantiate prototype, filling in parameters from flags, and
           emitting the generated code to stdout.
  describe Display documentation and details about a prototype
  search   Search for a prototype`,

	Example: `  # Display documentation about prototype
  # 'io.ksonnet.pkg.prototype.simple-deployment', including:
  #
  #   (1) a description of what gets generated during instantiation
  #   (2) a list of parameters that are required to be passed in with CLI flags
  #
  # NOTE: Many subcommands only require the user to specify enough of the
  # identifier to disambiguate it among other known prototypes, which is why
  # 'simple-deployment' is given as argument instead of the fully-qualified
  # name.
  ksonnet prototype describe simple-deployment

  # Instantiate prototype 'io.ksonnet.pkg.prototype.simple-deployment', using
  # the 'nginx' image, and port 80 exposed.
  #
  # SEE ALSO: Note above for a description of why this subcommand can take
  # 'simple-deployment' instead of the fully-qualified prototype name.
  ksonnet prototype use simple-deployment \
    --name=nginx                          \
    --image=nginx                         \
    --port=80                             \
    --portName=http

  # Search known prototype metadata for the string 'deployment'.
  #
  # SEE ALSO: Note above for a description of why this subcommand can take
  # 'simple-deployment' instead of the fully-qualified prototype name.
  ksonnet prototype search deployment`,
}
