Users
    Field	Type
    id	uuid
    name	varchar
    email	varchar
    password	varchar
    role	varchar

Categories
    Field	Type
    id	uuid
    name	varchar

Products
    Field	Type
    id	uuid
    category_id	uuid
    name	varchar
    description	text
    price	decimal
    stock	int

Product Images
    Field	Type
    id	uuid
    product_id	uuid
    image_url	varchar

Carts
    Field	Type
    id	uuid
    user_id	uuid

Cart Items
    Field	Type
    id	uuid
    cart_id	uuid
    product_id	uuid
    quantity	int

Orders
    Field	Type
    id	uuid
    user_id	uuid
    total_price	decimal
    status	varchar

Order Items
    Field	Type
    id	uuid
    order_id	uuid
    product_id	uuid
    quantity	int
    price	decimal