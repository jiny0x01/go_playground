package lotto

import "testing"

func TestCompareLottoNumber(t *testing.T) {
	type args struct {
		userLotto     []int
		computerLotto computerLottoNumber
	}
	tests := []struct {
		name string `json:name`
		args args `json:args`
		want int `json:want`
	}{
		{ // 구조체 순서에 맞춰 초기화
			"1등 테스트",
			args{
				[]int{1, 2, 3, 4, 5, 6},
				computerLottoNumber{
					[]int{1, 2, 3, 4, 5, 6},
					7,
				},
			},
			1,
		},
		{ // json으로 구조체 초기화
			name: "2등 테스트",
			args: args{
				[]int{1, 2, 3, 4, 5, 7},
				computerLottoNumber{
					[]int{1, 2, 3, 4, 5, 6},
					7,
				},
			},
			want: 2,
		},
		{ 
			name: "3등 테스트",
			args: args{
				[]int{1, 2, 3, 4, 5, 7},
				computerLottoNumber{
					[]int{1, 2, 3, 4, 5, 6},
					8,
				},
			},
			want: 3,
		},
		{ 
			name: "4등 테스트",
			args: args{
				[]int{1, 2, 3, 4, 9, 7},
				computerLottoNumber{
					[]int{1, 2, 3, 4, 5, 6},
					8,
				},
			},
			want: 4,
		},
		{ 
			name: "5등 테스트",
			args: args{
				[]int{1, 2, 3, 11,12,13},
				computerLottoNumber{
					[]int{1, 2, 3, 4, 5, 6},
					8,
				},
			},
			want: 5,
		},
		{ 
			name: "낙첨 테스트1",
			args: args{
				[]int{1, 2, 14, 11,12,13},
				computerLottoNumber{
					[]int{1, 2, 3, 4, 5, 6},
					8,
				},
			},
			want: 0,
		},
		{ 
			name: "낙첨 테스트2",
			args: args{
				[]int{1, 15, 16, 11,12,13},
				computerLottoNumber{
					[]int{1, 2, 3, 4, 5, 6},
					8,
				},
			},
			want: 0,
		},
		{ 
			name: "낙첨 테스트3",
			args: args{
				[]int{40, 15, 16, 11,12,13},
				computerLottoNumber{
					[]int{1, 2, 3, 4, 5, 6},
					8,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLottoNumber(tt.args.userLotto, tt.args.computerLotto); got != tt.want {
				t.Errorf("CompareLottoNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
