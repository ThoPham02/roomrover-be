CREATE TABLE IF NOT EXISTS "homes_tbl" (
    "id" BIGINT,
    "owner_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "province_code" INT NOT NULL,
    "district_code" INT NOT NULL,
    "ward_code" INT NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "wifi_service" BIGINT NOT NULL DEFAULT -1,
    "electricity_service" BIGINT NOT NULL DEFAULT -1,
    "water_service" BIGINT NOT NULL DEFAULT -1,
    "cleaning_service" BIGINT NOT NULL DEFAULT -1,
    "parking_service" BIGINT NOT NULL DEFAULT -1,
    "security_service" BIGINT NOT NULL DEFAULT -1,
    "other_service" TEXT,
    "status" INT NOT NULL DEFAULT 0,
    "created_at" BIGINT,
    "created_by" BIGINT,
    "updated_at" BIGINT,
    "updated_by" BIGINT,
    PRIMARY KEY ("id")
);

COMMENT ON COLUMN "homes_tbl"."wifi_service" IS 'VND/MONTH -1: No, 0: Free';
COMMENT ON COLUMN "homes_tbl"."electricity_service" IS 'VND/MONTH -1: No, 0: Free';
COMMENT ON COLUMN "homes_tbl"."water_service" IS 'VND/MONTH -1: No, 0: Free';
COMMENT ON COLUMN "homes_tbl"."cleaning_service" IS 'VND/MONTH -1: No, 0: Free';
COMMENT ON COLUMN "homes_tbl"."parking_service" IS 'VND/MONTH -1: No, 0: Free';
COMMENT ON COLUMN "homes_tbl"."security_service" IS 'VND/MONTH -1: No, 0: Free';
COMMENT ON COLUMN "homes_tbl"."status" IS '1: Draft, 2: Active, 4: Inactive';

CREATE TABLE IF NOT EXISTS "room_class_tbl" (
    "id" BIGINT,
    "owner_id" BIGINT NOT NULL,
    "home_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "unit_price" DOUBLE PRECISION NOT NULL,
    "acreage" DOUBLE PRECISION NOT NULL,
    "room_info" TEXT,
    "type" INT NOT NULL DEFAULT 0,
    "status" INT NOT NULL DEFAULT 0,
    "created_at" BIGINT,
    "created_by" BIGINT,
    "updated_at" BIGINT,
    "updated_by" BIGINT,
    PRIMARY KEY ("id")
);

COMMENT ON COLUMN "room_class_tbl"."type" IS '1: Apartment, 2: House, 4: Villa';
COMMENT ON COLUMN "room_class_tbl"."status" IS '1: Draft, 2: Active, 4: Inactive';

CREATE TABLE IF NOT EXISTS "rooms_tbl" (
    "id" BIGINT,
    "owner_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "room_class_id" BIGINT NOT NULL,
    "status" INT NOT NULL DEFAULT 0,
    "created_at" BIGINT,
    "created_by" BIGINT,
    "updated_at" BIGINT,
    "updated_by" BIGINT,
    PRIMARY KEY ("id")
);

COMMENT ON COLUMN "rooms_tbl"."status" IS '1: Unavailable, 2: Available, 4: Rented, 8: Inactive';

CREATE TABLE IF NOT EXISTS "room_albums_tbl" (
    "id" BIGINT,
    "owner_id" BIGINT NOT NULL,
    "room_class_id" BIGINT NOT NULL,
    "url" VARCHAR(255) NOT NULL,
    "status" INT NOT NULL DEFAULT 0,
    "created_at" BIGINT,
    "created_by" BIGINT,
    "updated_at" BIGINT,
    "updated_by" BIGINT,
    PRIMARY KEY ("id")
);
