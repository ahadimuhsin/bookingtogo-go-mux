/*
 Navicat Premium Data Transfer

 Source Server         : local-postgre
 Source Server Type    : PostgreSQL
 Source Server Version : 110000
 Source Host           : localhost:5432
 Source Catalog        : backend-test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110000
 File Encoding         : 65001

 Date: 31/12/2022 16:05:49
*/


-- ----------------------------
-- Sequence structure for next_cust_val
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."next_cust_val";
CREATE SEQUENCE "public"."next_cust_val" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for next_family_list
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."next_family_list";
CREATE SEQUENCE "public"."next_family_list" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for next_nationality
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."next_nationality";
CREATE SEQUENCE "public"."next_nationality" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS "public"."customer";
CREATE TABLE "public"."customer" (
  "cst_id" int4 NOT NULL DEFAULT nextval('next_cust_val'::regclass),
  "nationality_id" int4 NOT NULL,
  "cst_name" char(50) COLLATE "pg_catalog"."default" NOT NULL,
  "cst_dob" date NOT NULL,
  "cst_phone" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "cst_email" varchar(50) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Records of customer
-- ----------------------------
INSERT INTO "public"."customer" VALUES (1, 1, 'Cek                                               ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (3, 1, 'Cek                                               ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (4, 1, 'Cek                                               ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (6, 1, 'Cek                                               ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (7, 1, 'Cek                                               ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (8, 1, 'Cek                                               ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (9, 1, 'Cek                                               ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (10, 1, 'Cek                                               ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (13, 1, 'Cek                                               ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (2, 1, 'CekAda                                            ', '1996-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (5, 1, 'Cek Muhsin                                        ', '1986-02-02', '088782323', 'cek@mail.com');
INSERT INTO "public"."customer" VALUES (14, 1, 'Haikal                                            ', '1996-02-02', '088782323', 'cek@mail.com');

-- ----------------------------
-- Table structure for family_list
-- ----------------------------
DROP TABLE IF EXISTS "public"."family_list";
CREATE TABLE "public"."family_list" (
  "fl_id" int4 NOT NULL DEFAULT nextval('next_family_list'::regclass),
  "cst_id" int4 NOT NULL,
  "fl_relation" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "fl_name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "fl_dob" date NOT NULL
)
;

-- ----------------------------
-- Records of family_list
-- ----------------------------

-- ----------------------------
-- Table structure for nationality
-- ----------------------------
DROP TABLE IF EXISTS "public"."nationality";
CREATE TABLE "public"."nationality" (
  "nationality_id" int4 NOT NULL DEFAULT nextval('next_nationality'::regclass),
  "nationality_name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "nationality_code" char(2) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Records of nationality
-- ----------------------------
INSERT INTO "public"."nationality" VALUES (1, 'Indonesia', 'ID');
INSERT INTO "public"."nationality" VALUES (2, 'Malaysia', 'MY');

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."next_cust_val"
OWNED BY "public"."customer"."cst_id";
SELECT setval('"public"."next_cust_val"', 29, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."next_family_list"
OWNED BY "public"."family_list"."fl_id";
SELECT setval('"public"."next_family_list"', 8, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."next_nationality"
OWNED BY "public"."nationality"."nationality_id";
SELECT setval('"public"."next_nationality"', 3, true);

-- ----------------------------
-- Primary Key structure for table customer
-- ----------------------------
ALTER TABLE "public"."customer" ADD CONSTRAINT "customer_pkey" PRIMARY KEY ("cst_id");

-- ----------------------------
-- Primary Key structure for table family_list
-- ----------------------------
ALTER TABLE "public"."family_list" ADD CONSTRAINT "family_list_pkey" PRIMARY KEY ("fl_id");

-- ----------------------------
-- Primary Key structure for table nationality
-- ----------------------------
ALTER TABLE "public"."nationality" ADD CONSTRAINT "nationality_pkey" PRIMARY KEY ("nationality_id");

-- ----------------------------
-- Foreign Keys structure for table customer
-- ----------------------------
ALTER TABLE "public"."customer" ADD CONSTRAINT "nationality" FOREIGN KEY ("nationality_id") REFERENCES "public"."nationality" ("nationality_id") ON DELETE RESTRICT ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table family_list
-- ----------------------------
ALTER TABLE "public"."family_list" ADD CONSTRAINT "customer" FOREIGN KEY ("cst_id") REFERENCES "public"."customer" ("cst_id") ON DELETE CASCADE ON UPDATE NO ACTION;
