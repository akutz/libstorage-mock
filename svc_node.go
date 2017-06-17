// +build server plugin

package main

import (
	"golang.org/x/net/context"

	"github.com/codedellemc/libstorage-mock/csi"
)

type nodeServer struct{}

func (s *nodeServer) GetSupportedVersions(
	context.Context,
	*csi.GetSupportedVersionsRequest) (*csi.GetSupportedVersionsResponse, error) {

	return nil, nil
}

func (s *nodeServer) GetPluginInfo(
	context.Context,
	*csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {

	return nil, nil
}

func (s *nodeServer) NodePublishVolume(
	context.Context,
	*csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {

	return nil, nil
}

func (s *nodeServer) NodeUnpublishVolume(
	context.Context,
	*csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {

	return nil, nil
}

func (s *nodeServer) GetNodeID(
	context.Context,
	*csi.GetNodeIDRequest) (*csi.GetNodeIDResponse, error) {

	return nil, nil
}

func (s *nodeServer) ProbeNode(
	context.Context,
	*csi.ProbeNodeRequest) (*csi.ProbeNodeResponse, error) {

	return nil, nil
}

func (s *nodeServer) NodeGetCapabilities(
	context.Context,
	*csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {

	return nil, nil
}
