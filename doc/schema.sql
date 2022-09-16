CREATE TABLE "temgo" WITH ENCODING="UTF-8";

\c "temgo";
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION "pg_trgm";

/* 内部人员 */
CREATE TABLE "staff"(
    "id" BIGSERIAL NOT NULL PRIMARY KEY,
    "username" VARCHAR(32) NOT NULL,
    "name" VARCHAR(32) NOT NULL,
    "salt" VARCHAR(32) NOT NULL,
    "ptype" VARCHAR(16) NOT NULL,
    "password" VARCHAR(256) NOT NULL,
    "status" VARCHAR(32) NOT NULL DEFAULT 'OK',
    "is_superuser" BOOLEAN NOT NULL DEFAULT FALSE,
    "phone" VARCHAR(32) NOT NULL DEFAULT '',
    "email" VARCHAR(128) NOT NULL DEFAULT '',
    "created_by" BIGINT,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE("username")
);
CREATE INDEX ON "staff" USING GIN("username" gin_trgm_ops);
CREATE INDEX ON "staff" USING GIN("name" gin_trgm_ops);
CREATE INDEX ON "staff"("status");

CREATE TABLE "staff_token"(
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "staff_id" BIGINT NOT NULL,
    "device" VARCHAR(1024) NOT NULL,
    "ip" VARCHAR(256) NOT NULL,
    "expires_at" TIMESTAMP WITH TIME ZONE,
    "status" VARCHAR(32) NOT NULL DEFAULT 'OK',
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX ON "staff_token"("staff_id","status");
CREATE INDEX ON "staff_token"("staff_id","expires_at");
CREATE INDEX ON "staff_token"("created_at");
