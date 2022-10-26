
CREATE TABLE "Orders" (
  "order_id" int PRIMARY KEY NOT NULL,
  "costumer_id" int NOT NULL,
  "product_id" int NOT NULL,
  "order_date" date NOT NULL,
  "total" double precision NOT NULL
);

CREATE TABLE "customers" (
  "id" int PRIMARY KEY NOT NULL,
  "costumer_name" char(50) NOT NULL
);

CREATE TABLE "products" (
  "id" int PRIMARY KEY NOT NULL,
  "name" char(50) NOT NULL
);

ALTER TABLE "Orders" ADD FOREIGN KEY ("costumer_id") REFERENCES "customers" ("id");

ALTER TABLE "Orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

/* tabel Aqua.products */
INSERT INTO products(id, name) VALUES (1, 'ikan');
UPDATE "products" SET "name" = 'udang' WHERE "name" = 'ikan';
DELETE FROM "products" WHERE "id" = 1;


/* tabel Aqua.costumers */
INSERT INTO "customers" (id, "costumer_name") VALUES (1, 'dandi');
UPDATE customers SET "costumer_name" = 'dodo' WHERE "costumer_name" = 'dandi';
DELETE FROM customers WHERE "costumer_name" = 'dodo';


/* tabel Aqua.Orders */
INSERT INTO "Orders" (order_id, costumer_id, product_id, order_date, "total") VALUES (1, 1, 1, '2022-10-26', 20000);
UPDATE "Orders" SET "total" = 40000 WHERE "total" = 20000;
DELETE FROM "Orders" WHERE order_id  = 1;


/* Bonus Round */
SELECT
  "Orders",
  "customers",
  "products"
FROM "Orders"
JOIN "customers"
  ON "customers".id = "Orders".costumer_id
JOIN "products"
  ON "products".id = "Orders".product_id;