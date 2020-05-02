# go interface sample

## 1 Go Interface

    once we define a interface

    for example

```golang==
type Man interface {
    Run()
    Say()
}
```

    then every struct implement Run , Say function

    will be type of Man

## 2 Stringer interface

fmt have an interface call Stringer

with a function String(): string

once you implement String(): string for your struct

then fmt.Print will show your String() value

## 3 Storage with interface

    使用 type assertion to 判斷不同的type

    value.(type) 方式