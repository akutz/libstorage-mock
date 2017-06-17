# the name of the program being built
PROG := mock

# the name server binary to build
SERVER_BIN := $(PROG)-server

# the name client binary to build
CLIENT_BIN := $(PROG)-client

# adjust the server & client binary names for windows
ifeq (windows,$(GOOS))
SERVER_BIN := $(SERVER_BIN).exe
CLIENT_BIN := $(CLIENT_BIN).exe
endif

# the name of the plug-in binary to build
PLUGIN_BIN := $(PROG)-plugin.so

all: $(SERVER_BIN) $(CLIENT_BIN)
# add the plug-in target to all if building on linux for linux
ifeq (linux_linux,$(GOOS)_$(GOHOSTOS))
all: $(PLUGIN_BIN)
endif

# the libStorage protobuf files
LIBS_PROTO := csi/libs.proto
LIBS_GOSRC := csi/libs.pb.go

# configure and load the csi protobuf generator
CSI_PROTO_ADD += $(LIBS_PROTO)
CSI_PROTO_DIR := csi
include csi.mk

# get information about the host
ifeq (,$(strip $(GOHOSTOS)))
GOHOSTOS := $(shell go env | grep GOHOSTOS | sed 's/GOHOSTOS="\(.*\)"/\1/')
endif
ifeq (,$(strip $(GOOS)))
GOOS := $(shell go env | grep GOOS | sed 's/GOOS="\(.*\)"/\1/')
endif

SERVICES_SRCS := controller.go identity.go libstorage.go node.go

$(SERVER_BIN): $(CSI_GOSRC) server.go svc_*.go
	go build -o $@ -tags server

$(CLIENT_BIN): $(CSI_GOSRC) client.go
	go build -o $@ -tags client

$(PLUGIN_BIN): $(CSI_GOSRC) plugin.go svc_*.go
	go build -o $@ -buildmode plugin -tags plugin

$(LIBS_GOSRC): $(CSI_GOSRC)

server: $(SERVER_BIN)
client: $(CLIENT_BIN)
plugin: $(PLUGIN_BIN)

clean:
	rm -f $(SERVER_BIN) $(CLIENT_BIN) $(PLUGIN_BIN)

clobber: clean
	rm -f $(CSI_PROTO) $(LIBS_GOSRC) $(CSI_GOSRC)

.PHONY: clean clobber
