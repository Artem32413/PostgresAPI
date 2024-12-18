CREATE TABLE "Flowers"(
    id SERIAL NOT NULL,
    "Name" varchar(100) NOT NULL,
    "Quantity" integer NOT NULL DEFAULT 0,
    "Price" double precision NOT NULL,
    "ArrivalDate" varchar(100) NOT NULL,
    PRIMARY KEY(id)
);