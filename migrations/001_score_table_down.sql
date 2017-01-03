/**
 * Drops "score" table and indexes
 */
DROP INDEX "ix_score_user_id";
DROP INDEX "ix_score_user_id_score";
DROP INDEX "ix_score_create_date";
DROP INDEX "ix_score_update_date";
DROP TABLE "public"."score";
