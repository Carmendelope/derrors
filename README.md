# derrors - Extended Errors

This repository contains the definition of an extended error structure for Go applications.

## General overview

The main purpose of this repository is to improve error reporting facilitating the communication of error states to the
users and allowing deeper reporting of the errors for the developers at the same time.

The Error interfaces defines a set of basic methods that makes a Error compatible with the GolangError but
provides extra functions to track the error origin.


## How to use it

Defining a new error automatically extracts the StackTrace

```
return derrors.NewEntityError(descriptor, errors.NetworkDoesNotExists, err)
```

Will print the following message when calling `String()`.

```
[Operation] network does not exists
```

And a detailed one when calling `DebugReport()`.

```
[Operation] network does not exists
Parameters:
P0: []interface {}{"n1b59e008-a9f2-4a25-866a-ace0cabc38b2asdf"}

StackTrace:
<stack trace from the caller>
```

## Transforming a Go error

Use the automatic extraction, notice that if the error if nil, the result is nil to facilitate `return` constructs.

```go
func (h * Manager) SampleFunction() derrors.Error {
    err := myClassicGoFunction()
    return derrors.AsError(err, "optional message")
}
```

## Contributing
​
Please read [contributing.md](contributing.md) and [code-of-conduct.md](code-of-conduct.md) for details on our code of conduct, and the process for submitting pull requests to us.
​
## Versioning
​
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/nalej/derrors/tags). 
​
## Authors
​
See also the list of [contributors](https://github.com/nalej/derrors/contributors) who participated in this project.
​
## License
​
This project is licensed under the Apache 2.0 License - see the [LICENSE-2.0.txt](LICENSE-2.0.txt) file for details.