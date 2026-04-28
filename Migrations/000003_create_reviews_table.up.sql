CREATE TABLE IF NOT EXISTS reviews(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    movie_id INT NOT NULL REFERENCES movies(id),
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 10),
    text TEXT NOT NULL
)