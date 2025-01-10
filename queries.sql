-- name: ListAlbums :many
SELECT * FROM album;

-- name: ListAlbumsWithArtist :many
SELECT sqlc.embed(album), sqlc.embed(artist) FROM album join artist on album.artist_id = artist.artist_id;

-- name: ListAlbumsCustom :many
SELECT album_id, title FROM album;

-- name: FindTracks :many
select track.track_id, track.name, track.composer, album.title
from track
join album on track.album_id = album.album_id
where
    (NOT @by_album_title::boolean or album.title = @album_title) AND
    (NOT @by_composer::boolean or track.composer = @composer);
