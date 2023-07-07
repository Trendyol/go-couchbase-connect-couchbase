package gocouchbaseconnectcouchbase

import "github.com/Trendyol/go-couchbase-connect-couchbase/couchbase"

type Mapper func(event couchbase.Event) []couchbase.CBActionDocument

func DefaultMapper(event couchbase.Event) []couchbase.CBActionDocument {
	if event.IsMutated {
		return []couchbase.CBActionDocument{couchbase.NewIndexAction(event.Key, event.Value)}
	}
	return []couchbase.CBActionDocument{couchbase.NewDeleteAction(event.Key)}
}
