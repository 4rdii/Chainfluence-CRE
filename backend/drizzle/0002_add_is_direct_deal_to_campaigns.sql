ALTER TABLE "campaigns" ALTER COLUMN "amount" SET DATA TYPE varchar(78);--> statement-breakpoint
ALTER TABLE "campaigns" ADD COLUMN "is_direct_deal" boolean DEFAULT false;