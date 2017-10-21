# ecom

* <bold>Ecom</bold> is a less usefull ecommerce platform.
* It provides an API (at least thats the  goal)
* Ensure that config details are provied
## POST - /api/login
#### request  
```json
POST /api/login HTTP/1.1
Content-Type: application/json
{
  "username":"",
  "password":""
}
 ```
 #### response
 ```json
 HTTP/1.1 200 OK
 Content-Type: application/json

 {
   "token":"some jwt token",
   "uuid":"user uuuid"
 }
 ```  

## POST - /api/register
#### request
```json
POST /api/register
Content-Type: application/json

{
  "firstname":"",
  "lastname":"",
  "username":"",
  "location":"",
  "phonenumber":"",
  "email":"",
  "password":""
}

```
 #### response
 ```json

 HTTP/1.1 201 Created
 Content-Type: application/json
 {
   "token":"some jwt token"
 }
 ```

 ## GET /api/products
 Returns alist of avali products.
 This endpoint is not protected and can be accessed by anyone
 - [ ] Add pagination

 ```json
GET /api/products
HOST
HTTP/1.1 200 OK
{
     "count": 1,  
     "items": [{
       "id": 1,
       "uuid": "",
       "name": "",
       "photoid": "",
       "description": "",
       "price": 0,
       "stock": 0,
       "update_date":"0001-01-01T00:00:00Z",
       "quantitypreunit": 0,
       "category": 1,
       "Category": {
         "id": 1,
         "uuid": 0,
         "name": "",
         "description": "",
         "picture": ""
       }
     }]
   }
 ```

 ## GET /api/product/{productid}

 ```json
 GET /api/products
 HOST:
 Content-Type: application/json
 HTTP/1.1 200 OK
 {
    "id": 1,
    "uuid": "",
    "name": "",
    "photoid": "",
    "description": "",
    "price": 0,
    "stock": 0,
    "update_date":"0001-01-01T00:00:00Z",
    "quantitypreunit": 0,
    "category": 1,
    "Category": {
      "id": 1,
      "uuid": 0,
      "name": "Drinks",
      "description": "Thinks to drink",
      "picture": ""
    }
 }
 ```

## Get /api/payments
* Gets a list of payments methods
* No Authentication needed
Request
```json
GET /api/payments HTTP/1.1
```

Response
```json
HTTP/1.1 200 OK
Content-Type: application/json
{
  "payments_methods":[
    {
      "name":"",
      "id":"",
    }
  ]
}
```
## POST /api/order
* Places and order  
* Must be authenticated
```json
POST /api/checkout
Authentication:"jwt auth key"
Content-Type:application/json
{
  "items":[
    {
      "quantity":1,
      "productid":1234,
    }
  ],
  "paymentid":123,
  "shippingadress":"",
}
```

## GET /api/checkout
This is work in progress
```json
POST /api/checkout
Authentication:"jwt auth string"
Content-Type: application/json
 {
  "checkout":{
    "email": "duncan.smith@example.com",
    "line_items": [{
      "item_id": 111,
      "quantity": 1
    }],
    "shipping_address": {
      "first_name": "duncan",
      "last_name": "Smith",
      "address1": "126 .",
      "city": "Nairobi",
      "phone": "+(254)79000000",
      "zip": "4556"
    }
  }
}
 ```

# Work in progress. Feel free to help :+1:
