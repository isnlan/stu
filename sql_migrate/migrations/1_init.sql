
-- +migrate Up
DROP TABLE IF EXISTS `block`;
CREATE TABLE `block`  (
  `number` bigint(20) NOT NULL COMMENT '当前块高度',
  `previous_hash` varbinary(128) NULL DEFAULT NULL COMMENT '前块Hash',
  `hash` varbinary(128) NOT NULL COMMENT '当前块Hash',
  `data_hash` varbinary(128) NOT NULL COMMENT '数据Hash',
  `transaction_count` int(11) NOT NULL COMMENT '当前块中交易数量',
  `channel_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '链ID',
  `timestamp` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '时间戳',
  PRIMARY KEY (`number`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- +migrate Down
drop table block;
