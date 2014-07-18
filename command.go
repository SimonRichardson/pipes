package main

type Command interface {
	Execute(Note) Note
}
