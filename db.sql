CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "full_name" varchar NOT NULL,
    "email_id" varchar UNIQUE NOT NULL,
    "created_at" timestampz NOT NULL DEFAULT(now())
);

CREATE TABLE "group" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL,
    "created_at" timestampz NOT NULL DEFAULT(now()),
    "created_by" bigserial
);

CREATE TABLE "users_groups" (
    "user_id" bigserial,
    "group_id" bigserial
);
