 CREATE TABLE "Flowers" (  
    id SERIAL PRIMARY KEY,  
    "Name" VARCHAR(100) NOT NULL UNIQUE,  
    "Quantity" INTEGER NOT NULL DEFAULT 0,  
    "Price" DOUBLE PRECISION NOT NULL,  
    "ArrivalDate" VARCHAR(100) NOT NULL  
);  

CREATE TABLE "Cars" (  
    id SERIAL PRIMARY KEY,  
    "Brand" VARCHAR(100) NOT NULL,  
    "Model" VARCHAR(100) NOT NULL,  
    "Mileage" INT NOT NULL,  
    "Owners" INT NOT NULL  
);  

CREATE TABLE "Furnitures" (  
    id SERIAL PRIMARY KEY,  
    "Name" VARCHAR(100) NOT NULL,  
    "Manufacturer" VARCHAR(100) NOT NULL,  
    "Height" INT NOT NULL,  
    "Width" INT NOT NULL,  
    "Length" INT NOT NULL  
);