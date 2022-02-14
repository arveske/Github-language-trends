# Github trending languages API

This API provides list of most popular github languages, used in 100 most stared repositories for last 30 days.

The entire application is contained within the `main.go` file.

Application also uses Github API route `https://api.github.com/search/repositories?q=created:>{date}&sort=stars&order=desc&per_page=100`.

## Install

    go get .

## Run the app

    go run main.go

# API ROUTES

## Get list of Languages

### Request

`GET /languages/`

    curl -i -H 'Accept: application/json' http://localhost:8080/languages

### Response

    Status: 200 OK
    Content-Type: application/json

    {
    "C": {
        "repository_count": 8,
        "repository_list": [
            "https://github.com/berdav/CVE-2021-4034",
            "https://github.com/arthepsy/CVE-2021-4034",
            "https://github.com/ly4k/PwnKit",
            "https://github.com/0voice/kernel_new_features",
            "https://github.com/bytedance/android-inline-hook",
            "https://github.com/Crusaders-of-Rust/CVE-2022-0185",
            "https://github.com/TheScienceElf/TI-84-CE-Raytracing",
            "https://github.com/thefLink/RecycledGate"
        ]
    },
    ...
    }

## Get a specific Language

### Request

`GET /languages/{name}`

    curl -i -H 'Accept: application/json' http://localhost:8080/languages/Go

### Response

    Status: 200 OK
    Content-Type: application/json

    {
    "Go": {
        "repository_count": 5,
        "repository_list": [
            "https://github.com/utkusen/wholeaked",
            "https://github.com/nikolaydubina/go-binsize-treemap",
            "https://github.com/mergestat/timediff",
            "https://github.com/ajeetdsouza/clidle",
            "https://github.com/cmars/oniongrok"
        ]
        }
    }

## Get a non-existent Language

### Request

`GET /languages/{name}`

    curl -i -H 'Accept: application/json' http://localhost:8080/languages/Pod

### Response

    Status: 404 Not Found
    Content-Type: application/json

    {
    "message": "language not found"
    }
