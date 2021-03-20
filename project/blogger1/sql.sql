create database BloggerGin;
use BloggerGin;

create table  BlogTag(
                         Id int(10) unsigned not null auto_increment primary key comment '标签id',
                         Name varchar(100) default '' comment '标签名字',
                         CreatedDt int(10) unsigned not null  default 0 comment '创建时间',
                         CreateBy varchar(100) default '' comment '创建人',
                         UpdateDt int(10) unsigned not null  default 0  comment '修改时间',
                         UpdateBy varchar(100) default '' comment '更新人',
                         State tinyint(3) unsigned default 1 comment '状态 0/禁用  1/启用 2/删除'
)comment '文章标签表';

CREATE TABLE BlogArticle (
                             Id bigint unsigned primary key NOT NULL AUTO_INCREMENT COMMENT '文章ID',
                             TagId int(10) unsigned default 0 NOT NULL COMMENT '分类Id',
                             Title varchar(1024) default '' NOT NULL COMMENT '文章标题',
                             Summary varchar(256) default '' NOT NULL COMMENT '文章摘要',
                             Content longtext  NOT NULL COMMENT '文章内容',
                             CreatedDt int(10) unsigned not null  default 0  comment '创建时间',
                             CreateBy varchar(100) default '' comment '创建人',
                             UpdateDt int(10) unsigned not null  default 0  comment '修改时间',
                             UpdateBy varchar(100) default '' comment '更新人',
                             State tinyint(3) unsigned default 1 comment '状态 0/禁用  1/启用 2/删除'
) comment '文章管理';

create table BlogAuth(
                         Id int(10) unsigned primary key  auto_increment comment 'Id',
                         UserName varchar(50) default '' comment '账号',
                         Password varchar(50) default '' comment '密码',
                         State tinyint(3) unsigned default 1 comment '状态 0/禁用  1/启用 2/删除'
)comment '认证表';

-- 增加认证表信息
insert into BlogAuth (UserName,Password) values('admin','admin')
