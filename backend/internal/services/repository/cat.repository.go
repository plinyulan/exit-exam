package repository

type CatRepository interface {
	GetCats() []string
}

type catRepository struct{}

func NewCatRepository() CatRepository {
	return &catRepository{}
}

func (r *catRepository) GetCats() []string {
	return []string{"Whiskers", "Fluffy", "Mittens"}
}
