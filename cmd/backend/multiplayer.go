package main

import (
	"io"
	"log"

	pb "github.com/sgielen/vierkantle/pkg/proto"
)

type vierkantleService struct {
	log   *log.Logger
	teams VierkantleTeams
}

func NewVierkantleService(log *log.Logger) *vierkantleService {
	return &vierkantleService{
		log:   log,
		teams: *NewVierkantleTeams(),
	}
}

func (s *vierkantleService) TeamStream(stream pb.VierkantleService_TeamStreamServer) error {
	// The first message should be create / join
	msg, err := stream.Recv()
	if err == io.EOF {
		s.log.Printf("New stream closed immediately")
		return nil
	} else if err != nil {
		s.log.Printf("New stream closed with error: %s", err.Error())
		return err
	}

	var team *VierkantleTeam
	var playerName string

	defer func() {
		if team != nil && playerName != "" {
			team.PlayerLeaves(playerName)
			team.BroadcastTeamInfo()
			if team.IsEmpty() {
				s.teams.ForgetTeam(team.GetToken())
			}
		}
	}()

	if create := msg.GetCreate(); create != nil {
		team = s.teams.NewTeam()
		s.log.Printf("New stream creates new team for %s, token %s", create.Name, team.GetToken())
		if err := team.PlayerJoins(create.Name, stream); err != nil {
			return stream.Send(&pb.TeamStreamServerMessage{
				Response: &pb.TeamStreamServerMessage_Error{
					Error: &pb.ErrorResponse{
						Error: err.Error(),
					},
				},
			})
		}
		playerName = create.Name
	} else if join := msg.GetJoin(); join != nil {
		s.log.Printf("New stream joins team %s for %s", join.Token, join.Name)
		team = s.teams.GetTeam(join.Token)
		if team == nil {
			return stream.Send(&pb.TeamStreamServerMessage{
				Response: &pb.TeamStreamServerMessage_Error{
					Error: &pb.ErrorResponse{
						Error: "team not found",
					},
				},
			})
		}
		if err := team.PlayerJoins(join.Name, stream); err != nil {
			return stream.Send(&pb.TeamStreamServerMessage{
				Response: &pb.TeamStreamServerMessage_Error{
					Error: &pb.ErrorResponse{
						Error: err.Error(),
					},
				},
			})
		}
		playerName = join.Name
	} else {
		s.log.Printf("New stream sends invalid message %+v", msg)
		return stream.Send(&pb.TeamStreamServerMessage{
			Response: &pb.TeamStreamServerMessage_Error{
				Error: &pb.ErrorResponse{
					Error: "must first create/join team",
				},
			},
		})
	}

	team.BroadcastTeamInfo()
	// TODO: also send all words already guessed before in this team

	for {
		msg, err = stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if changeName := msg.GetName(); changeName != nil {
			if err := team.PlayerJoins(changeName.Name, stream); err == nil {
				team.PlayerLeaves(playerName)
				playerName = changeName.Name
				team.BroadcastTeamInfo()
			} else {
				if err := stream.Send(&pb.TeamStreamServerMessage{
					Response: &pb.TeamStreamServerMessage_Error{
						Error: &pb.ErrorResponse{
							Error: err.Error(),
						},
					},
				}); err != nil {
					return err
				}
			}
			continue
		} else if word := msg.GetWord(); word != nil {
			team.BroadcastToEveryoneExcept(&pb.TeamStreamServerMessage{
				Response: &pb.TeamStreamServerMessage_Word{
					Word: &pb.WordGuessedResponse{
						Player: playerName,
						Word:   word.Word,
					},
				},
			}, playerName)
		} else {
			if err := stream.Send(&pb.TeamStreamServerMessage{
				Response: &pb.TeamStreamServerMessage_Error{
					Error: &pb.ErrorResponse{
						Error: "invalid message",
					},
				},
			}); err != nil {
				return err
			}
		}
	}
}
