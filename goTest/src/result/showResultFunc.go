package result

import (
	"fmt"
	//	"github.com/Centny/Cny4go/log"
)

func ShowResult() {
	//show   the   test  result
	for i := 0; i < len(TotalResult); i++ {
		fmt.Println("===================", "prosessName:", TotalResult[i].Process.Name, "||", "prosessNumber:", TotalResult[i].Process.Number, "====================", "\n")
		fmt.Println("- -logPersonAnnount:", TotalResult[i].Person.Name)
		fmt.Println("- -logPersonNumber:", TotalResult[i].Person.Number)
		fmt.Println("- - - -testName:", TotalResult[i].Test.Name)
		fmt.Println("- - - -testNumber:", TotalResult[i].Test.Number)
		fmt.Println("- - - - - - - - caseNumber:", TotalResult[i].CaseResult.Number)
		fmt.Println("- - - - - - - -caseResult:", TotalResult[i].CaseResult.Result)
		for {
			if i+1 < len(TotalResult) && TotalResult[i+1].Process.Number == TotalResult[i].Process.Number {
				if TotalResult[i+1].Person.Number == TotalResult[i].Person.Number {
					if TotalResult[i+1].Test.Number == TotalResult[i].Test.Number {
						fmt.Println("- - - - - - - -caseNumber:", TotalResult[i+1].CaseResult.Number)
						fmt.Println("- - - - - - - -caseResult:", TotalResult[i+1].CaseResult.Result)

					}else {
						fmt.Println("- - - -testName:", TotalResult[i+1].Test.Name)
						fmt.Println("- - - -testNumber:", TotalResult[i+1].Test.Number)
						fmt.Println("- - - - - - - -caseNumber:", TotalResult[i+1].CaseResult.Number)
						fmt.Println("- - - - - - - -caseResult:", TotalResult[i+1].CaseResult.Result)

					}

				}else {
					fmt.Println("\n")
					fmt.Println("- -logPersonAnnount:", TotalResult[i+1].Person.Name)
					fmt.Println("- -logPersonNumber:", TotalResult[i+1].Person.Number)
					fmt.Println("- - - -testName:", TotalResult[i+1].Test.Name)
					fmt.Println("- - - -testNumber:", TotalResult[i+1].Test.Number)
					fmt.Println("- - - - - - - -caseNumber:", TotalResult[i+1].CaseResult.Number)
					fmt.Println("- - - - - - - -caseResult:", TotalResult[i+1].CaseResult.Result)

				}
				i++
			}else {
				break
			}

		}

	}
	        fmt.Println("\n")
			fmt.Println("....................................repot of the test.................................")
	        for j:=0;j<len(TotalReported);j++{
				if TotalReported[j].TotalTestPeople>0{
					TotalReported[j].TotalTest=TotalReported[j].TotalTest/TotalReported[j].TotalTestPeople
					TotalReported[j].PreTestTotalCase=TotalReported[j].PreTestTotalCase/TotalReported[j].TotalTestPeople
					TotalReported[j].PreTestTrueCaseNum=TotalReported[j].PreTestTrueCaseNum/TotalReported[j].TotalTestPeople
					TotalReported[j].PreTestFalseCaseNum=TotalReported[j].PreTestFalseCaseNum/TotalReported[j].TotalTestPeople
					fmt.Println("共有",ProcessNum,"个流程")
					fmt.Println("流程",TotalReported[j].ProcessNumber,":共有\n",TotalReported[j].TotalTestPeople,"个测试人员",
						TotalReported[j].TotalTest,"个测试",TotalReported[j].PreTestTotalCase,"个测试用例",TotalReported[j].PreTestTrueCaseNum,"个成功",TotalReported[j].PreTestFalseCaseNum,"个失败")
				}


			}
}

