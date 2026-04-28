CREATE TABLE IF NOT EXISTS watchlist(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    movie_id INT NOT NULL REFERENCES movies(id)
)