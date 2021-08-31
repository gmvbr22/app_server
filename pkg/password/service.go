package password

type Service interface {
	GenerateFromPassword(password []byte) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

type service struct {
	adapter Adapter
}

func NewService(adapter Adapter) Service {
	return &service{adapter: adapter}
}

func (service *service) GenerateFromPassword(password []byte) ([]byte, error){
	return service.adapter.GenerateFromPassword(password)
}

func (service * service) CompareHashAndPassword(hashedPassword, password []byte) error {
	return service.adapter.CompareHashAndPassword(hashedPassword, password)
}
