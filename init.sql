-- CREATE TABLE "Flowers"(
--     id SERIAL NOT NULL,
--     "Name" varchar(100) NOT NULL,
--     "Quantity" integer NOT NULL DEFAULT 0,
--     "Price" double precision NOT NULL,
--     "ArrivalDate" varchar(100) NOT NULL,
--     PRIMARY KEY(id)
-- );
-- CREATE TABLE "Cars"(
--     id SERIAL NOT NULL,
--     "Brand" varchar(100) NOT NULL,
--     "Model" varchar(100) NOT NULL,
--     "Mileage" INT NOT NULL,
--     "Owners" INT NOT NULL,
--     PRIMARY KEY(id)
-- );
-- CREATE TABLE "Furnitures"(
--     id SERIAL NOT NULL,
--     "Name" varchar(100) NOT NULL,
--     "Manufacturer" varchar(100) NOT NULL,
--     "Height" INT NOT NULL,
--     "Width" INT NOT NULL,
--     "Length" INT NOT NULL,
--     PRIMARY KEY(id)
-- );
-- Создание таблицы "Flowers"  
-- Создание таблицы "Flowers"  
CREATE TABLE "Flowers" (  
    id SERIAL PRIMARY KEY,  
    "Name" VARCHAR(100) NOT NULL UNIQUE,  
    "Quantity" INTEGER NOT NULL DEFAULT 0,  
    "Price" DOUBLE PRECISION NOT NULL,  
    "ArrivalDate" VARCHAR(100) NOT NULL  
);  

-- Создание таблицы "Cars"  
CREATE TABLE "Cars" (  
    id SERIAL PRIMARY KEY,  
    "Brand" VARCHAR(100) NOT NULL,  
    "Model" VARCHAR(100) NOT NULL,  
    "Mileage" INT NOT NULL,  
    "Owners" INT NOT NULL  
);  

-- Создание таблицы "Furnitures"  
CREATE TABLE "Furnitures" (  
    id SERIAL PRIMARY KEY,  
    "Name" VARCHAR(100) NOT NULL,  
    "Manufacturer" VARCHAR(100) NOT NULL,  
    "Height" INT NOT NULL,  
    "Width" INT NOT NULL,  
    "Length" INT NOT NULL  
);