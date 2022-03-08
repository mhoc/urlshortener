package util

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNewID_Sanity(t *testing.T) {
	ids := []string{NewID(), NewID(), NewID(), NewID(), NewID()}
	for i, id := range ids {
		if len(id) != 8 {
			t.Errorf("expected id to have length 8, but got: %v", id)
		}
		for _, id2 := range ids[i+1:] {
			if id == id2 {
				t.Errorf("expected each generated id to be different, but two were the same")
			}
		}
	}
}

func TestIDToShortlink(t *testing.T) {
	id := NewID()
	exampleUrl, _ := url.Parse("https://example.com")
	shortlink := IDToShortlink(exampleUrl, id)
	if shortlink != fmt.Sprintf("https://example.com/%v", id) {
		t.Errorf("expected id-to-shortlink to result in a sane url, but instead got: %v", shortlink)
	}
}

func TestShortlinkToID(t *testing.T) {
	shortlink := "https://example.com/12345abc"
	id, err := ShortlinkToID(shortlink)
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
	if id != "12345abc" {
		t.Errorf("expected shortlink-to-id to result in '12345abc', but instead got: %v", id)
	}
}
