# Go — Loops Quick Reference

> Quick reference for `for`/`range` patterns in Go.
> C# equivalent noted where useful.

---

## Classic for (C# style)

```go
for i := 0; i < 10; i++ {
    // ...
}
```

## for as while

```go
for condition {
    // ...
}
```

## Infinite loop

```go
for {
    // break to exit
}
```

---

## for range — Slice/Array

```go
// index + value
for i, v := range items {
    // i = index, v = value
}

// index only
for i := range items {
    // i = index
}

// value only (discard index)
for _, v := range items {
    // v = value
}
```

## for range — Map

```go
// key + value
for k, v := range m {
    // k = key, v = value
}

// key only
for k := range m {
    // k = key
}

// value only (discard key)
for _, v := range m {
    // v = value
}
```

**Important:** With a single variable in `range map`, you get the **key**, not the value.

## for range — String

```go
for i, r := range s {
    // i = byte index
    // r = rune (Unicode code point)
}
```

---

## Common Gotchas

- Map iteration order is **not guaranteed**
- Single variable in `range slice` = **index**, not value
- Single variable in `range map` = **key**, not value
- `_` means "discard this value" (blank identifier)

---

## Mental Model

| Collection | Two variables | Single variable |
|------------|--------------|-----------------|
| slice/array | index, value | index |
| map | key, value | key |
| string | byte index, rune | byte index |

**Rule:** Single variable always returns the **first** element of the pair.
