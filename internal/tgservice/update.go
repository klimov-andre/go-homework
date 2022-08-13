package tgservice

func (s *Service) Update(args string) (string, error) {
	m, err := s.robot.Update(args)
	if err != nil {
		return "", err
	}

	return m.String(), nil
}
