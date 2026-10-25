package spotify

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/FrankKair/spootify/track"
)

func TestParsePreservesMetadata(t *testing.T) {
	want := track.Track{
		Artist: "Artist|||Name",
		Album:  "Album \"Deluxe\" (Remastered)",
		Title:  "İstanbul\\Tokyo\n雪",
	}
	out, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := parse(append(out, '\n'))
	if err != nil || got != want {
		t.Fatalf("parse() = %#v, %v; want %#v", got, err, want)
	}
}

func TestParseInvalidMetadata(t *testing.T) {
	for _, input := range []string{
		"not JSON", "", "{}", "null",
		`{"artist":"Artist","title":" "}`,
		`{"artist":null,"title":"Song"}`,
	} {
		t.Run(input, func(t *testing.T) {
			if _, err := parse([]byte(input)); err == nil {
				t.Fatalf("parse(%q) should fail", input)
			}
		})
	}
}

func TestParseAllowsMissingAlbum(t *testing.T) {
	got, err := parse([]byte(`{"artist":"Artist","title":"Song"}`))
	if err != nil || got.Album != "" || got.Title != "Song" {
		t.Fatalf("parse() = %#v, %v", got, err)
	}
}

func commandFixture(t *testing.T, script string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "osascript")
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"+script+"\n"), 0o700); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestRunSeparatesStderr(t *testing.T) {
	executable := commandFixture(t, `printf '%s\n' '{"artist":"Artist","album":"Album (Deluxe)","title":"Song"}'
printf '%s\n' 'diagnostic message' >&2`)
	got, err := run(context.Background(), executable)
	if err != nil || got.Album != "Album (Deluxe)" {
		t.Fatalf("run() = %#v, %v", got, err)
	}
}

func TestRunIncludesExecutionError(t *testing.T) {
	executable := commandFixture(t, "echo 'Automation permission denied' >&2\nexit 1")
	_, err := run(context.Background(), executable)
	if err == nil || !strings.Contains(err.Error(), "Automation permission denied") || errors.Unwrap(err) == nil {
		t.Fatalf("expected wrapped error with AppleScript diagnostic, got %v", err)
	}
}

func TestRunMissingExecutable(t *testing.T) {
	_, err := run(context.Background(), filepath.Join(t.TempDir(), "missing"))
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected missing executable error, got %v", err)
	}
}

func TestRunTimeout(t *testing.T) {
	executable := commandFixture(t, "exec sleep 5")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	_, err := run(ctx, executable)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected timeout, got %v", err)
	}
}
