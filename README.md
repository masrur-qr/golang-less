# golang-less

bson.M{"date": "2024-07-01", "item": "Laptop", "quantity": 1, "price": 1000, "store": "A"},
		bson.M{"date": "2024-07-01", "item": "Keyboard", "quantity": 2, "price": 50, "store": "A"},
		bson.M{"date": "2024-07-01", "item": "Mouse", "quantity": 3, "price": 25, "store": "B"},
		bson.M{"date": "2024-07-02", "item": "Monitor", "quantity": 1, "price": 200, "store": "B"},
		bson.M{"date": "2024-07-02", "item": "Laptop", "quantity": 2, "price": 1000, "store": "A"},
		bson.M{"date": "2024-07-03", "item": "Keyboard", "quantity": 1, "price": 50, "store": "B"},
		bson.M{"date": "2024-07-03", "item": "Mouse", "quantity": 1, "price": 25, "store": "A"},

++=====================
[
    { "_id": 1, "order_id": "ORD1001", "customer_id": "C001", "item": "Laptop", "quantity": 1, "price": 1000 },
    { "_id": 2, "order_id": "ORD1002", "customer_id": "C002", "item": "Mouse", "quantity": 2, "price": 50 },
    { "_id": 3, "order_id": "ORD1003", "customer_id": "C001", "item": "Keyboard", "quantity": 1, "price": 100 },
    { "_id": 4, "order_id": "ORD1004", "customer_id": "C003", "item": "Monitor", "quantity": 1, "price": 300 }
]
[
    { "_id": "C001", "name": "John Doe", "email": "john.doe@example.com" },
    { "_id": "C002", "name": "Jane Smith", "email": "jane.smith@example.com" },
    { "_id": "C003", "name": "Jim Brown", "email": "jim.brown@example.com" }
]
