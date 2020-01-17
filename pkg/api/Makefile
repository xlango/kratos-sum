#
# Syntax example: make GOPATH=e:/Go/GOPATH
#
# Choose GOPATH
GOPATH ?= $(shell go env GOPATH | tr "\\" "/" | tr ";" "\n")
# Choose protoc binary
ROOT ?= $(shell pwd | sed 's/^\/drives\/\([a-z]\)/\1:/')/../..
GOGO_PROTOBUF:= pkg/mod/github.com/gogo/protobuf@v1.3.1

FLAGS+= -I $(ROOT)/third_party/googleapis
FLAGS+= -I. $(patsubst %,-I$(GOPATH)/%,$(GOGO_PROTOBUF))
FLAGS+= --gogofast_out=plugins=grpc,
FLAGS:= $(FLAGS)Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,
#FLAGS:= $(FLAGS)Mgoogle/type/date.proto=github.com/gogo/googleapis/google/type,
#FLAGS:= $(FLAGS)Mgoogle/type/timeofday.proto=github.com/gogo/googleapis/google/type,
#FLAGS:= $(FLAGS)Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,
FLAGS:= $(FLAGS)paths=source_relative:.
FLAGS+= --grpc-gateway_out=logtostderr=true:. 
FLAGS+= --swagger_out=logtostderr=true:.

SUFFIX:= pb.go
GW_SUFFIX:= pb.gw.go
SWAGGER_SUFFIX:= swagger.json

DEPS:= $(shell find . -type f -name '*.proto' | sed "s/proto$$/$(SUFFIX)/")
OUTS:=$(DEPS)
OUTS+= $(patsubst %.$(SUFFIX),%.$(GW_SUFFIX),$(DEPS))
OUTS+= $(patsubst %.$(SUFFIX),%.$(SWAGGER_SUFFIX),$(DEPS))

all: $(DEPS)

%.$(SUFFIX) :  %.proto
	protoc $(FLAGS) $*.proto

clean:
	rm -f $(OUTS) 2> /dev/null