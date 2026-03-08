# Go Loops Quick Reference

Cola rápida para não travar em `for`/`range` no Go.

## 1) `for` clássico (estilo C#)

```go
for i := 0; i < 10; i++ {
    // ...
}
```

## 2) `for` como `while`

```go
for condition {
    // ...
}
```

## 3) `for` infinito

```go
for {
    // ...
}
```

## 4) `for range` em slice/array

```go
for i, v := range items {
    // i = indice, v = valor
}
```

```go
for i := range items {
    // i = indice
}
```

```go
for _, v := range items {
    // v = valor (indice descartado)
}
```

## 5) `for range` em map

```go
for k, v := range m {
    // k = chave, v = valor
}
```

```go
for k := range m {
    // k = chave
}
```

```go
for _, v := range m {
    // v = valor (chave descartada)
}
```

Importante: com **1 variavel** no `range map`, voce recebe a **chave**.

## 6) `for range` em string

```go
for i, r := range s {
    // i = indice em bytes
    // r = rune (code point Unicode)
}
```

## 7) Pegadinhas que mais causam erro

- Em `map`, ordem de iteracao nao e garantida.
- Em `slice`, 1 variavel no `range` = indice, nao valor.
- Em `map`, 1 variavel no `range` = chave, nao valor.
- `_` significa "descartar valor que nao vou usar".

## 8) Regra mental rapida

- `range map` -> `key, value`
- `range slice/array` -> `index, value`
- `range string` -> `index, rune`

Se usar so uma variavel, sempre vem o primeiro item do par.
