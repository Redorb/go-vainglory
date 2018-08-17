# go-vainglory
Vainglory API Wrapper

## documentation
Documentation can be found at [here](https://godoc.org/github.com/redorb/go-vainglory).

## testing
In order to run the tests locally you will need to add a conf.json file in the root folder

In the file you will need to add the following:
> {
>   "key"       : "[API key]",
>   "rateLimit" : [# of requests per minute]
>}

## dependencies
Though this client only has one [dependency](https://github.com/google/go-querystring) the project uses the new go module system. More information on them can be found [here](https://github.com/golang/go/wiki/Modules).
