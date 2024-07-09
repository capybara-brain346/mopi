package models

type MovieSchema struct {
	Name     string
	Rating   string
	Genre    string
	Year     string
	Released string
	Score    string
	Votes    string
	Director string
	Writer   string
	Star     string
	Country  string
	Budget   string
	Gross    string
	Company  string
	Runtime  string
}

// var MovieDB = map[int]MovieSchema{
// 	1:  {Title: "Indiana Jones", ReleaseDate: "10/12/1981"},
// 	2:  {Title: "The Matrix", ReleaseDate: "31/3/1999"},
// 	3:  {Title: "The Godfather", ReleaseDate: "24/3/1972"},
// 	4:  {Title: "The Dark Knight", ReleaseDate: "18/7/2008"},
// 	5:  {Title: "Pulp Fiction", ReleaseDate: "14/10/1994"},
// 	6:  {Title: "The Lord of the Rings: The Return of the King", ReleaseDate: "17/12/2003"},
// 	7:  {Title: "Star Wars: Episode V - The Empire Strikes Back", ReleaseDate: "21/5/1980"},
// 	8:  {Title: "Forrest Gump", ReleaseDate: "6/7/1994"},
// 	9:  {Title: "Inception", ReleaseDate: "16/7/2010"},
// 	10: {Title: "Fight Club", ReleaseDate: "15/10/1999"},
// 	11: {Title: "The Shawshank Redemption", ReleaseDate: "14/10/1994"},
// 	12: {Title: "Goodfellas", ReleaseDate: "19/9/1990"},
// 	13: {Title: "Schindler's List", ReleaseDate: "30/11/1993"},
// 	14: {Title: "The Usual Suspects", ReleaseDate: "19/8/1995"},
// 	15: {Title: "Gladiator", ReleaseDate: "5/5/2000"},
// 	16: {Title: "The Departed", ReleaseDate: "6/10/2006"},
// 	17: {Title: "The Prestige", ReleaseDate: "20/10/2006"},
// 	18: {Title: "The Green Mile", ReleaseDate: "10/12/1999"},
// 	19: {Title: "Avatar", ReleaseDate: "18/12/2009"},
// 	20: {Title: "The Silence of the Lambs", ReleaseDate: "14/2/1991"},
// }
