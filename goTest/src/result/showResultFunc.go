package result

import (
	"fmt"
	"log"
)

func ShowResult() {
	//show   the   test  result
	for i := 0; i < len(TotalResult); i++ {
		fmt.Println("===================", "processName:", TotalResult[i].Process.Name, "||", "processNumber:", TotalResult[i].Process.Number, "====================", "\n")
		fmt.Println("- -logPersonAccount:", TotalResult[i].Person.Name)
		fmt.Println("- -logPersonNumber:", TotalResult[i].Person.Number)
		fmt.Println("- - - -testName:", TotalResult[i].Test.Name)
		fmt.Println("- - - -testNumber:", TotalResult[i].Test.Number)
		switch TotalResult[i].CaseResult.Result{
		case"true":
			fmt.Println("- - - - - - - - caseNumber:", TotalResult[i].CaseResult.Number)
			fmt.Println("- - - - - - - -caseResult:", TotalResult[i].CaseResult.Result)
		case"false":
			log.Println("------------------False position------------------------------")
			log.Println("err processNumber: ",TotalResult[i].Process.Number,"|| err ProcessName: ",TotalResult[i].Process.Name)
			log.Println("err logPersonNumber: ",TotalResult[i].Person.Number,"|| err logPersonAccount: ",TotalResult[i].Person.Name)
			log.Println("err testNumber: ",TotalResult[i].Test.Number,"|| err  testName:",TotalResult[i].Test.Name)
			log.Println("err caseNumber: ",TotalResult[i].CaseResult.Number,"|| err caseresult: ",TotalResult[i].CaseResult.Result)

		}

		for {
			if i+1 < len(TotalResult) && TotalResult[i+1].Process.Number == TotalResult[i].Process.Number {
				if TotalResult[i+1].Person.Number == TotalResult[i].Person.Number {
					if TotalResult[i+1].Test.Number == TotalResult[i].Test.Number {
						switch TotalResult[i+1].CaseResult.Result{
						case"true":
							fmt.Println("- - - - - - - - caseNumber:", TotalResult[i+1].CaseResult.Number)
							fmt.Println("- - - - - - - -caseResult:", TotalResult[i+1].CaseResult.Result)
						case"false":
							log.Println("------------------False position------------------------------")
							log.Println("err processNumber: ",TotalResult[i+1].Process.Number,"|| err ProcessName: ",TotalResult[i+1].Process.Name)
							log.Println("err logPersonNumber: ",TotalResult[i+1].Person.Number,"|| err logPersonAccount: ",TotalResult[i+1].Person.Name)
							log.Println("err testNumber: ",TotalResult[i+1].Test.Number,"|| err  testName:",TotalResult[i+1].Test.Name)
							log.Println("err caseNumber: ",TotalResult[i+1].CaseResult.Number,"|| err caseresult: ",TotalResult[i+1].CaseResult.Result)
						}

					}else {
						fmt.Println("- - - -testName:", TotalResult[i+1].Test.Name)
						fmt.Println("- - - -testNumber:", TotalResult[i+1].Test.Number)
						switch TotalResult[i+1].CaseResult.Result{
						case"true":
							fmt.Println("- - - - - - - - caseNumber:", TotalResult[i+1].CaseResult.Number)
							fmt.Println("- - - - - - - -caseResult:", TotalResult[i+1].CaseResult.Result)
						case"false":
							log.Println("------------------False position------------------------------")
							log.Println("err processNumber: ",TotalResult[i+1].Process.Number,"|| err ProcessName: ",TotalResult[i+1].Process.Name)
							log.Println("err logPersonNumber: ",TotalResult[i+1].Person.Number,"|| err logPersonAccount: ",TotalResult[i+1].Person.Name)
							log.Println("err testNumber: ",TotalResult[i+1].Test.Number,"|| err  testName:",TotalResult[i+1].Test.Name)
							log.Println("err caseNumber: ",TotalResult[i+1].CaseResult.Number,"|| err caseresult: ",TotalResult[i+1].CaseResult.Result)
						}

					}

				}else {
					fmt.Println("\n")
					fmt.Println("- -logPersonAccount:", TotalResult[i+1].Person.Name)
					fmt.Println("- -logPersonNumber:", TotalResult[i+1].Person.Number)
					fmt.Println("- - - -testName:", TotalResult[i+1].Test.Name)
					fmt.Println("- - - -testNumber:", TotalResult[i+1].Test.Number)
					switch TotalResult[i+1].CaseResult.Result{
					case"true":
						fmt.Println("- - - - - - - - caseNumber:", TotalResult[i+1].CaseResult.Number)
						fmt.Println("- - - - - - - -caseResult:", TotalResult[i+1].CaseResult.Result)
					case"false":
						log.Println("------------------False position------------------------------")
						log.Println("err processNumber: ",TotalResult[i+1].Process.Number,"|| err ProcessName: ",TotalResult[i+1].Process.Name)
						log.Println("err logPersonNumber: ",TotalResult[i+1].Person.Number,"|| err logPersonAccount: ",TotalResult[i+1].Person.Name)
						log.Println("err testNumber: ",TotalResult[i+1].Test.Number,"|| err  testName:",TotalResult[i+1].Test.Name)
						log.Println("err caseNumber: ",TotalResult[i+1].CaseResult.Number,"|| err caseresult: ",TotalResult[i+1].CaseResult.Result)
					}

				}
				i++
			}else {
				break
			}

		}

	}
	        fmt.Println("\n")
			fmt.Println("....................................report of the test.................................")
	        fmt.Println("共有",ProcessNum,"个流程")
	          for j:=0;j<len(TotalReported);j++{
				if TotalReported[j].TotalTestPeople>0{
					TotalReported[j].TotalTest=TotalReported[j].TotalTest/TotalReported[j].TotalTestPeople
					TotalReported[j].PreTestTotalCase=TotalReported[j].PreTestTotalCase/TotalReported[j].TotalTestPeople
					TotalReported[j].PreTestTrueCaseNum=TotalReported[j].PreTestTrueCaseNum/TotalReported[j].TotalTestPeople
					TotalReported[j].PreTestFalseCaseNum=TotalReported[j].PreTestFalseCaseNum/TotalReported[j].TotalTestPeople
					fmt.Println("\n流程",TotalReported[j].ProcessNumber,":\n",TotalReported[j].TotalTestPeople,"个测试人员",
						TotalReported[j].TotalTest,"个测试",TotalReported[j].PreTestTotalCase,"个测试用例",TotalReported[j].PreTestTrueCaseNum,"个成功",TotalReported[j].PreTestFalseCaseNum,"个失败")
				}


			}
}

func TesDetailResult(arg_testName  string){
	//增加测试结果,一个测试写一个
	TestNum++
	PreResult.Test.Name=arg_testName
	PreResult.Test.Number=TestNum
}
