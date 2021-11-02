#!/bin/bash


protoc --inkerpc_out=plugins=rpc:. echo.proto
