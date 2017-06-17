// +build server plugin

package main

import (
	"golang.org/x/net/context"

	"github.com/codedellemc/libstorage-mock/csi"
)

type libsServer struct{}

func (s *libsServer) InitPlugin(
	context.Context,
	*csi.InitPluginRequest) (*csi.InitPluginResponse, error) {

	return nil, nil
}
