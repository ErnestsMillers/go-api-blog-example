# golang-api-blog-example
  A simple Blog CRUD API made with Go, Postgres, Fiber, Gorm and Docker. Manage posts and categories.
  
- ## How to install and run (docker)
  Inside the newly cloned project folder run the following command:
  ```bash
  $ docker compose up
  ```

- ### Running the application
  Inside the project, run the following command to run the application:
  ```bash
  $ go run main.go
  ```

- ## API
  #### Create a new post:
  ```bash
  # POST /api/v1/post

  $ curl --request POST \
  --url http://localhost:3000/api/v1/post \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "blog title",
	"text": "blog test",
	"category": 1
  }'

  ```

  #### Get all posts:
  ```bash
  # GET /api/v1/post

  $ curl --request GET \
  --url http://localhost:3000/api/v1/post

  ```

  #### Get a post by id:
  ```bash
  # GET /api/v1/post/1

  $ curl --request GET \
  --url http://localhost:3000/api/v1/post/1

  ```
  
  #### Update a post by id:
  ```bash
  # PUT /api/v1/post/1
 
  $ curl --request PUT \
  --url http://localhost:3000/api/v1/post/1 \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "a post",
	"text": "post text"
  }'

  ```

  #### Delete a post by id:
  ```bash
  # DELETE /api/v1/post/1

  $ curl --request DELETE \
  --url http://localhost:3000/api/v1/post/1

  ```

#### Categories have the same setup such as

  #### Delete a cateogry by id:
  ```bash
  # DELETE /api/v1/category/1

  $ curl --request DELETE \
  --url http://localhost:3000/api/v1/category/1

  ```