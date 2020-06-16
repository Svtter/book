package books

// Book book models
type Book struct {
	Title  string
	Author string
	Pages  int
}

// CategoryByLength category book by length
func (b Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	} else {
		return "SHORT STORY"
	}
}
