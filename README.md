# stringdiff

This is a golang package that impliments a string diffing algorithm based on Myers diff.

```
go get github.com/ntbosscher/stringdiff
```

```golang

func main(t *testing.T) {
	df := stringdiff.New("Hello world, this is a string", "Hello world, we have something new")
	fmt.Println(df)
	// > Hello world,
	// - this is a string
	// + we have something new

	df = stringdiff.New("Hello John, this is a string", "Hello Kennith Watson, this is a string")
	fmt.Println(df)
	// > Hello
	// - John,
	// + Kennith Watson,
	// > this is a string

	df = stringdiff.New("Hello John, how are you?", "Hello Kennith Watson, how is your fiance?")
	fmt.Println(df)
	// > Hello
	// - John,
	// + Kennith Watson,
	// > how
	// - are you?
	// + is your fiance?
}
```
