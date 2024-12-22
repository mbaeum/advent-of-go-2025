package challenges

import (
	"fmt"

	"github.com/mbaeum/advent-of-go-2025/pkg/util"
)

type Challenge interface {
	GetId() int
	GetTestData() string
	RunPartOneTest() error
	RunPartOne() error
	RunPartTwoTest() error
	RunPartTwo() error
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

func (cf *ChallengeFactory) RegisterChallenge(c Challenge) error {
	if c.GetId() > len(cf.m) {
		return fmt.Errorf("Challenge factory can only hold %d challenges, tried adding %d", len(cf.m), c.GetId())
	}
	cf.m[c.GetId()] = c

	return nil
}

func (cf *ChallengeFactory) GetChallenge(id int) (Challenge, error) {
	if c, ok := cf.m[id]; !ok {
		return nil, fmt.Errorf("could not retrieve challenge %d", id)
	} else {
		return c, nil
	}
}

func NewChallengeFactory() ChallengeFactory {
	m := make(map[int]Challenge, 24)
	return ChallengeFactory{m}
}
