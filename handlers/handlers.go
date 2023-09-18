package handlers

import (
	"context"
	"html/template"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gopheramol/notesapp/models"
	"github.com/gorilla/mux"
)

func GetNotes(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	collection := client.Database("notes").Collection("notes")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var notes []models.Note
	for cursor.Next(context.Background()) {
		var note models.Note
		cursor.Decode(&note)
		notes = append(notes, note)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, notes)
}

func AddNote(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	note := models.Note{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	collection := client.Database("notes").Collection("notes")
	_, err := collection.InsertOne(context.TODO(), note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteNote(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	vars := mux.Vars(r)
	title := vars["title"]

	collection := client.Database("notes").Collection("notes")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"title": title})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
