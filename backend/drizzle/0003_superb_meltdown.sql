ALTER TABLE "channels" ADD COLUMN "world_id_nullifier" varchar(100);--> statement-breakpoint
ALTER TABLE "channels" ADD COLUMN "world_id_verified_at" timestamp;