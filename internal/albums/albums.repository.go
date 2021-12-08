package albums

import "gorm.io/gorm"

type AlbumRepository struct {
	db *gorm.DB
}

func (repo *AlbumRepository) FindAll() ([]Album, error) {
	var items []Album
	result := repo.db.Find(&items)
	return items, result.Error
}
func (repo *AlbumRepository) FindById(id string) (*Album, error) {
	var item Album
	result := repo.db.First(&item, id)
	return &item, result.Error
}

func (repo *AlbumRepository) AddOne(item *Album) error {
	result := repo.db.Create(item)
	return result.Error
}

func NewRepository(db *gorm.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}
