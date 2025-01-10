// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package sqlc

import (
	"context"
)

const findTracks = `-- name: FindTracks :many
select track.track_id, track.name, track.composer, album.title
from track
join album on track.album_id = album.album_id
where
    (NOT $1::boolean or album.title = $2) AND
    (NOT $3::boolean or track.composer = $4)
`

type FindTracksParams struct {
	ByAlbumTitle bool    `json:"by_album_title"`
	AlbumTitle   string  `json:"album_title"`
	ByComposer   bool    `json:"by_composer"`
	Composer     *string `json:"composer"`
}

type FindTracksRow struct {
	TrackID  int32   `json:"track_id"`
	Name     string  `json:"name"`
	Composer *string `json:"composer"`
	Title    string  `json:"title"`
}

func (q *Queries) FindTracks(ctx context.Context, arg FindTracksParams) ([]FindTracksRow, error) {
	rows, err := q.db.Query(ctx, findTracks,
		arg.ByAlbumTitle,
		arg.AlbumTitle,
		arg.ByComposer,
		arg.Composer,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindTracksRow
	for rows.Next() {
		var i FindTracksRow
		if err := rows.Scan(
			&i.TrackID,
			&i.Name,
			&i.Composer,
			&i.Title,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAlbums = `-- name: ListAlbums :many
SELECT album_id, title, artist_id FROM album
`

func (q *Queries) ListAlbums(ctx context.Context) ([]Album, error) {
	rows, err := q.db.Query(ctx, listAlbums)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Album
	for rows.Next() {
		var i Album
		if err := rows.Scan(&i.AlbumID, &i.Title, &i.ArtistID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAlbumsCustom = `-- name: ListAlbumsCustom :many
SELECT album_id, title FROM album
`

type ListAlbumsCustomRow struct {
	AlbumID int32  `json:"album_id"`
	Title   string `json:"title"`
}

func (q *Queries) ListAlbumsCustom(ctx context.Context) ([]ListAlbumsCustomRow, error) {
	rows, err := q.db.Query(ctx, listAlbumsCustom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAlbumsCustomRow
	for rows.Next() {
		var i ListAlbumsCustomRow
		if err := rows.Scan(&i.AlbumID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAlbumsWithArtist = `-- name: ListAlbumsWithArtist :many
SELECT album.album_id, album.title, album.artist_id, artist.artist_id, artist.name FROM album join artist on album.artist_id = artist.artist_id
`

type ListAlbumsWithArtistRow struct {
	Album  Album  `json:"album"`
	Artist Artist `json:"artist"`
}

func (q *Queries) ListAlbumsWithArtist(ctx context.Context) ([]ListAlbumsWithArtistRow, error) {
	rows, err := q.db.Query(ctx, listAlbumsWithArtist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAlbumsWithArtistRow
	for rows.Next() {
		var i ListAlbumsWithArtistRow
		if err := rows.Scan(
			&i.Album.AlbumID,
			&i.Album.Title,
			&i.Album.ArtistID,
			&i.Artist.ArtistID,
			&i.Artist.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
