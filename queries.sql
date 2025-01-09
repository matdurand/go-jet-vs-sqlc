-- name: ListAlbums :many
SELECT * FROM album;

-- name: ListAlbumsWithArtist :many
SELECT sqlc.embed(album), sqlc.embed(artist) FROM album join artist on album.artist_id = artist.artist_id;

-- name: ListAlbumsCustom :many
SELECT album_id, title FROM album;
