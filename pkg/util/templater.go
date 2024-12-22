package util

import (
	"bytes"
	"text/template"
)

const challengeFileTemplate = `package challenges

import (
	"errors"

	"github.com/mbaeum/advent-of-go-2025/pkg/util"
)

type Challenge{{printf "%02d" .Id}} struct{ sc string }

func NewChallenge{{printf "%02d" .Id}}(cfg *util.Config) Challenge{{printf "%02d" .Id}} {
	return Challenge{{printf "%02d" .Id}}{cfg.SessionCookie}
}

func (c *Challenge{{printf "%02d" .Id}}) GetId() int { return {{.Id}} }

func (c *Challenge{{printf "%02d" .Id}}) GetData(m Mode) (string, error) {
	switch m {
	case TestMode:
		s := ` + "`" + `test` + "`" + `
		return s, nil
	case MainMode:
		return GetData(c, c.sc)
	default:
		return "", errors.New("mode not supported")
	}
}

func (c *Challenge{{printf "%02d" .Id}}) RunPartOne(m Mode) (string, error) { return "", nil }

func (c *Challenge{{printf "%02d" .Id}}) RunPartTwo(m Mode) (string, error) { return "", nil }

`

const challengeFileTestTemplate = `package challenges_test

import (
	"testing"

	"github.com/mbaeum/advent-of-go-2025/challenges"
	"github.com/mbaeum/advent-of-go-2025/pkg/util"
)

func TestRunPartOne(t *testing.T) {
	expected := ""
	mockConfig := util.Config{}
	c := challenges.NewChallenge{{printf "%02d" .Id}}(&mockConfig)
	res, err := c.RunPartOne(challenges.TestMode)
	if err != nil {
		t.Fatalf("Expected no error got: %v", err)
	}

	if !(expected == res) {
		t.Fatalf("Expected '%s' but got '%s'", expected, res)
	}

}

func TestRunPartTwo(t *testing.T) {
	expected := ""
	mockConfig := util.Config{}
	c := challenges.NewChallenge{{printf "%02d" .Id}}(&mockConfig)
	res, err := c.RunPartTwo(challenges.TestMode)
	if err != nil {
		t.Fatalf("Expected no error got: %v", err)
	}

	if !(expected == res) {
		t.Fatalf("Expected '%s' but got '%s'", expected, res)
	}

}

`

type templaterInput struct {
	Id int
}

func createChallengeFileString(i templaterInput, t string) (string, error) {
	tmpl, err := template.New("ChallengeFile").Parse(t)
	if err != nil {
		return "", err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, i)
	if err != nil {
		return "", err
	}
	return output.String(), nil

}
