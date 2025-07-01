CREATE TABLE user_info
(
    id          SERIAL NOT NULL PRIMARY KEY,
    phone       VARCHAR(24) ,
    name        VARCHAR(24),
    department  VARCHAR(32),
    position    VARCHAR(32),
    gender      varchar(24),
    is_admin    BOOLEAN,
    is_disable    BOOLEAN,
    create_time DATE DEFAULT NOW()
);

COMMENT ON COLUMN user_info.id IS '主键';
COMMENT ON COLUMN user_info.phone IS '手机号';
COMMENT ON COLUMN user_info.name IS '名字';
COMMENT ON COLUMN user_info.department IS '部门';
COMMENT ON COLUMN user_info.position IS '职位';
COMMENT ON COLUMN user_info.gender IS '性别';
COMMENT ON COLUMN user_info.is_admin IS '是否管理员';
COMMENT ON COLUMN user_info.is_admin IS '是否禁用';
COMMENT ON COLUMN user_info.create_time IS '创建时间';


CREATE TABLE user_foreigner_id
(
    user_id      INT NOT NULL,
    foreign_code VARCHAR(64),
    foreigner_id VARCHAR(64),
    create_time  DATE DEFAULT NOW()
);

CREATE UNIQUE INDEX u_idx_user_id_foreign_code
    on user_foreigner_id (USER_ID, FOREIGN_CODE);


CREATE UNIQUE INDEX u_idx_foreign_code_foreigner_id
    on user_foreigner_id (FOREIGN_CODE, FOREIGNER_ID);

COMMENT ON COLUMN user_foreigner_id.user_id IS '用户id';
COMMENT ON COLUMN user_foreigner_id.foreign_code IS '外部码- 可以理解为微信小程序 appid';
COMMENT ON COLUMN user_foreigner_id.foreigner_id IS '外部用户id- 可以理解为微信小程序 openid';
COMMENT ON COLUMN user_foreigner_id.create_time IS '创建时间';




CREATE TABLE device
(
    id                     VARCHAR(24),
    type                   VARCHAR(24),
    rules                  VARCHAR(24),
    state                  TEXT,
    owner_user_id          INT,
    owner_use_name         TEXT,
    annual_inspection_time DATE,
    validity_period_time   DATE,
    create_time            DATE DEFAULT NOW(),
    CHECK (state IN ('RENTABLE', 'RENTED', 'MAINTAIN'))
);

CREATE TABLE rental
(
    id          SERIAL PRIMARY KEY ,
    user_id     INT,
    user_name   VARCHAR(24),
    device_id   VARCHAR(24),
    device_type VARCHAR(24),
    rent_time   DATE,
    return_time DATE,
    create_time DATE DEFAULT NOW()
);