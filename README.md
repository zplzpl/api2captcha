# 2captcha.com V2 API

Go library for accessing the 2captcha.com v2 API

# Install

```go
go get github.com/zplzpl/api2captcha
```

# Usage

```go
import "github.com/zplzpl/api2captcha"

func main() {
    client := api2captcha.NewClient("your-key-here")
    
    balance, _ := client.GetBalance(context.Background())
    
    fmt.Println(balance)
}
```
