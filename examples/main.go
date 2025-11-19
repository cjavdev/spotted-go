package main

import (
	"context"
	"fmt"

	"github.com/cjavdev/spotted-go"
	// "github.com/cjavdev/spotted-go/option"
)

func main() {
	client := spotted.NewClient(
	// option.WithClientID("My Client ID"),         // defaults to os.LookupEnv("SPOTIFY_CLIENT_ID")
	// option.WithClientSecret("My Client Secret"), // defaults to os.LookupEnv("SPOTIFY_CLIENT_SECRET")
	)
	album, err := client.Albums.Get(
		context.Background(),
		"4aawyAB9vmqN3uQ7FjRGTy",
		spotted.AlbumGetParams{},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", album.ID)
	fmt.Printf("%+v\n", album.Href)
	fmt.Printf("%+v\n", album.Name)
	fmt.Printf("%+v\n", album.Images)
	fmt.Printf("%+v\n", album.Label)
	fmt.Printf("%+v\n", album.ReleaseDate)
	fmt.Printf("%+v\n", album.ReleaseDatePrecision)
	fmt.Printf("%+v\n", album.TotalTracks)
	fmt.Printf("%+v\n", album.Type)
	fmt.Printf("%+v\n", album.Uri)
}
