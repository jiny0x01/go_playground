package lotto

import (
	"fmt"
	"math/rand"
	"sort"

	"lotto_simulator/src/lotto_util"
)

type ComputerLottoNumber struct {
	Numbers []int
	Bonus   int
}

func NewGame() []int {
	fmt.Println("1~45중에 숫자 6개를 중복되지 않도록 뽑으세요.")
	numbers := []int{}

	for len(numbers) < 6 {
		var num int
		fmt.Printf("현재 입력된 번호들\n\t - %v\n", numbers)
		fmt.Printf("%d번째 번호 입력:", len(numbers)+1)
		_, err := fmt.Scanf("%d", &num)
		if err != nil || num < 1 || num > 45 {
			fmt.Println("잘못된 입력")
			continue
		}
		if lotto_util.Contain(numbers, num) {
			fmt.Println("이미 존재하는 번호입니다. 재입력하세요.")
			continue
		}
		numbers = append(numbers, num)
		sort.Ints(numbers)
	}
	fmt.Printf("입력한 로또 번호: %v\n", numbers)
	return numbers
}

func GenerateRandomLottoNumberForUser() []int {
	fmt.Println("자동 로또 번호 생성중 . . .")
	var newLottoNumber []int
	for len(newLottoNumber) < 6 {
		// 1 <= num <= 45
		num := rand.Intn(45) + 1
		if lotto_util.Contain(newLottoNumber, num) {
			continue
		}
		newLottoNumber = append(newLottoNumber, num)
	}
	sort.Ints(newLottoNumber)
	fmt.Printf("\t생성된 사용자 로또 번호: %v\n", newLottoNumber)
	return newLottoNumber
}

// 컴퓨터가 생성한 로또 번호
func GenerateRandomLottoNumberForComputer() ComputerLottoNumber {
	fmt.Println("컴퓨터 로또 번호 생성중 . . .")
	var newLottoNumber ComputerLottoNumber
	for len(newLottoNumber.Numbers) < 6 {
		// 1 <= num <= 45
		num := rand.Intn(45) + 1
		if lotto_util.Contain(newLottoNumber.Numbers, num) {
			continue
		}
		newLottoNumber.Numbers = append(newLottoNumber.Numbers, num)
	}
	sort.Ints(newLottoNumber.Numbers)
	for {
		// 1 <= num <= 45
		num := rand.Intn(45) + 1
		if lotto_util.Contain(newLottoNumber.Numbers, num) {
			continue
		}
		newLottoNumber.Bonus = num
		break
	}

	fmt.Printf("\t생성된 로또 번호: %v, 보너스 번호:%d\n", newLottoNumber.Numbers, newLottoNumber.Bonus)
	return newLottoNumber
}

func CompareLottoNumber(userLotto []int, computerLotto *ComputerLottoNumber) int {
	if len(userLotto) != 6 || len(computerLotto.Numbers) != 6 {
		panic("lotto number of Numbers should be 6")
	}

	same := 0
	for _, num := range userLotto {
		if lotto_util.Contain(computerLotto.Numbers, num) {
			same++
		}
	}
	if same == 6 {
		return 1
	}
	if same == 5 && lotto_util.Contain(userLotto, computerLotto.Bonus) {
		return 2
	}
	if same <= 2 { //낙첨
		return 0
	}
	return 8 - same // 3~5등
}
