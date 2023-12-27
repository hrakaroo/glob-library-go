# glob-library-go

Glob style pattern matcher for strings in Go.

## Why

DRAFT

This is a work in process.  I'm attempting to port my 
https://github.com/hrakaroo/glob-library-java library
to Go.  So far I've gotten all of the code to compile and I'm just pulling over the
tests now.  I still want to verify code coverage and run some benchmarks before
it will be "done".

I also fully acknowledge that this library isn't as "needed" for Go since Go ships
with it's own https://go.dev/src/path/filepath/match.go and https://pkg.go.dev/v.io/v23/glob 
which looks like it operates pretty similar.
