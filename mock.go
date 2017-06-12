package mock

import (
	apitypes "github.com/codedellemc/libstorage/api/types/v1"
)

// Name is the name of the driver.
const Name = "mock"

// New returns a new driver instance.
func New() interface{} {
	return &driver{}
}

type driver struct {
	ctx    apitypes.Context
	config apitypes.Config
}

// Init initializes the storage driver with context and configuration
// objects.
//
// Returning a non-nill error from this function will mark the driver
// instance as in error and unusable.
func (d *driver) Init(ctxObj, configObj interface{}) error {
	var ok bool

	if d.ctx, ok = ctxObj.(apitypes.Context); !ok {
		return apitypes.ErrInvalidContext
	}

	if d.config, ok = configObj.(apitypes.Config); !ok {
		return apitypes.ErrInvalidConfig
	}

	return nil
}

// Name returns the name of the driver.
func (d *driver) Name() string {
	return Name
}

// Supports returns a mask of the operations supported by the driver.
func (d *driver) Supports() uint64 {
	return apitypes.SdSupported
}
