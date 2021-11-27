	package testP

	import (
		"testing"
	)


	func TestPremiers(t *testing.T) {
    		if premiers_2(5) == false {
        		t.Error("Le résultat doit être True car 5 est premier")
    		}	
	}