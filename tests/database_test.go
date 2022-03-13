package test

import (
	"testing"

	dbase "github.com/jaysonmulwa/jumia/internal/database"
)

func TestConnect(t *testing.T) {

	db, err := dbase.Connect()

	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}

	if db == nil {
		t.Errorf("Expected db to be not nil, got nil")
	}

}
