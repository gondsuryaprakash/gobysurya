package test

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
)

func TestParseFilePattern(t *testing.T) {

	testCases := []struct {
		scenario           string
		pattern            string
		expextedApiKey     string
		expectedTimestamp  string
		expectedFileset    string
		expectedUniqueID   string
		expectedPosition   string
		expectedDivisionID string
		expectedError      bool
	}{
		{
			scenario:           "valid pattern with all fields",
			pattern:            "new_division_all_bf47e9ff08c583bbce3572317f2df68611172069_20240601052223_auto_b757675b4ecd4bd79667ad864e96cdae-0",
			expextedApiKey:     "bf47e9ff08c583bbce3572317f2df68611172069",
			expectedTimestamp:  "20240601052223",
			expectedFileset:    "auto",
			expectedUniqueID:   "b757675b4ecd4bd79667ad864e96cdae-0",
			expectedPosition:   "0",
			expectedDivisionID: "all",
			expectedError:      false,
		},
		{
			scenario:           "valid Pattern without division ID",
			pattern:            "new_bf47e9ff08c583bbce3572317f2df68611172069_20240601052223_auto_b757675b4ecd4bd79667ad864e96cdae-0",
			expextedApiKey:     "bf47e9ff08c583bbce3572317f2df68611172069",
			expectedTimestamp:  "20240601052223",
			expectedFileset:    "auto",
			expectedUniqueID:   "b757675b4ecd4bd79667ad864e96cdae-0",
			expectedPosition:   "0",
			expectedDivisionID: "",
			expectedError:      false,
		},
		{
			scenario:           "valid pattern without position",
			pattern:            "new_division_all_bf47e9ff08c583bbce3572317f2df68611172069_20240601052223_auto_b757675b4ecd4bd79667ad864e96cdae",
			expextedApiKey:     "bf47e9ff08c583bbce3572317f2df68611172069",
			expectedTimestamp:  "20240601052223",
			expectedFileset:    "auto",
			expectedUniqueID:   "b757675b4ecd4bd79667ad864e96cdae",
			expectedPosition:   "0",
			expectedDivisionID: "all",
			expectedError:      false,
		},
		{
			scenario:           "invalid pattern",
			pattern:            "invalid_pattern",
			expextedApiKey:     "",
			expectedTimestamp:  "",
			expectedFileset:    "",
			expectedUniqueID:   "",
			expectedPosition:   "",
			expectedDivisionID: "",
			expectedError:      true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.scenario, func(t *testing.T) {
			apikey, timestamp, fileset, uniqueID, position, divisionID, err := ParseFilePattern(tt.pattern)
			if tt.expectedError {
				if err == nil {
					t.Errorf("expected error but not got error")
				}
				return
			}
			if err != nil {
				t.Errorf("didn't expected error %v", err)
			}

			if tt.expectedUniqueID == "" {
				h := sha1.New()
				_, _ = io.WriteString(h, tt.pattern)
				tt.expectedUniqueID = fmt.Sprintf("%x", h.Sum(nil))
			}

			if apikey != tt.expextedApiKey {
				t.Errorf("expected apikey %s, got %s", tt.expextedApiKey, apikey)
			}
			if timestamp != tt.expectedTimestamp {
				t.Errorf("expected timestamp %s, got %s", tt.expectedTimestamp, timestamp)
			}
			if fileset != tt.expectedFileset {
				t.Errorf("expected fileset %s, got %s", tt.expectedFileset, fileset)
			}
			if uniqueID != tt.expectedUniqueID {
				t.Errorf("expected uniqueID %s, got %s", tt.expectedUniqueID, uniqueID)
			}
			if position != tt.expectedPosition {
				t.Errorf("expected position %s, got %s", tt.expectedPosition, position)
			}
			if divisionID != tt.expectedDivisionID {
				t.Errorf("expected divisionID %s, got %s", tt.expectedDivisionID, divisionID)
			}
		})

	}

}
