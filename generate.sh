#:/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc calculator/calculatepb/calculate.proto --go_out=plugins=grpc:.
