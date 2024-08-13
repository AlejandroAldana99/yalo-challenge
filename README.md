# Yalo-Challenge
Repository to complete the Yalo Challenge :rocket:

## Pre-requisites

1. Golang (1.21 current version for this project)
2. Postman
3. Docker

## Setup

1. Clone the repository:
   
````
git clone https://github.com/AlejandroAldana99/yalo-challenge.git
````


2. Install all the dependencies:

````
go get

go mod tidy
````

3. Execute any of the options to turn on the service:
   
````
go run ./main.go

make run

run the debug (needs .vscode folder)

build the image
````

4. Test the service:
   
It's possible to test the service by the [postman collection](https://github.com/AlejandroAldana99/yalo-challenge/blob/main/yalo.postman_collection) (attached in the repo) or doing a curl, here are some examples:

> Collect interactions
````
curl --location --request POST 'http://localhost:8000/interactions/collect' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "user_id":"c0110dc1-26bf-440f-9879-e0f3a7f5f0e0",
        "product_sku":"cdfa157f-d6e1-4d2a-8b81-3afab457d8d8",
        "action":"add_to_cart",
        "timestamp":"2024-08-11T04:27:51+00:00",
        "duration":9
    },
    {
        "user_id":"c0110dc1-26bf-440f-9879-e0f3a7f5f0e0",
        "product_sku":"82d09729-c725-47c7-9938-3ff666eb094d",
        "action":"add_to_cart",
        "timestamp":"2024-08-11T04:27:51+00:00",
        "duration":5
    }
]'
````

> Get recommendations
````
curl --location --request GET 'http://localhost:8000/recommendations/c0110dc1-26bf-440f-9879-e0f3a7f5f0e0'
````

## Considerations

1. Port: 

By default, the service is running in the 8000 port for all enviroments, you can change this on `.env` or `Dockerfile` files
   
2. .vscode folder:

This folder contains a simple `launch.json` file to debug the service

3. Makefile:
   
If you are using Windows, you're going to need to install `make` to compile this kind of files, [here](https://gnuwin32.sourceforge.net/packages/make.htm) is an option.

4. .env file

This file just contains the port and the enviroment