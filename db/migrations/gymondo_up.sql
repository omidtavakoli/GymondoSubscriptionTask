CREATE TABLE "users" (
                         "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
                         "user_name" varchar UNIQUE,
                         "email" varchar UNIQUE NOT NULL,
                         "full_name" varchar,
                         "created_at" timestamp DEFAULT (now()),
                         "updated_at" timestamp DEFAULT (now()),
                         "deleted_at" timestamp,
                         "active_voucher" int
);

CREATE TABLE "products" (
                            "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
                            "name" varchar,
                            "created_at" timestamp DEFAULT (now()),
                            "updated_at" timestamp DEFAULT (now()),
                            "deleted_at" timestamp
);

CREATE TABLE "plans" (
                         "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
                         "name" varchar,
                         "price" int,
                         "discount" int,
                         "created_at" timestamp DEFAULT (now()),
                         "updated_at" timestamp DEFAULT (now()),
                         "deleted_at" timestamp,
                         "product_id" int,
                         "duration" int
);

CREATE TABLE "user_plans" (
                              "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
                              "user_id" int,
                              "plan_id" int,
                              "plan_status" varchar,
                              "start_date" timestamp DEFAULT (now()),
                              "end_date" timestamp DEFAULT (now()),
                              "voucher" int,
                              "tax" int,
                              "created_at" timestamp DEFAULT (now()),
                              "updated_at" timestamp DEFAULT (now()),
                              "deleted_at" timestamp
);

CREATE TABLE "vouchers" (
                            "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
                            "name" varchar,
                            "discount" int,
                            "discount_type" varchar,
                            "valid" boolean,
                            "created_at" timestamp DEFAULT (now()),
                            "updated_at" timestamp DEFAULT (now()),
                            "start" timestamp,
                            "end" timestamp,
                            "deleted_at" timestamp
);

CREATE TABLE "voucher_plan" (
                                "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
                                "voucher_id" int,
                                "plan_id" int,
                                "created_at" timestamp DEFAULT (now()),
                                "updated_at" timestamp DEFAULT (now()),
                                "deleted_at" timestamp
);

ALTER TABLE "users" ADD FOREIGN KEY ("active_voucher") REFERENCES "vouchers" ("id");

ALTER TABLE "plans" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "user_plans" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_plans" ADD FOREIGN KEY ("plan_id") REFERENCES "plans" ("id");

ALTER TABLE "user_plans" ADD FOREIGN KEY ("voucher") REFERENCES "vouchers" ("id");

ALTER TABLE "voucher_plan" ADD FOREIGN KEY ("voucher_id") REFERENCES "vouchers" ("id");

ALTER TABLE "voucher_plan" ADD FOREIGN KEY ("plan_id") REFERENCES "plans" ("id");