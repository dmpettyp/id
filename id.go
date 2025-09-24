package id

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	google_uuid "github.com/google/uuid"
)

// ID in the uuid package provides a type that represents UUIDs and common
// functionality for them. The ID type is intended to be embedded in other
// specific identifier types (e.g. address ID, domain entity IDs, etc) to
// provide type-safe identifiers.
//
// The ID type is implemented using a google/uuid ID, but offers a level
// of indirection that prevents clients from having to import the google/uuid
// package directly and depend on it.
type ID struct {
	GoogleUUID google_uuid.UUID
}

// New creates a new random UUID and returns an error if a
// UUID can't be generated.
func New() (ID, error) {
	id, err := google_uuid.NewRandom()

	if err != nil {
		return ID{}, errors.New("could not generate a UUID")
	}

	return ID{GoogleUUID: id}, nil
}

// MustNew creates a new random UUID or panics
func MustNew() ID {
	return ID{GoogleUUID: google_uuid.New()}
}

// String returns the string form of the UUID
func (uuid ID) String() string {
	return uuid.GoogleUUID.String()
}

// Parse constructs a UUID by parsing a UUID string and storing the resulting
// value in a UUID struct
func Parse(s string) (ID, error) {
	id, err := google_uuid.Parse(s)

	if err != nil {
		return ID{}, fmt.Errorf("could not parse UUID from %q: %w", s, err)
	}

	return ID{GoogleUUID: id}, nil
}

// Scan implements sql.Scanner so UUIDs can be read from databases
// transparently
func (uuid *ID) Scan(src any) error {
	return uuid.GoogleUUID.Scan(src)
}

// Value implements sql.Valuer so that UUIDs can be written to databases
// transparently
func (uuid ID) Value() (driver.Value, error) {
	return uuid.GoogleUUID.Value()
}

// MarshallJSON ensures that a UUID is serialized as the string-representation
// of the underlying UUID instead of a structre containing the underlying UUID
func (uuid ID) MarshalJSON() ([]byte, error) {
	marshaled, err := json.Marshal(uuid.GoogleUUID.String())

	if err != nil {
		return nil, fmt.Errorf("could not marshal JSON uuid: %w", err)
	}

	return marshaled, nil
}

// UnmarshallJSON deserializes a string UUID and assigns the parsed UUID to the
// underlying UUID structure
func (uuid *ID) UnmarshalJSON(data []byte) error {
	var err error
	var uuidString string

	if err = json.Unmarshal(data, &uuidString); err != nil {
		return fmt.Errorf("could not unmarshall JSON uuid into a string: %w", err)
	}

	uuid.GoogleUUID, err = google_uuid.Parse(uuidString)

	if err != nil {
		return fmt.Errorf("could not parse uuid string %q: %w", uuidString, err)
	}

	return nil
}

// IsNil determines if a UUID is nil (i.e. all 0s)
func (uuid ID) IsNil() bool {
	return uuid == ID{}
}
