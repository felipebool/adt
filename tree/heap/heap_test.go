package heap

import "testing"

func Test_getChildrenIndexes(t *testing.T) {
	type args struct {
		index int
		last  int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getChildrenIndexes(tt.args.index, tt.args.last)
			if got != tt.want {
				t.Errorf("getChildrenIndexes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getChildrenIndexes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getParentIndex(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getParentIndex(tt.args.index); got != tt.want {
				t.Errorf("getParentIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
