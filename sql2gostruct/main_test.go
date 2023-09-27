package main

import (
	"fmt"
	"testing"
)

func TestConvertDDL2Struct(t *testing.T) {
	tests := []struct {
		name string
		ddl  string
	}{
		{
			"normal",
			`create table if not exists ` + "`sp-databus`.t_product_line" + `
(
    c_id           char(32)    not null comment '主键' primary key,
    c_code varchar(32) not null comment '产品线编码',
    c_name       varchar(200) not null comment '名称',
    c_remark       varchar(200) default null comment '备注',
    c_add_dt       datetime    not null,
    c_add_by_id    char(32)    not null,
    c_update_dt    datetime    null,
    c_update_by_id char(32)    null,
    c_is_deleted  tinyint(1) default 0 comment '是否删除：1-删除，0-正常，默认0'
)  comment '产品线信息表';`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("#########")
			got := ConvertDDL2Struct(tt.ddl)
			fmt.Println(got)
			fmt.Println("#########")
		})
	}
}
