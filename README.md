# go-nba-project

client will handle making api calls, returning the bytes, and handle unmarshaling the json into representation structs. 

domain will handle validation, translating to domain structs, and passing down to the repo for persistence.  

repo/db layer will handle persisting domain structs to the DB and retrieving/updating records when necessary

handler layer? maybe? would just be responsible for retrieving records from the db and displaying them to the user