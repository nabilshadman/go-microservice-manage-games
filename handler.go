package main

import (
  "encoding/json"
  "net/http"
  "time"

  "github.com/gorilla/mux"
)

// CreateGameHandler reads in JSON data and creates a new Game within the database
// Returns HTTP Code 201 and JSON data of the newly created game
func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
  var game Game
  defer r.Body.Close()
  err := json.NewDecoder(r.Body).Decode(&game)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(
      &JsonErr{Error: "unable to create game, please check your data"},
    )
    return
  }

  var lastInsertId int
  err = db.QueryRow(
    "INSERT INTO games(title, console, rating, complete, created, updated) VALUES($1, $2, $3, $4, $5, $5) returning id",
    game.Title,
    game.Console,
    game.Rating,
    game.Complete,
    time.Now(),
  ).Scan(&lastInsertId)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(
      &JsonErr{Error: "unable to write to database: " + err.Error()},
    )
    return
  }

  game = Game{}
  db.QueryRow("SELECT * FROM games g WHERE g.id = $1", lastInsertId).Scan(
    &game.ID,
    &game.Title,
    &game.Console,
    &game.Rating,
    &game.Complete,
    &game.Created,
    &game.Updated,
  )

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(&game)
}

// RetrieveGameHandler reads the ID from the URL (using mux) and queries the database for
// the game in question.
// Returns 200 and JSON data if the game exists, otherwise returns a 404
func RetrieveGameHandler(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]

  var game Game
  err := db.QueryRow("SELECT * FROM games g WHERE g.id = $1", id).Scan(
    &game.ID,
    &game.Title,
    &game.Console,
    &game.Rating,
    &game.Complete,
    &game.Created,
    &game.Updated,
  )

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(&game)
}

// UpdateGameHandler queries the database for the ID of the game to be updated.
// Replaces the writeable data within the game with the JSON data,
// and rewrites the game to the database.
// Returns HTTP Code 200 with newly update game data
func UpdateGameHandler(w http.ResponseWriter, r *http.Request) {
  	params := mux.Vars(r)
	id := params["id"]

	var game Game
	err := db.QueryRow("SELECT * FROM games g WHERE g.id = $1", id).Scan(
		&game.ID,
		&game.Title,
		&game.Console,
		&game.Rating,
		&game.Complete,
		&game.Created,
		&game.Updated,
	)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// handle "read-only" attributes
	created := game.Created
	updated := time.Now()
	gid := game.ID

	// overwrite data with provided JSON
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			&JsonErr{Error: "bad JSON data, please check the update data"},
		)
		return
	}

	// update game and then run query
	game.ID = gid
	game.Created = created
	game.Updated = updated


	_, err = db.Exec(
		"UPDATE games g SET title = $1, console = $2, rating = $3, complete = $4, created = $5, updated = $6 WHERE g.id = $7",
		game.Title,
		game.Console,
		game.Rating,
		game.Complete,
		game.Created,
		game.Updated,
		game.ID,
	)

	json.NewEncoder(w).Encode(&game)

}

// DeleteGameHandler removes the game from the database based off the ID passed in through the URL.
// Returns HTTP Code 204 if successful, 404 if no game exists for ID
func DeleteGameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec("DELETE FROM games g WHERE g.id = $1", id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			&JsonErr{Error: "unable to delete game: " + err.Error()},
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RetrieveGamesHandler returns all games in collection as a JSON array
func RetrieveGamesHandler(w http.ResponseWriter, r *http.Request) {
  games := make([]Game, 0)

  rows, err := db.Query("SELECT * FROM games g ORDER BY g.id")
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(
      &JsonErr{Error: "unable to retreive games at this time: " + err.Error()},
    )
    return
  }

  for rows.Next() {
    var game Game
    rows.Scan(
      &game.ID,
      &game.Title,
      &game.Console,
      &game.Rating,
      &game.Complete,
      &game.Created,
      &game.Updated,
    )

    games = append(games, game)
  }

  json.NewEncoder(w).Encode(games)
}

// RetrieveGamesByConsoleHandler is a simple filter/search method to return all games of a specific console
// Reads the console from the URL and queries the database.
// Returns HTTP Code 200 with all game data, if sent an "invalid" console, an empty array is returned
func RetrieveGamesByConsoleHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("not implemented\n"))
}
