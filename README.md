# ik-app

REST API Structure

Endpoints:

ik.kz/
ik.kz/merch/
ik.kz/media/
ik.kz/audio/
ik.kz/video/
ik.kz/show/
ik.kz/about/
ik.kz/contacts/

GET		  ik.kz/merch/		  - get from all categories
GET		  ik.kz/merch/{id}	- get from specific category
POST	  ik.kz/merch/		  - create category
PATCH	  ik.kz/merch/{id}	- update category
DELETE	ik.kz/merch/{id}	- delete category 
...

Database:

Categories table
id	          integer
title	        varchar(255)
description   varchar(255)

Item table
id	          integer
title	        varchar(255)
description	  varchar(255)
quantity	    integer
size	        integer
color	        integer

List table
id	          integer
category_id   integer
item_id	      integer
