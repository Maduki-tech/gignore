package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

var sut model

func TestMain(t *testing.M) {
	sut = initialModel()
}

func testInitialModel(t *testing.T) {
	sut = initialModel()

	if sut.choices == nil {
		t.Fatal("Expected choices to be not nil")
	}

	if sut.cursor != 0 {
		t.Fatal("Expected cursor to be 0, got ", sut.cursor)
	}
}

func testInit(t *testing.T) {
	sut = initialModel()
	cmd := sut.Init()

	if cmd != nil {
		t.Fatal("Expected cmd to be nil")
	}
}

func testUpdate(t *testing.T) {
	msg := tea.Msg("enter")
	model, cmd := sut.Update(msg)

	if model != nil {
		t.Fatal("Expected model to be the same")
	}

	if cmd != nil {
		t.Fatal("Expected cmd to be nil")
	}
}

func testView(t *testing.T) {
	view := sut.View()

	if view == "" {
		t.Fatal("Expected view to be not empty")
	}

	if view != "Java\nGo\nPython\nJavaScript\nTypeScript\n" {
		t.Fatal("Expected view to be Java\nGo\nPython\nJavaScript\nTypeScript\n, got ", view)
	}
}
