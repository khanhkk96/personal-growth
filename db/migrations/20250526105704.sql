-- Create "payments" table
CREATE TABLE "public"."payments" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "pay_by" text NULL DEFAULT 'momo',
  "amount" bigint NOT NULL,
  "bank_code" text NULL,
  "transaction_no" text NOT NULL,
  "bank_transaction_no" text NULL,
  "pay_date" timestamptz NULL,
  "transaction_status" text NULL DEFAULT 'success',
  "order_info" text NULL,
  "txn_ref" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_payments_deleted_at" to table: "payments"
CREATE INDEX "idx_payments_deleted_at" ON "public"."payments" ("deleted_at");
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "username" text NOT NULL,
  "password" text NOT NULL,
  "full_name" text NOT NULL,
  "email" text NOT NULL,
  "phone" text NULL,
  "otp" text NULL,
  "otp_expired_at" timestamptz NULL,
  "otp_counter" bigint NULL DEFAULT 0,
  "is_active" boolean NULL DEFAULT false,
  "role" text NULL DEFAULT 'user',
  "avatar" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_email" UNIQUE ("email"),
  CONSTRAINT "uni_users_phone" UNIQUE ("phone"),
  CONSTRAINT "uni_users_username" UNIQUE ("username")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create "articles" table
CREATE TABLE "public"."articles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "created_by_id" uuid NULL,
  "title" text NOT NULL,
  "summary" text NOT NULL,
  "content" text NOT NULL,
  "quote" text NULL,
  "read_turns" bigint NULL DEFAULT 0,
  "is_published" boolean NULL DEFAULT true,
  "is_featured" boolean NULL DEFAULT false,
  "published_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_articles_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_articles_deleted_at" to table: "articles"
CREATE INDEX "idx_articles_deleted_at" ON "public"."articles" ("deleted_at");
-- Create "base_models" table
CREATE TABLE "public"."base_models" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "created_by_id" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_base_models_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_base_models_deleted_at" to table: "base_models"
CREATE INDEX "idx_base_models_deleted_at" ON "public"."base_models" ("deleted_at");
-- Create "projects" table
CREATE TABLE "public"."projects" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "created_by_id" uuid NULL,
  "name" text NOT NULL,
  "type" text NOT NULL,
  "summary" text NOT NULL,
  "stack" text NOT NULL,
  "description" text NULL,
  "start_at" timestamptz NULL,
  "end_at" timestamptz NULL,
  "status" text NOT NULL DEFAULT 'ongoing',
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_projects_name" UNIQUE ("name"),
  CONSTRAINT "fk_projects_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_projects_deleted_at" to table: "projects"
CREATE INDEX "idx_projects_deleted_at" ON "public"."projects" ("deleted_at");
-- Create "issues" table
CREATE TABLE "public"."issues" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "created_by_id" uuid NULL,
  "name" text NOT NULL,
  "description" text NULL,
  "files" text NULL,
  "project_id" uuid NULL,
  "status" text NOT NULL,
  "priority" text NOT NULL,
  "issued_at" timestamptz NULL,
  "need_to_solve" numeric NULL,
  "references" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_issues_name" UNIQUE ("name"),
  CONSTRAINT "fk_issues_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_issues_project" FOREIGN KEY ("project_id") REFERENCES "public"."projects" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_issues_deleted_at" to table: "issues"
CREATE INDEX "idx_issues_deleted_at" ON "public"."issues" ("deleted_at");
-- Create "plans" table
CREATE TABLE "public"."plans" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "created_by_id" uuid NULL,
  "name" text NOT NULL,
  "description" text NULL,
  "objective" text NOT NULL,
  "expected_start_time" timestamptz NOT NULL,
  "expected_end_time" timestamptz NULL,
  "actual_start_time" timestamptz NULL,
  "actual_end_time" timestamptz NULL,
  "progress" bigint NULL DEFAULT 0,
  "status" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_plans_name" UNIQUE ("name"),
  CONSTRAINT "fk_plans_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_plans_deleted_at" to table: "plans"
CREATE INDEX "idx_plans_deleted_at" ON "public"."plans" ("deleted_at");
-- Create "schedules" table
CREATE TABLE "public"."schedules" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "created_by_id" uuid NULL,
  "name" text NOT NULL,
  "description" text NULL,
  "note" text NULL,
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz NULL,
  "is_required" boolean NULL DEFAULT true,
  "plan_id" uuid NULL,
  "status" text NULL DEFAULT 'pending',
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_schedules_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_schedules_plan" FOREIGN KEY ("plan_id") REFERENCES "public"."plans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_schedules_deleted_at" to table: "schedules"
CREATE INDEX "idx_schedules_deleted_at" ON "public"."schedules" ("deleted_at");
