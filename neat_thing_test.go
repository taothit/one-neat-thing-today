package discover_test

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"code.soquee.net/testlog"
	bolt "github.com/etcd-io/bbolt"
	"github.com/google/go-cmp/cmp"
	discover "github.com/taothit/one-neat-thing-today"
	neatthing "github.com/taothit/one-neat-thing-today/gen/neat_thing"
)

var want = &neatthing.NeatThing{
	Name:         strptr("One Neat Thing Today"),
	Definition:   strptr("A daily neat thing provider"),
	Date:         strptr(todayAtMidnight()),
	Link:         strptr("https://github.com/taothit/one-neat-thing-today"),
	Bibliography: nil,
}

func strptr(str string) *string {
	out := str
	return &out
}

func TestNeatThingsrvc_NeatThingToday(t *testing.T) {
	todayAtMidnight()

	testDatastore, err := initializeTestDb(t, want)
	if err != nil {
		t.Fatalf("failed to store test records: %v", err)
	}

	logger := testlog.New(t)
	nts, _ := discover.NewNeatThing(logger, strptr(testDatastore))

	if got, view, err := nts.NeatThingToday(context.Background()); got == nil || !cmp.Equal(view, "default") || err != nil {
		t.Errorf("neatThingsrvc.NeatThingToday()=%s; got=[nt=%s|view=%s]; err=%v", want, got, view, err)
	}
}

func initializeTestDb(t *testing.T, records ...*neatthing.NeatThing) (string, error) {
	var testDatastore = generateDatastoreName()
	_ = os.Remove(testDatastore) // Remove if it is there; otherwise, proceed.
	db := openTestDatastore(testDatastore, t)

	if err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(discover.BucketName))
		if err != nil {
			return fmt.Errorf("failed to create test bucket for records: %v", err)
		}

		for _, record := range records {
			buf := encode(record, t)
			if err := b.Put([]byte(*record.Date), buf.Bytes()); err != nil {
				return fmt.Errorf("failed to write %v to %s: %v", record, testDatastore, err)
			}
		}
		return nil
	}); err != nil {
		t.Fatalf("failed to write records to %s: %v", testDatastore, err)
	}

	if err := db.Close(); err != nil {
		return "", fmt.Errorf("failed to close test datastore: %v", err)
	}

	return testDatastore, nil
}

func openTestDatastore(testDatastore string, t *testing.T) *bolt.DB {
	db, err := bolt.Open(testDatastore, 0600, bolt.DefaultOptions)
	if err != nil {
		t.Fatalf("failed to initialize test copy of datastore: %v", err)
	}
	return db
}

func generateDatastoreName() string {
	p := fmt.Sprintf("%s/test-%d.db", os.TempDir(), rand.Int())
	return strings.Replace(p, strings.Repeat(string(os.PathSeparator), 2), string(os.PathSeparator), -1)
}

func encode(record *neatthing.NeatThing, t *testing.T) *bytes.Buffer {
	buf := bytes.NewBuffer(make([]byte, 0))
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(record); err != nil {
		t.Errorf("failed to encode %v: %v", record, err)
	}
	return buf
}

func todayAtMidnight() string {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Format(discover.KeyDateFormat)
}

func TestNeatThingsrvc_NewNeatThing(t *testing.T) {
	logger := testlog.New(t)
	testDatastore := generateDatastoreName()
	nts, _ := discover.NewNeatThing(logger, strptr(testDatastore))

	if got, view, err := nts.NewNeatThing(context.Background(), want); got == nil || !cmp.Equal(*got, *want) || view != "full" || err != nil {
		t.Errorf("NewNeatThing(%s)=%s; got=[nt=%s|view=%s]; err=%v", want, want, got, view, err)
	} else {
		db := openTestDatastore(testDatastore, t)
		if got, err := discover.Fetch(todayAtMidnight(), db); got == nil || !cmp.Equal(*got, *want) || err != nil {
			t.Errorf("NewNeatThing(%s)=%s; got=[nt=%s|view=%s]; err=%v", want, want, got, view, err)
		}
	}
}

func TestNewNeatThing(t *testing.T) {
	tests := []struct {
		name string
		path *string
		err  error
	}{
		{"File only", strptr("file.db"), nil},
		{"Directory exists", strptr(os.TempDir() + "/file.db"), nil},
		{"Directory doesn't exist", strptr("./.db/file.db"), nil},
	}
	var logger = testlog.New(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := discover.NewNeatThing(logger, test.path); err != nil && err != test.err {
				t.Errorf("discover.NewNeatThing(logger, %s); err=%v", *test.path, err)
			} else {
				if _, err := os.Stat(*test.path); os.IsNotExist(err) {
					t.Errorf("discover.NewNeatThing(logger, %s); err=%v", *test.path, err)
				}
			}
			p, f := path.Split(*test.path)
			_ = os.RemoveAll(p)
			_ = os.Remove(f)
		})
	}
}
