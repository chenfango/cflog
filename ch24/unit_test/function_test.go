/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 11:05:33
 * @LastEditTime: 2019-08-16 11:21:15
 * @LastEditors: Please set LastEditors
 */
package testing

import (
	"testing"

)
func TestSqure(t *testing.T){
	 inputs := [...]int{1,2,3}
	 expected := [...]int{1,4,9}
	 for i:=0;i<len(inputs);i++{
		 ret:=square(inputs[i])
		 if ret!=expected[i]{
			t.Errorf("inputs is %d, the expected is %d, the actual is %d", inputs[i], expected[i], ret)
		 }
	 }

}

