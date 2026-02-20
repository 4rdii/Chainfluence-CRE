CREATE TABLE "campaigns" (
	"id" serial PRIMARY KEY NOT NULL,
	"onchain_campaign_id" integer,
	"advertiser_id" integer NOT NULL,
	"title" varchar(200) NOT NULL,
	"content_text" text NOT NULL,
	"content_hash" varchar(66),
	"token_address" varchar(42) NOT NULL,
	"amount" bigint NOT NULL,
	"min_views" integer DEFAULT 0,
	"campaign_duration" integer DEFAULT 0,
	"expiry_deadline" timestamp NOT NULL,
	"chain_id" integer DEFAULT 11155111,
	"status" varchar(20) DEFAULT 'draft' NOT NULL,
	"tx_hash" varchar(66),
	"categories" jsonb DEFAULT '[]'::jsonb,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "chain_events" (
	"id" serial PRIMARY KEY NOT NULL,
	"event_name" varchar(100) NOT NULL,
	"tx_hash" varchar(66) NOT NULL,
	"block_number" bigint,
	"raw_data" jsonb,
	"processed" boolean DEFAULT false,
	"created_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "channels" (
	"id" serial PRIMARY KEY NOT NULL,
	"user_id" integer NOT NULL,
	"platform" varchar(20) DEFAULT 'twitter' NOT NULL,
	"platform_handle" varchar(100) NOT NULL,
	"follower_count" integer DEFAULT 0,
	"categories" jsonb DEFAULT '[]'::jsonb,
	"price_per_post" integer DEFAULT 0,
	"is_active" boolean DEFAULT true,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "deals" (
	"id" serial PRIMARY KEY NOT NULL,
	"campaign_id" integer NOT NULL,
	"channel_id" integer,
	"influencer_id" integer NOT NULL,
	"status" varchar(20) DEFAULT 'proposed' NOT NULL,
	"proposed_by" varchar(42) NOT NULL,
	"post_url" text,
	"onchain_campaign_id" integer,
	"accepted_at" timestamp,
	"funded_at" timestamp,
	"posted_at" timestamp,
	"completed_at" timestamp,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "users" (
	"id" serial PRIMARY KEY NOT NULL,
	"wallet_address" varchar(42) NOT NULL,
	"role" varchar(20) DEFAULT 'both' NOT NULL,
	"display_name" varchar(100),
	"nonce" varchar(64),
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "verifications" (
	"id" serial PRIMARY KEY NOT NULL,
	"deal_id" integer NOT NULL,
	"status" varchar(20) DEFAULT 'pending' NOT NULL,
	"views_checked" integer,
	"content_matched" boolean,
	"duration_checked" boolean,
	"is_edited" boolean,
	"action" varchar(20),
	"message" text,
	"created_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
ALTER TABLE "campaigns" ADD CONSTRAINT "campaigns_advertiser_id_users_id_fk" FOREIGN KEY ("advertiser_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "channels" ADD CONSTRAINT "channels_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "deals" ADD CONSTRAINT "deals_campaign_id_campaigns_id_fk" FOREIGN KEY ("campaign_id") REFERENCES "public"."campaigns"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "deals" ADD CONSTRAINT "deals_channel_id_channels_id_fk" FOREIGN KEY ("channel_id") REFERENCES "public"."channels"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "deals" ADD CONSTRAINT "deals_influencer_id_users_id_fk" FOREIGN KEY ("influencer_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "verifications" ADD CONSTRAINT "verifications_deal_id_deals_id_fk" FOREIGN KEY ("deal_id") REFERENCES "public"."deals"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
CREATE UNIQUE INDEX "wallet_address_idx" ON "users" USING btree ("wallet_address");