# go-nba-project

client will handle making api calls and returning the bytes

service will handle unmarshaling the json into representation structs
then those representation structs will be translated to domain structs

repo/db layer will handle persisting domain structs to the DB and retrieving/updating records when necessary

handler layer? maybe? would just be responsible for retrieving records from the db and displaying them to the user