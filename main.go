package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var cmd = Command{}
var commandProcess = Runner{}
var storage = Storage{capacity: 5}

type Runner struct{}

type Command struct {
	command string
	data    string
}

type Storage struct {
	capacity int
	notes    []Note
}

type Note struct {
	data string
}

func (*Runner) run() {
	cmd.initializeStorage()
	for {
		cmd.getCommand()
		cmd.processCommand()
		if cmd.command == "exit" {
			return
		}
	}
}

func (storage *Storage) setCapacity(newCapacity int) {
	storage.capacity = newCapacity
}

func (storage *Storage) listNotes() {
	for index, note := range storage.notes {
		index++
		if !note.isEmpty() {
			fmt.Printf("[Info] %d: %s\n", index, note.data)
		}
	}
}

func (storage *Storage) hasNoteAtPosition(index int) bool {
	if index <= len(storage.notes)-1 {
		return true
	}
	return false
}

func (storage *Storage) clearNotes() {
	storage.notes = []Note{}
}

func (storage *Storage) hasNotes() bool {
	return len(storage.notes) > 0
}

func (storage *Storage) isFull() bool {
	return len(storage.notes) == storage.capacity
}

func (storage *Storage) addNote(note Note) {
	storage.notes = append(storage.notes, note)
}

func (cmd *Command) initializeStorage() {
	fmt.Print("Enter the maximum number of notes: ")
	var newCapacity int
	_, _ = fmt.Scan(&newCapacity)
	storage.setCapacity(newCapacity)
}

func (cmd *Command) createNote() {
	if storage.isFull() {
		fmt.Println("[Error] Notepad is full")
		return
	}

	note := Note{data: cmd.data}
	if note.isEmpty() {
		fmt.Println("[Error] Missing note argument")
		return
	}

	storage.addNote(Note{data: cmd.data})
	fmt.Println("[OK] The note was successfully created")
}

func (cmd *Command) deleteNote() {
	if cmd.data == "" {
		fmt.Println("[Error] Missing position argument")
		return
	}

	noteIndex, err := strconv.Atoi(cmd.data)
	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", cmd.data)
		return
	}

	if noteIndex < 1 || noteIndex > storage.capacity {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", noteIndex, storage.capacity)
	}
	noteIndex--

	if storage.hasNoteAtPosition(noteIndex) {
		storage.notes = append(storage.notes[:noteIndex], storage.notes[noteIndex+1:]...)
		noteIndex++
		fmt.Printf("[OK] The note at position %d was successfully deleted\n", noteIndex)
	} else {
		fmt.Println("[Error] There is nothing to delete")
		return
	}
}

func (cmd *Command) updateNote() {
	input := strings.Split(cmd.data, " ")

	notePseudoIndex, updatedNoteData := strings.TrimSpace(input[0]), strings.TrimSpace(strings.Join(input[1:], " "))

	if notePseudoIndex == "" {
		fmt.Println("[Error] Missing position argument")
		return
	}

	if updatedNoteData == "" {
		fmt.Println("[Error] Missing note argument")
		return
	}

	noteIndex, err := strconv.Atoi(notePseudoIndex)
	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", notePseudoIndex)
		return
	}

	if noteIndex < 1 || noteIndex > storage.capacity {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", noteIndex, storage.capacity)
		return
	}
	noteIndex--

	updatedNote := Note{data: updatedNoteData}
	if storage.hasNoteAtPosition(noteIndex) {
		storage.notes[noteIndex] = updatedNote
		noteIndex++
		fmt.Printf("[OK] The note at position %d was successfully updated\n", noteIndex)
	} else {
		fmt.Println("[Error] There is nothing to update")
		return
	}
}

func (cmd *Command) clearNotes() {
	if storage.hasNotes() {
		storage.clearNotes()
	}
	fmt.Println("[OK] All notes were successfully deleted")
}

func (cmd *Command) listNotes() {
	if storage.hasNotes() {
		storage.listNotes()
	} else {
		fmt.Println("[Info] Notepad is empty")
	}
}

func (cmd *Command) getCommand() {
	fmt.Print("Enter a command and data:")
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	cmd.command, cmd.data = strings.TrimSpace(input[0]), strings.TrimSpace(strings.Join(input[1:], " "))
}

func (cmd *Command) processCommand() {
	switch cmd.command {
	case "exit":
		fmt.Print("[Info] Bye!\n")
		return
	case "create":
		cmd.createNote()
	case "update":
		cmd.updateNote()
	case "delete":
		cmd.deleteNote()
	case "list":
		cmd.listNotes()
	case "clear":
		cmd.clearNotes()
	default:
		fmt.Print("[Error] Unknown command\n")
	}
}

func (note *Note) isEmpty() bool {
	return strings.TrimSpace(note.data) == ""
}

func main() {
	commandProcess.run()
}
