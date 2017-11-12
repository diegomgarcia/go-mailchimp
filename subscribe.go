package mailchimp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/diegomgarcia/go-mailchimp/status"
)

// Subscribe ...
func (c *Client) Subscribe(listID string, email string, directSubscribe bool, mergeFields map[string]interface{}) (*MemberResponse, error) {

	subscribeStatus := status.Pending

	if directSubscribe {
		subscribeStatus = status.Subscribed
	}

	// Make request
	params := map[string]interface{}{
		"email_address": email,
		"status":        subscribeStatus,
		"merge_fields":  mergeFields,
	}
	resp, err := c.do(
		"POST",
		fmt.Sprintf("/lists/%s/members/", listID),
		&params,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Allow any success status (2xx)
	if resp.StatusCode/100 == 2 {
		// Unmarshal response into MemberResponse struct
		memberResponse := new(MemberResponse)
		if err := json.Unmarshal(data, memberResponse); err != nil {
			return nil, err
		}
		return memberResponse, nil
	}

	// Request failed
	errorResponse, err := extractError(data)
	if err != nil {
		return nil, err
	}
	return nil, errorResponse
}
