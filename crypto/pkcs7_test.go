package crypto

import (
	"testing"
	"fmt"
	"github.com/smartystreets/assertions"
)

func TestPadding(t *testing.T){
	src:=[]byte{1,2,3,1}
	result:=PKCS7.padding(src,4)
	fmt.Println(result)

	result,err:=PKCS7.unpadding(result,4)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	fmt.Println(assertions.ShouldEqual(src,result))
}