package service

import (
	"strings"
)

func (s *Service) List(args string) (string, error) {
	list, err := s.robot.List()
	if err != nil {
		return "", err
	}

	res := make([]string, 0, len(list))
	for _, v := range list {
		res = append(res, v.String())
	}
	return strings.Join(res, "\n"), nil
}
