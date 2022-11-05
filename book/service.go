package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	// merubah nilai dari request menjadi int sebelumnya json.Number
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
		Email: bookRequest.Email,
		Description: bookRequest.Description,
		Rating: int(rating),
		Discount: int(discount),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}
