## Maps: pointer-receiver methods on map values (unaddressable values)

- Cause: When a method has a pointer receiver (e.g., `func (m *User) PrintFirstName()`), calling it on a map element like `userMap[id]` fails because map index expressions yield unaddressable values. The compiler reports an error such as `cannot call pointer method PrintFirstName on userMap[id]`.

- Example:

```go
type User struct {
    FirstName string
}

func (m *User) PrintFirstName() {
    fmt.Println(m.FirstName)
}

userMap := make(map[string]User)
userMap["100"] = User{FirstName: "Hariom"}
// compile error: cannot call pointer method PrintFirstName on userMap["100"]
userMap["100"].PrintFirstName() // ❌
```

- Solutions:

1. Use a value receiver if the method doesn't need to mutate the receiver:

```go
func (m User) PrintFirstName() { // value receiver
    fmt.Println(m.FirstName)
}
// then userMap["100"].PrintFirstName() works
```

2. Store pointers in the map:

```go
userMap := make(map[string]*User)
user := &User{FirstName: "Hariom"}
userMap["100"] = user
userMap["100"].PrintFirstName() // works
```

3. Copy the map value to an addressable variable and call the pointer method (note: changes to the copy won't update the map unless you reassign):

```go
u := userMap["100"]
(&u).PrintFirstName()
// if you modified u and want to persist it back:
userMap["100"] = u
```

- Tip: Choose pointers if you need to mutate the stored value or want to avoid copying large structs; use value receivers when methods are read-only and cheap to copy.
