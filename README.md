# freeforex

Client for freeforexapi.com API.

## Install

````bash
go get github.com/will-evil/freeforex
````

## Usage
````
import github.com/will-evil/freeforex

func main() {
    client := freeforex.Client{}
    
    rate, err := client.Rate("EURGBP")
    if err != nil {
        panic(err)
    }
    
    ...
}
````