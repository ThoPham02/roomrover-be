CREATE TABLE "homes_tbl" (
    "id" BIGINT,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "province_code" INT NOT NULL,
    "district_code" INT NOT NULL,
    "ward_code" INT NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "wifi_service" BIGINT NOT NULL DEFAULT -1 COMMENT 'VND/MONTH -1: No, 0: Free',
    "electricity_service" BIGINT NOT NULL DEFAULT -1 COMMENT 'VND/MONTH -1: No, 0: Free',
    "water_service" BIGINT NOT NULL DEFAULT -1 COMMENT 'VND/MONTH -1: No, 0: Free',
    "cleaning_service" BIGINT NOT NULL DEFAULT -1 COMMENT 'VND/MONTH -1: No, 0: Free',
    "parking_service" BIGINT NOT NULL DEFAULT -1 COMMENT 'VND/MONTH -1: No, 0: Free',
    "security_service" BIGINT NOT NULL DEFAULT -1 COMMENT 'VND/MONTH -1: No, 0: Free',
    "other_service" TEXT,
    "status" INT NOT NULL DEFAULT 0 COMMENT '1: Draf, 2: Active, 4: Inactive',
    "created_at" BIGINT,
    "created_by" BIGINT,
    "updated_at" BIGINT,
    "updated_by" BIGINT,
    PRIMARY KEY ("id")
);

CREATE TABLE "room_types_tbl" (
    "id" BIGINT,
    "home_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "unit_price" DOUBLE NOT NULL,
    "acreage" DOUBLE NOT NULL,
    "room_info" TEXT,
    "type" INT NOT NULL DEFAULT 0 COMMENT '1: Apartment, 2: House, 4: Villa',
    "status" INT NOT NULL DEFAULT 0 COMMENT '1: Draf, 2: Active, 4: Inactive',
    "created_at" BIGINT,
    "created_by" BIGINT,
    "updated_at" BIGINT,
    "updated_by" BIGINT,
    PRIMARY KEY ("id")
);

CREATE TABLE "rooms_tbl" (
    "id" BIGINT,
    "name" VARCHAR(255) NOT NULL,
    "room_type_id" BIGINT NOT NULL,
    "status" INT NOT NULL DEFAULT 0 COMMENT '1: Unavable, 2: Avable, 4: Renter, 8: Inactive',
    "created_at" BIGINT,
    "created_by" BIGINT,
    "updated_at" BIGINT,
    "updated_by" BIGINT,
    PRIMARY KEY ("id")
);