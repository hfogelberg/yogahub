package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kyokomi/cloudinary"
)

func cloudinaryUpload(src string, fileName string) error {
	cloudinaryKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudinarySecret := os.Getenv("CLOUDINARY_API_SECRET")
	cloudname := "golizzard"

	ctx := context.Background()
	con := fmt.Sprintf("cloudinary://%s:%s@%s", cloudinaryKey, cloudinarySecret, cloudname)

	fmt.Printf("Uploading to cloudinary %s\n", con)

	ctx = cloudinary.NewContext(ctx, con)
	data, _ := ioutil.ReadFile(src)

	if err := cloudinary.UploadStaticImage(ctx, fileName, bytes.NewBuffer(data)); err != nil {
		log.Println("Error uploading image to cloudinary")
		return err
	}

	fmt.Println("Upload to cloudinary OK")

	_ = os.Remove(src)

	return nil
}
