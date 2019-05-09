package behavior

import (
	"github.com/havoc-io/mutagen/pkg/filesystem"
)

// DecomposesUnicodeByPath determines whether or not the filesystem on which the
// directory at the specified path resides decomposes Unicode filenames. On
// Windows this function always returns false since Windows filesystems preserve
// Unicode filename normalization.
func DecomposesUnicodeByPath(_ string, probeMode ProbeMode) (bool, error) {
	// Check for invalid probe modes.
	if !probeMode.Supported() {
		panic("invalid probe mode")
	}

	// Return the well-known behavior.
	return false, nil
}

// DecomposesUnicode determines whether or not the specified directory (and its
// underlying filesystem) decomposes Unicode filenames. On Windows this function
// always returns false since Windows filesystems preserve Unicode filename
// normalization.
func DecomposesUnicode(_ *filesystem.Directory, probeMode ProbeMode) (bool, error) {
	// Check for invalid probe modes.
	if !probeMode.Supported() {
		panic("invalid probe mode")
	}

	// Return the well-known behavior.
	return false, nil
}