package main

import (
	"fmt"
	"sync"

	pb "github.com/sgielen/vierkantle/pkg/proto"
)

type VierkantleTeam struct {
	mtx     sync.Mutex
	token   string
	players map[string]pb.VierkantleService_TeamStreamServer
}

type VierkantleTeams struct {
	mtx   sync.Mutex
	teams map[string]*VierkantleTeam
}

func NewVierkantleTeams() *VierkantleTeams {
	return &VierkantleTeams{
		teams: map[string]*VierkantleTeam{},
	}
}

func (t *VierkantleTeams) NewTeam() *VierkantleTeam {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	name := RandomString(8)
	team := &VierkantleTeam{
		token:   name,
		players: map[string]pb.VierkantleService_TeamStreamServer{},
	}
	t.teams[name] = team
	return team
}

func (t *VierkantleTeams) GetTeam(token string) *VierkantleTeam {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if team, ok := t.teams[token]; ok {
		return team
	} else {
		return nil
	}
}

func (t *VierkantleTeams) ForgetTeam(token string) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	delete(t.teams, token)
}

func (t *VierkantleTeam) GetPlayerNames() []string {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	return t.getPlayerNames()
}

func (t *VierkantleTeam) getPlayerNames() []string {
	res := make([]string, len(t.players))
	i := 0
	for p := range t.players {
		res[i] = p
		i++
	}
	return res
}

func (t *VierkantleTeam) BroadcastTeamInfo() {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	for p, stream := range t.players {
		/* ignore err */
		_ = stream.Send(&pb.TeamStreamServerMessage{
			Response: &pb.TeamStreamServerMessage_Team{
				Team: t.getInfo(p),
			},
		})
	}
}

func (t *VierkantleTeam) BroadcastToEveryoneExcept(message *pb.TeamStreamServerMessage, except string) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	for p, stream := range t.players {
		if p == except {
			continue
		}
		_ = stream.Send(message)
	}
}

func (t *VierkantleTeam) GetToken() string {
	// no need to lock, token never changes
	return t.token
}

func (t *VierkantleTeam) GetInfo(player string) *pb.TeamInfoResponse {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	return t.getInfo(player)
}

func (t *VierkantleTeam) getInfo(player string) *pb.TeamInfoResponse {
	return &pb.TeamInfoResponse{
		Token:    t.token,
		YourName: player,
		Players:  t.getPlayerNames(),
	}
}

func (t *VierkantleTeam) PlayerJoins(name string, stream pb.VierkantleService_TeamStreamServer) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if name == "" {
		return fmt.Errorf("player name cannot be empty")
	}
	if _, ok := t.players[name]; ok {
		return fmt.Errorf("player name is in use")
	}
	t.players[name] = stream
	return nil
}

func (t *VierkantleTeam) PlayerLeaves(name string) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	delete(t.players, name)
}

func (t *VierkantleTeam) IsEmpty() bool {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	return len(t.players) == 0
}
