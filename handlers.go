package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("").ParseFiles("templates/index.html", "templates/layout.html")
	err = tpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Fatalln("Error serving index template ", err.Error())
		return
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("").ParseFiles("templates/admin.html", "templates/layout.html")
	err = tpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Fatalln("Error serving admin template ", err.Error())
		return
	}
}

func (conn *Connection) posesHandler(w http.ResponseWriter, r *http.Request) {
	poses, err := conn.getPoses()
	if err != nil {
		return
	}

	tpl, err := template.New("").ParseFiles("templates/poses.html", "templates/layout.html")
	err = tpl.ExecuteTemplate(w, "layout", poses)
	if err != nil {
		log.Fatalln("Error serving poses template ", err.Error())
		return
	}
}

func createPoseHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("").ParseFiles("templates/createPose.html", "templates/layout.html")
	err = tpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Fatalln("Error serving index template ", err.Error())
		return
	}
}

func (conn *Connection) postCreatePoseHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error reading form image %s\n", err.Error())
		return
	}
	defer file.Close()

	root := "./public/temp/"
	os.Mkdir(root, 0700)

	filename := r.FormValue("title-en") + ".svg"
	path := root + "/" + filename

	out, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating file in public/temp %s", err.Error())
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Printf("Error writing file to public/tmp %s", err.Error())
		return
	}

	if err := cloudinaryUpload(path, filename); err != nil {
		log.Printf("Error uploading to Cloudinary %s\n", err.Error())
		return
	}

	pose := Pose{
		TitleEn:     r.FormValue("title-en"),
		TitleSa:     r.FormValue("title-se"),
		Description: r.FormValue("description"),
		File:        filename,
		When:        time.Now(),
	}

	fmt.Println("Saving pose to Db")
	fmt.Println(pose)

	if err := conn.createPose(&pose); err != nil {
		fmt.Printf("Error creating pose %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println("SAVED OK")
	http.Redirect(w, r, "/admin/poses", http.StatusSeeOther)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/favicon.ico")
}
