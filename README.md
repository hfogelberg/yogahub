# Yogahub

## Build css from sass
````
$ wt watch sass/main.scss -s compressed -b public/
````

## Build and Deploy on Zeit
````
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o yogahub .
$ docker build -t yogahub .
$ now -e GOOGLE_CALLBACK_URL="..." -e GOOGLE_CLIENT_ID="..." -e GOOGLE_CLIENT_SECRET="..." -e YOGA_DB_HOST="..."
````


## Run Dockerized app locally
````
$ docker run -d -p 3000:3000 yogahub
````