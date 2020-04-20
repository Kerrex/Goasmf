#!/bin/bash

GOOS=js GOARCH=wasm go build -o sample/static/main.wasm && ./sample/server