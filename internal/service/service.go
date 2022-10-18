package service

type Service struct {
	db       Storage
	password PasswordService
}

func NewService(db Storage, password PasswordService) *Service {
	return &Service{db: db, password: password}
}

func (s *Service) WithStorage(db Storage) *Service {
	nc := &Service{}
	*nc = *s
	nc.db = db
	return nc
}
