# pycallgo
experiment calling go from python

I want to use AWS lambda, but I like coding in Go

This example shows how to pass, return and free memory of strings when using
a go shared library from python

Compiling shared library:
 go build -o libdoubler.so -buildmode=c-shared .

