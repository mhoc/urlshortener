package store

import (
	"testing"
	"time"
)

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

func TestInMemoryStore_CreateMulti_WithExpiration(t *testing.T) {
	ims := NewInMemoryStore()
	_, err := ims.Create("https://example.com/1", 60*time.Second)
	if err != nil {
		t.Error(err)
	}
	_, err = ims.Create("https://example.com/2", 60*time.Second)
	if err != nil {
		t.Error(err)
	}
	_, err = ims.Create("https://example.com/3", 120*time.Second)
	if err != nil {
		t.Error(err)
	}
	_, err = ims.Create("https://example.com/4", 140*time.Second)
	if err != nil {
		t.Error(err)
	}
	_, err = ims.Create("https://example.com/5", 2*time.Second)
	if err != nil {
		t.Error(err)
	}
}

func TestInMemoryStore_Expiration(t *testing.T) {
	ims := NewInMemoryStore()
	idShouldBeDeleted, err := ims.Create("https://example.com/1", 1*time.Second)
	if err != nil {
		t.Error(err)
	}
	idShouldRemain, err := ims.Create("https://example.com/2", 60*time.Second)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(3 * time.Second)
	redirectToUrl, err := ims.Get(idShouldBeDeleted)
	if err != nil {
		t.Error(err)
	}
	if redirectToUrl != "" {
		t.Error("expected the provided key to be expired, but it wasn't")
	}
	redirectToUrl, err = ims.Get(idShouldRemain)
	if err != nil {
		t.Error(err)
	}
	if redirectToUrl == "" {
		t.Error("expected the provided key to not be expired, but it wasn't")
	}
}
