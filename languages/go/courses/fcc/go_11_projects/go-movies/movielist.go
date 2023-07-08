package main

var movies []Movie

func populateMovieList() {
	movies = append(movies, Movie{ID: "1", ISBN: "438", Title: "Die Smart", Director: &Director{FirstName: "John", LastName: "McClain"}})
	movies = append(movies, Movie{ID: "2", ISBN: "138", Title: "Straight Outta Bellville", Director: &Director{FirstName: "Koos", LastName: "Kaalarm"}})
	movies = append(movies, Movie{ID: "3", ISBN: "736", Title: "Tassies For Life", Director: &Director{FirstName: "Adam", LastName: "Tas"}})
}
