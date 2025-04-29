package mock_infra

//go:generate mockgen -source=../redis.go -destination=./redis.go
//go:generate mockgen -source=../websocket.go -destination=./websocket.go
