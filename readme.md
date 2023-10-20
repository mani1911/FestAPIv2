# FestAPI

---

### Requirements

- [Go](https://go.dev/)
- [golangci](https://golangci-lint.run/usage/install/)
- [Docker](https://www.docker.com/)
- [Reflex](https://github.com/cespare/reflex)
- [Swagger](https://github.com/swaggo/swag)


### Setup

- Configure .vscode/settings.json
  ```
  {
      "go.lintTool":"golangci-lint",
      "go.lintFlags": [
      "--fast"
      ],
      "go.lintOnSave": "package",
      "go.formatTool": "goimports",
      "go.useLanguageServer": true,
      "[go]": {
          "editor.formatOnSave": true,
          "editor.codeActionsOnSave": {
              "source.organizeImports": true
          }
      },
      "go.docsTool": "gogetdoc"
  }
  ```
- Create .env file
  ```sh
  cp .env.example .env
  ```
- Enable githooks
  ```sh
  git config core.hooksPath .githooks
  ```

### Swagger

- To update swag docs, Run 
  ```sh
    make swag_init
  ```

- To format swaggo docs, Run
  ```sh
    make swag_fmt
  ```

### Seeding Database

- to seed admin table

  ```
  make seed_admin
  ```
- to seed other tables

  ```
  make seed_database
  ```

### Run

- #### On Local
  - To run the server

    ```sh
    make run
    ```
  
  - To run the server in watch mode
  
    ```sh
    make watch
    ```

- #### On Docker
  ```sh
  docker compose up
  ```

### Linting Errors

- #### Golangci-lint Not Found

  - Install golangci-lint from [here](https://golangci-lint.run/usage/install/#local-installation)
  - Add golangci-lint to PATH

- #### Permission denied for some folder (usually with docker volumes)
  - Change the permission of the folder
    ```sh
    sudo chmod -R 777 <folder>
    ```
