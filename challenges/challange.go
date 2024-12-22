package challenges

import "fmt"

type Challenge interface {
	RunTest() error
	Run() error
}

type ChallengeFactory struct {
	m map[int]Challenge
}

func (cf *ChallengeFactory) RegisterChallenge(id int, c Challenge) error {
	if id > len(cf.m) {
		return fmt.Errorf("Challenge factory can only hold %d challenges, tried adding %d", len(cf.m), id)
	}
	cf.m[id] = c

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
