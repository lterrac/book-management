package cmd

import (
	"book-management/pkg/apis"
	"book-management/pkg/book-cli/pkg/options"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// RunCommand performs http request to the book server for all operations.
func RunCommand(opts *options.CommandOptions) error {
	httpClient := http.Client{
		Timeout: 20 * time.Second,
	}

	req, err := http.NewRequest(string(opts.Operation), opts.URL(), bytes.NewBuffer([]byte(opts.Object)))

	if err != nil {
		return fmt.Errorf("something went wrong while creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("Sending request to %v\n", opts.URL())
	res, err := httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("something went wrong while creating resource: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return fmt.Errorf("something went wrong while reading response: %v", err)
	}

	responseBody := apis.Message{}
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		return fmt.Errorf("something went wrong while unmarshalling response: %v", err)
	}

	displayResponse(opts, responseBody)
	return nil
}

func displayResponse(opts *options.CommandOptions, responseBody apis.Message) {
	fmt.Printf("%v %v: %v\n", opts.Operation, opts.Resource, responseBody)
}
