# Store-Manager-Pro
Store-Manager-Pro is the server for the Store Manager Pro system. This program runs on the servers and then can be accessed through a web browser or using the Store Manager Pro Inventory windows app.


# Running
## Dev
To run this for development create a data folder in the root directory of the project.
```bash
mkdir data
```

Next go to the `src` folder and copy the `.env.example` to `.env`.
```bash
cp .env.example .env
```

Finally run `go run .`.
```bash
go run .
```

Then 

## Docker Compose
Download the create a `store-manager-pro` folder and create a file called `docker-compose.yaml`. Open this folder in your favorite text editor and paste in the `docker-compose.yaml` file.
```bash
mkdir store-manager-pro
cd store-manager-pro
touch docker-compose.yaml

nano docker-compose.yaml
```

Next start the docker image.
```bash
sudo docker compose up -d
```


# Building
Here is the information on how to build this with docker and with go cli.

## Go CLI
To build this into a deployable binary use this command. Make be sure to update the build version

```bash
go build -ldflags "-X gitea.locker98.com/locker98/Store-Manager-Pro/utils.Version=v0.8.1 -X gitea.locker98.com/locker98/Store-Manager-Pro/utils.Commit=$(git rev-parse --short HEAD)"
```

### Docker
Here are the commands to build the storemanagerpro-server docker image.

```bash
# Build for multi platform deployment
docker buildx build --platform linux/arm64,linux/amd64 --build-arg VERSION=v0.8.1 --build-arg COMMIT=$(git rev-parse --short HEAD) -t gitea.locker98.com/locker98/store-manager-pro:latest -t gitea.locker98.com/locker98/store-manager-pro:v0.8.1 . --push
```
