
CREATE TABLE UserGlobalDoNotDisturb (
  user_id varchar(255) NOT NULL,
  do_not_disturb_enable tinyint(1) DEFAULT 0,
  do_not_disturb_from int,   -- GMT 로 0000 ~ 2400
  do_not_disturb_to int,     -- 예를들면 PM:10:00 ~ AM:07:00 까지면,
                             -- 24시간 변환 22:00 ~ 07:00
                             -- GMT 변환: 13:00 ~ 22:00
                             -- DB record:
                               -- do_not_disturb_from: 1300
                               -- do_not_disturb_to: 2200
);

CREATE TABLE UserEndpointDoNotDisturb (
  user_id varchar(255) NOT NULL,
  endpoint_id varchar(255) NOT NULL,
  do_not_disturb_enable tinyint(1) DEFAULT 0,
  do_not_disturb_from int,
  do_not_disturb_to int,
  follow_global_setting tinyint(1) DEFAULT 0,
);

CREATE TABLE UserNotiMethod (
  user_id varchar(255) NOT NULL,
  noti_method_id varchar(255) NOT NULL,
  type varchar(255) NULL NULL,       -- email, slack, sms
  identifier varchar(255) NULL NULL, -- email address, slack channel id, phone number
  verified tinyint(1) DEFAULT 0,
  PRIMARY KEY (user_id, noti_method_id)
);

CREATE TABLE NotiMethodGlobalSetting (
  noti_method_id varchar(255) NOT NULL,
  enabled tinyint(1) DEFAULT 1,
  throttle_second int DEFAULT 0,
  PRIMARY KEY (noti_method_id)
);


CREATE TABLE NotiMethodEndpointSetting (
  noti_method_id varchar(255) NOT NULL,
  endpoint_id varchar(255) NOT NULL,
  enabled tinyint(1) DEFAULT 1,
  throttle_second int DEFAULT 0,
  follow_global_setting tinyint(1) DEFAULT 0,
  PRIMARY KEY (noti_method_id, endpoint_id)
);

INSERT INTO Users VALUES ('a', 'yoo', 20);
INSERT INTO Users VALUES ('b', 'kim', 21);
INSERT INTO Users VALUES ('c', 'lee', 22);
INSERT INTO Users VALUES ('d', 'park', 23);


