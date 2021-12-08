package albums

type AlbumRepository struct {
}

var itemsStore = []Album{
	{ID: "0", Title: "In Rainbows", Artist: "Radiohead", Price: 99.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrance", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (repo *AlbumRepository) FindAll() []Album {
	return itemsStore
}
func (repo *AlbumRepository) FindById(id string) *Album {
	for _, item := range itemsStore {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func (repo *AlbumRepository) AddOne(newAlbum *Album) {
	itemsStore = append(itemsStore, *newAlbum)
}

func NewRepository() *AlbumRepository {
	return &AlbumRepository{}
}
