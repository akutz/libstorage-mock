package mock

import (
	apitypes "github.com/codedellemc/libstorage/api/types/v1"
)

// Do does the operation specified by the opID.
func (d *driver) Do(
	ctxObj interface{},
	opID uint64,
	args ...interface{}) (result interface{}, err error) {

	ctx, ok := ctxObj.(apitypes.Context)
	if !ok {
		return nil, apitypes.ErrInvalidContext
	}

	// do instance ops
	if opID >= apitypes.SdInstOpMin && opID <= apitypes.SdInstOpMax {
		return d.doInstanceOps(ctx, opID, args...)
	}

	// do volume ops
	if opID >= apitypes.SdVolOpIDMin && opID <= apitypes.SdVolOpIDMax {
		return d.doVolumeOps(ctx, opID, args...)
	}

	// do snapshot ops
	if opID >= apitypes.SdSnapOpIDMin && opID <= apitypes.SdSnapOpIDMax {
		return d.doSnapshotOps(ctx, opID, args...)
	}

	// do the remaining operations
	return d.doUncatOps(ctx, opID, args...)
}

func (d *driver) doUncatOps(
	ctx apitypes.Context,
	opID uint64,
	args ...interface{}) (result interface{}, err error) {

	switch opID {
	// TODO
	}

	return nil, apitypes.ErrInvalidOp
}

func (d *driver) doInstanceOps(
	ctx apitypes.Context,
	opID uint64,
	args ...interface{}) (result interface{}, err error) {

	switch opID {

	//     func(ctx) (InstanceID, error)
	case apitypes.SdInstanceID:
		return d.instanceID(ctx)
	}

	return nil, apitypes.ErrInvalidOp
}

func (d *driver) doVolumeOps(
	ctx apitypes.Context,
	opID uint64,
	args ...interface{}) (result interface{}, err error) {

	switch opID {

	//     func(ctx, VolumesOpts) ([]Volume, error)
	case apitypes.SdVolumes:
		if len(args) != 1 {
			return nil, apitypes.ErrInvalidArgsLen
		}
		opts, ok := args[0].(apitypes.VolumesOpts)
		if !ok {
			return nil, apitypes.ErrInvalidArgs
		}
		r, err := d.volumes(ctx, opts)
		if err != nil {
			return nil, err
		}

		// cast elements of the array into
		// interface{} types so they may
		// be passed through the Go plug-in
		// framework without type errors
		i := make([]interface{}, len(r))
		for x, v := range r {
			i[x] = v
		}

		return i, nil

	//     func(
	//         ctx,
	//         volumeID string,
	//         VolumeInspectOpts) (Volume, error)
	case apitypes.SdVolumeInspect:
		if len(args) != 2 {
			return nil, apitypes.ErrInvalidArgsLen
		}
		var (
			ok       bool
			volumeID string
			opts     apitypes.VolumeInspectOpts
		)
		if volumeID, ok = args[0].(string); !ok {
			return nil, apitypes.ErrInvalidArgs
		}
		if opts, ok = args[1].(apitypes.VolumeInspectOpts); !ok {
			return nil, apitypes.ErrInvalidArgs
		}
		return d.volumeInspect(ctx, volumeID, opts)

	//     func(
	//         ctx,
	//         volumeName string,
	//         VolumeInspectOpts) (Volume, error)
	case apitypes.SdVolumeInspectByName:
		if len(args) != 2 {
			return nil, apitypes.ErrInvalidArgsLen
		}
		var (
			ok         bool
			volumeName string
			opts       apitypes.VolumeInspectOpts
		)
		if volumeName, ok = args[0].(string); !ok {
			return nil, apitypes.ErrInvalidArgs
		}
		if opts, ok = args[1].(apitypes.VolumeInspectOpts); !ok {
			return nil, apitypes.ErrInvalidArgs
		}
		return d.volumeInspectByName(ctx, volumeName, opts)
	}

	return nil, apitypes.ErrInvalidOp
}

func (d *driver) doSnapshotOps(
	ctx apitypes.Context,
	opID uint64,
	args ...interface{}) (result interface{}, err error) {

	switch opID {

	//     func(ctx, Store) ([]Snapshot, error)
	case apitypes.SdSnapshots:

		if len(args) != 1 {
			return nil, apitypes.ErrInvalidArgsLen
		}
		store, ok := args[0].(apitypes.Store)
		if !ok {
			return nil, apitypes.ErrInvalidArgs
		}
		r, err := d.snapshots(ctx, store)
		if err != nil {
			return nil, err
		}

		// cast elements of the array into
		// interface{} types so they may
		// be passed through the Go plug-in
		// framework without type errors
		i := make([]interface{}, len(r))
		for x, v := range r {
			i[x] = v
		}

		return i, nil
	}

	return nil, apitypes.ErrInvalidOp
}
