package db

// data base interface
// --------------------
type Store interface {
	Save(code, uri string)
	Fetch(code string) (string, bool)
	FindByURL(url string) (string, bool)
	Exists(code string) bool
}
