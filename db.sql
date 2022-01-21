CREATE TABLE "users" (
    "id" uuid gen_random_uuid () DEFAULT PRIMARY KEY,
    "full_name" varchar NOT NULL,
    "user_name" varchar NOT NULL UNIQUE,
    "email_id" varchar UNIQUE NOT NULL,
    "created_at" timestampz NOT NULL DEFAULT(now())
);

CREATE TABLE "group" (
    "id" uuid gen_random_uuid () DEFAULT PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL,
    "created_at" timestampz NOT NULL DEFAULT(now()),
    "created_by" bigserial
);

CREATE TABLE "users_groups" (
    "user_id" uuid,
    "group_id" uuid
);
