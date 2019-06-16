package discover

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	bolt "github.com/etcd-io/bbolt"
	neatthing "github.com/taothit/one-neat-thing-today/gen/neat_thing"
)

const (
	dbStore = "./.db/ontt.db"

	BucketName = "neatThings"
)

var (
	KeyDateFormat = time.RFC3339
)

// neatThing service example implementation.
// The example methods log the requests and return zero values.
type neatThingsrvc struct {
	logger *log.Logger
	db     *bolt.DB
}

// NewNeatThing returns the neatThing service implementation.
// Panics when persistent store cannot be opened.
func NewNeatThing(logger *log.Logger, dbFilePath *string) neatthing.Service {
	return &neatThingsrvc{logger, loadDb(stringOrDefault(dbFilePath, dbStore))}
}

func loadDb(dbPath string) (datastore *bolt.DB) {
	if db, err := bolt.Open(dbPath, 0600, bolt.DefaultOptions); err != nil {
		panic(fmt.Sprintf("could not open neat things datastore: %v", err))
	} else {
		datastore = db
	}
	return
}

func stringOrDefault(dbFilePath *string, defaultValue string) string {
	dataStore := defaultValue
	if dbFilePath != nil && *dbFilePath != "" {
		dataStore = *dbFilePath
	}
	return dataStore
}

// NeatThingToday implements neatThingToday.
func (s *neatThingsrvc) NeatThingToday(ctx context.Context) (res *neatthing.NeatThing, view string, err error) {
	res, err = Fetch(todayAtMidnight(time.UTC).Format(KeyDateFormat), s.db)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get today's neat thing: %v", err)
	} else {
		view = "default"
	}
	return
}

func Fetch(key string, db *bolt.DB) (nt *neatthing.NeatThing, err error) {
	nt = &neatthing.NeatThing{}
	if err := db.View(func(tx *bolt.Tx) error {
		bucket, err := openNeatThingBucket(tx)
		if err != nil {
			return fmt.Errorf("failed to open bucket full of neat things: %v", err)
		}

		enc := bucket.Get([]byte(key))
		if enc == nil {
			return fmt.Errorf("failed to find neat thing for key[%s]", key)
		}

		decoder := gob.NewDecoder(bytes.NewReader(enc))
		if err := decoder.Decode(&nt); err != nil {
			return fmt.Errorf("failed to decode neat thing: %v", err)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to fetch neat thing: %v", err)
	}
	return nt, nil
}

func todayAtMidnight(loc *time.Location) *time.Time {
	year, month, day := time.Now().Date()
	t := time.Date(year, month, day, 0, 0, 0, 0, loc)
	return &t
}

func openNeatThingBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	bucket := tx.Bucket([]byte(BucketName))
	if bucket == nil {
		return nil, fmt.Errorf("failed to open bucket (%s)", BucketName)
	}
	return bucket, nil
}

// NewNeatThing stores a neat thing to be viewed at a later date.
func (s *neatThingsrvc) NewNeatThing(ctx context.Context, nt *neatthing.NeatThing) (res *neatthing.NeatThing, view string, err error) {
	res = nt
	view = "full"
	if err := s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return fmt.Errorf("failed to create test bucket for records: %v", err)
		}

		buf, err := encode(nt)
		if err := b.Put([]byte(*nt.Date), buf.Bytes()); err != nil {
			return fmt.Errorf("failed to write %s to [%s/%s]: %v", nt, BucketName, tx.DB().Path(), err)
		}
		return nil
	}); err != nil {
		return nil, "", fmt.Errorf("failed to write %s to db [%v]: %v", nt, s.db, err)
	}
	if err := s.db.Close(); err != nil {
		return nil, "", fmt.Errorf("failed to close connection to db [%v]: %v", s.db, err)
	}
	return
}

func encode(record *neatthing.NeatThing) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(make([]byte, 0))
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(record); err != nil {
		return nil, fmt.Errorf("failed to encode %v: %v", record, err)
	}
	return buf, nil
}
