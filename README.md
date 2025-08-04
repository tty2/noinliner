# noinliner

⚠️ Disclaimer

  This linter is opinionated.

  The **noinliner** linter reflects the author's personal preference for writing Go code — specifically the belief that separating variable declarations from if conditions improves clarity and readability, belief that function call and error check are two separate steps (actions).

  This is not part of the official Go style guide, and many Go developers (and tools like go fmt, golint, or staticcheck) accept or even encourage inline declarations in if statements.

  Use this tool if it aligns with your team's style philosophy — or ignore it if it doesn’t.


**noinliner** is a custom Go linter that disallows inline variable declarations in `if` statements.

It flags patterns like:

```go
if val, err := doSomething(); err != nil {
    // ❌ flagged by noinliner
}

if _, ok := myMap[key]; ok {
    // ❌ flagged by noinliner
}
```

🚫 Why Avoid Inline Declarations in if?

Inline declarations can reduce readability and make debugging harder, especially in complex control flows. By enforcing separation of declaration and condition, you get:

- Cleaner, easier-to-read code
- More consistent error handling
- Simpler debugging (variables exist outside the if scope)

Real app code that made me to write this linter

```
config := []byte("{}")
if configString := c.String("apns_config"); configString != "" {
	config = []byte(configString)
} else if configFile := c.String("apns_config_file"); configFile != "" {
	contents, err := os.ReadFile(filepath.Clean(configFile))
	if err != nil {
		log.Fatalf("read apns config file %q: %s", configFile, err)
	}

	config = contents
}
```


✅ Allowed Patterns

```
val, err := doSomething()
if err != nil {
    // ✅ not flagged
}

result, ok := myMap[key]
if ok {
    // ✅ not flagged
}
```

🔍 What It Checks

- Flags any if statement with an inline Init declaration.

- Does not check switch statements (by design).

- Ignores the variable names or types — any inline init is flagged.

🛠 Installation

    go install github.com/tty2/noinliner@latest

🚀 Usage

Run noinliner on a package or directory:

    noinliner ./...

Or target a specific file:

    noinliner myfile.go

License

MIT
