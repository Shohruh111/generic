CREATE TABLE "roles"(
    "id" UUID PRIMARY KEY,
    "type" VARCHAR(20) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "users"(
    "id" UUID PRIMARY KEY,
    "role_id" UUID REFERENCES "roles" ("id"),
    "first_name" VARCHAR (50) NOT NULL,
    "last_name" VARCHAR(50) NOT NULL,
    "email" VARCHAR(100) NOT NULL,
    "phone_number" VARCHAR(20) NOT NULL,
    "passsword" VARCHAR(8) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "courses"(
    "id" UUID PRIMARY KEY,
    "user_id" UUID REFERENCES "users"("id"),
    "name" VARCHAR(30) NOT NULL,
    "description" VARCHAR(50) NOT NULL,
    "weekly_number" NUMERIC NOT NULL,
    "duration" VARCHAR(10) NOT NULL,
    "price" VARCHAR(20) NOT NULL,
    "date_of_beginning" VARCHAR(20) NOT NULL,

); 