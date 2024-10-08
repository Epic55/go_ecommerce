##**Ecommerce With Golang Project By Epic55**
**Code-Along video series on Youtube**

```bash
# You can start the project with below commands
docker-compose up -d
go run main.go
```

1) **SIGNUP FUNCTION API CALL (POST REQUEST)**

http://localhost:8000/users/signup

```json
{
  "first_name": "Epic55",
  "last_name": "Mazur",
  "email": "Epic55@gmail.com",
  "password": "Epic55Mazur",
  "phone": "+4534545435"
}
```

Response :"Successfully Signed Up!!"

2) **LOGIN FUNCTION API CALL (POST REQUEST)**

  http://localhost:8000/users/login

```json
{
  "email": "Epic55@gmail.com",
  "password": "Epic55Mazur"
}
```
3) **Admin add Product Function (POST REQUEST)**

  http://localhost:8000/admin/addproduct

```json
{
  "product_name": "Alienware x15",
  "price": 2500,
  "rating": 10,
  "image": "alienware.jpg"
}
```

Response : "Successfully added our Product Admin!!"

4) **View all the Products in db GET REQUEST**

  http://localhost:8000/users/productview

Response

```json
[
  {
    "Product_ID": "6153ff8edef2c3c0a02ae39a",
    "product_name": "alienwarex15",
    "price": 1500,
    "rating": 10,
    "image": "alienware.jpg"
  },
  {
    "Product_ID": "616152679f29be942bd9df8f",
    "product_name": "giner ale",
    "price": 900,
    "rating": 5,
    "image": "gin.jpg"
  },
  {
    "Product_ID": "616152ee9f29be942bd9df90",
    "product_name": "iphone 13",
    "price": 1700,
    "rating": 4,
    "image": "ipho.jpg"
  },
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "whiskey",
    "price": 100,
    "rating": 7,
    "image": "whis.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "acer predator",
    "price": 3000,
    "rating": 10,
    "image": "acer.jpg"
  }
]
```

5) **Search Product by regex function (GET REQUEST)**

defines the word search sorting
http://localhost:8000/users/search?name=al

response:

```json
[
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "Alienware x15",
    "price": 1500,
    "rating": 10,
    "image": "1.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "ginger Ale",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  }
]
```

6) **Adding the Products to the Cart (GET REQUEST)**

  http://localhost:8000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx

  Corresponding mongodb query

7) **Removing Item From the Cart (GET REQUEST)**

  http://localhost:8000/removeitem?id=xxxxxxx&userID=xxxxxxxxxxxx

8) **Listing the item in the users cart (GET REQUEST) and total price**  CHECK

  http://localhost:8000/listcart?id=xxxxxxuser_idxxxxxxxxxx

9) **Addding the Address (POST REQUEST)**  CHECK

  http://localhost:8000/addadress?id=user_id**\*\***\***\*\***

  The Address array is limited to two values home and work address more than two address is not acceptable

```json
{
  "house_name": "10",
  "street_name": "Alfarabi",
  "city_name": "Almaty",
  "pin_code": "1234"
}
```

10) **Editing the Home Address(PUT REQUEST)**

  http://localhost:8000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

11) **Editing the Work Address(PUT REQUEST)**

  http://localhost:8000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

12) **Delete Addresses(GET REQUEST)**

  http://localhost:8000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx

  delete both addresses

13) **Cart Checkout Function and placing the order(GET REQUEST)**

  After placing the order the items have to be deleted from cart functonality added

  http://localhost:8000/cartcheckout?id=xxuser_idxxx

14) **Instantly Buying the Products(GET REQUEST)**

  http://localhost:8000/instantbuy?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx
