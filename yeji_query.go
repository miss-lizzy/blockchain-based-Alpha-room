package main

import(
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct{
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response{
	_, args := stub.GetFunctionAndParameters()
	var name,id,webmail,studentNum,phoneNum string
	var seatNum int //0이면 좌석 사용x, 0이외의 수면 해당 번호의 좌석 사용 중
	var enterBit int //입장했으면 1, 입장하지 않았으면 0
	var uniqueNum int // 고유번호 
	var studentData string
    //var studentId string
	var err error
//+studentID=id
//키 설정 필요 ex) studentId=id

	if len(args)!=7{
		return shim.Error("Incorrect number of arguments, ExPECTING 7")
	}

	name=args[0]
	id=args[1]
	webmail=args[2]
	studentNum=args[3]
	phoneNum=args[4]
	seatNum=args[5]
	enterBit=args[6]

	/*고유번호 발급 function*/

	fmt.Printf("dear %s. your unique number is %d, don't tell anyone.\n",id,uniqueNum)

	studentData="\n"+"name = "+name+"\n"+"id = "+id+"\n"+"webmail = "+webmail+"\n"+"stuentNum = "+studentNum+"\n"+"phoneNum = "+phoneNum+"\n" +"seatNum = "+seatNum +"\n" +"enterBit = "+ enterBit+"\n"+"uniqueNum = "+uniqueNum

	err=stub.PutState(id,[]byte(studentData))
	if err!=nil{
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	function, args:=stub.GetFunctionAndParameters()
	if function == "invoke"{

		return t.invoke(stub,args)
	} else if function =="delete"{
		return t.delete(stub,args)
	} else if function == "query"{
		return t.query(stub,args)
	}
	return shim.Error("Invaild inovke function name. Expecting \"invoke\"\"delete\"\"query\"")
}

//invoke
func(t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	var uniqueNum int	//사용자의 고유번호
	var enterBit, seatNum int	//입장 비트와 좌석 번호
	var err error

	if len(args)!=3{
		return shim.Error("Incorrect number of arguments, Expecting 3")
	}

	uniqueNum = args[0]
	enterBit = args[1]
	seatNum = args[2]

	//원장의 정보 불러오기
	valbytes, err:=stub.GetState(uniqueNum)
	if err!= nil {
		return shim.Error("failed to get state")
	}

	stringData := string(valbytes)

	stringData += "\n" + "enterBit = "+enterBit+"\n"+"seatNum = " + seatNum
	
	
	//원장에 정보 업데이트하기
	err=stub.PutState(eduId, []byte(stringData))
	if err!=nil{
		return shim.Error("Failed to put state")
	}

	return shim.Success(nil)
}

func(t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response{

	if len(args)!=1{
		return shim.Error("Incorrect number of arguments, Expecting 1")
	}
	eduId :=args[0]

	err := stub.DelState(eduId)
	if err != nil{
		return shim.Error("Failed to delete state")
	}
	return shim.Success(nil)
}

func(t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response{

	var studentId string
	var err error

	if len(args)!=1{
		return shim.Error("Incorrect number of arguments, Expecting name of the person to query")
	}

    studentId = args[0]

	valbytes, err := stub.GetState(studentId) // studentId에 따른 정수,문자열 값을 원장에서 가져와서 valbytes에 저장
	if err != nil{
		jsonResp :="{\"Error\":\"Failed to get State for "+seatNum+"\"}"
		return shim.Error(jsonResp)	}
	
		jsonResp:="\n"+"{\"studentId\":\"" + studentId + "\",\n\"Data&State\":\""+string(valbytes)+"\"}" 
		//여기서 스트링만 출력이 아니라 방법 생각해 봐야함, init할때 int형 자료들도 string으로 같이..? 키에 대한 정보 선택이 어떻게 되는지
        fmt.Printf("Query Response:%s\n",jsonResp)
		return shim.Success(valbytes) 
	}

	func main() {
		err:=shim.Start(new(SimpleChaincode))
		if err!=nil{
			fmt.Printf("Error starting Simple chaincode: %s",err)
		}
	}