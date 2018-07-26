package word

import "testing"

var kor_str string = "가나다라마마바사"
var ahpha_str string = "abcedftABCDefg"
var number_str string = "1233213213120321991"
var chinese_str string = "阿耨多羅三藐三菩堤"
var japanese_str string = "ナムウィキ、みんなでてるの"
var japanese_str2 string = "ナムウィキ、みんなで育てる知識の木"
var thai_str string = "นามุวิกิ ต้นไม้แห่งความรู้ที่พวกคุณปลูก"

func TestKorean(t *testing.T) {

	for _, r := range kor_str {
		if isKorean(r) == false {
			t.Error("invalid korean!", r, string(r))
		}
	}
	for _, r := range ahpha_str {
		if isKorean(r) == true {
			t.Error("invalid korean!", r, string(r))
		}
	}
	for _, r := range number_str {
		if isKorean(r) == true {
			t.Error("invalid korean!", r, string(r))
		}
	}
	for _, r := range chinese_str {
		if isKorean(r) == true {
			t.Error("invalid korean!", r, string(r))
		}
	}
	for _, r := range japanese_str {
		if isKorean(r) == true {
			t.Error("invalid korean!", r, string(r))
		}
	}
	for _, r := range thai_str {
		if isKorean(r) == true {
			t.Error("invalid korean!", r, string(r))
		}
	}
}

func TestAhphabet(t *testing.T) {
	for _, r := range kor_str {
		if isAlphabet(r) == true {
			t.Error("invalid Ahphabet!", r, string(r))
		}
	}
	for _, r := range ahpha_str {
		if isAlphabet(r) == false {
			t.Error("invalid Ahphabet!", r, string(r))
		}
	}
	for _, r := range number_str {
		if isAlphabet(r) == true {
			t.Error("invalid Ahphabet!", r, string(r))
		}
	}
	for _, r := range chinese_str {
		if isAlphabet(r) == true {
			t.Error("invalid Ahphabet!", r, string(r))
		}
	}
	for _, r := range japanese_str {
		if isAlphabet(r) == true {
			t.Error("invalid Ahphabet!", r, string(r))
		}
	}
	for _, r := range thai_str {
		if isAlphabet(r) == true {
			t.Error("invalid Ahphabet!", r, string(r))
		}
	}
}

func TestNumber(t *testing.T) {
	for _, r := range kor_str {
		if isNumber(r) == true {
			t.Error("invalid Number!", r, string(r))
		}
	}
	for _, r := range ahpha_str {
		if isNumber(r) == true {
			t.Error("invalid Number!", r, string(r))
		}
	}
	for _, r := range number_str {
		if isNumber(r) == false {
			t.Error("invalid Number!", r, string(r))
		}
	}
	for _, r := range chinese_str {
		if isNumber(r) == true {
			t.Error("invalid Number!", r, string(r))
		}
	}
	for _, r := range japanese_str {
		if isNumber(r) == true {
			t.Error("invalid Number!", r, string(r))
		}
	}
	for _, r := range thai_str {
		if isNumber(r) == true {
			t.Error("invalid Number!", r, string(r))
		}
	}
}

func TestChinese(t *testing.T) {
	for _, r := range kor_str {
		if isChinese(r) == true {
			t.Error("invalid Chinese1!", r, string(r))
		}
	}
	for _, r := range ahpha_str {
		if isChinese(r) == true {
			t.Error("invalid Chinese2!", r, string(r))
		}
	}
	for _, r := range number_str {
		if isChinese(r) == true {
			t.Error("invalid Chinese3!", r, string(r))
		}
	}
	for _, r := range chinese_str {
		if isChinese(r) == false {
			t.Error("invalid Chinese4!", r, string(r))
		}
	}

	for _, r := range japanese_str {
		if isChinese(r) == true {
			t.Error("invalid Chinese5!", r, string(r))
		}
	}
	for _, r := range thai_str {
		if isChinese(r) == true {
			t.Error("invalid Chinese6!", r, string(r))
		}
	}
}

func TestJapanese(t *testing.T) {
	for _, r := range kor_str {
		if isJapanese(r) == true {
			t.Error("invalid Japanese1!", r, string(r))
		}
	}
	for _, r := range ahpha_str {
		if isJapanese(r) == true {
			t.Error("invalid Japanese2!", r, string(r))
		}
	}
	for _, r := range number_str {
		if isJapanese(r) == true {
			t.Error("invalid Japanese3!", r, string(r))
		}
	}

	for _, r := range japanese_str2 {
		if isJapanese(r) == false {
			t.Error("invalid Japanese5!", r, string(r))
		}
	}
	for _, r := range thai_str {
		if isJapanese(r) == true {
			t.Error("invalid Japanese6!", r, string(r))
		}
	}
}

func TestThai(t *testing.T) {
	for _, r := range kor_str {
		if isThai(r) == true {
			t.Error("invalid Thai1!", r, string(r))
		}
	}
	for _, r := range ahpha_str {
		if isThai(r) == true {
			t.Error("invalid Thai2!", r, string(r))
		}
	}
	for _, r := range number_str {
		if isThai(r) == true {
			t.Error("invalid Thai3!", r, string(r))
		}
	}
	for _, r := range chinese_str {
		if isThai(r) == true {
			t.Error("invalid Thai4!", r, string(r))
		}
	}
	for _, r := range japanese_str {
		if isThai(r) == true {
			t.Error("invalid Thai5!", r, string(r))
		}
	}
	/*
		for _, r := range thai_str {
			if isThai(r) == false {
				t.Error("invalid Thai6!", r, string(r))
			}
		}
	*/
}
