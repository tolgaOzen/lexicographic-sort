# lexicographic-sort

![GitHub](https://img.shields.io/github/go-mod/go-version/tolgaOzen/lexicographic-sort)
![GitHub](https://img.shields.io/github/last-commit/tolgaOzen/lexicographic-sort)
![GitHub](https://img.shields.io/github/license/tolgaOzen/lexicographic-sort)

Generate new string between two given strings for lexicographic sorting

## Install
```
go get -u github.com/tolgaOzen/lexicographic-sort
```

## Test
```
go test
```

## Examples

```go
func main() {
    fmt.Println(GenerateBetween("abcde", "abchi"))
    fmt.Println(GenerateBetween("abc", "abchi"))
}

// abcf
// abcd
```

```go
func main() {
    fmt.Println(GenerateBetween("abhs", "abit"))
    fmt.Println(GenerateBetween("abh", "abit"))
}

// abhw
// abhn
```


```go
func main() {
    fmt.Println(GenerateBetween("abhz", "abit"))
    fmt.Println(GenerateBetween("abhzs", "abit"))
    fmt.Println(GenerateBetween("abhzz", "abit"))
}

// abhzn
// abhzw
// abhzzn
```

```go
func main() {
    fmt.Println(GenerateBetween("abc", "abcah"))
    fmt.Println(GenerateBetween("abc", "abcab"))
    fmt.Println(GenerateBetween("abc", "abcaah"))
    fmt.Println(GenerateBetween("abc", "abcb"))
}

// abcad
// abcaan
// abcaad
// abcan
```

## Author
>Tolga Özen

>mtolgaozen@gmail.com

## License

MIT License

Copyright (c) 2021 Tolga Ozen

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
