package main

import (
	"fmt"
)

func main() {
	res,e:=doubler(-5);
	if(e!=nil){
		fmt.Println("Error: ",e);
	}else{
		fmt.Println("-5 : ",res);
	}

	res,e=doubler("hello");
	if(e!=nil){
		fmt.Println("Error: ",e);
	}else{
		fmt.Println("hello : ",res);
	}

	res,e=doubler(true);
	if(e!=nil){
		fmt.Println("Error: ",e);
	}else{
		fmt.Println("true : ",res);
	}

	res,e=doubler(3.14);
	if(e!=nil){
		fmt.Println("Error: ",e);
	}else{
		fmt.Println("3.14 : ",res);
	}

	res,e=doubler(uint(42));
	if(e!=nil){
		fmt.Println("Error: ",e);
	}else{
		fmt.Println("42 : ",res);
	}
	fmt.Println("42 : ",res);

	res,e=doubler(uint8(42));
	if(e!=nil){
		fmt.Println("Error: ",e);
	}else{
		fmt.Println("42 : ",res);
	}

	res,e=doubler(uint16(42));
	if(e!=nil){
		fmt.Println("Error: ",e);
	}else{
		fmt.Println("42 : ",res);
	}

	res,e=doubler(uint32(42));
	if(e!=nil){
		fmt.Println("Error: ",e);
	}else{
		fmt.Println("42 : ",res);
	}

	res,e=doubler(uint64(42));
	if(e!=nil){
		fmt.Println("Error: ",e);
	}else{
		fmt.Println("42 : ",res);
	}
}

func doubler(v interface{}) (string,error){
	switch t:=v.(type){
	case string:
		return t+t,nil
	case bool:
		if t{
			return "truetrue",nil
		}else{
			return "falsefalse",nil
		}
	case float32,float64:
		if f,ok:=t.(float64);ok{
			return fmt.Sprintf("float64 %f",f*2),nil
		}else{
			return fmt.Sprintf("float32 %f",t.(float32)*2),nil
		}
	case int:
		return fmt.Sprintf("int %d",t*2),nil
	case int8:
		return fmt.Sprintf("int8 %d",(t)*2),nil
	case int16:
		return fmt.Sprintf("int16 %d",(t)*2),nil
	case int32:
		return fmt.Sprintf("int32 %d",(t)*2),nil
	case int64:
		return fmt.Sprintf("int64 %d",(t)*2),nil
	case uint:
		return fmt.Sprintf("uint %d",(t)*2),nil
	case uint8:
		return fmt.Sprintf("uint8 %d",(t)*2),nil
	case uint16:
		return fmt.Sprintf("uint16 %d",(t)*2),nil
	case uint32:
		return fmt.Sprintf("uint32 %d",(t)*2),nil
	case uint64:
		return fmt.Sprintf("uint64 %d",(t)*2),nil
	default:
		return "",fmt.Errorf("unsupported type: %T",v)
	}
}