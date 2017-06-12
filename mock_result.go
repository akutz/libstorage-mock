package mock

import (
	apitypes "github.com/codedellemc/libstorage/api/types/v1"
)

type result map[string]interface{}

const rID = "id"

func (r result) GetID() string {
	if v, ok := r[rID].(string); ok {
		return v
	}
	return ""
}

const rDriver = "driver"

func (r result) GetDriver() string {
	if v, ok := r[rDriver].(string); ok {
		return v
	}
	return ""
}

const rService = "service"

func (r result) GetService() string {
	if v, ok := r[rService].(string); ok {
		return v
	}
	return ""
}

const rFields = "fields"

func (r result) GetFields() map[string]string {
	if v, ok := r[rFields].(map[string]string); ok {
		return v
	}
	return nil
}

const rAttachments = "attachments"

func (r result) GetAttachments() []interface{} {
	if v, ok := r[rAttachments].([]interface{}); ok {
		return v
	}
	return nil
}

const rAttachmentState = "attachmentState"

func (r result) GetAttachmentState() uint64 {
	if v, ok := r[rAttachmentState].(uint64); ok {
		return v
	}
	return 0
}

const rAvailabilityZone = "availabilityZone"

func (r result) GetAvailabilityZone() string {
	if v, ok := r[rAvailabilityZone].(string); ok {
		return v
	}
	return ""
}

const rIsEncrypted = "isEncrypted"

func (r result) IsEncrypted() bool {
	if v, ok := r[rIsEncrypted].(bool); ok {
		return v
	}
	return false
}

const rIOPS = "iops"

func (r result) GetIOPS() int64 {
	if v, ok := r[rIOPS].(int64); ok {
		return v
	}
	return 0
}

const rName = "name"

func (r result) GetName() string {
	if v, ok := r[rName].(string); ok {
		return v
	}
	return ""
}

const rNetworkName = "networkName"

func (r result) GetNetworkName() string {
	if v, ok := r[rNetworkName].(string); ok {
		return v
	}
	return ""
}

const rSize = "size"

func (r result) GetSize() int64 {
	if v, ok := r[rSize].(int64); ok {
		return v
	}
	return 0
}

const rStatus = "status"

func (r result) GetStatus() string {
	if v, ok := r[rStatus].(string); ok {
		return v
	}
	return ""
}

const rType = "type"

func (r result) GetType() uint8 {
	if v, ok := r[rType].(uint8); ok {
		return v
	}
	return 0
}

const rDeviceName = "deviceName"

func (r result) GetDeviceName() string {
	if v, ok := r[rDeviceName].(string); ok {
		return v
	}
	return ""
}

const rMountPoint = "mountPoint"

func (r result) GetMountPoint() string {
	if v, ok := r[rMountPoint].(string); ok {
		return v
	}
	return ""
}

const rInstanceID = "instanceID"

func (r result) GetInstanceID() interface{} {
	if v, ok := r[rInstanceID].(apitypes.InstanceID); ok {
		return v
	}
	return nil
}

const rVolumeID = "volumeID"

func (r result) GetVolumeID() string {
	if v, ok := r[rVolumeID].(string); ok {
		return v
	}
	return ""
}

const rDescription = "description"

func (r result) GetDescription() string {
	if v, ok := r[rDescription].(string); ok {
		return v
	}
	return ""
}

const rStartTime = "startTime"

func (r result) GetStartTime() int64 {
	if v, ok := r[rStartTime].(int64); ok {
		return v
	}
	return 0
}

const rVolumeSize = "volumeSize"

func (r result) GetVolumeSize() int64 {
	if v, ok := r[rVolumeSize].(int64); ok {
		return v
	}
	return 0
}

const rMountID = "mountID"

func (r result) GetMountID() int {
	if v, ok := r[rMountID].(int); ok {
		return v
	}
	return 0
}

const rParent = "parent"

func (r result) GetParent() int {
	if v, ok := r[rParent].(int); ok {
		return v
	}
	return 0
}

const rMajor = "major"

func (r result) GetMajor() int {
	if v, ok := r[rMajor].(int); ok {
		return v
	}
	return 0
}

const rMinor = "minor"

func (r result) GetMinor() int {
	if v, ok := r[rMinor].(int); ok {
		return v
	}
	return 0
}

const rRoot = "root"

func (r result) GetRoot() string {
	if v, ok := r[rRoot].(string); ok {
		return v
	}
	return ""
}

const rOpts = "opts"

func (r result) GetOpts() string {
	if v, ok := r[rOpts].(string); ok {
		return v
	}
	return ""
}

const rOptional = "optional"

func (r result) GetOptional() string {
	if v, ok := r[rOptional].(string); ok {
		return v
	}
	return ""
}

const rFSType = "fsType"

func (r result) GetFSType() string {
	if v, ok := r[rFSType].(string); ok {
		return v
	}
	return ""
}

const rSource = "source"

func (r result) GetSource() string {
	if v, ok := r[rSource].(string); ok {
		return v
	}
	return ""
}

const rVFSOpts = "vfsOpts"

func (r result) GetVFSOpts() string {
	if v, ok := r[rVFSOpts].(string); ok {
		return v
	}
	return ""
}

const rIgnored = "ignored"

func (r result) IsIgnored() bool {
	if v, ok := r[rIgnored].(bool); ok {
		return v
	}
	return false
}

const rPrefix = "prefix"

func (r result) GetPrefix() string {
	if v, ok := r[rPrefix].(string); ok {
		return v
	}
	return ""
}

const rPattern = "pattern"

func (r result) GetPattern() string {
	if v, ok := r[rPattern].(string); ok {
		return v
	}
	return ""
}
