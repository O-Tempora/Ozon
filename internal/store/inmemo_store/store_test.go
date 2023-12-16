package inmemostore

import (
	"context"
	"testing"
)

var store *InmemStore = CreateInmemoStore()

func TestGetUrlExists(t *testing.T) {
	var long = "https://phoenixnap.com/kb/linux-network-commands"
	var short = "12345abCD_"
	store.urls.Store(long, struct{}{})
	store.pairs.Store(short, long)

	url, err := store.GetOriginalURL(context.Background(), short)
	if err != nil {
		t.Fatalf("get original url error: %s", err.Error())
	}
	if url != long {
		t.Errorf("Invalid get url result for %s: want %s, got %s", short, long, url)
	}
	t.Cleanup(func() {
		store.urls.Delete(long)
		store.pairs.Delete(short)
	})
}

func TestGetUrlNotExist(t *testing.T) {
	var long = "https://phoenixnap.com/kb/linux-network-commands"
	var short = "12345abCD_"
	var wrongShort = "69420"
	store.urls.Store(long, struct{}{})
	store.pairs.Store(short, long)

	url, err := store.GetOriginalURL(context.Background(), wrongShort)
	if url != "" {
		t.Fatalf("get original url error: %s", err.Error())
	}
	if err == nil {
		t.Errorf("Invalid get url result for %s: want non-nil err, got nil", wrongShort)
	}
	t.Cleanup(func() {
		store.urls.Delete(long)
		store.pairs.Delete(short)
	})
}
