/**
 * Create Table Structure and Indexes for the "score" table
 * This table stores all the score values for a user
 */
CREATE TABLE "public"."score" (
	"id" uuid NOT NULL,
	"create_date" timestamp NOT NULL,
	"update_date" timestamp NOT NULL,
	"user_id" uuid NOT NULL,
	"score" int2 NOT NULL,
	PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE
)
WITH (OIDS=FALSE);
CREATE INDEX "ix_score_user_id" ON "public"."score" USING btree(user_id);
CREATE INDEX "ix_score_user_id_score" ON "public"."score" USING btree(user_id, score);
CREATE INDEX "ix_score_create_date" ON "public"."score" USING btree(create_date);
CREATE INDEX "ix_score_update_date" ON "public"."score" USING btree(update_date);
