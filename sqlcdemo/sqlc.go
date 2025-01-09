package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"matdurand/go-get-demo/sqlc"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func main() {
	ctx := context.Background()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := pgx.Connect(ctx, psqlInfo)
	panicOnError(err)
	defer conn.Close(ctx)

	queries := sqlc.New(conn)

	//Compared to Jet, this query returns a custom struct with just 2 fields, since I'm not select * from Albums
	//This is way less error-prone. However that could lead to the proliferation of structs, but that's controllable
	albumsCustom, err := queries.ListAlbumsCustom(ctx)
	panicOnError(err)
	jsonSave("./sqlc-albums-custom.json", albumsCustom)

	//And this is the full version that used the generated model for Album, with all the fields
	albums, err := queries.ListAlbums(ctx)
	panicOnError(err)
	jsonSave("./sqlc-albums.json", albums)

	//And with a join
	//We can have all the fields flat, or use embedded struct like
	//{
	//	"Album": {
	//		"AlbumID": 1,
	//		"Title": "For Those About To Rock We Salute You",
	//		"ArtistID": 1
	//	},
	//	"Artist": {
	//		"ArtistID": 1,
	//		"Name": "AC/DC"
	//	}
	//}
	albumsAndArtists, err := queries.ListAlbumsWithArtist(ctx)
	panicOnError(err)
	jsonSave("./sqlc-albums-and-artists.json", albumsAndArtists)
}

func jsonSave(path string, v interface{}) {
	jsonText, _ := json.MarshalIndent(v, "", "\t")

	err := os.WriteFile(path, jsonText, 0600)

	panicOnError(err)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
