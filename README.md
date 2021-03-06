# libstorage-mock
A mock storage driver for libStorage.

## Getting Started
This section demonstrates how to build and use the CSI client and server:

```bash
# build the client and server programs
$ make
go build -o mock-server -tags server
go build -o mock-client -tags client

# export the csi endpoint used by the client and server
export CSI_ENDPOINT=tcp://127.0.0.1:9000

# launch the csi server in the background
$ ./mock-server &
[1] 75208

# query the csi server using the client
$ ./mock-client
VolumeID=vol-001
VolumeID=vol-002
VolumeID=vol-003

# kill the csi server
$ kill $(ps aux | grep "[m]ock-server" | awk '{print $2}')
[1]+  Terminated: 15          ./mock-server
```

## Build Reference
This section describes how to build the project.

### Generated Sources
While the CSI protobuf file and source are versioned with the project
at `csi/csi.proto` and `csi.pb.go`, it is still possible to generate
them from a remote CSI specification file.

First, remove the generated files:

```bash
$ make clobber
```

Next, make the target `csi/csi.pb.go`:

```bash
$ make csi/csi.pb.go
```

It's also possible to influence the location from which the CSI specification
is retrieved with the following environment variables:

| Name | Description | Default |
|------|-------------|---------|
| `CSI_SPEC_FILE` | The path to a local spec file used to generate the protobuf | |
| `CSI_GIT_OWNER` | The GitHub user or organization that owns the git repository that contains the CSI spec file | `container-storage-interface` |
| `CSI_GIT_REPO` | The GitHub repository that contains the CSI spec file | `spec` |
| `CSI_GIT_REF` | The git ref to use when getting the CSI spec file. This value can be a branch name, a tag, or a git commit ID | `master` |
| `CSI_SPEC_NAME` | The name of the CSI spec markdown file | `spec.md` |
| `CSI_SPEC_PATH` | The remote path of the CSI markdown file | |
| `CSI_PROTO_NAME` | The name of the protobuf file to generate. This value should not include the file extension | `csi` |
| `CSI_PROTO_DIR` | The path of the directory in which the protobuf and Go source files will be generated. If this directory does not exist it will be created. | `.` |
| `CSI_PROTO_ADD` | A list of additional protobuf files used when building the Go source file | |
| `CSI_IMPORT_PATH` | The package of the generated Go source | `csi` |

### CSI Server
The CSI server binary may be built with the following command:

```bash
$ make server
```

The above command builds the `mock-server` binary (`mock-server.exe`
on Windows).

### CSI Client
The CSI client binary may be built with the following command:

```bash
$ make client
```

The above command builds the `mock-client` binary (`mock-client.exe`
on Windows).

### libStorage Plug-in
The libStorage plug-in binary may be built on a Linux host with the
following command:

```bash
$ make plugin
```

The above command builds the `mock-plugin.so` binary. Go plug-ins may
be built only on Linux hosts with `GOOS=linux`.
