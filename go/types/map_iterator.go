// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

// mapIterator can efficiently iterate through a Noms Map.
type mapIterator struct {
	cursor       *sequenceCursor
	currentKey   Value
	currentValue Value
}

func (mi *mapIterator) Valid() bool {
	return mi.cursor.valid()
}

func (mi *mapIterator) Entry() (k Value, v Value) {
	return mi.Key(), mi.Value()
}

func (mi *mapIterator) Key() Value {
	if !mi.cursor.valid() {
		return nil
	}
	return mi.cursor.current().(mapEntry).key
}

func (mi *mapIterator) Value() Value {
	if !mi.cursor.valid() {
		return nil
	}
	return mi.cursor.current().(mapEntry).value
}

func (mi *mapIterator) Position() uint64 {
	if !mi.cursor.valid() {
		return 0
	}
	return uint64(mi.cursor.idx)
}

// Prev returns the previous entry from the Map. If there is no previous entry, Prev() returns nils.
func (mi *mapIterator) Prev() bool {
	if !mi.cursor.valid() {
		return false
	}
	return mi.cursor.retreat()
}

// Next returns the subsequent entries from the Map, starting with the entry at which the iterator
// was created. If there are no more entries, Next() returns nils.
func (mi *mapIterator) Next() bool {
	if !mi.cursor.valid() {
		return false
	}
	return mi.cursor.advance()
}
