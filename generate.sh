#!/bin/bash

#Shell script to compile .proto files

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.