# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete users

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```

``` bash
go build
./go_restapi
```

## Endpoints

### Get All users
``` bash
GET api/users
```
### Get Single user
``` bash
GET api/users/{id}
```

### Delete user
``` bash
DELETE api/users/{id}
```

### Create user
``` bash
POST api/users

# Request sample
# {
#   "isbn":"4545454",
#   "title":"user Three",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }
```

### Update user
``` bash
PUT api/users/{id}

# Request sample
# {
#   "name":"Name",
#   "author":{"example@gmail.com}
# }

```


``
