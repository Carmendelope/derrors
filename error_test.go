//
// Copyright (C) 2017 Daisho Group - All Rights Reserved
//
// Error tests

package derrors

import (
    "errors"
    "fmt"
    "strings"
    "testing"
)

func callerGetStackTrace() []StackEntry {
    return GetStackTrace()
}

func TestGetStackTrace(t *testing.T) {
    stackTrace := callerGetStackTrace()
    assertTrue(t, len(stackTrace) > 0, "expecting stack")
    parent := stackTrace[0]
    parentFunctionName := strings.Split(parent.FunctionName, ".")[2]
    assertEquals(t, "TestGetStackTrace",
        parentFunctionName, "Expecting parent function")
}

type testPrettyStruct struct {
    msg string
}
func (ss * testPrettyStruct) String() string {
    return "PRETTY " + ss.msg
}

func TestNewGenericError(t *testing.T) {
    var err DaishoError = NewGenericError("msg")
    assertTrue(t, err != nil, "Expecting new error")
}

func TestNewEntityError(t *testing.T) {
    var err DaishoError = NewEntityError("entity", "msg")
    assertTrue(t, err != nil, "Expecting new error")
}

func TestNewOperationError(t *testing.T) {
    var err DaishoError = NewOperationError("msg").WithParams("param1")
    assertTrue(t, err != nil, "expecting new error")
}

func TestPrettyPrintStruct(t *testing.T) {
    basicElement := "string"
    r1 := PrettyPrintStruct(basicElement)
    assertEquals(t, "\"string\"", r1, "expecting same message")
    type basicStruct struct {
        msg string
    }
    structElement := &basicStruct{"string2"}
    r2 := PrettyPrintStruct(structElement)
    assertEquals(t, "&derrors.basicStruct{msg:\"string2\"}", r2, "expecting struct message")

    stringElement := &testPrettyStruct{"PRINT"}
    r3 := PrettyPrintStruct(stringElement)
    assertEquals(t, "PRETTY PRINT", r3, "expecting pretty print")

}

func TestEntityError(t *testing.T) {
    basicEntity := "basicEntity"
    parent := errors.New("golang error")
    parent2 := errors.New("previous error")
    e := NewEntityError(basicEntity, IOError, parent, parent2)
    errorMsg := e.Error()
    assertEquals(t, "[Entity] I/O error", errorMsg, "Message should match")
    detailedError := e.DebugReport()
    fmt.Println(detailedError)
    fmt.Println("Error(): " + e.Error())
}

func TestAsDaishoError(t *testing.T) {
    err := errors.New("some golang error")
    derror := AsDaishoError(err, "msg")
    assertEquals(t, "msg", derror.Message, "Expecting message")

    derrorFromNil := AsDaishoError(nil, "msg")
    assertTrue(t, derrorFromNil == nil, "Should be nil")

    derrorWithParam := AsDaishoError(err, "msg").WithParams("param1")
    assertTrue(t, derrorWithParam != nil, "should not be nil")
    assertEquals(t, 1, len(derrorWithParam.Parameters), "expecting one parameter")
}