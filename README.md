# Antlr Go Tut

## Motivation
Doing the Advent of Code a few years back, one thing that frustrated me was that it took me a while to manhandle their input text into a data structure I could actually use to do the problem.  Here's a simple example: https://adventofcode.com/2017/day/12 and my javascript code to accomplish the parsing (https://github.com/sesquipedalian-dev/Advent2017/blob/master/js/containers/day12.js#L16).  I wondered if it would be easier / faster to use a grammar generator like ANTLR.  There is a cool environment of tools to debug the grammar itself that would make it easy I think to write and debug the grammar, then you can translate from your grammar's tokens / nodes to the data structure you want for the problem by walking the tree or directly handling the AST. I grabbed the vscode extension for antlr (https://marketplace.visualstudio.com/items?itemName=mike-lischke.vscode-antlr4) but there are a number of tools.

## Running it
### dependencies
- Go 1.14+
- Java RE for running antlr
- VS Code extension for generating go parser / vscode to use settings.json

### build
`go build *.go`

### run
- Put the input to be tested in `hello_test.txt`
- `./main`
- expected output should reflect the nodes / edges graph structure in `struct_i_want.go`. Example (using checked-in hello_test.txt):
```
output: {[%!s(int=0) %!s(int=1) %!s(int=2) %!s(int=3) %!s(int=4) %!s(int=5) %!s(int=6) %!s(int=7) %!s(int=8) %!s(int=9)] map[%!s(int=0):[%!s(int=412) %!s(int=480) %!s(int=777) %!s(int=1453)] %!s(int=1):[%!s(int=132) %!s(int=1209)] %!s(int=2):[%!s(int=1614)] %!s(int=3):[%!s(int=3)] %!s(int=4):[%!s(int=1146)] %!s(int=5):[%!s(int=5) %!s(int=528)] %!s(int=6):[%!s(int=107) %!s(int=136) %!s(int=366) %!s(int=1148) %!s(int=1875)] %!s(int=7):[%!s(int=452) %!s(int=701) %!s(int=1975)] %!s(int=8):[%!s(int=164)] %!s(int=9):[%!s(int=912) %!s(int=920)]]}
```

## Evaluation
I still think that this might be quicker to write a quick grammar for future puzzle inputs (as compared to splitting / regex operations in Go), but the listener code is not very elegant.  What I really wanted was something like https://github.com/caarlos0/env where you can specify what you want somehow in the struct tags, but this didn't really get there.