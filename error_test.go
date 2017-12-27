//
// Copyright (C) 2017 Daisho Group - All Rights Reserved
//
// Error tests
package derrors

import (
    "testing"
    "strings"
)

func TestGetStackTrace(t *testing.T) {
    stackTrace := GetStackTrace()
    AssertTrue(t, len(stackTrace) > 0, "expecting stack")
    parent := stackTrace[0]
    parentFunctionName := strings.Split(parent.FunctionName, ".")[2]
    AssertEquals(t, "TestGetStackTrace",
        parentFunctionName, "Expecting parent function")
}

type testPrettyStruct struct {
    msg string
}
func (ss * testPrettyStruct) String() string {
    return "PRETTY " + ss.msg
}


func TestPrettyPrintStruct(t *testing.T) {
    basicElement := "string"
    r1 := PrettyPrintStruct(basicElement)
    AssertEquals(t, "\"string\"", r1, "expecting same message")
    type basicStruct struct {
        msg string
    }
    structElement := &basicStruct{"string2"}
    r2 := PrettyPrintStruct(structElement)
    AssertEquals(t, "&derrors.basicStruct{msg:\"string2\"}", r2, "expecting struct message")

    stringElement := &testPrettyStruct{"PRINT"}
    r3 := PrettyPrintStruct(stringElement)
    AssertEquals(t, "PRETTY PRINT", r3, "expecting pretty print")


}