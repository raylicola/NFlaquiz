set global local_infile = 1;

USE nflaquiz;

CREATE TABLE areas
(
  id VARCHAR(30),
  name VARCHAR(10),
  PRIMARY KEY(id)
);

CREATE TABLE colors
(
  id VARCHAR(10),
  name VARCHAR(10),
  PRIMARY KEY(id)
);

CREATE TABLE countries
(
  id VARCHAR(2),
  area_id VARCHAR(30),
  name VARCHAR(20),
  description VARCHAR(200),
  FOREIGN KEY (area_id) REFERENCES areas(id),
  PRIMARY KEY(id)
);

CREATE TABLE flag_colors
(
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  country_id VARCHAR(2),
  color_id VARCHAR(10),
  FOREIGN KEY (country_id) REFERENCES countries(id),
  FOREIGN KEY (color_id) REFERENCES colors(id),
  PRIMARY KEY(id)
);

CREATE TABLE quizzes
(
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  hiragana VARCHAR(30),
  country_id VARCHAR(2),
  hint1 VARCHAR(50),
  hint2 VARCHAR(50),
  hint3 VARCHAR(50),
  FOREIGN KEY (country_id) REFERENCES countries(id),
  PRIMARY KEY(id)
);

CREATE TABLE users
(
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  email VARCHAR(100),
  password LONGBLOB,
  PRIMARY KEY(id)
);

CREATE TABLE bookmarks
(
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  country_id VARCHAR(2),
  FOREIGN KEY (country_id) REFERENCES countries(id),
  PRIMARY KEY(id)
);

CREATE TABLE quiz_results
(
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  country_id VARCHAR(2),
  user_id INT,
  FOREIGN KEY (country_id) REFERENCES countries(id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  PRIMARY KEY(id)
);