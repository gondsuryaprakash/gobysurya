package test

import (
	"crypto/sha1"
	"fmt"
	"io"
	"regexp"
)

var reFilePattern = regexp.MustCompile(`^(?:new_division_([a-zA-Z0-9]+)_|new_)?([0-9a-f]{40})_(\d{14})_(\S+?)(?:_([a-z0-9]+)-(\d+))?_(\S+).*$`)

func ParseFilePattern(pattern string) (apikey, timestamp, fileset, uniqueID, position, divisionID string, err error) {

	match := reFilePattern.FindStringSubmatch(pattern)
	if match == nil || len(match) < 8 {
		return "", "", "", "", "", "", errFilePattern(pattern)
	}
	if match[5] == "" {
		h := sha1.New()

		_, err = io.WriteString(h, pattern)
		if err != nil {
			return "", "", "", "", "", "", err
		}

		match[5] = fmt.Sprintf("%x", h.Sum(nil))
	}

	if match[6] == "" {
		match[6] = "0"
	}

	// The divisionID is captured in match[1] when present, otherwise it should be empty
	divisionID = match[1]

	return match[2], match[3], match[4], match[7], match[6], divisionID, nil
}

var errFilePattern = func(file string) error {
	return fmt.Errorf("file %s does not match the file pattern '(api-key)_(timestamp)_(file-set)_.*'", file)
}
