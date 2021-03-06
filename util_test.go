package main

import "testing"

func Test_toOneWhereClause(t *testing.T) {
	type args struct {
		a FilterData
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"equal numeric",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: Numeric,
					Operator:  Equal,
					Arg0:      "12",
					Arg1:      "",
				}},
			"t.\"f1\" = 12",
		},
		{"less numeric",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: Numeric,
					Operator:  Less,
					Arg0:      "12",
					Arg1:      "",
				}},
			"t.\"f1\" < 12",
		},
		{"greater numeric",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: Numeric,
					Operator:  Greater,
					Arg0:      "12",
					Arg1:      "",
				}},
			"t.\"f1\" > 12",
		},
		{"between numeric",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: Numeric,
					Operator:  Between,
					Arg0:      "12",
					Arg1:      "13",
				}},
			"t.\"f1\" BETWEEN 12 AND 13",
		},
		{"between string",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: String,
					Operator:  Between,
					Arg0:      "12",
					Arg1:      "13",
				}},
			"",
		},
		{"like string",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: String,
					Operator:  Like,
					Arg0:      "12",
				}},
			"t.\"f1\" LIKE '%12%'",
		},
		{"no equal bool",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: Bool,
					Operator:  NotEqual,
					Arg0:      "true",
				}},
			"t.\"f1\" <> true",
		},
		{"like bool",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: Bool,
					Operator:  Like,
					Arg0:      "12",
				}},
			"",
		},
		{"not null",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: String,
					Operator:  NotNull,
				}},
			"t.\"f1\" IS NOT NULL",
		},
		{"is null",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: String,
					Operator:  Null,
				}},
			"t.\"f1\" IS NULL",
		},
		{"in",
			args{
				a: FilterData{
					FieldName: "f1",
					FieldType: String,
					Operator:  In,
					Args:      "(\"q\", \"w\", \"3\")",
				}},
			"t.\"f1\" IN (\"q\", \"w\", \"3\")",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertFilterDataToSql(tt.args.a); got != tt.want {
				t.Errorf("convertFilterDataToSql() = %v, want %v", got, tt.want)
			}
		})
	}
}
