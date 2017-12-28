//
// Copyright (C) 2017 Daisho Group - All Rights Reserved
//
// Serialization tests

package derrors

import (
    "encoding/json"
    "errors"
    "testing"
)

func TestFromJsonGenericError(t *testing.T) {
    cause := errors.New("error cause")
    msg := "Error message"
    toSend := NewGenericError(msg, cause)

    data, err := json.Marshal(toSend)
    assertEquals(t, nil, err, "expecting no error")
    retrieved, err := FromJSON(data)

    assertEquals(t, nil, err, "message should be deserialized")
    //assertEquals(t, GenericErrorType, retrieved.Type(), "type mismatch")
    assertEquals(t, toSend, retrieved, "structure should match")
}

func TestFromJsonEntityError(t *testing.T){

    entity := "serializableEntity"
    cause := errors.New("another cause")
    msg := "Other message"
    toSend := NewEntityError(entity, msg, cause)

    data, err := json.Marshal(toSend)
    assertEquals(t, nil, err, "expecting no error")
    retrieved, err := FromJSON(data)

    assertEquals(t, nil, err, "message should be deserialized")
    //assertEquals(t, GenericErrorType, retrieved.Type(), "type mismatch")
    assertEquals(t, toSend, retrieved, "structure should match")

}

func TestFromJsonConnectionError(t *testing.T){

    URL := "http://url-that-fails.com"
    cause := errors.New("yet another cause")
    msg := "Yet another message"
    toSend := NewConnectionError(URL, msg, cause)

    data, err := json.Marshal(toSend)
    assertEquals(t, nil, err, "expecting no error")
    retrieved, err := FromJSON(data)

    assertEquals(t, nil, err, "message should be deserialized")
    //assertEquals(t, GenericErrorType, retrieved.Type(), "type mismatch")
    assertEquals(t, toSend, retrieved, "structure should match")

}