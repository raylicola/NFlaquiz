set global local_infile = 1;

USE nflaquiz;

CREATE TABLE area
(
  id VARCHAR(30),
  name VARCHAR(10),
  PRIMARY KEY(id)
);

CREATE TABLE color
(
  id VARCHAR(10),
  name VARCHAR(10),
  PRIMARY KEY(id)
);

CREATE TABLE country
(
  id VARCHAR(2),
  area_id VARCHAR(30),
  name VARCHAR(20),
  description VARCHAR(200),
  FOREIGN KEY (area_id) REFERENCES area(id),
  PRIMARY KEY(id)
);

CREATE TABLE flag_color
(
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  country_id VARCHAR(2),
  color_id VARCHAR(10),
  FOREIGN KEY (country_id) REFERENCES country(id),
  FOREIGN KEY (color_id) REFERENCES color(id),
  PRIMARY KEY(id)
);

CREATE TABLE quiz
(
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  hiragana VARCHAR(30),
  country_id VARCHAR(2),
  FOREIGN KEY (country_id) REFERENCES country(id),
  PRIMARY KEY(id)
);

CREATE TABLE hint
(
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  content VARCHAR(200),
  country_id VARCHAR(2),
  FOREIGN KEY (country_id) REFERENCES country(id),
  PRIMARY KEY(id)
);