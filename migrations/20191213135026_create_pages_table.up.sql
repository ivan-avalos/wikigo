CREATE TABLE IF NOT EXISTS pages(
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    body TEXT NOT NULL,
    PRIMARY KEY (id)
);