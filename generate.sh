#:/bin/bash

protoc great/greatpb/great.proto --go_out=plugins=grpc:.
