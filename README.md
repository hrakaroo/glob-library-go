# glob-library-go

Glob style pattern matcher for strings in Go.

## Why

This is a work in process.  I'm attempting to port my 
https://github.com/hrakaroo/glob-library-java library
to Go.  So far I've gotten all of the code to compile and pulled over all tests.  
I still want to verify code coverage and run some benchmarks before
it will be "done".

I also fully acknowledge that this library isn't as "needed" for Go since Go ships
with it's own https://go.dev/src/path/filepath/match.go and https://pkg.go.dev/v.io/v23/glob 
which looks like it operates pretty similar.

v0.1.0-alpha

So I've created a v0.1.0-alpha release of this library. To use it just put

`require github.com/hrakaroo/glob-library-go v0.1.0-alpha`

in your `go.mod` file and then you can import the library with

`import glob "github.com/hrakaroo/glob-library-go"`

in your code.  So far it seems to work pretty well.  I still need to build the
benchmarks, but otherwise my simple test worked

```func main() {
	fmt.Println("hello")
	m, err := glob.Compile("*foo*")
	if err != nil {
		panic(err)
	}

	fmt.Println(m.Matches("bogfoo"))
}```
