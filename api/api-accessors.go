// Copyright 2020
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by gen-accessors; DO NOT EDIT.
package api

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (k *Key) GetID() string {
	if k == nil || k.ID == nil {
		return ""
	}
	return *k.ID
}

// GetKey returns the Key field if it's non-nil, zero value otherwise.
func (k *Key) GetKey() string {
	if k == nil || k.Key == nil {
		return ""
	}
	return *k.Key
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (k *Key) GetName() string {
	if k == nil || k.Name == nil {
		return ""
	}
	return *k.Name
}

// HasScopes checks if Key has any Scopes.
func (k *Key) HasScopes() bool {
	if k == nil || k.Scopes == nil {
		return false
	}
	if len(k.Scopes) == 0 {
		return false
	}
	return true
}

// HasScopes checks if KeyRequest has any Scopes.
func (k *KeyRequest) HasScopes() bool {
	if k == nil || k.Scopes == nil {
		return false
	}
	if len(k.Scopes) == 0 {
		return false
	}
	return true
}