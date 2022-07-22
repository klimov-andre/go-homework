package service

func (s *Service) Remove(args string) (string, error) {
	err := s.robot.Remove(args)
	if err != nil {
		return "", err
	}

	return "", nil
}
