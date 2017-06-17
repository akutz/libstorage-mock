// +build server plugin

package main

import (
	"golang.org/x/net/context"

	"github.com/codedellemc/libstorage-mock/csi"
)

type ctrlrServer struct{}

func (s *ctrlrServer) CreateVolume(
	context.Context,
	*csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {

	return nil, nil
}

func (s *ctrlrServer) DeleteVolume(
	context.Context,
	*csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {

	return nil, nil
}

func (s *ctrlrServer) ControllerPublishVolume(
	context.Context,
	*csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {

	return nil, nil
}

func (s *ctrlrServer) ControllerUnpublishVolume(
	context.Context,
	*csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {

	return nil, nil
}

func (s *ctrlrServer) ValidateVolumeCapabilities(
	context.Context,
	*csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {

	return nil, nil
}

var volInfos = []*csi.VolumeInfo{
	&csi.VolumeInfo{
		Id: &csi.VolumeID{
			Values: map[string]string{"id": "vol-001"},
		},
	},
	&csi.VolumeInfo{
		Id: &csi.VolumeID{
			Values: map[string]string{"id": "vol-002"},
		},
	},
	&csi.VolumeInfo{
		Id: &csi.VolumeID{
			Values: map[string]string{"id": "vol-003"},
		},
	},
}

func (s *ctrlrServer) ListVolumes(
	context.Context,
	*csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {

	return &csi.ListVolumesResponse{
		Reply: &csi.ListVolumesResponse_Result_{
			Result: &csi.ListVolumesResponse_Result{
				Entries: []*csi.ListVolumesResponse_Result_Entry{
					&csi.ListVolumesResponse_Result_Entry{VolumeInfo: volInfos[0]},
					&csi.ListVolumesResponse_Result_Entry{VolumeInfo: volInfos[1]},
					&csi.ListVolumesResponse_Result_Entry{VolumeInfo: volInfos[2]},
				},
			},
		},
	}, nil
}

func (s *ctrlrServer) GetCapacity(
	context.Context,
	*csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {

	return nil, nil
}

func (s *ctrlrServer) ControllerGetCapabilities(
	context.Context,
	*csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {

	return nil, nil
}
