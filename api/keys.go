package api

import "github.com/davidji99/simpleresty"

// Keys handles communication with the API keys related
// methods of the Sendgrid APIv3.
type KeysService service

// Key represents an API key.
type Key struct {
	ID     *string  `json:"api_key_id,omitempty"`
	Key    *string  `json:"api_key,omitempty"`
	Name   *string  `json:"name,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

// KeyCreateRequest represents a request to create an API key.
type KeyRequest struct {
	Name   string   `json:"name,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	Sample string   `json:"sample,omitempty"`
}

// List all API keys.
func (k *KeysService) List() ([]*Key, *simpleresty.Response, error) {
	var result []*Key
	urlStr := k.client.http.RequestURL("/api_keys")

	// Execute the request
	response, getErr := k.client.http.Get(urlStr, &result, nil)

	return result, response, getErr
}

// Get a single API key.
func (k *KeysService) Get(id string) (*Key, *simpleresty.Response, error) {
	var result *Key
	urlStr := k.client.http.RequestURL("/api_keys/%s", id)

	// Execute the request
	response, err := k.client.http.Get(urlStr, &result, nil)

	return result, response, err
}

// Create an API key.
func (k *KeysService) Create(opts *KeyRequest) (*Key, *simpleresty.Response, error) {
	var result *Key
	urlStr := k.client.http.RequestURL("/api_keys")

	// Execute the request
	response, err := k.client.http.Post(urlStr, &result, opts)

	return result, response, err
}

// UpdateName updates only the name of an existing API Key.
func (k *KeysService) UpdateName(id, name string) (*Key, *simpleresty.Response, error) {
	var result *Key
	urlStr := k.client.http.RequestURL("/api_keys/%s", id)

	opts := struct {
		Name string `json:"name,omitempty"`
	}{Name: name}

	// Execute the request
	response, err := k.client.http.Patch(urlStr, &result, opts)

	return result, response, err
}

// UpdateNameScopes updates only the name and scopes of an existing API Key.
//
// Name parameter is required even if only changing scopes. Scope parameter is required
// even if you're just changing name.
func (k *KeysService) UpdateNameScopes(id, name string, scopes []string) (*Key, *simpleresty.Response, error) {
	var result *Key
	urlStr := k.client.http.RequestURL("/api_keys/%s", id)

	opts := struct {
		Name   string   `json:"name,omitempty"`
		Scopes []string `json:"scopes,omitempty"`
	}{
		Name: name, Scopes: scopes,
	}

	// Execute the request
	response, err := k.client.http.Put(urlStr, &result, opts)

	return result, response, err
}

// Delete an existing API key.
func (k *KeysService) Delete(id string) (*simpleresty.Response, error) {
	urlStr := k.client.http.RequestURL("/api_keys/%s", id)

	// Execute the request
	response, err := k.client.http.Delete(urlStr, nil, nil)

	return response, err
}
