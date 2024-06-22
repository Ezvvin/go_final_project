package models

const DatePattern string = "20060102"

type Task struct {
	ID      uint   `json:"id,string"` // автоинкрементный идентификатор
	Date    string `json:"date"`      // дата задачи, которая будет хранится в формате YYYYMMDD или в Go-представлении 20060102;
	Title   string `json:"title"`     // заголовок задачи;
	Comment string `json:"comment"`   // комментарий к задаче;
	Repeat  string `json:"repeat"`    // строковое поле не более 128 символов
}
type Tasks struct {
	Tasks []Task `json:"tasks"`
}
