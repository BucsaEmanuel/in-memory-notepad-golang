# Hyperskill Go - In-Memory Notepad Project

- [Hyperskill Go - In-Memory Notepad project](#hyperskill-go-in-memory-notepad-project)
    - [About the project](#about-the-project)
        - [Status](#status)
        - [See also](#see-also)
    - [Getting started](#getting-started)

# Hyperskill Go - In-Memory Notepad project

## About the project

This is the implementation of the Go In-Memory Notepad project on Hyperskill.

### Status

The project is posted as it currently is, having passed the last step of the project ( 4 / 4 ).

### See also

* [Hyperskill - In-Memory Notepad (Go)](https://hyperskill.org/projects/238)

## Getting started

Either
`go build main.go`
then `./main`

or
`go run main.go`

## Examples
The greater-than symbol followed by a space (> ) represents the user input. Note that it's not part of the input.

### Example 1: below is the program execution of the `update` and `delete` commands followed by a few spaces.

```
Enter the maximum number of notes: > 3

Enter a command and data: > list
[Info] Notepad is empty

Enter a command and data: > create This is my first record!
[OK] The note was successfully created

Enter a command and data: > list
[Info] 1: This is my first record!

Enter a command and data: > update 1
[Error] Missing note argument

Enter a command and data: > update one Updated first note!
[Error] Invalid position: one

Enter a command and data: > update
[Error] Missing position argument

Enter a command and data: > update 1 Updated first note!
[OK] The note at position 1 was successfully updated

Enter a command and data: > list
[Info] 1: Updated first note!

Enter a command and data: > delete one
[Error] Invalid position: one

Enter a command and data: > delete
[Error] Missing position argument

Enter a command and data: > delete 1
[OK] The note at position 1 was successfully deleted

Enter a command and data: > list
[Info] Notepad is empty

Enter a command and data: > exit
[Info] Bye!
```

### Example 2: below is the program execution of the `create`, `update` and `delete` commands followed by a few spaces.

```
Enter the maximum number of notes: > 3

Enter a command and data: > create This is my first record!
[OK] The note was successfully created

Enter a command and data: > create This is my second record!
[OK] The note was successfully created

Enter a command and data: > list
[Info] 1: This is my first record!
[Info] 2: This is my second record!

Enter a command and data: > delete 1
[OK] The note at position 1 was successfully deleted

Enter a command and data: > list
[Info] 1: This is my second record!

Enter a command and data: > create This is my third record!
[OK] The note was successfully created

Enter a command and data: > create This is my forth record!
[OK] The note was successfully created

Enter a command and data: > update 2 Updated third record!
[OK] The note at position 2 was successfully updated

Enter a command and data: > list
[Info] 1: This is my second record!
[Info] 2: Updated third record!
[Info] 3: This is my forth record!

Enter a command and data: > exit
[Info] Bye!
```