source ./.env

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/color.csv' INTO TABLE colors FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n'"

mysql -u$MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE --local-infile=1 -e "LOAD DATA LOCAL INFILE './csv/area.csv' INTO TABLE areas FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n'"