package challenges

import (
	"fmt"

	"github.com/mbaeum/advent-of-go-2025/pkg/util"
)

type Mode int

const (
	TestMode Mode = iota
	MainMode
)

type Challenge interface {
	GetId() int
	SetSessionCookie(sc string)
	GetData(m Mode) (string, error)
	RunPartOne(m Mode) (string, error)
	RunPartTwo(m Mode) (string, error)
}

func GetData(c Challenge, sc string) (string, error) {
	d, err := util.GetChallengeData(c.GetId(), sc)
	if err != nil {
		return "", err
	}
	return d, nil

}

type ChallengeFactory struct {
	m map[int]Challenge
}

func (cf *ChallengeFactory) RegisterChallenge(c Challenge) {
	cf.m[c.GetId()-1] = c
	fmt.Printf("registered challenge %d: %v", c.GetId(), cf.m)
}

func (cf *ChallengeFactory) GetChallenge(id int) (Challenge, error) {
	if c, ok := cf.m[id-1]; !ok {
		return nil, fmt.Errorf("could not retrieve challenge %d from %v", id, cf.m)
	} else {
		return c, nil
	}
}

func NewChallengeFactory() *ChallengeFactory {
	m := make(map[int]Challenge)
	cf := ChallengeFactory{m}

	registerChallenges(&cf)
	return &cf
}

// This function is to register new challenges
// by adding a line
// f.RegisterChallenge(&ChallengeX{})
func registerChallenges(f *ChallengeFactory) {
	f.RegisterChallenge(&Challenge01{})
}
