package core_test

import (
	"fmt"
	"testing"

	core "github.com/jacob-ebey/graphql-core"
)

func TestWrappedErrorDefaultMessage(t *testing.T) {
	err := &core.WrappedError{}

	if err.Error() != core.DefaultErrorMessage {
		t.Fatal("Default error message was not correct.")
	}
}

func TestWrappedErrorCustomMessage(t *testing.T) {
	message := "Custom message"
	err := &core.WrappedError{
		Message: message,
	}

	if err.Error() != message {
		t.Fatal("Custom message was not used.")
	}
}

func TestWrappedErrorInternalMessage(t *testing.T) {
	internalError := fmt.Errorf("Internal message")
	err := &core.WrappedError{
		InternalError: internalError,
	}

	if err.Error() != internalError.Error() {
		t.Fatal("Internal error message was not used.")
	}
}

func TestWrappedErrorExtensionsNilValues(t *testing.T) {
	err := &core.WrappedError{
		Message: "Custom message",
	}

	extensions := err.Extensions()

	if extensions["internalError"].(*string) != nil {
		t.Fatal("internalError extension should be nil.")
	}

	if extensions["code"].(*string) != nil {
		t.Fatal("code extension should be nil.")
	}
}

func TestWrappedErrorExtensionsValues(t *testing.T) {
	internalError := fmt.Errorf("Internal message")
	code := "CODE"
	err := &core.WrappedError{
		Message:       "Custom message",
		InternalError: internalError,
		Code:          code,
	}

	extensions := err.Extensions()

	internalErrorExtension := extensions["internalError"].(*string)

	if *internalErrorExtension != internalError.Error() {
		t.Fatal("internalError extension was not used.")
	}

	codeExtension := extensions["code"].(*string)

	if *codeExtension != code {
		t.Fatal("code extension was not used.")
	}
}
