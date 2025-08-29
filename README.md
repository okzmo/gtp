# Translate golang struct into ts types!

![banner](./images/gtp_banner.png)

GTP stands for: "Generate TyPes".

Install it like this:
```go install github.com/okzmo/gtp@latest```

Use it like this: 
```gpt --in="./model.go" --out="./types" --namespace="MyProject"```

The flags you can use:
- `--namespace` followed by a String will tell the program that you want all of your structs to be part of a namespace. Very useful if you use JSDoc on the frontend, here's an example: `/** @type {MyProject.User} */`

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

Now we run ```gtp --in="./model.go" --out="./types"```

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

Pretty cool right ? By the way you can use `//?` for optional types.

If you have any issue or find a bug please open an issue and I'll fix it.
