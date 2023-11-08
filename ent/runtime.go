// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/ochanoco/torima/ent/requestlog"
	"github.com/ochanoco/torima/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	requestlogFields := schema.RequestLog{}.Fields()
	_ = requestlogFields
	// requestlogDescFlag is the schema descriptor for flag field.
	requestlogDescFlag := requestlogFields[3].Descriptor()
	// requestlog.DefaultFlag holds the default value on creation for the flag field.
	requestlog.DefaultFlag = requestlogDescFlag.Default.(string)
}
