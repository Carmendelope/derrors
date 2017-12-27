//
// Copyright (C) 2017 Daisho Group - All Rights Reserved
//
// Definition of the generic Daisho error.
// Notice that stack traces are not serialized.

package derrors

import (
    "fmt"
    "strings"
    "runtime"
    "reflect"
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
    return fmt.Sprintf("%s:%d - %s", se.FunctionName, se.Line, se.File)
}

// GenericError structure that defines the basic elements shared by all DaishoErrors.
type GenericError struct {
    // ErrorType from the enumeration.
    ErrorType ErrorType `json:"errorType"`
    // Message contains the error message.
    Message string `json:"message"`
    // stackTrace contains the calling stack trace.
    stackTrace [] StackEntry
}

// StackToString generates a string with a stack entry per line.
func (ge * GenericError) StackToString() string {
    toString := make([] string, len(ge.stackTrace))
    for i, v := range ge.stackTrace {
        toString[i] = v.String()
    }
    return strings.Join(toString, "\n")
}

func (ge * GenericError) basicError() string {
    return fmt.Sprintf("[%s] %s", ge.ErrorType, ge.Message)
}

// Type returns the ErrorType associated with the current DaishoError.
func (ge * GenericError) Type() ErrorType {
    return ge.ErrorType
}

// StackTrace returns an array with the calling stack that created the error.
func (ge * GenericError) StackTrace() [] StackEntry {
    return ge.stackTrace
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
    entity interface {}
}

// NewEntityError creates a new DaishoError of type Entity.
//   params:
//     entity The associated entity.
//     msg The error message.
//   returns:
//     An EntityError.
func NewEntityError(entity interface{}, msg string) * EntityError {
    return &EntityError{
        GenericError{EntityErrorType, msg, GetStackTrace()}, entity}
}

// Error returns the string representation of the error.
func (ee * EntityError) Error() string {
    return fmt.Sprintf("%s, entity: %s", ee.basicError(), PrettyPrintStruct(ee.entity))
}

// DebugReport returns a detailed error report including the stack information.
func (ee * EntityError) DebugReport() string {
    return fmt.Sprintf("%s\nentity: %s\nStackTrace:\n%s",
        ee.basicError(), PrettyPrintStruct(ee.entity), ee.StackToString())
}