CREATE DATABASE "temgo" WITH ENCODING = "UTF-8";

\ c "temgo";

CREATE EXTENSION "uuid-ossp";

CREATE EXTENSION "pg_trgm";


CREATE TABLE "temtem_user_team"(
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" VARCHAR(256) NOT NULL DEFAULT '',
    -- Team Name
    "temtems" JSONB NOT NULL DEFAULT '[]',
    -- Team Members
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);