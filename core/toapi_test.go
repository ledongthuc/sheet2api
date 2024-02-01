package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_getHeader(t *testing.T) {
	rows := []string{"First Name", "Last Name", "Age"}
	actualHeaders := getHeader(rows)

	assert.Equal(t, rows[0], actualHeaders[0])
	assert.Equal(t, rows[1], actualHeaders[1])
	assert.Equal(t, rows[2], actualHeaders[2])
}

func TestGetRows(t *testing.T) {
	path := "../test/users_test.xlsx"
	require.FileExists(t, path)

	rows, err := GetRows(path, "users")
	require.NoError(t, err)

	expectedRows := []map[string]interface{}{
		{
			"First Name": "Dulce",
			"Last Name":  "Abril",
			"Gender":     "Female",
			"Country":    "United States",
			"Age":        "32",
			"Date":       "15/10/2017",
			"Id":         "1562",
		},
		{
			"First Name": "Mara",
			"Last Name":  "Hashimoto",
			"Gender":     "Female",
			"Country":    "Great Britain",
			"Age":        "25",
			"Date":       "16/08/2016",
			"Id":         "1582",
		},
		{
			"First Name": "Philip",
			"Last Name":  "Gent",
			"Gender":     "Male",
			"Country":    "France",
			"Age":        "36",
			"Date":       "21/05/2015",
			"Id":         "2587",
		},
	}

	assert.Equal(t, expectedRows, rows)
}
