CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    email       text NOT NULL UNIQUE ,
    nickname    text NOT NULL UNIQUE,
    password    text NOT NULL,
    name        text NOT NULL,
    phone_number text NOT NULL UNIQUE,

    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_ad timestamp(0) with time zone NOT NULL DEFAULT NOW()
);