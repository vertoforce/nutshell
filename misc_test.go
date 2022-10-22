package nutshell

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalName(t *testing.T) {
	s := `{"name":"test"}`
	var st Stub
	err := json.Unmarshal([]byte(s), &st)
	require.NoError(t, err)
	require.Equal(t, "test", st.Name.DisplayName)

	s = `{"name":{"displayName":"test2"}}`
	err = json.Unmarshal([]byte(s), &st)
	require.NoError(t, err)
	require.Equal(t, "test2", st.Name.DisplayName)
}
