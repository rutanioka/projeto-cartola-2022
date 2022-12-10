package entity

type Team struct{
    ID string
    Name string
    Player []*Player
}

func NewTeam (id, name string) *Team{
    return &Team{
        ID: id,
        Name: name,
    }
}

func (t *Team) AddPlayer(p *Player){
    t.Player = append(t.Player, player)   
}

func (t *Team) RemovePlayer(player *Player){
    for i, p := range t.Player{
        if p.ID == player.ID{
            t.Player = append(t.Player[:i],t.Player[i+1:]...)
            return
        }
    }
}