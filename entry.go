package gonx

import (
	"fmt"
	"strconv"
)

// Shortcut for the map of strings
type Fields map[string]string

// Parsed log record. Use Get method to retrieve a value by name instead of
// threating this as a map, because inner representation is in design.
type Entry struct {
	fields Fields
}

// Creates an empty Entry to be filled later
func NewEmptyEntry() *Entry {
	return &Entry{make(Fields)}
}

// Creates an Entry with fiven fields
func NewEntry(fields Fields) *Entry {
	return &Entry{fields}
}

// Return entry field value by name or empty string and error if it
// does not exist.
func (entry *Entry) Field(name string) (value string, err error) {
	value, ok := entry.fields[name]
	if !ok {
		err = fmt.Errorf("Field '%v' does not found in record %+v", name, *entry)
	}
	return
}

// Return entry field value as float64. Rutuen nil if field does not exists
// and convertion error if cannot cast a type.
func (entry *Entry) FloatField(name string) (value float64, err error) {
	tmp, err := entry.Field(name)
	if err == nil {
		value, err = strconv.ParseFloat(tmp, 64)
	}
	return
}

// Field value setter
func (entry *Entry) SetField(name string, value string) {
	entry.fields[name] = value
}

// Float field value setter. It accepts float64, but still store it as a
// string in the same fields map. The precision is 2, its enough for log
// parsing task
func (entry *Entry) SetFloatField(name string, value float64) {
	entry.SetField(name, strconv.FormatFloat(value, 'f', 2, 64))
}

// Integer field value setter. It accepts float64, but still store it as a
// string in the same fields map.
func (entry *Entry) SetUintField(name string, value uint64) {
	entry.SetField(name, strconv.FormatUint(uint64(value), 10))
}
