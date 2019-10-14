package main

import (
	"testing"

	"github.com/vancelongwill/snapshot"
)

func TestMain(t *testing.T) {
	example := "example"
	err := snapshot.WithLabel("equals 'example'").Matches(&example)
	if err != nil {
		t.Fatal(err.Error())
	}
}

type Something struct {
	MoreThings []string
}

func TestSomething(t *testing.T) {
	something := Something{}
	something.MoreThings = []string{"a", "b", "c"}
	label := "should be something"
	err := snapshot.WithLabel(label).Matches(&something)
	// returns an error with a pretty diff if the snapshot doesnt match
	if err != nil {
		// print the label and the diff
		t.Fatalf("Snapshot didn't match.\n - '%s'\n%s",
			label,
			err.Error())
	}
}
