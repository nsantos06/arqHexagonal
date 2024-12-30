package application

type ProductService struct {
	Persistance ProductPersistanceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error){
	product, err := s.Persistance.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

