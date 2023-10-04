### Folder Structure :

Project structure reference: [link](https://github.com/achimonchi/simple-ddd)

```
starfish
|--- cmd
|    |--- api
|       |--- main.go
|--- domain
|    |--- auth
|       |--- auth.go
|       |--- controller.go
|       |--- service.go
|       |--- repository.go
|--- config
|--- helpers
|--- migrations
```

### How to run:

1. Create `.env` file and copy content from `.env.example`.

2. Run the command below to initialize the database container.

```
make up
```

3. Install the `migrate` command if you haven't installed it.

- [Windows](https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-windows/)
- [Ubuntu](https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/)
- [Mac](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md#macos)

4. Run the command below to run the database migration.

```
make migrate-up
```

5. Run the command below to run the application

```
make run
```

6. You're ready to go. You can test the server at `http://localhost:3000`
