# noinliner
---
**noinliner** is a custom Go linter plugin for golangci-lint that disallows inline variable declarations in `if` statements.

⚠️ Disclaimer

  This linter is opinionated.

  The **noinliner** linter reflects the author's personal preference for writing Go code — specifically the belief that separating variable declarations from if conditions improves clarity and readability, belief that function call and error check are two separate steps (actions).

  This is not part of the official Go style guide, and many Go developers (and tools like go fmt, golint, or staticcheck) accept or even encourage inline declarations in if statements.

  Use this tool if it aligns with your team's style philosophy — or ignore it if it doesn’t.

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

## Behavior

❌ It flags patterns like:

```go
if val, err := doSomething(); err != nil {
    // ❌ flagged by noinliner
}

if _, ok := myMap[key]; ok {
    // ❌ flagged by noinliner
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

🛠 Installation

It uses golangci-lint [Module Plugin System](https://golangci-lint.run/plugins/module-plugins/)

Inside `.custom-gcl.yaml` you can change the version (it's better to have the same golangci-lint has on the machine).

Run:

    golangci-lint custom -v

the result binary can be called instead of golangci-lint or you can change this binary to golangci-lint. It is in your `$(go env GOPATH)/bin`

Better way. Use Makefile:

    make build

If you want to replace golangci-lint with the one with this linter, call:

    make replace


License

MIT
