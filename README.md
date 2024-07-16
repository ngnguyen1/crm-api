# CRM backend

## Get all customers

```bash
curl http://localhost:3000/customers
```

## Get a customers

```bash
curl http://localhost:3000/customers/{id}
```

## Create new customer

```bash
curl -H "Content-Type: application/json" -X POST --data '{"id": 3, "FirstName":"Nga","LastName":"Nguyen", "Email": "ngand1@fpt.com"}' http://localhost:3000/customers
```

## Update a customer

```bash
curl -H "Content-Type: application/json" -X PUT --data '{"id": 2, "FirstName":"Nga-updated","LastName":"Nguyen", "Email": "ngand1@fpt.vn"}' http://localhost:3000/customers/2
```

## Delete a customer

```bash
curl -X DELETE http://localhost:3000/customers/{id}
```