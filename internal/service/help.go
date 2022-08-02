package service

func (s *Service) Help(_ string) (string, error) {
	return "/help - list commands\n" +
		"/list - list movies\n" +
		"/add <title> <year> - add new movie with title and year\n" +
		"/remove <id> - remove movie with id\n" +
		"/update <id> <title> {<year>} - remove movie with id and title, year is optional", nil
}
