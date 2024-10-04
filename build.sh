#! /bin/bash

cd cppsrc && g++ -shared -o libmeme.so -fPIC *.cpp && cp libmeme.so ..
cd -
go build src/main.go
