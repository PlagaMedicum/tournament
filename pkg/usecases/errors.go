package usecases

import "errors"

var (
	ErrNotEnoughPoints    = errors.New("user have not enough points to join the tournament")
	ErrParticipantExists  = errors.New("this user already joined the tournament")
	ErrFinishedTournament = errors.New("this tournament is already finished")
	ErrNoParticipants     = errors.New("cannot assign winner while there is no participants")
)
