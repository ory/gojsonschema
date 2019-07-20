package gojsonschema

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONLoaderLocal(t *testing.T) {
	for _, p := range []string{
		"file://./testdata/extra/fragment_schema.json",
		"file://testdata/extra/fragment_schema.json",
	} {
		t.Run("path="+p, func(t *testing.T) {
			_, err := NewSchema(NewReferenceLoader(p))
			require.NoError(t, err)
		})
	}
}
