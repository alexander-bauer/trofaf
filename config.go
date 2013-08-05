package main

import (
	"encoding/json"
	"os"
)

const (
	DefaultOptionsPath = "options.json"
)

// Save uses encoding/json to marshal opt into a human-readable JSON
// format, and saves it to the file of the given path. If the file
// doesn't yet exist, it will be created.
func (opt *options) Save(path string) (err error) {
	// Open the file and truncate it exists, or create it with umask
	// 0666 if it doesn't.
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	// Here, we need to MarshalIndent to a temporary buffer, so that
	// the output file is readable.
	b, err := json.MarshalIndent(opt, "", "\t")
	if err != nil {
		return
	}

	// Finally, write all the bytes, and return any errors.
	_, err = f.Write(b)
	return err
}

// Load reads options from the given path in JSON format, and puts
// them in the given options type. If the file cannot be opened for
// reading, such as if it doesn't exist, there will be no error, and
// opt will remain unchanged.
func (opt *options) Load(path string) error {
	// Open the file for reading only.
	f, err := os.Open(path)
	if err != nil {
		// If it can't be opened, return no error.
		return nil
	}
	defer f.Close()

	// Return any errors when decoding. We pass opt here because it's
	// already a pointer to the options block that we want to operate
	// on. This way, we don't need to create a new options object and
	// return it.
	return json.NewDecoder(f).Decode(opt)
}
