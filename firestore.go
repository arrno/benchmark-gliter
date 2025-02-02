package main

import (
	"context"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// makeDB - Obscure the db implementation
func makeDB() Database {
	f, err := NewFirestore("./keys/firestore.json")
	expectNil(err)
	return f
}

type Firestore struct {
	client *firestore.Client
	ctx    context.Context
}

func NewFirestore(credPath string) (*Firestore, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	f := Firestore{
		client: client,
		ctx:    ctx,
	}
	return &f, nil
}

func (f *Firestore) SetBatch(data []DocBundle) {
	bw := f.client.BulkWriter(f.ctx)
	results := make([]*firestore.BulkWriterJob, len(data))
	for i, db := range data {
		res, err := bw.Set(f.client.Doc(f.parsePath(db.Path)), db.Data)
		expectNil(err)
		results[i] = res
	}
	bw.End()
	for _, res := range results {
		_, err := res.Results()
		expectNil(err)
	}
}

func (f *Firestore) FetchBatch(path string) []DocBundle {
	results := make([]DocBundle, 100)
	iter := f.client.Collection(f.parsePath(path)).Limit(100).Documents(f.ctx)
	idx := 0
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		expectNil(err)
		var d MyData
		doc.DataTo(&d)
		results[idx] = DocBundle{Path: f.parsePath(doc.Ref.Path), Data: d}
		idx++
	}
	return results
}

func (f *Firestore) DeleteCollection(path string) {
	col := f.client.Collection(f.parsePath(path))
	bulkwriter := f.client.BulkWriter(f.ctx)

	for {
		iter := col.Limit(100).Documents(f.ctx)
		numDeleted := 0

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			expectNil(err)

			bulkwriter.Delete(doc.Ref)
			numDeleted++
		}

		if numDeleted == 0 {
			bulkwriter.End()
			break
		}

		bulkwriter.Flush()
	}
}

func (f *Firestore) parsePath(path string) string {
	tokens := strings.Split(path, "(default)/documents/")
	if len(tokens) == 1 {
		return path
	}
	return tokens[len(tokens)-1]
}
