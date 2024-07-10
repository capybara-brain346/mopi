package database

import (
	"database/sql"
	"log"

	"github.com/capybarabrain346/mopi/models"
	_ "github.com/mattn/go-sqlite3"
)

const file = "handlers/database/movieDB.db"

func DBGet(movieName string) (models.MovieSchema, error) {
	database, err := sql.Open("sqlite3", file)
	if err != nil {
		return models.MovieSchema{}, err
	}
	defer database.Close()

	rows, err := database.Query("SELECT * FROM movie_table WHERE name=?", movieName)
	if err != nil {
		return models.MovieSchema{}, err
	}
	defer rows.Close()

	result := models.MovieSchema{}
	if rows.Next() {
		if err := rows.Scan(&result.Name, &result.Rating, &result.Genre,
			&result.Year, &result.Released, &result.Score,
			&result.Votes, &result.Director, &result.Writer,
			&result.Star, &result.Country, &result.Budget,
			&result.Gross, &result.Company, &result.Runtime); err != nil {
			return models.MovieSchema{}, err
		}
	} else {
		log.Printf("Row not found.")
		return result, sql.ErrNoRows
	}

	return result, nil
}

func DBWrite(create_movie models.MovieSchema) (int64, error) {
	var database, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Println(err)
	}

	statement, err := database.Prepare(`INSERT INTO movie_table (Name, Rating, Genre, Year, Released, 
									Score, Votes, Director, Writer, Star, Country, Budget, Gross, Company, Runtime) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err)
	}
	result, err := statement.Exec(create_movie.Name, create_movie.Rating, create_movie.Genre, create_movie.Year, create_movie.Released,
		create_movie.Score, create_movie.Votes, create_movie.Director, create_movie.Writer, create_movie.Star, create_movie.Country,
		create_movie.Budget, create_movie.Gross, create_movie.Company, create_movie.Runtime)

	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func DBDelete(delete_movie string) (int64, error) {
	var database, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	statement, err := database.Prepare("DELETE FROM movie_table WHERE name=?")
	if err != nil {
		log.Println(err)
		return -1, err
	}

	result, err := statement.Exec(delete_movie)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return result.RowsAffected()

}
