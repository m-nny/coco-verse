package albums

import "gorm.io/gorm"

type AlbumRepository struct {
	db *gorm.DB
}

func (repo *AlbumRepository) FindAll() ([]Album, error) {
	var items []Album
	err := repo.db.Find(&items).Error
	return items, err
}

func (repo *AlbumRepository) FindById(id string) (*Album, error) {
	var item Album
	err := repo.db.First(&item, id).Error
	return &item, err
}

func (repo *AlbumRepository) AddOne(item *Album) error {
	err := repo.db.Create(item).Error
	return err
}

func NewRepository(db *gorm.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}
