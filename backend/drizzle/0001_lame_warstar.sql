ALTER TABLE "channels" ADD COLUMN "platform_user_id" varchar(100);--> statement-breakpoint
ALTER TABLE "channels" ADD COLUMN "profile_image_url" varchar(500);--> statement-breakpoint
ALTER TABLE "channels" ADD COLUMN "verified_at" timestamp;