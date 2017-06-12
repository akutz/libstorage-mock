package main

import "C"

import (
	driver "github.com/codedellemc/libstorage-mock"
)

// APIVersion is the version of the API supported by this mod.
var APIVersion = "v1.0.0"

// Types is the symbol the host process uses to retrieve the mod's type map.
var Types = map[string]func() interface{}{
	driver.Name: driver.New,
}
