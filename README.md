# ik-app

## REST API Structure

### Endpoints:

```
GET     ik.kz/category/      - get from all categories
GET     ik.kz/category/{id}  - get from specific category
POST    ik.kz/category/      - create category
PATCH   ik.kz/category/{id}  - update category
DELETE  ik.kz/category/{id}  - delete category
...
```

### Database:

#### Category table
```
id            integer
title         varchar(255)
description   varchar(255)
```

#### Item table
```
id            integer
title         varchar(255)
description   varchar(255)
quantity      integer
size          integer
color         integer
```

#### List table
```
id            integer
category_id   integer
item_id       integer
```
