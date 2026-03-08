package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGeolocation(t *testing.T) {
	tests := []struct {
		name         string
		ipAddress    string
		countryCode  string
		country      string
		city         string
		latitude     string
		longitude    string
		mysteryValue string
		expectNil    bool
	}{
		{
			name:         "ValidInput",
			ipAddress:    "10.0.0.1",
			countryCode:  "SI",
			country:      "Nepal",
			city:         "DuBuquemouth",
			latitude:     "-84.87503094689836",
			longitude:    "7.206435933364332",
			mysteryValue: "7823011346",
			expectNil:    false,
		},
		{
			name:         "InvalidIP",
			ipAddress:    "10.0.0.",
			countryCode:  "SI",
			country:      "Nepal",
			city:         "DuBuquemouth",
			latitude:     "-84.87503094689836",
			longitude:    "7.206435933364332",
			mysteryValue: "7823011346",
			expectNil:    true,
		},
		{
			name:         "EmptyCountryCode",
			ipAddress:    "10.0.0.1",
			countryCode:  "",
			country:      "Nepal",
			city:         "DuBuquemouth",
			latitude:     "-84.87503094689836",
			longitude:    "7.206435933364332",
			mysteryValue: "7823011346",
			expectNil:    true,
		},
		{
			name:         "EmptyCountry",
			ipAddress:    "10.0.0.1",
			countryCode:  "SI",
			country:      "",
			city:         "DuBuquemouth",
			latitude:     "-84.87503094689836",
			longitude:    "7.206435933364332",
			mysteryValue: "7823011346",
			expectNil:    true,
		},
		{
			name:         "EmptyCity",
			ipAddress:    "10.0.0.1",
			countryCode:  "SI",
			country:      "Nepal",
			city:         "",
			latitude:     "-84.87503094689836",
			longitude:    "7.206435933364332",
			mysteryValue: "7823011346",
			expectNil:    true,
		},
		{
			name:         "LatitudeOutOfRange",
			ipAddress:    "10.0.0.1",
			countryCode:  "SI",
			country:      "Nepal",
			city:         "DuBuquemouth",
			latitude:     "-91.87503094689836",
			longitude:    "7.206435933364332",
			mysteryValue: "7823011346",
			expectNil:    true,
		},
		{
			name:         "EmptyLatitude",
			ipAddress:    "10.0.0.1",
			countryCode:  "SI",
			country:      "Nepal",
			city:         "DuBuquemouth",
			latitude:     "",
			longitude:    "7.206435933364332",
			mysteryValue: "7823011346",
			expectNil:    true,
		},
		{
			name:         "LongitudeOutOfRange",
			ipAddress:    "10.0.0.1",
			countryCode:  "SI",
			country:      "Nepal",
			city:         "DuBuquemouth",
			latitude:     "-84.87503094689836",
			longitude:    "181.206435933364332",
			mysteryValue: "7823011346",
			expectNil:    true,
		},
		{
			name:         "EmptyLongitude",
			ipAddress:    "10.0.0.1",
			countryCode:  "SI",
			country:      "Nepal",
			city:         "DuBuquemouth",
			latitude:     "-84.87503094689836",
			longitude:    "",
			mysteryValue: "7823011346",
			expectNil:    true,
		},
		{
			name:         "EmptyMysteryValue",
			ipAddress:    "10.0.0.1",
			countryCode:  "SI",
			country:      "Nepal",
			city:         "DuBuquemouth",
			latitude:     "-84.87503094689836",
			longitude:    "7.206435933364332",
			mysteryValue: "",
			expectNil:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewGeolocation(tc.ipAddress, tc.countryCode, tc.country, tc.city, tc.latitude, tc.longitude, tc.mysteryValue)
			if tc.expectNil {
				assert.Nil(t, g)
			} else {
				assert.NotNil(t, g)
			}
		})
	}
}
