# Contributing

Contributing and pull requests are very welcome.

## Requirements

- go 1.19
- node v18.13.0
- yarn

There is a `.nvmrc` in the `ui` directory. If you use nvm you can easily run `nvm use` to configure the correct node version.

## Project structure

tvhgo contains a backend server written in golang which is located in the root directory of the repository. The web ui is located at `ui`.
The production build of the web ui is served by the backend server.

## Development

Make sure you have correctly configured the backend server using environment variables or the `config.yml`. Further information about configuring the application can be found [here](https://github.com/davidborzek/tvhgo/wiki/Configuration).

Start the backend server:

```bash
# Install dependencies
go mod download

# Start the server
go run main.go
```

Now go to the `ui` directory and run:

```bash
# Install dependencies
yarn

# Start the development server
yarn dev
```

## Linting and formatting

You can format the backend code using the standard go formatter:

```bash
go fmt ./...
```

To format the frontend code use the following command in the `ui` directory:

```bash
yarn format
```

## Tests

Run the backend tests using:

```bash
go test ./...
```

## Building manually

There are github workflows to build a docker image containing the backend server and the static frontend.
But you can also test the build with these manual steps.

Go to the `ui` directory and build the frontend:

```bash
yarn build
```

Now you can go back to the root and build the backend server:

```bash
go build -o tvhgo main.go
```

The static frontend production build in `ui/dist` is embedded into the server binary using `go:embed`.
