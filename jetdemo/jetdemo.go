package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"matdurand/go-get-demo/.gen/postgres/public/model"
	"os"

	. "github.com/go-jet/jet/v2/postgres"
	. "matdurand/go-get-demo/.gen/postgres/public/table"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func main() {
	var db *sql.DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt :=
		SELECT(
			Album.AlbumID, Album.Title, Artist.Name,
		).FROM(
			Album.INNER_JOIN(Artist, Album.ArtistID.EQ(Artist.ArtistID)),
		).ORDER_BY(Album.AlbumID)
	printStatementInfo(stmt)

	// This is a bit error-prone. If I forget the alias, the query doesn't fail but I get nothing
	// If some aliases are wrong, these fields are ignored
	// This seems like something that would generate errors in the long run
	var dest []struct {
		AlbumID int32  `alias:"album.album_id"`
		Title   string `alias:"album.title"`
		Artist  string `alias:"artist.name"`
	}
	err = stmt.Query(db, &dest)
	panicOnError(err)
	jsonSave("./jet-albums-custom-struct.json", dest)

	// While this is ok, in this query, I'm not selecting the Artist, so I get a bunch of Albums with ArtistID = 0
	// This is also error-prone as someone downstream using the returned struct wouldn't know if the data is just
	// empty or missing
	var dest2 []struct {
		model.Album
		model.Artist
	}
	err = stmt.Query(db, &dest2)
	panicOnError(err)
	jsonSave("./jet-albums-model.json", dest2)
}

func printStatementInfo(stmt SelectStatement) {
	query, args := stmt.Sql()

	fmt.Println("Parameterized query: ")
	fmt.Println("==============================")
	fmt.Println(query)
	fmt.Println("Arguments: ")
	fmt.Println(args)

	debugSQL := stmt.DebugSql()

	fmt.Println("\n\nDebug sql: ")
	fmt.Println("==============================")
	fmt.Println(debugSQL)
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
