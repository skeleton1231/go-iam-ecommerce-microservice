// Copyright 2020 Talhuang<talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

// Make sure to replace the import path with your actual package import path

//go:generate mockgen -self_package=github.com/marmotedu/iam/internal/apiserver/store -destination mock_store.go -package store github.com/marmotedu/iam/internal/apiserver/store Factory,UserStore,SecretStore,PolicyStore

var client Factory

// Factory defines the iam platform storage interface.
type Factory interface {
	Users() UserStore
	Secrets() SecretStore
	Policies() PolicyStore
	PolicyAudits() PolicyAuditStore
	Items() ItemStore
	ItemAttributes() ItemAttributesStore
	ItemImage() ItemImageStore
	Close() error
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}
