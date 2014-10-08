pipes
=====

## Commands

Monadic commands

A very simple, but powerful monadic command runner, where commands declare 
functions to be run.

Since the Go language lacks defining total immutable state, it's thought of
as best practice. If using the command runner the commands are created once, 
so any state created in there could effect the next run, so be very careful.

###  Examples

The [example](example.go) is a lot more up to date by it's very nature of being 
code, but to give you an idea, see the following:

Create a set of commands that we want to use:

```go
type AddCommand struct {
    x Int
}

func NewAddCommand() AddCommand {
    return AddCommand{
        x: NewInt(1),
    }
}

func (c AddCommand) Execute(note pipes.Note) pipes.CommandResult {
    return pipes.ContinueResult(note.(Conf).Map(func(x Val) Val {
        return NewVal(x.Sum().Run(c.x).(Int))
    }))
}
```

Create a successful program:

```
conf := NewConf(NewVal(NewInt(0)))

program := pipes.EitherT{}.Of(conf).
    Eff(pipes.Do(NewAddCommand())).
    Eff(pipes.Do(NewAddCommand())).
    Eff(pipes.Do(NewAddCommand()))
program.Run.Run()
// (Right (3), [Right (0) Right (1) Right (2)])
```

For a unsuccessful program:

```
conf := NewConf(NewVal(NewInt(0)))

program := pipes.EitherT{}.Of(conf).
    Eff(pipes.Do(NewAddCommand())).
    Eff(pipes.Do(NewBadCommand())).
    Eff(pipes.Do(NewAddCommand()))
program.Run.Run()
// (Left (1), [Right (0) Right (1) Left (1)])
```

### Notes

Although some of the structs are named after other functional language types 
(Either, etc), they are not in fact the same. They're close, but each type has
been locked down to a tighter set of types, to prevent the need for type casting.
Which means they'll not pass any monadic laws!!
