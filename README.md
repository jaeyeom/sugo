# sugo
Sugo offers handy syntactic sugar (though not technically a syntactic sugar) for
Go, designed to simplify your code and save you time. Whether you're dealing
with repetitive tasks or complex operations, `sugo` provides intuitive solutions
that enhance readability and efficiency in Go programming.

## Features
 - **par**: Run multiple goroutines concurrently, possibly for parellelizing it.
 - **errors/must**: If you prefer not to manually handle errors, you may use it.
 - **ptr/ref**: Convenient way to create a pointer to a literal value.
 - **ptr/deref**: Convenient way to dereference a pointer with a default value
   for a `nil` pointer.
 - **arg**: ArgMin/ArgMax.
 - **ranger**: To deal with integer indices safely. It's more useful with `par`.

## Getting Started

### Installation
```bash
go get github.com/jaeyeom/sugo
```

## Usage
Please take a look at examples in https://pkg.go.dev/github.com/jaeyeom/sugo

## Versions
Backward compatibility is guaranteed for the same major version since v1.0.0.

## Contributing
We welcome contributions from everyone. If you have a suggestion that could
improve `sugo`, please open an issue and tag it with "enhancement" for
discussion. Don't forget to star the project! Thank you again!

## License
Sugo is licensed under the MIT License - see the LICENSE file for details.
