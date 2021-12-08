package albums

// func TestGetAll(t *testing.T) {
// 	repo := NewRepository()

// 	got := repo.FindAll()
// 	want := itemsStore

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }

// func TestGetById(t *testing.T) {
// 	repo := NewRepository()

// 	t.Run("find existing item", func(t *testing.T) {
// 		got := repo.FindById("1")
// 		want := &itemsStore[1]

// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v want %v", got, want)
// 		}
// 	})
// 	t.Run("find non existsing item", func(t *testing.T) {
// 		got := repo.FindById("-1")

// 		if got != nil {
// 			t.Errorf(`repo.FindById("-1") != nil`)
// 		}
// 	})

// }

// func TestAddOne(t *testing.T) {
// 	repo := NewRepository()

// 	t.Run("add one item", func(t *testing.T) {
// 		newAlbum := Album{ID: "10", Title: "Black Album", Artist: "Metallica", Price: 99.99}

// 		repo.AddOne(&newAlbum)

// 		got := itemsStore[len(itemsStore)-1]
// 		if !reflect.DeepEqual(got, newAlbum) {
// 			t.Errorf("got %v want %v", got, newAlbum)
// 		}
// 	})
// }
