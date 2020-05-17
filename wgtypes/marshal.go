package wgtypes

import (
	"encoding/base64"
	"encoding/json"
)

// MarshalJSON marshals a Key
func (k Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(base64.StdEncoding.EncodeToString(k[:]))
}

// UnmarshalJSON unmarshals a Key
func (k Key) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	k, err = NewKey([]byte(s))
	return err
}
