# live-reloading-code

### How do you execute your go projects?

Maybe your first option is:
* Execute with go run <main>
* I have a Makefile, so I only need to run make build

The previous isn’t bad, it works but you lost time in start, stop the process, move to the terminal, etc.

I present [AIR](https://github.com/cosmtrek/air) that reloads your changes when you codding. Air it's used in your machine or inside the docker.

### Usage

* Install air in your machine
* Inside docker


### Test

Example: 1

```sh
# air is installed in your machine
$ air main.go
```

Example: 2

```sh
# air is installed in your docker image
$ docker-compose up
```

__For each change that you make applies the refresh__

Happy coding!