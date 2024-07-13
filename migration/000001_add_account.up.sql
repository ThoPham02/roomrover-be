CREATE TABLE IF NOT EXISTS "users_tbl" (
    "user_id" BIGINT,
    "profile_id" BIGINT,
    "username" VARCHAR(255) UNIQUE NOT NULL,
    "password_hash" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "role" BIGINT,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS "profiles_tbl" (
    "profile_id" BIGINT,
    "fullname" VARCHAR(255),
    "dob" BIGINT,
    "avatar_url" VARCHAR(255),
    "address" VARCHAR(255),
    "phone" VARCHAR(255),
    "created_at" BIGINT,
    "created_by" BIGINT,
    "updated_at" BIGINT,
    "updated_by" BIGINT,
    PRIMARY KEY ("profile_id")
);

