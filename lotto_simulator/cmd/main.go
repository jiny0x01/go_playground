package main

import (
	"fmt"

	"lotto_simulator/src/lotto"
)

func main() {
	fmt.Println("Lotto Simulator")
	for {
		fmt.Println("1. 자동 구매\n2. 수동 구매\n3. 종료")
		var menu int
		for {
			fmt.Print("메뉴 선택(1 or 2 or 3):")
			_, err := fmt.Scanf("%d", &menu)
			if err != nil {
				fmt.Println("\t잘못된 입력")
				continue
			}
			if menu < 1 || menu > 3 {
				continue
			} 
			if menu == 3 {
				return
			}
			break
		}

		fmt.Print("게임 횟수 입력(최대 10000회):")
		var times int
		for {
			_, err := fmt.Scanf("%d", &times)
			if err != nil {
				fmt.Print("잘못된 입력 다시 입력하세요:")
				continue
			}
			if times < 1 || times > 10000 {
				fmt.Print("잘못된 입력 다시 입력하세요:")
				continue
			} 
			break

		}

		computerLotto := lotto.GenerateRandomLottoNumberForComputer()
		var gameResult [6]int
		switch menu {
		case 1:
			var autoGameResult [][]int
			for i:=1; i<=times; i++ {
				fmt.Println(i, "번째 게임")
				autoGameResult = append(autoGameResult, lotto.GenerateRandomLottoNumberForUser())
			}

			for i, game := range autoGameResult {
				result := lotto.CompareLottoNumber(game, &computerLotto)
				fmt.Printf("당첨 번호:%v 보너스 번호:%d\n", computerLotto.Numbers, computerLotto.Bonus)	
				if result == 0 {
					fmt.Println(i+1, "번째 게임:", game, "결과:낙첨")
				} else {
					fmt.Println(i+1, "번째 게임:", game, "결과:", result, "등")
				}
				gameResult[result]++
			}
		case 2:
			var selfGameResult [][]int
			for i:=1; i<=times; i++ {
				fmt.Println(i, "번째 게임")
				selfGameResult = append(selfGameResult, lotto.NewGame())
			}

			for i, game := range selfGameResult {
				result := lotto.CompareLottoNumber(game, &computerLotto)
				fmt.Printf("당첨 번호:%v 보너스 번호:%d\n", computerLotto.Numbers, computerLotto.Bonus)	
				if result == 0 {
					fmt.Println(i+1, "번째 게임:", game, "결과:낙첨")
				} else {
					fmt.Println(i+1, "번째 게임:", game, "결과:", result, "등")
				}
				gameResult[result]++
			}
		}
		fmt.Println("결과 정산")
		fmt.Println("구매한 게임 수:", times)
		for i:=1; i<=5; i++ {
			fmt.Printf("%d등: %d번\n", i, gameResult[i])
		}
		fmt.Printf("낙첨: %d번\n", gameResult[0])
	}
}
