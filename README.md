# [GO - CRUD] person-api

![Image](assets/images/golang.png)
> A simple CRUD made with GO and Mux
> Application that runs the API [React Native](https://)

## Endpoints
### Persons
- GET - /persons
- GET - /person/{id}
- POST - /person/create
- PUT - /person/update/{id}
- DELETE - /person/delete/{id}

## Installation
```sh
$ git clone https://github.com/TacioAntonio/go-crud-person-api.git
$ cd go-crud-person-api
$ go mod tidy
$ go run main.go
```

# Build
```sh
$ go build main.go
```

# Create a tunnel
Create a tunnel to enable access to a locally hosted API.
```sh
$ start ngrok.sh
```

## Contribution
Please read [CONTRIBUTING.md](https://github.com/TacioAntonio/go-crud-person-api/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/TacioAntonio/go-crud-person-api/tags).

## Authors
| ![Tácio Antônio](https://avatars2.githubusercontent.com/u/44682965?s=150&=4)
| -
| [Tácio Antônio](https://github.com/TacioAntonio/)

See also the list of [contributors](https://github.com/TacioAntonio/go-crud-person-api/graphs/contributors) who participated in this project.

## License
This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/TacioAntonio/go-crud-person-api/blob/master/LICENSE.md) file for details.
