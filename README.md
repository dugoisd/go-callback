# Golang Callback Registration from C and execution from Go

## Description

    this is a simple example used for my personnal GO exploration
    this code use a Go software as library ( buildmode=c-shared ).
    User Apps (C Code) register a callback function with context.
    Go Library will call this callback in a go routine.

## compile and test

`go build -buildmode=c-shared -o libtest main.go`

`cd test`

`gcc -o test main.c -Wl,--whole-archive -L. -l:../libtest -lpthread -Wl,-no-whole-archive`