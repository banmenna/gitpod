// Copyright (c) 2024 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

var jetBrainsProjectBuildCmd = &cobra.Command{
	Use:   "project-build",
	Short: "Interact with Gitpod's OIDC identity provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		return callJetBrainsBackendCLI(cmd.Context(), operatorProjectBuild, nil)
	},
}

func init() {
	jetBrainsCmd.AddCommand(jetBrainsProjectBuildCmd)
}