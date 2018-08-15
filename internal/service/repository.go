package service

import (
	"fmt"

	"github.com/PrakharSrivastav/album-service-grpc/internal/model"
	"github.com/brianvoe/gofakeit"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type repository struct {
	db *sqlx.DB
}

func (r *repository) get(ID string) (*model.Album, error) {
	var album model.Album
	if err := r.db.Get(&album, "Select * from albums where id = $1", ID); err != nil {
		return nil, err
	}
	return &album, nil
}

func (r *repository) getAll() ([]*model.Album, error) {
	var albums []*model.Album
	if err := r.db.Select(&albums, "Select * from albums"); err != nil {
		return nil, err
	}
	return albums, nil
}

func (r *repository) getIn(albumIds []string) ([]*model.Album, error) {
	// {"SELECT * FROM foo WHERE x in (?)",[]interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8},8}
	var albums []*model.Album
	if err := r.db.Select(&albums, "Select * from albums  where id in (?) ", albumIds); err != nil {
		return nil, err
	}
	return albums, nil
}

func (r *repository) setupDatabase() error {
	r.db.MustExec(schema)
	var albums []model.Album

	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_1", ID: "album_id_1", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_2", ID: "album_id_2", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_1", ID: "album_id_3", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_5", ID: "album_id_4", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_10", ID: "album_id_5", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_10", ID: "album_id_6", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_5", ID: "album_id_7", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_1", ID: "album_id_8", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_10", ID: "album_id_9", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_3", ID: "album_id_10", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_1", ID: "album_id_11", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_4", ID: "album_id_12", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_9", ID: "album_id_13", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_9", ID: "album_id_14", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_2", ID: "album_id_15", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_1", ID: "album_id_16", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_5", ID: "album_id_17", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_8", ID: "album_id_18", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_8", ID: "album_id_19", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_2", ID: "album_id_20", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_8", ID: "album_id_21", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_5", ID: "album_id_22", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_1", ID: "album_id_23", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_7", ID: "album_id_24", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_3", ID: "album_id_25", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_7", ID: "album_id_26", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_3", ID: "album_id_27", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_4", ID: "album_id_28", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_6", ID: "album_id_29", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_4", ID: "album_id_30", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_6", ID: "album_id_31", Name: gofakeit.HipsterSentence(4)})
	albums = append(albums, model.Album{Description: gofakeit.Sentence(5), ArtistID: "artist_id_4", ID: "album_id_32", Name: gofakeit.HipsterSentence(4)})
	tx := r.db.MustBegin()
	for _, u := range albums {
		fmt.Printf("%#v\n", &u)
		_, err := tx.NamedExec("INSERT into albums (id,name,description,artistId) values(:id,:name,:description,:artistId)", u)
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return nil

}

var schema = `
Drop Table if exists albums;
CREATE TABLE albums (
    id text,
	name text,
	artistId text,
    description text
);`
