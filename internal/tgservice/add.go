package tgservice

func (s *Service) Add(args string) (string, error) {
	m, err := s.robot.Add(args)
	if err != nil {
		return "", err
	}

	return m.String(), nil
}
