//go:build !prod
// +build !prod

package mock_core

//go:generate mockgen -destination=mock_gen.go github.com/davidborzek/tvhgo/core UserRepository,SessionRepository,Clock
