// +build server plugin

package main

import (
	"golang.org/x/net/context"

	"github.com/codedellemc/libstorage-mock/csi"
)

type identServer struct{}

func (s *identServer) GetSupportedVersions(
	context.Context,
	*csi.GetSupportedVersionsRequest) (*csi.GetSupportedVersionsResponse, error) {

	return nil, nil
}

func (s *identServer) GetPluginInfo(
	context.Context,
	*csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {

	return nil, nil
}
