/*
Copyright 2016 - 2017 Huawei Technologies Co., Ltd. All rights reserved.

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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Huawei/containerops/common/utils"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy ContainerOps, Kubernetes, etcd, flannel and others.",
	Long: `Deploy ContainerOps modules, Kubernetes masters and nodes, etcd cluster, flannel plugin,
CoreDNS, Prometheus and others in Cloud Native stack.`,
}

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "deploy stack with a template file",
	Long:  `Deploy ContainerOps modules or others with deploy template file, the template.`,
	Run:   templateRun,
}

// init()
func init() {
	RootCmd.AddCommand(deployCmd)

	// Add sub command.
	deployCmd.AddCommand(templateCmd)
}

// Deploy the Cloud Native stack with a template file.
func templateRun(cmd *cobra.Command, args []string) {
	if len(args) <= 0 || utils.IsFileExist(args[0]) == false {
		fmt.Println("The template file is required.")
		os.Exit(1)
	}

}
