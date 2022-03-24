// Carly Hutson
// Lab 3b
// GET account/profile

package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountProfile struct {
	Member      bool   `json:"member"`
	Credits     int    `json:"credits"`
	DisplayName string `json:"display_name"`
	Created     string `json:"created"`
}

func (s *Client) AccountProfile() (*AccountProfile, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret AccountProfile
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
