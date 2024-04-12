Disclaimer: This is mostly a toy project made in an afternoon so don't expect miracles kind reader but it's actually the first time I built something that's actually useful (for me at least) so I'm quite happy with the result!

# Translate golang struct into ts types!

Basically you can install it like this:
```go install github.com/okzmo/gtp```

And then you can simply use it like this: 
```gpt --in="./model.go" --out="./types" --namespace="MyProject"```

But what are these flags I'm using ? Glad you ask:
- `--namespace` followed by a String will tell the program that you want all of your structs to be part of a namespace. That's actually really useful if like me you like JSDoc in SvelteKit that way you can just do `/** @type {MyProject.User} */ and there your go it's typed!

- `--in` the input for the program, it needs to be .go file to work otherwise you'll get an error.

- `--out` the output of the program, you can enter a whole path either towards a directory or you can also add the file name at the end like `models.ts` (if you don't it'll automatically create a file named `types.d.ts` where you're pointing).