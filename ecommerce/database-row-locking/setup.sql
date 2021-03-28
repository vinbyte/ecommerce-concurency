SET timezone = 'Asia/Jakarta';

-- setup products and data
CREATE TABLE public.products (
	id serial NOT NULL,
	code varchar(200) NOT NULL,
	"name" varchar(200) NOT NULL,
	description text NOT NULL,
	stock int4 NOT NULL,
	price int8 NOT NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id),
	CONSTRAINT products_un UNIQUE (code)
);
INSERT INTO products (code, "name", description, stock, price) VALUES('P1', 'iPhone 12 64GB Red', 'Iphone 12 with 64GB Storage. Product Red', 3, 1000);
INSERT INTO products (code, "name", description, stock, price) VALUES('P2', 'Macbook Pro M1 8GB Ram 246GB', 'The new macbook pro M1 8GB RAM with 256GB SSD', 2, 5000);

-- setup users dan data
CREATE TABLE public.users (
	id serial NOT NULL,
	"name" varchar NOT NULL,
	email varchar NOT NULL,
	address text NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);
INSERT INTO users ("name", email, address) VALUES('User 1', 'user1@mailinator.com', 'Address User 1');
INSERT INTO users ("name", email, address) VALUES('User 2', 'user2@mailinator.com', 'Address User 2');
INSERT INTO users ("name", email, address) VALUES('User 3', 'user3@mailinator.com', 'Address User 3');
INSERT INTO users ("name", email, address) VALUES('User 4', 'user4@mailinator.com', 'Address User 4');
INSERT INTO users ("name", email, address) VALUES('User 5', 'user5@mailinator.com', 'Address User 5');
INSERT INTO users ("name", email, address) VALUES('User 6', 'user6@mailinator.com', 'Address User 6');
INSERT INTO users ("name", email, address) VALUES('User 7', 'user7@mailinator.com', 'Address User 7');
INSERT INTO users ("name", email, address) VALUES('User 8', 'user8@mailinator.com', 'Address User 8');
INSERT INTO users ("name", email, address) VALUES('User 9', 'user9@mailinator.com', 'Address User 9');
INSERT INTO users ("name", email, address) VALUES('User 10', 'user10@mailinator.com', 'Address User 10');

-- setup cart
CREATE TABLE public.cart (
	id serial NOT NULL,
	user_id int4 NOT NULL,
	"date" timestamptz(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT cart_pkey PRIMARY KEY (id),
	CONSTRAINT cart_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- setup cart_items
CREATE TABLE public.cart_items (
	id serial NOT NULL,
	cart_id int2 NOT NULL,
	product_code varchar NOT NULL,
	quantity int4 NOT NULL,
	"date" timestamptz(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT cart_items_pkey PRIMARY KEY (id),
	CONSTRAINT cart_items_fk FOREIGN KEY (cart_id) REFERENCES cart(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- setup orders
CREATE TABLE public.orders (
	id serial NOT NULL,
	user_id int4 NOT NULL,
	is_paid bool NOT NULL DEFAULT false,
	"date" timestamptz(0) NOT NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (id),
	CONSTRAINT orders_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- setup order_items
CREATE TABLE public.order_items (
	id serial NOT NULL,
	product_code varchar NOT NULL,
	quantity int4 NOT NULL,
	order_id int4 NOT NULL,
	"date" timestamptz(0) NOT NULL,
	CONSTRAINT order_items_pkey PRIMARY KEY (id),
	CONSTRAINT order_items_fk FOREIGN KEY (order_id) REFERENCES orders(id) ON UPDATE CASCADE ON DELETE CASCADE
);
