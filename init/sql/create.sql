CREATE TABLE IF NOT EXISTS website (
    id INT GENERATED ALWAYS AS IDENTITY,
    url VARCHAR,
    title_path VARCHAR,
    news_path VARCHAR,
    description_path VARCHAR,
    link_path VARCHAR,
    pubDate_path VARCHAR
);

CREATE TABLE IF NOT EXISTS news (
    id INT GENERATED ALWAYS AS IDENTITY,
    title VARCHAR NOT NULL,
    link VARCHAR NOT NULL,
    description VARCHAR,
    pubDate VARCHAR
);