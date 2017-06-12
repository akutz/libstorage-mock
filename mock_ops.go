package mock

import (
	"strings"

	apitypes "github.com/codedellemc/libstorage/api/types/v1"
)

// InstanceID is the value of the ID field for the Mock InstanceID.
const InstanceID = "mock.InstanceID"

func (d *driver) instanceID(
	ctx apitypes.Context) (apitypes.InstanceID, error) {

	return result{
		rID:      InstanceID,
		rDriver:  Name,
		rService: Name,
	}, nil
}

func (d *driver) nextDeviceInfo(
	ctx apitypes.Context) (apitypes.NextDeviceInfo, error) {

	return nil, nil
}

func (d *driver) getType(
	ctx apitypes.Context) (uint8, error) {

	return 0, nil
}

func (d *driver) instanceInspect(
	ctx apitypes.Context,
	opts apitypes.Store) (apitypes.Instance, error) {

	return nil, nil
}

const tenGiBInBytes = 10737418240

var vols = []apitypes.Volume{
	&result{
		rID:   "mock-vol-001",
		rName: "Mock Volume 001",
		rSize: tenGiBInBytes,
	},
	&result{
		rID:   "mock-vol-002",
		rName: "Mock Volume 002",
		rSize: tenGiBInBytes,
	},
	&result{
		rID:   "mock-vol-003",
		rName: "Mock Volume 003",
		rSize: tenGiBInBytes,
	},
}

func (d *driver) volumes(
	ctx apitypes.Context,
	opts apitypes.VolumesOpts) ([]apitypes.Volume, error) {

	return vols, nil
}

func (d *driver) volumeInspect(
	ctx apitypes.Context,
	volumeID string,
	opts apitypes.VolumeInspectOpts) (apitypes.Volume, error) {

	for _, v := range vols {
		if strings.EqualFold(v.GetID(), volumeID) {
			return v, nil
		}
	}

	return nil, nil
}

func (d *driver) volumeInspectByName(
	ctx apitypes.Context,
	volumeName string,
	opts apitypes.VolumeInspectOpts) (apitypes.Volume, error) {

	for _, v := range vols {
		if strings.EqualFold(v.GetName(), volumeName) {
			return v, nil
		}
	}

	return nil, nil
}

func (d *driver) volumeCreate(
	ctx apitypes.Context,
	name string,
	opts apitypes.VolumeCreateOpts) (apitypes.Volume, error) {

	return nil, nil
}
func (d *driver) volumeCreateFromSnapshot(
	ctx apitypes.Context,
	snapshotID,
	volumeName string,
	opts apitypes.VolumeCreateOpts) (apitypes.Volume, error) {

	return nil, nil
}

func (d *driver) volumeCopy(
	ctx apitypes.Context,
	volumeID,
	volumeName string,
	opts apitypes.Store) (apitypes.Volume, error) {

	return nil, nil
}

func (d *driver) volumeSnapshot(
	ctx apitypes.Context,
	volumeID,
	snapshotName string,
	opts apitypes.Store) (apitypes.Snapshot, error) {

	return nil, nil
}

func (d *driver) volumeRemove(
	ctx apitypes.Context,
	volumeID string,
	opts apitypes.VolumeRemoveOpts) error {

	return nil
}

func (d *driver) volumeAttach(
	ctx apitypes.Context,
	volumeID string,
	opts apitypes.VolumeAttachOpts) (apitypes.Volume, string, error) {

	return nil, "", nil
}

func (d *driver) volumeDetach(
	ctx apitypes.Context,
	volumeID string,
	opts apitypes.VolumeDetachOpts) (apitypes.Volume, error) {

	return nil, nil
}

func (d *driver) snapshots(
	ctx apitypes.Context,
	opts apitypes.Store) ([]apitypes.Snapshot, error) {

	return nil, nil
}

func (d *driver) snapshotInspect(
	ctx apitypes.Context,
	snapshotID string,
	opts apitypes.Store) (apitypes.Snapshot, error) {

	return nil, nil
}

func (d *driver) snapshotCopy(
	ctx apitypes.Context,
	snapshotID,
	snapshotName,
	destinationID string,
	opts apitypes.Store) (apitypes.Snapshot, error) {

	return nil, nil
}

func (d *driver) snapshotRemove(
	ctx apitypes.Context,
	snapshotID string,
	opts apitypes.Store) error {

	return nil
}
