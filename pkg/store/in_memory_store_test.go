package store

import "testing"

func TestInMemoryStore_CreateAndGetBasic(t *testing.T) {
	ims := NewInMemoryStore()
	expectedUrl := "https://example.com/"
	id, err := ims.Create(expectedUrl, 0)
	if err != nil {
		t.Error(err)
	}
	retrievedUrl, err := ims.Get(id)
	if err != nil {
		t.Error(err)
	}
	if expectedUrl != retrievedUrl {
		t.Errorf("urls did not match: '%v' != '%v'", expectedUrl, retrievedUrl)
	}
}

func TestInMemoryStore_CreateAndGet_Duplicate(t *testing.T) {
	ims := NewInMemoryStore()
	expectedUrl := "https://example.com/2"
	id1, err := ims.Create(expectedUrl, 0)
	if err != nil {
		t.Error(err)
	}
	id2, err := ims.Create(expectedUrl, 0)
	if err != nil {
		t.Error(err)
	}
	if id1 != id2 {
		t.Errorf("in memory store generated duplicate ids for same URL: %v %v", id1, id2)
	}
}

func TestInMemoryStore_Get_Nonexistant(t *testing.T) {
	ims := NewInMemoryStore()
	redirectUrl, err := ims.Get("not_a_real_key")
	if err != nil {
		t.Error(err)
	}
	if redirectUrl != "" {
		t.Errorf("in memory store did not return empty string for nonexistant key: %v", redirectUrl)
	}
}

func TestInMemoryStore_CreateAndRemove_Basic(t *testing.T) {
	ims := NewInMemoryStore()
	expectedUrl := "https://example.com/"
	id, err := ims.Create(expectedUrl, 0)
	if err != nil {
		t.Error(err)
	}
	removed, err := ims.Remove(id)
	if err != nil {
		t.Error(err)
	}
	if !removed {
		t.Errorf("expected removal of key extant in in-memory store, but store reports key did not exist")
	}
	retrievedUrl, err := ims.Get(id)
	if err != nil {
		t.Error(err)
	}
	if retrievedUrl != "" {
		t.Errorf("expected removal of key extant in in-memory store, but key still exists in the store")
	}
}

func TestInMemoryStore_Remove_Nonexisting(t *testing.T) {
	ims := NewInMemoryStore()
	removed, err := ims.Remove("nonexistant_key")
	if err != nil {
		t.Error(err)
	}
	if removed {
		t.Errorf("expected negative removal of key nonexistent in in-memory store, but store reports key did exist")
	}
}
