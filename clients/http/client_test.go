package http

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestClient(t *testing.T) {
	c, err := New(
		"https://invest-public-api.tinkoff.ru/rest",
		"t.ZPRR_Dd7YJQi4rkxApdEWyK1JcxcB7fr_hUlUWmdOqYzb8E_UgI83zwvL2SRHRw8s2fN3-LmYqBBV91KSsC1sQ",
	)
	if err != nil {
		t.Fatal(err)
	}
	id, err := uuid.Parse("affad2be-53f2-4203-bd12-2214e7397dec")
	if err != nil {
		t.Fatal(err)
	}
	err = c.ShareByUID(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
}
