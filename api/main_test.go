package main

import (
	"api/mock"
	"context"
	"testing"
)

var ctx = context.TODO()

func TestGet(t *testing.T) {
	response, err := HandleRequest(ctx, mock.Get)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if response.StatusCode != 200 {
		t.Fail()
	}
}
