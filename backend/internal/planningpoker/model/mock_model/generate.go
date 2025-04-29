package mock_model

//go:generate mockgen -source=../participant.go -destination=./participant.go
//go:generate mockgen -source=../round.go -destination=./round.go
//go:generate mockgen -source=../session.go -destination=./session.go
//go:generate mockgen -source=../vote.go -destination=./vote.go
