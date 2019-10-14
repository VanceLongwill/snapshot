## snapshot

Provides basic methods for snapshot testing, inspired by [jest's snapshots](https://jestjs.io/docs/en/snapshot-testing)

This package doesn't aim to be anything more than a simple utility to be used in conjunction with
golang's own testing package, or any other testing framework which is also compatible with `go test`
snapshots are auto-generated and managed by this package:

### Getting started

1. Grab the package

```bash
go get -u github.com/vancelongwill/snapshot
```

2. Use it in tests

```golang
import (
  "testing"

  "github.com/VanceLongwill/snapshot"
)

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
```

3. Update failing snapshots

```bash
go test ./... -u
```

#### Contributing

Contributions are absolutely welcome!

##### Steps

- Make sure tests are passing

  ```bash
  make test
  ```

- Make sure it passes the linter

  ```bash
  make lint
  ```

- Open a pull request
