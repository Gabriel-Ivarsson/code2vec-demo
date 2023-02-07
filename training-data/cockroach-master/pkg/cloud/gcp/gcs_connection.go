// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package gcp

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/cloud/externalconn"
	"github.com/cockroachdb/cockroach/pkg/cloud/externalconn/connectionpb"
	"github.com/cockroachdb/cockroach/pkg/cloud/externalconn/utils"
	"github.com/cockroachdb/cockroach/pkg/security/username"
	"github.com/cockroachdb/errors"
)

func validateGCSConnectionURI(
	ctx context.Context, execCfg interface{}, user username.SQLUsername, uri string,
) error {
	if err := utils.CheckExternalStorageConnection(ctx, execCfg, user, uri); err != nil {
		return errors.Wrap(err, "failed to create gs external connection")
	}

	return nil
}

func init() {
	externalconn.RegisterConnectionDetailsFromURIFactory(
		gcsScheme,
		connectionpb.ConnectionProvider_gs,
		externalconn.SimpleURIFactory,
	)

	externalconn.RegisterDefaultValidation(gcsScheme, validateGCSConnectionURI)
}
