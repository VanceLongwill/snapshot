package snapshot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vancelongwill/snapshot/errors"
	"github.com/vancelongwill/snapshot/snaps"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	// SnapshotDir is where snapshots are stored
	SnapshotDir string = "__snapshots__"
	// SnapshotExtension is the extension for snapshot files
	SnapshotExtension string = ".snap.xml"
	// Indent - 4 spaces
	Indent string = "    "
)

// Comparison compares snapshots
type Comparison interface {
	// Matches runs a comparison with the saved snapshot
	// overwrites the saved snapshot if CLI flag "-u" is passed to `go test`
	Matches(v interface{}) error
}

type snapshotCompare struct {
	prev         []byte
	shouldUpdate bool
	update       func([]byte) error
}

func (s *snapshotCompare) Matches(v interface{}) error {
	var isMatch = true
	// serialise the data
	vSerialised := prettyJSON(v)

	var snapshotEmpty = !(len(s.prev) > 0)

	// compare snapshot if there is one
	if !snapshotEmpty {
		isMatch = bytes.Equal(vSerialised, s.prev)
	}

	if snapshotEmpty || (s.shouldUpdate && !isMatch) {
		defer func() {
			err := s.update(vSerialised)
			if err != nil {
				panic(err)
			}
		}()
		return nil
	}

	if !isMatch {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(string(s.prev), string(vSerialised), false)
		diff := dmp.DiffPrettyText(diffs)
		return errors.SnapshotMismatch(fmt.Errorf(diff))
	}

	return nil
}

// WithLabel reads the existing snapshot and allows comparison
func WithLabel(label string) Comparison {
	// Get the file name of the invoker
	calledFromFilename := getCallerFilename()

	snapshotDir, snapshotPath := snapshotPath(calledFromFilename)

	var snapshotData []byte
	snapshotFile, err := os.OpenFile(snapshotPath, os.O_RDWR, 0666)

	defer func() {
		_ = snapshotFile.Close()
	}()

	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		_ = os.Mkdir(snapshotDir, os.ModePerm)
	} else {
		snapshotData, err = ioutil.ReadFile(snapshotPath)
		if err != nil {
			panic(err)
		}
	}

	s := snaps.New(snapshotData)

	foundSnap, err := s.Find(label)
	if err != nil {
		switch err.(type) {
		case errors.ErrSnapsEmpty, errors.ErrSnapNotFound:
			foundSnap = snaps.Snap{Label: label}
			if err := s.Add(foundSnap); err != nil {
				panic(err)
			}
		default:
			panic(err)
		}
	}

	var update = func(content []byte) error {
		foundSnap.Content = content
		err = s.Update(foundSnap)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(snapshotPath, s.Serialize(), 0644)
		if err == nil {
			fmt.Printf("Snapshot '%s' written to %s\n", foundSnap.Label, snapshotPath)
		}
		return err
	}

	return &snapshotCompare{
		prev:         foundSnap.Content,
		shouldUpdate: *shouldUpdate,
		update:       update,
	}
}
