\c "temgo";


CREATE INDEX ON "temtem_type" USING GIN("name" gin_trgm_ops);
CREATE INDEX ON "temtem_type" USING GIN("effective_against");
CREATE INDEX ON "temtem_type" USING GIN("ineffective_against");
CREATE INDEX ON "temtem_type" USING GIN("resistant_to");
CREATE INDEX ON "temtem_type" USING GIN("weak_to");
CREATE INDEX ON "temtem_type"("sort" ASC);

CREATE INDEX ON "temtem"("no");
CREATE INDEX ON "temtem" USING GIN("name" gin_trgm_ops);
CREATE INDEX ON "temtem" USING GIN("type");
CREATE INDEX ON "temtem" USING GIN("traits");
CREATE INDEX ON "temtem" USING GIN("evolves_to");
CREATE INDEX ON "temtem"(("stats"->'HP'->'base'));
CREATE INDEX ON "temtem"(("stats"->'STA'->'base'));
CREATE INDEX ON "temtem"(("stats"->'SPD'->'base'));
CREATE INDEX ON "temtem"(("stats"->'ATK'->'base'));
CREATE INDEX ON "temtem"(("stats"->'DEF'->'base'));
CREATE INDEX ON "temtem"(("stats"->'SPATK'->'base'));
CREATE INDEX ON "temtem"(("stats"->'SPDEF'->'base'));


CREATE INDEX ON "temtem_trait" USING GIN("name" gin_trgm_ops);

CREATE INDEX ON "temtem_leveling_up_technique"("temtem");
CREATE INDEX ON "temtem_leveling_up_technique"("level");
CREATE INDEX ON "temtem_leveling_up_technique"("technique_name");

CREATE INDEX ON "temtem_course_technique"("temtem");
CREATE INDEX ON "temtem_course_technique" USING GIN("course" gin_trgm_ops);
CREATE INDEX ON "temtem_course_technique"("technique_name");

CREATE INDEX ON "temtem_breeding_technique"("temtem");
CREATE INDEX ON "temtem_breeding_technique" USING GIN("parents");
CREATE INDEX ON "temtem_breeding_technique"("technique_name");

CREATE INDEX ON "temtem_location" USING GIN("name" gin_trgm_ops);
CREATE INDEX ON "temtem_location_area"("location");
CREATE INDEX ON "temtem_location_area" USING GIN("temtems");

CREATE INDEX ON "temtem_status_condition" USING GIN("name" gin_trgm_ops);
CREATE INDEX ON "temtem_status_condition" USING GIN("techniques");
CREATE INDEX ON "temtem_status_condition" USING GIN("traits");
CREATE INDEX ON "temtem_status_condition"("group");

CREATE INDEX ON "temtem_course_item" USING GIN("no" gin_trgm_ops);
CREATE INDEX ON "temtem_course_item" USING GIN("technique" gin_trgm_ops);

CREATE INDEX ON "temtem_item_category"("parent");
CREATE INDEX ON "temtem_item_category"("sort");

CREATE INDEX ON "temtem_item" USING GIN("name" gin_trgm_ops);
CREATE INDEX ON "temtem_item"("category");
CREATE INDEX ON "temtem_item"("sort");