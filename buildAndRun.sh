#!/bin/bash

GOOS=js GOARCH=wasm go build -o sample/main.wasm && goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("sample/")))'