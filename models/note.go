// note.go

package models

type Note struct {
	Title   string `bson:"title"`
	Content string `bson:"content"`
}
