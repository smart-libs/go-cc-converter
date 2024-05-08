# The `go-cc-converter` package

The `go-cc-converter` package belongs to the `go-cc` package family that provides crosscutting layer services.
Its sole responsibility is to provide data conversion services through a registry, and converter façade.

A _crosscutting layer_ is a layer within a software architecture that is responsible for handling crosscutting concerns. It provides a centralized location for implementing and managing aspects of the system that cut across multiple components or modules.

## Overview

When starting the quick proposal for the conversion function signature is:

`func[From any, To any](From) (To, error)`

It basically receives the data to be converted, apply the conversion logic and returns the converted value or an error.
Soon you will realize this signature works fine when you know how to create an instance of the `To` data type. 
For instance, if the `To` type is an interface and you have an array of bytes, which implementation to use?

Although the signature is fine for most of the cases, you cannot create in Golang a façade to convert any kind of data. 
The Golang Generics is only available for struct or functions, but not for methods. You cannot create a structure whose
methods uses Generics like below:

```go
Facade interface {
	Convert[From any, To any](From) (To, error)
}
```
The alternative is to remember all conversion functions by name and use it when needed. It is not too bad.
Even if you can remember all conversion function names, this approach does not work for libs providing
specialized services that have to deal with data conversions where the From and To type are not known upfront and so
the generic feature is not feasible.

To help, the `go-cc-converter` package offers a mixed alternative with a façade conversion function without Generics but
with helping functions using Generics. It also provides a registry where all conversion functions with distinct
signatures can be registered avoiding you to remember the names of all these functions.

## What does the package have?

### Interfaces
#### The [`converter.Converters`](converters.go) interface
The [`converter.Converters`](converters.go) interface is the façade to have access to all conversion functions available. 
```go
	Converters interface {
		Convert(from any, to any) error
		ConvertToType(from any, toType reflect.Type) (any, error)
	}
```
The `Convert()/ConvertToType()` methods do the following steps:
1. Identifies the `from` and `to` types;
2. Looks up for a conversion function using the `from` and `to` types;
3. If it finds the conversion function, then it executes the conversion and returns the result;
4. If no conversion function is found, then it tries many fallbacks to manage a way to do the conversion;
5. If no fallback succeeds, then returns a not found error.

The library preload many conversion functions but you can add your own functions too. You can also have a special 
`Converters` if your app or library need.

The conversion functions provided by the library are in the [funcs](funcs) folder, or the `converterfuncs` package
#### The `converter.Registry` interface
The `converter.Registry` interface is the contract the registry of conversion function shall implement. 
It has two basic methods:

```go
package converter

type Registry interface {
    Add(handler Handler) *Handler
    GetConverter(from, to reflect.Type) FromToFunc
	GetConverterTo(from, to reflect.Type) ToFunc
}
```
The `Add()` method adds a new conversion function specified by the given `Handler`. The `Handler` is the strategy used to
avoid mistakes when providing the conversion function to be registered. `Registry` stores function with the
following signature defined by `FromToFunc`: 

```go
package converter

type FromToFunc func(any, any) error
```
This signature is too wide, it allows any `from` and any `to` which it is not real. Even using reflection in Golang it 
is not possible to identify the real argument types (well, they are any!). To better capture the `from` and `to` type
and ensure they are not incorrect, the package provides following Handler factories:

```go
package converter

func CreateHandler[From any, To any](converter func(From, *To) error) Handler
func CreateHandlerWithReturn[From any, To any](converter func(From) (To, error)) Handler
func CreateHandlerWithReturnNoError[From any, To any](converter func(From) To) Handler
```

All these factories capture the `from` and the `to` types and make any adjustment needed. They will support almost all 
conversion functions.

The package provides default `Registry` instance that can be accessed as shown below:

```go
converterdefault.DefaultRegistry
```

The `converterdefault` package provides helper functions to quickly add Handler instances to the default `Registry` as 
shown below:

```go
package converterdefault

func AddHandler[From any, To any](fromToFunc func(From, *To) error) {
	DefaultRegistry.Add(converter.CreateHandler[From, To](fromToFunc))
}

func AddHandlerWithReturn[From any, To any](toFunc func(From) (To, error)) {
	DefaultRegistry.Add(converter.CreateHandlerWithReturn[From, To](toFunc))
}

func AddHandlerWithReturnNoError[From any, To any](toFunc func(From) To) {
	DefaultRegistry.Add(converter.CreateHandlerWithReturnNoError[From, To](toFunc))
}
```

If you need to create a special `Registry` instance for your application or package with special converters, then
you can use the function below that will return the default implementation of `Registry`:

```go
converterdefault.NewRegistry()
```
