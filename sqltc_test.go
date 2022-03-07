package sqltc

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Convert(t *testing.T) {
	t.Parallel()

	type args struct {
		i    string
		want Columns
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "Include Comment",
			args: args{
				i: "CREATE TABLE IF NOT EXISTS test.testdata(  name VARCHAR(255) NOT NULL,  info VARCHAR(200) NOT NULL,  PRIMARY KEY (name))",
				want: Columns{
					{Name: "name", Type: "VARCHAR", IsNULL: true},
					{Name: "info", Type: "VARCHAR", IsNULL: true},
				},
			},
		},
	}

	for n, tt := range cases {
		tt := tt
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			got := Convert(tt.args.i)
			if diff := cmp.Diff(tt.args.want, got); diff != "" {
				t.Errorf("Convert does not notice content: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_excludeComment(t *testing.T) {
	t.Parallel()

	type args struct {
		i    string
		want string
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "Include Comment",
			args: args{
				i:    "SELECT FROM * users -- this is a comment",
				want: "SELECT FROM * users ",
			},
		},
		{
			name: "Not Comment",
			args: args{
				i:    "SELECT FROM * users",
				want: "SELECT FROM * users",
			},
		},
	}

	for n, tt := range cases {
		tt := tt
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			got := excludeComment(tt.args.i)
			if diff := cmp.Diff(tt.args.want, got); diff != "" {
				t.Errorf("excludeComment does not notice content: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_lineFromFile(t *testing.T) {
	t.Parallel()

	type args struct {
		i    string
		want []string
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "one line file",
			args: args{
				i:    "./testdata/oneline.sql",
				want: []string{"SELECT FROM * users", ""},
			},
		},
		{
			name: "multi line file",
			args: args{
				i:    "./testdata/multiline.sql",
				want: []string{"SELECT FROM * users", "DELETE FROM * users", ""},
			},
		},
	}

	for n, tt := range cases {
		tt := tt
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			got, _ := lineFromFile(tt.args.i)
			if diff := cmp.Diff(tt.args.want, got); diff != "" {
				t.Errorf("lineFromFile does not notice content: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_load(t *testing.T) {
	t.Parallel()

	type args struct {
		i    string
		want []string
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "testdata create table query",
			args: args{
				i: "./testdata/create_table.sql",
				want: []string{"CREATE DATABASE IF NOT EXISTS test",
					"CREATE TABLE IF NOT EXISTS test.testdata(  name VARCHAR(255) NOT NULL,  info VARCHAR(200) NOT NULL,  PRIMARY KEY (name))",
				},
			},
		},
	}

	for n, tt := range cases {
		tt := tt
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			got, _ := load(tt.args.i)
			if diff := cmp.Diff(tt.args.want, got); diff != "" {
				t.Errorf("load does not notice content: (+got -want)\n%s", diff)
			}
		})
	}
}

func Test_File(t *testing.T) {
	t.Parallel()

	type args struct {
		i    string
		want []string
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "testdata",
			args: args{
				i: "./testdata/create_table.sql",
				want: []string{"CREATE DATABASE IF NOT EXISTS test",
					"CREATE TABLE IF NOT EXISTS test.testdata(  name VARCHAR(255) NOT NULL,  info VARCHAR(200) NOT NULL,  PRIMARY KEY (name))",
				},
			},
		},
	}

	for n, tt := range cases {
		tt := tt
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			s := SqlFile{}
			_ = s.File(tt.args.i)
			got := s.Queries
			if diff := cmp.Diff(tt.args.want, got); diff != "" {
				t.Errorf("File does not notice content: (+got -want)\n%s", diff)
			}
		})
	}
}

func Test_Directory(t *testing.T) {
	t.Parallel()

	type args struct {
		i    string
		want []string
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "testdata directory",
			args: args{
				i: "./testdata",
				want: []string{"CREATE DATABASE IF NOT EXISTS test",
					"CREATE TABLE IF NOT EXISTS test.testdata(  name VARCHAR(255) NOT NULL,  info VARCHAR(200) NOT NULL,  PRIMARY KEY (name))",
				},
			},
		},
	}

	for n, tt := range cases {
		tt := tt
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			s := SqlFile{}
			_ = s.Directory(tt.args.i)
			got := s.Queries
			if diff := cmp.Diff(tt.args.want, got); diff != "" {
				t.Errorf("Directory does not notice content: (+got -want)\n%s", diff)
			}
		})
	}
}
