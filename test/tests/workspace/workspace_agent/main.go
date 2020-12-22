// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package main

import (
	"io/ioutil"

	"github.com/gitpod-io/gitpod/test/pkg/integration"
	"github.com/gitpod-io/gitpod/test/tests/workspace/workspace_agent/api"
)

func main() {
	integration.ServeAgent(new(WorkspaceAgent))
}

// WorkspaceAgent provides ingteration test services from within a workspace
type WorkspaceAgent struct {
}

// ListDir lists a directory's content
func (*WorkspaceAgent) ListDir(req *api.ListDirRequest, resp *api.ListDirResponse) error {
	dc, err := ioutil.ReadDir(req.Dir)
	if err != nil {
		return err
	}

	*resp = api.ListDirResponse{}
	for _, c := range dc {
		resp.Files = append(resp.Files, c.Name())
	}
	return nil
}
