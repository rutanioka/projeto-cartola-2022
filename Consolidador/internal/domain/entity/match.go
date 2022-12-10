package entity

import "time"

type MatchResult struct{
    teamAScore int
    teamBScore int
}

func (m *MatchResult)

type Match struct {
    ID string
    TeamA *Team
    TeamB *Team
    TeamAID string
    TeamBID string
    Date time.time
    Status string
    Result 
}
