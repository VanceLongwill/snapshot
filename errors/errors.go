package errors

// ErrSnapshotMismatch is returned when the snapshot generated doesn't match that on file
type ErrSnapshotMismatch struct{ error }

// SnapshotMismatch wraps error
func SnapshotMismatch(err error) ErrSnapshotMismatch {
	return ErrSnapshotMismatch{err}
}

// ErrSnapNotFound indicates that a snapshot cannot be located
type ErrSnapNotFound struct{ error }

// SnapNotFound wraps error
func SnapNotFound(err error) ErrSnapNotFound {
	return ErrSnapNotFound{err}
}

// ErrSnapAlreadyExists is used to indicate when a snapshot is thought to be knew, but in fact is already present
type ErrSnapAlreadyExists struct{ error }

// SnapAlreadyExists wraps error
func SnapAlreadyExists(err error) ErrSnapAlreadyExists {
	return ErrSnapAlreadyExists{err}
}

// ErrSnapsEmpty indicates that no snaps were loaded when it was expected that they would be
type ErrSnapsEmpty struct{ error }

// SnapsEmpty wraps error
func SnapsEmpty(err error) ErrSnapsEmpty {
	return ErrSnapsEmpty{err}
}
