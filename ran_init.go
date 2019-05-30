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
	var err error

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

func(t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	var eduId string
	var subjectName, subjectGrade string
	var err error

	if len(args)!=3{
		return shim.Error("Incorrect number of arguments, Expecting 3")
	}

	eduId = args[0]
	subjectName = args[1]
	subjectGrade = args[2]

	
	valbytes, err:=stub.GetState(eduId)
	if err!= nil {
		return shim.Error("failed to get state")
	}

	stringData := string(valbytes)

	stringData += "\n" + "subject = "+subjectName+"\n"+"Grade = " + subjectGrade
	
	
	
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

	var eduId string
	var err error
	if len(args)!=1{
		return shim.Error("Incorrect number of arguments, Expecting name of the person to query")
	}
	eduId = args[0]

	valbytes, err := stub.GetState(eduId)
	if err != nil{
		jsonResp :="{\"Error\":\"Failed to get State for "+eduId+"\"}"
		return shim.Error(jsonResp)	}
	
		jsonResp:="\n"+"{\"EduId\":\"" + eduId + "\",\n\"Data\":\""+string(valbytes)+"\"}"
		fmt.Printf("Query Response:%s\n",jsonResp)
		return shim.Success(valbytes) 
	}

	func main() {
		err:=shim.Start(new(SimpleChaincode))
		if err!=nil{
			fmt.Printf("Error starting Simple chaincode: %s",err)
		}
	}