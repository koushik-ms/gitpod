// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package api

// ListDirRequest is the argument for Exec
type ListDirRequest struct {
	Dir string
}

// ListDirResponse is the response for Exec
type ListDirResponse struct {
	Files []string
}
