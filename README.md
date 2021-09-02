# Unit Test in GoLang

### [Medium blog link](https://medium.com/@karthikkn1997/unit-testing-in-go-language-using-mocks-3b873ce1348d)

This repo contains the sample code for implenting unit test in go language and unit test implementions using following methods
  - Simple unit test 
  - Unit test using tabular method
  - Unit test using custom method for interface function
  - Unit test using mocks

##Running the test locally
```bash
git clone https://github.com/Karthik-K-N/goUT.git
cd goUT
go mod vendor
go test ./...
```

## Generating the mock methods

- Add the silent import in the file

```bash
     _ "github.com/golang/mock/mockgen/model"
```
- Download the mockgen package into vendor folder

```bash
    go mod vendor
```
- Add the mockgen command to the interface

```bash
    //go:generate go run ../vendor/github.com/golang/mock/mockgen -source=./client.go -destination=./mock/client_generated.go -package=mock
    type Client interface {
      GetSum(int, int) (int, error)
    }
```
- Generate the mock methods

```bash
    go generate ./...
```
