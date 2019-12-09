package model

//Data thing
type Data struct {
	Articles []Article
}

//Article Thing
type Article struct {
	AId     string `json:"aid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//Hello message
type Hello struct {
	Message string `json:message`
}

//Error
type Error struct {
	Message string `json:message`
}
