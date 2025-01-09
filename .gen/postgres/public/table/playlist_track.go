//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var PlaylistTrack = newPlaylistTrackTable("public", "playlist_track", "")

type playlistTrackTable struct {
	postgres.Table

	// Columns
	PlaylistID postgres.ColumnInteger
	TrackID    postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PlaylistTrackTable struct {
	playlistTrackTable

	EXCLUDED playlistTrackTable
}

// AS creates new PlaylistTrackTable with assigned alias
func (a PlaylistTrackTable) AS(alias string) *PlaylistTrackTable {
	return newPlaylistTrackTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PlaylistTrackTable with assigned schema name
func (a PlaylistTrackTable) FromSchema(schemaName string) *PlaylistTrackTable {
	return newPlaylistTrackTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PlaylistTrackTable with assigned table prefix
func (a PlaylistTrackTable) WithPrefix(prefix string) *PlaylistTrackTable {
	return newPlaylistTrackTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PlaylistTrackTable with assigned table suffix
func (a PlaylistTrackTable) WithSuffix(suffix string) *PlaylistTrackTable {
	return newPlaylistTrackTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPlaylistTrackTable(schemaName, tableName, alias string) *PlaylistTrackTable {
	return &PlaylistTrackTable{
		playlistTrackTable: newPlaylistTrackTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newPlaylistTrackTableImpl("", "excluded", ""),
	}
}

func newPlaylistTrackTableImpl(schemaName, tableName, alias string) playlistTrackTable {
	var (
		PlaylistIDColumn = postgres.IntegerColumn("playlist_id")
		TrackIDColumn    = postgres.IntegerColumn("track_id")
		allColumns       = postgres.ColumnList{PlaylistIDColumn, TrackIDColumn}
		mutableColumns   = postgres.ColumnList{}
	)

	return playlistTrackTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		PlaylistID: PlaylistIDColumn,
		TrackID:    TrackIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
