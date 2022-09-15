\c "temgo";

CREATE EXTENSION "pg_trgm";

CREATE INDEX ON "temtem_type" USING GIN("name" gin_trgm_ops);
CREATE INDEX ON "temtem_type" USING GIN("effective_against");
CREATE INDEX ON "temtem_type" USING GIN("ineffective_against");
CREATE INDEX ON "temtem_type" USING GIN("resistant_to");
CREATE INDEX ON "temtem_type" USING GIN("weak_to");
