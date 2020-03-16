package newsfeed

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

// Item struct : simple json with two fields
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post`
}

// Repo : slice of Item structs
type Repo struct {
	Items []Item
}

// New : add new item to the repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// Add : Adds new item
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// GetAll : Gets all the items
func (r *Repo) GetAll() []Item {
	return r.Items
}
