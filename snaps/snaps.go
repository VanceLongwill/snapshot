package snaps

import (
	"encoding/xml"
	"fmt"

	"github.com/vancelongwill/snapshot/errors"
)

// Snaps represents a set of snapshots from a snapshot file
type Snaps struct {
	XMLName xml.Name `xml:"Snaps"`
	Snaps   []Snap   `xml:"Snap"`
}

// New reads in byte snapshot data
func New(raw []byte) *Snaps {
	var snaps Snaps
	if len(raw) == 0 {
		return &snaps
	}
	err := xml.Unmarshal(raw, &snaps)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse snapshot data: %s", err.Error()))
	}
	return &snaps
}

// Serialize converts a set of snapshots to byte data
func (s *Snaps) Serialize() []byte {
	bytes, err := xml.MarshalIndent(&s, "", "    ")
	if err != nil {
		panic(err)
	}
	return bytes
}

func (s *Snaps) isEmpty() bool {
	return s.Snaps == nil || len(s.Snaps) == 0
}

func (s *Snaps) findIndex(label string) (int, error) {
	if s.isEmpty() {
		return -1, errors.SnapsEmpty(
			fmt.Errorf("Can't find snap with label '%s' in empty list", label))
	}
	for i, snap := range s.Snaps {
		if snap.Label == label {
			return i, nil
		}
	}
	return -1, errors.SnapNotFound(
		fmt.Errorf("Unable to find snap with label '%s'", label))
}

// Find searches for a snapshot which matches the given label
func (s *Snaps) Find(label string) (Snap, error) {
	ind, err := s.findIndex(label)
	if err != nil {
		return Snap{}, err
	}
	return s.Snaps[ind], nil
}

// Update replaces a snapshot which is already present with it's new version
func (s *Snaps) Update(snap Snap) error {
	ind, err := s.findIndex(snap.Label)
	if err != nil {
		return err
	}
	s.Snaps[ind] = snap
	return nil
}

// Add appends a snapshot to the set
func (s *Snaps) Add(snap Snap) error {
	ind, _ := s.findIndex(snap.Label)
	if ind != -1 {
		return errors.SnapAlreadyExists(
			fmt.Errorf("Can't add snap with label which already exists: %s", snap.Label))
	}
	s.Snaps = append(s.Snaps, snap)
	return nil
}
