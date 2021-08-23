package client

//go:generate go run ../vendor/github.com/golang/mock/mockgen -source=./client.go -destination=./mock/client_generated.go -package=mock
type Client interface {
	GetSum(int, int) (int, error)
}
