package database

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/mesg-foundation/engine/database/store"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/ownership"
)

var (
	errCannotSaveOwnershipWithoutHash = errors.New("database: can't save ownership without hash")
)

// OwnershipDB is a database for storing ownership definition.
type OwnershipDB struct {
	s   store.Store
	cdc *codec.Codec
}

// NewOwnershipDB returns the database which is located under given path.
func NewOwnershipDB(s store.Store, cdc *codec.Codec) *OwnershipDB {
	return &OwnershipDB{
		s:   s,
		cdc: cdc,
	}
}

// unmarshal returns the ownership from byte slice.
func (d *OwnershipDB) unmarshalOwnership(hash hash.Hash, value []byte) (*ownership.Ownership, error) {
	var s ownership.Ownership
	if err := d.cdc.UnmarshalBinaryBare(value, &s); err != nil {
		return nil, fmt.Errorf("database: could not decode ownership %q: %w", hash.String(), err)
	}
	return &s, nil
}

// All returns every ownership in database.
func (d *OwnershipDB) All() ([]*ownership.Ownership, error) {
	var (
		ownerships []*ownership.Ownership
		iter       = d.s.NewIterator()
	)
	for iter.Next() {
		hash := hash.Hash(iter.Key())
		s, err := d.unmarshalOwnership(hash, iter.Value())
		if err != nil {
			return nil, err
		}
		ownerships = append(ownerships, s)
	}
	iter.Release()
	return ownerships, iter.Error()
}

// Save stores ownership in database.
// If there is an another ownership that uses the same sid, it'll be deleted.
func (d *OwnershipDB) Save(o *ownership.Ownership) error {
	if o.Hash.IsZero() {
		return errCannotSaveOwnershipWithoutHash
	}
	b, err := d.cdc.MarshalBinaryBare(o)
	if err != nil {
		return err
	}
	return d.s.Put(o.Hash, b)
}

// Close closes database.
func (d *OwnershipDB) Close() error {
	return d.s.Close()
}
