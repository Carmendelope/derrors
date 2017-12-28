//
// Copyright (C) 2017 Daisho Group - All Rights Reserved
//
// Definition of the generic Daisho error.
// Notice that stack traces are not serialized.

package derrors

import (
    "encoding/json"
    "errors"
    "fmt"
    "reflect"
    "runtime"
    "strings"
)

// StackEntry structure that contains information about an element in the calling stack.
type StackEntry struct {
    // FunctionName of the calling function.
    FunctionName string
    // File where the function is located.
    File string
    // Line of the file where the function is located.
    Line int
}

// NewStackEntry creates a new stack entry.
//   params:
//     functionName The name of the calling function.
//     file The name of the file where the function is located.
//     line The line in the file.
func NewStackEntry(functionName string, file string, line int) * StackEntry {
    return &StackEntry{functionName, file, line}
}

// String returns the string representation of an StackEntry.
func (se * StackEntry) String() string {
    return fmt.Sprintf("%s - %s:%d", se.FunctionName, se.File, se.Line)
}

// GenericError structure that defines the basic elements shared by all DaishoErrors.
type GenericError struct {
    // ErrorType from the enumeration.
    ErrorType ErrorType `json:"errorType"`
    // Message contains the error message.
    Message string `json:"message"`
    // causes contains the list of causes of the error.
    Causes [] string `json:"causes"`
    // stackTrace contains the calling stack trace.
    Stack [] StackEntry `json:"stackTrace"`
}

// NewGenericError returns a general purpose error.
func NewGenericError(msg string, causes ...error) * GenericError {
    return &GenericError{
        GenericErrorType,
        msg,
        ErrorsToString(causes),
        GetStackTrace()}
}

// StackToString generates a string with a stack entry per line.
func (ge * GenericError) StackToString() string {
    toString := make([] string, len(ge.Stack))
    for i, v := range ge.Stack {
        toString[i] = v.String()
    }
    return strings.Join(toString, "\n ")
}

func (ge * GenericError) basicError() string {
    return fmt.Sprintf("[%s] %s", ge.ErrorType, ge.Message)
}

func (ge * GenericError) causesToString() string {
    return strings.Join(ge.Causes, "\n> ")
}

func (ge * GenericError) Error() string {
    return ge.basicError()
}

// Type returns the ErrorType associated with the current DaishoError.
func (ge * GenericError) Type() ErrorType {
    return ge.ErrorType
}

// DebugReport returns a detailed error report including the stack information.
func (ge * GenericError) DebugReport() string {
    return fmt.Sprintf("%s\nCauses:\n%s\nStackTrace:\n%s",
        ge.Error(), ge.causesToString(), ge.StackToString())
}

// StackTrace returns an array with the calling stack that created the error.
func (ge * GenericError) StackTrace() [] StackEntry {
    return ge.Stack
}

// AsDaishoError checks an error. If it is nil, it returns nil, if not, it will create an equivalent GenericError
func AsDaishoError(err error, msg string) * GenericError {
    if err != nil {
        return NewGenericError(msg, err)
    }
    return nil
}

// GetStackTrace retrieves the calling stack and transform that information into an array of StackEntry.
func GetStackTrace() [] StackEntry {
    var programCounters [32] uintptr
    // Skip the two first callers.
    callersToSkip := 2
    callerCount := runtime.Callers(callersToSkip, programCounters[:])
    stackTrace := make([] StackEntry, callerCount)
    for i := 0; i < callerCount; i++ {
        function := runtime.FuncForPC(programCounters[i])
        filePath, line := function.FileLine(programCounters[i])
        stackTrace[i] = *NewStackEntry(function.Name(), filePath, line)
    }
    return stackTrace
}

// ErrorsToString transform a list of errors into a list of strings.
func ErrorsToString(errors []error) [] string {
    result := make([]string, len(errors))
    for i := 0; i < len(errors); i++ {
        result[i] = errors[i].Error()
    }
    return result
}

// PrettyPrintStruct aims to produce a pretty print of a giving structure by analyzing it.
func PrettyPrintStruct(data interface {}) string {
    v := reflect.ValueOf(data)
    method := v.MethodByName("String")
    if method.IsValid() && method.Type().NumIn() == 0 && method.Type().NumOut() == 1 &&
        method.Type().Out(0).Kind() == reflect.String{
        result := method.Call([]reflect.Value{})
        return result[0].String()
    }
    return fmt.Sprintf("%#v", data)
}

// EntityError defines a structure for entity related errors.
type EntityError struct {
    // Generic Error
    GenericError
    // associated entity.
    Entity interface {} `json:"entity"`
}

// NewEntityError creates a new DaishoError of type Entity.
//   params:
//     entity The associated entity.
//     msg The error message.
//   returns:
//     An EntityError.
func NewEntityError(entity interface{}, msg string, causes ...error) * EntityError {
    return &EntityError{
        GenericError{
            EntityErrorType,
            msg,
            ErrorsToString(causes),
            GetStackTrace()},
        entity}
}

// Error returns the string representation of the error.
func (ee * EntityError) Error() string {
    return fmt.Sprintf("%s, entity: %s", ee.basicError(), PrettyPrintStruct(ee.Entity))
}

// DebugReport returns a detailed error report including the stack information.
func (ee * EntityError) DebugReport() string {
    return fmt.Sprintf("%s\nentity: %s\nCauses:\n%s\nStackTrace:\n%s",
        ee.basicError(), PrettyPrintStruct(ee.Entity), ee.causesToString(), ee.StackToString())
}

// ConnectionError structure for connectivity related issues.
type ConnectionError struct {
    // GenericError structure.
    GenericError
    // URL associated with the error.
    URL string `json:"url"`
}

// NewConnectionError creates a new DaishoError of type Connection.
//   params:
//     entity The associated entity.
//     msg The error message.
//   returns:
//     An EntityError.
func NewConnectionError(URL string, msg string, causes ...error) * ConnectionError {
    return &ConnectionError{
        GenericError{
            ConnectionErrorType,
            msg,
            ErrorsToString(causes),
            GetStackTrace()},
        URL}
}

// Error returns the string representation of the error.
func (ce * ConnectionError) Error() string {
    return fmt.Sprintf("%s, URL: %s", ce.basicError(), PrettyPrintStruct(ce.URL))
}

// FromJSON unmarshalls a byte array with the JSON representation into a DaishoError of the correct type.
//   params:
//     data The byte array with the serialized JSON.
//   returns:
//     A DaishoError if the data can be unmarshalled.
//     A Golang error if the unmarshal operation fails.
func FromJSON(data [] byte) (DaishoError, error) {
    genericError := &GenericError{}
    err := json.Unmarshal(data, &genericError)
    if err != nil {
        return nil, err
    }
    if ValidErrorType(genericError.ErrorType) {
        switch genericError.ErrorType {
        case GenericErrorType : {
            var result DaishoError = genericError
            return result, nil
            }
        case EntityErrorType : {
            entityError := &EntityError{}
            err = json.Unmarshal(data, &entityError)
            if err != nil {
                return nil, err
            }
            var result DaishoError = entityError
            return result, nil
        }
        case ConnectionErrorType : {
            connectionError := &ConnectionError{}
            err = json.Unmarshal(data, &connectionError)
            if err != nil {
                return nil, err
            }
            var result DaishoError = connectionError
            return result, nil
        }
        default:
            return nil, errors.New("Unrecognized error type")
        }
    }
    return nil, errors.New("Unsupported error type in conversion")
}