package snapshot

import (
	"encoding/json"
	"path/filepath"
)

func snapshotPath(path string) (parentDir, fullPath string) {
	d, f := filepath.Split(path)
	parentDir = filepath.Join(d, SnapshotDir)
	fullPath = filepath.Join(parentDir, f+SnapshotExtension)
	return
}

func prettyJSON(v interface{}) []byte {
	vSerialised, err := json.MarshalIndent(&v, Indent+Indent, Indent)
	if err != nil {
		panic(err)
	}
	// preserve indentation in the snapshot file
	vSerialised = append([]byte("\n"+Indent+Indent), vSerialised...)
	vSerialised = append(vSerialised, []byte("\n"+Indent)...)
	return vSerialised
}
