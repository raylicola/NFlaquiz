source ./.env

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/Employee.csv' INTO TABLE employee FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n'"