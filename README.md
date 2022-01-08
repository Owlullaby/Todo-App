# todo-app
Simple Todo App

.env file will be created automatically through Github Actions using publish.yml file.

If you wish to try out the code on your own:
Create your own PostgreSQL database on ElephantSQL.
```
CREATE TABLE todos (
    todoid SERIAL PRIMARY KEY,
    description TEXT
)
```
Then adding your postgres url into .env file.
```
POSTGRES_URL=your_postgress_url
```
