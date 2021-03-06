FORMAT: 1A
HOST: http://user-service/

# User-Service

The user-Service is responsible to handle *User* related tasks.
This includes, creating and deleting a user using the *Database-Adapter*,
as well as handling all requests concerning *User* *Items*.

## Users Collection [/users]

### Create a new User [POST]

Use this endpoint to create a new User dataset.

+ Request (application/json)

        {
        }

+ Response 200 (application/json)

        {
        }
        
## Specific User Collection [/users/{userId}]

### Get a specific User [GET]

This endpoint will response the user data of the requested user id.

+ Request (application/json)

        {
            
        }

+ Response 200 (application/json)

        {
            
        }
            
### Delete a specific User [DELETE]

This endpoint deletes the user for the given id.

+ Request (application/json)

        {
        }

+ Response 200 (application/json)

        {

        }
        
## User Items Collection [/users/{userId}/items]

### Add a new Item to a User [POST]

Using this endpoint, a new item can be added to a specific user.

+ Request (application/json)

        {
            
        }

+ Response 200 (application/json)

        {
            
        }

### Get all Items of a User [GET]

This endpoint return all items that are associated to a specific user.

+ Response 200 (application/json)

        {
            
        }
        
## Specific User Items Collection [/users/{userId}/items/{itemId}]

### Update a specific Item [PUT]

Via this endpoint, an existing item, associated to a specific user, can be altered.

+ Request (application/json)

        {
            
        }

+ Response 200 (application/json)

        {
            
        }
        
### Delete a specific Item [DELETE]

This endpoint deletes the item with the given id.

+ Request (application/json)

        {
            
        }

+ Response 200 (application/json)

        {
            
        }

### Get a specific Item [GET]

This endpoint returns the user item with the given id.

+ Response 200 (application/json)

        {
            
        }
