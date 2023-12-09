package go_fake

import (
	"fmt"
	"testing"
)

func TestFake_LoadFakeFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "test_file/fund_batch_order_db.yaml",
			args: args{file: "test_file/fund_batch_order_db.yaml"},
			want: nil,
		},
		{
			name: "test_file/fund_collect_order_db.yaml",
			args: args{file: "test_file/fund_collect_order_db.yaml"},
			want: nil,
		},
		{
			name: "test_file/redis_test_config.yaml",
			args: args{file: "test_file/redis_test_config.yaml"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFake()
			if got := f.LoadFakeFile(tt.args.file); got.Error() == tt.want {
				t.Errorf("LoadFakeFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleFake() {
	f := NewFake().LoadFakeFile("test_file/fund_batch_order_db.yaml").
		LoadFakeFile("test_file/fund_collect_order_db.yaml").
		LoadFakeFile("test_file/redis_test_config.yaml")
	fmt.Println(f.Error())
	// output:nil
}
