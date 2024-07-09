# Mopi: Movie Information API

Mopi is a RESTful API designed to provide comprehensive information about movies. It allows users to retrieve details such as movie titles, release dates, genres, cast members, ratings, and more. This API is useful for developers building applications that require movie-related data, such as recommendation systems, movie catalogues, or entertainment platforms.

## Features

- **Search Movies:** Search for movies by title, genre, release year, or any combination thereof.
- **Retrieve Movie Details:** Get detailed information about a specific movie, including its synopsis, cast, director, and ratings.
- **Filter by Genre:** Filter movies based on genres such as action, comedy, drama, thriller, etc.
- **Rating Information:** Access ratings from various sources (e.g., IMDb, Rotten Tomatoes) to get a comprehensive view of the movie's reception.

## Technologies Used

- **Golang 1.22.5:** Programming language used for building the API with efficient performance and concurrency support.
- **SQLite3:** Lightweight relational database used for storing movie data during development.
- **Go Chi HTTP Router:** Lightweight, idiomatic router for building Go 1.7+ HTTP services with a focus on simplicity.

## Getting Started

To get started with Mopi, follow these steps:

### Prerequisites

- Golang 1.22.5 or higher installed on your machine.
- SQLite3 database installed locally (for development).
- OAuth2 credentials (if implementing authentication).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/mopi.git
   cd mopi
   ```
2. Build and run the application:

   ```
   go build -o mopi .
   ./mopi
   ```
