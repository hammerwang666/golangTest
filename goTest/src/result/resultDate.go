package result

//memory   the   test  result   date
var ProcessNum = 0
var TestNum = 0
var CaseNum = 0
var TotalResult []TestResult
var PreResult TestResult
type TestResult  struct {
	Person     ResultType
	Process   ResultType
	Test      ResultType
	CaseResult  CaseResultType
}
type ResultType  struct{
	Name    string
	Number  int
}
type  CaseResultType  struct {
	Result   string
	Number    int
}


//add  up  the   testResult
var TotalReported  []TotalReport
var PreTotalReport TotalReport
type   TotalReport struct {
	ProcessNumber   int
	TotalTestPeople   int
	TotalTest      int
	PreTestTotalCase   int
    PreTestTrueCaseNum  int
	PreTestFalseCaseNum  int
}
