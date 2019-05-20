/*
 Navicat PostgreSQL Data Transfer

 Source Server         : localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 110001
 Source Host           : localhost:5432
 Source Catalog        : postgres
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110001
 File Encoding         : 65001

 Date: 18/05/2019 20:43:37
*/


-- ----------------------------
-- Table structure for doctor_remote_node
-- ----------------------------
DROP TABLE IF EXISTS "public"."doctor_remote_node";
CREATE TABLE "public"."doctor_remote_node" (
  "ip" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "key" varchar(2048) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "comment" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."doctor_remote_node" OWNER TO "postgres";

-- ----------------------------
-- Primary Key structure for table doctor_remote_node
-- ----------------------------
ALTER TABLE "public"."doctor_remote_node" ADD CONSTRAINT "doctor_remote_node_pkey" PRIMARY KEY ("ip");
