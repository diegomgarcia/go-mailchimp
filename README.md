# go-mailchimp

A Golang SDK for Mailchimp API v3.

[![Travis Status for RichardKnop/go-mailchimp](https://travis-ci.org/RichardKnop/go-mailchimp.svg?branch=master)](https://travis-ci.org/RichardKnop/go-mailchimp)
[![godoc for RichardKnop/go-mailchimp](https://godoc.org/github.com/nathany/looper?status.svg)](http://godoc.org/github.com/RichardKnop/go-mailchimp)
[![goreportcard for RichardKnop/go-mailchimp](https://goreportcard.com/badge/github.com/RichardKnop/go-mailchimp)](https://goreportcard.com/report/RichardKnop/go-mailchimp)
[![codecov for RichardKnop/go-mailchimp](https://codeship.com/projects/fdac3010-3acd-0134-e9c5-06456b66cf53/status?branch=master)](https://codeship.com/projects/166426)
[![Codeship Status for RichardKnop/go-mailchimp](https://app.codeship.com/projects/35dc5880-71a7-0133-ec05-06b1c29ec1d7/status?branch=master)](https://app.codeship.com/projects/116961)

[![Sourcegraph for RichardKnop/go-mailchimp](https://sourcegraph.com/github.com/RichardKnop/go-mailchimp/-/badge.svg)](https://sourcegraph.com/github.com/RichardKnop/go-mailchimp?badge)
[![Donate Bitcoin](https://img.shields.io/badge/donate-bitcoin-orange.svg)](https://richardknop.github.io/donate/)

## Usage

```go
package main

import (
	"log"

	"github.com/diegomgarcia/go-mailchimp"
)

func main() {
	client, err := mailchimp.NewClient("the_api_key-us13", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the email is already subscribed
	memberResponse, err := client.CheckSubscription(
		"listID",
		"john@reese.com",
	)

	// User is already subscribed, update the subscription
	if err == nil {
		memberResponse, err = client.UpdateSubscription(
			"listID",
			"john@reese.com",
			map[string]interface{}{},
		)

		if err != nil {
			// Check the error response
			errResponse, ok := err.(*mailchimp.ErrorResponse)

			// Could not type assert error response
			if !ok {
				log.Fatal(err)
			}

			log.Fatal(errResponse)
		}

		log.Printf(
			"%s's subscription has been updated. Status: %s",
			memberResponse.EmailAddress,
			memberResponse.Status,
		)
		return
	}

	if err != nil {
		// Check the error response
		errResponse, ok := err.(*mailchimp.ErrorResponse)

		// Could not type assert error response
		if !ok {
			log.Fatal(err)
		}

		// 404 means we can process and subscribe user,
		// error other than 404 means we return error
		if errResponse.Status != http.StatusNotFound {
			log.Fatal(errResponse)
		}
	}

	// Indicates that this address should be marked as pending and request a confirmation for the user
	// to be added to the list, if you want to direct add to the list without confirmation set to true.
	directSubscribe := false; 
	
	// Subscribe the email
	memberResponse, err = client.Subscribe(
		"listID",
		"john@reese.com",
		directSubscribe,
		map[string]interface{}{},
	)

	if err != nil {
		// Check the error response
		errResponse, ok := err.(*mailchimp.ErrorResponse)

		// Could not type assert error response
		if !ok {
			log.Fatal(err)
		}

		log.Fatal(errResponse)
	}

	log.Printf(
		"%s has been subscribed successfully. Status: %s",
		memberResponse.EmailAddress,
		memberResponse.Status,
	)
}
```
