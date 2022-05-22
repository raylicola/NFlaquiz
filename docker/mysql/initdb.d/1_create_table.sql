set global local_infile = 1;

USE nflaquiz;

CREATE TABLE employee
(
  id           INT(10),
  name     VARCHAR(40)
);

INSERT INTO employee (id, name) VALUES (1, "Nagaoka");
INSERT INTO employee (id, name) VALUES (2, "Tanaka");
