package application

type ProductService struct {
	Persistance ProductPersistanceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistance.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func NewProductService (persistance ProductPersistanceInterface) *ProductService{
	return &ProductService{Persistance: persistance}
}

func (s *ProductService) Create(nome string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = nome
	product.Price = price
	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistance.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error){
	err := product.Enable()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistance.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}



func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error){
	err := product.Disable()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistance.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

