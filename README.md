pipes
=====

### Piped Commands

A very simple, but powerful monadic[1] command runner, where commands declare 
functions to be run.

Since the Go language lacks defining immutable state (other than probably 
closures), it's thought of as best practice. If using the command runner the
commands are created once, so any state created in there could effect the next
run, so be very careful.

Although some of the structs are named after other functional language types 
(Either, etc), they are not in fact the same. They're close, but each type has
been locked down to a tighter set of types, to prevent the need for type casting.
Which means they'll not pass any monadic laws!!

###  Example

The [example](example.go) is a lot more up to date by it's very nature of being 
code, but to give you an idea, see the following:

Create a set of commands that we want to use:

```
type AddCommand struct{}

func (c AddCommand) Execute(note pipes.Note) pipes.CommandResult {
    return pipes.ContinueResult(note.(Sum).Concat(NewSum(1)))
}

type BadCommand struct{}

func (c BadCommand) Execute(note pipes.Note) pipes.CommandResult {
    return pipes.BreakResult(note)
}
```

New create a runner:

```
commands := []pipes.Command{
        AddCommand{},
        AddCommand{},
        AddCommand{},
        AddCommand{},
}

runner := pipes.NewRunner(commands)
res := runner.Execute(NewSum(1))
```

Then get the result from the either:

```
res.Bimap(
    func(x pipes.Tuple) pipes.Tuple {
        fmt.Println("Failed: ", x)
        return x
    },
    func(x pipes.Tuple) pipes.Tuple {
        fmt.Println("Success: ", x)
        return x
    },
)
```

[1] Sort of!