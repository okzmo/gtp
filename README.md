# Translate golang struct into ts types!

![banner](./images/gtp_banner.png)

gtp stands for: "Generate TyPes", yes I know I suck at naming things.

Basically you can install it like this:
```go install github.com/okzmo/gtp@latest```

And then you can simply use it like this: 
```gpt --in="./model.go" --out="./types" --namespace="MyProject"```

But what are these flags I'm using ? Glad you ask:
- `--namespace` followed by a String will tell the program that you want all of your structs to be part of a namespace. That's actually really useful if like me you like JSDoc in SvelteKit that way you can just do `/** @type {MyProject.User} */ and there your go it's typed!

- `--in` the input for the program, it needs to be .go file to work otherwise you'll get an error.

- `--out` the output of the program, you can enter a whole path either towards a directory or you can also add the file name at the end like `models.ts` (if you don't it'll automatically create a file named `types.d.ts` where you're pointing).

## Usage
Here's what the translation looks like:
```go
// models.go
package models

type User struct {
  Username string
  Email string
  Age uint8 //?
  Posts []Post //?
}

type Post struct {
  Author User
  Content string
  Comments []string //?
}
```

Now we run ```gpt --in="./model.go" --out="./types"```

```ts
// types/types.d.ts
export type User = {
  username: string
  email: string
  age?: number
  posts?: Post[]
}

export type Post = {
  author: User
  content: string
  comments?: string[] 
}
```

Pretty cool right ? You just writing golang struct and they get translated for you. As for the `//?` that's just a random thing I came up with to easily have optional types.

That's it, hopefully it'll be of use to somebody, if you have any issue or find a bug please open an issue and I'll do my best to fix it.

PS: You can combine this with [air](https://github.com/cosmtrek/air) for a really cool DX where you just change your golang structs and don't even have to think about executing `gtp`.
