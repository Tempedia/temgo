\c "temgo";


CREATE INDEX ON "temtem_type" USING GIN("name" gin_trgm_ops);
CREATE INDEX ON "temtem_type" USING GIN("effective_against");
CREATE INDEX ON "temtem_type" USING GIN("ineffective_against");
CREATE INDEX ON "temtem_type" USING GIN("resistant_to");
CREATE INDEX ON "temtem_type" USING GIN("weak_to");
CREATE INDEX ON "temtem_type"("sort" ASC);

CREATE INDEX ON "temtem"("no");
CREATE INDEX ON "temtem" USING GIN("name" gin_trgm_ops);
CREATE INDEX ON "temtem" USING GIN("type" gin_trgm_ops);
