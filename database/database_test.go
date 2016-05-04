package db

import (
	"testing"
	"time"
	"fmt"
)

var DBCONN, _ = OpenConnection()

func TestOpenConn(t *testing.T) {
	dbConn, err := OpenConnection()
	if err != nil {
		t.Fatalf("%v", err)
	}
	InitDB(dbConn)
}

func TestAddWhitelistedTorrent(t *testing.T) {
	newTorrent := &Torrent{
		InfoHash:   "12345123451234512345",
		Name:       "Hello Kitty Island Adventure.exe",
		Downloaded: 0,
		Seeders:    0,
		Leechers:   0,
		AddedBy:    "127.0.0.1",
		DateAdded:  time.Now().Unix(),
	}

	newTorrent.AddWhitelistedTorrent()

	retval, err := GetTorrent(newTorrent.InfoHash)
	if err != nil {
		t.Fatalf("Failed to GetTorrent")
	}

	if newTorrent.InfoHash != retval.InfoHash {
		t.Fatalf("Expected %v, got %v", newTorrent.InfoHash, retval.InfoHash)
	}

	if newTorrent.DateAdded != retval.DateAdded {
		t.Fatalf("Expected %v, got %v", newTorrent.DateAdded, retval.DateAdded)
	}

	if newTorrent.Name != retval.Name {
		t.Fatalf("Expected %v, got %v", newTorrent.Name, retval.Name)
	}
}

func TestGetWhitelistedTorrents(t *testing.T) {
	retval, err := GetWhitelistedTorrents()
	if err != nil {
		t.Fatalf("Failed to get all whitelisted torrents")
	}

	fmt.Printf("%v", retval)
}

func TestGetWhitelistedTorrent(t *testing.T) {
	newTorrent := &Torrent{
		InfoHash:   "123451234512345123451234",
		Name:       "Hello Kitty Island Adventure.exe",
		Downloaded: 0,
		Seeders:    0,
		Leechers:   0,
		AddedBy:    "127.0.0.1",
		DateAdded:  time.Now().Unix(),
	}

	retval, err := GetWhitelistedTorrent(newTorrent.InfoHash)
	if err != nil {
		t.Fatalf("Failed to GetWhitelistedTorrent")
	}

	if retval.InfoHash != newTorrent.InfoHash {
		t.Fatalf("Expected %v, got %v", retval.InfoHash,
			newTorrent.InfoHash)
	}
}