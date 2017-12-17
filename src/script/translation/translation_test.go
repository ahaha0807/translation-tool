package translation

import (
	"testing"
)

func TestToEnglish(t *testing.T) {
	actual := ToEnglish("バナナ")
	expected := "Banana"

	AssertTranslation(actual, expected, t)

	actual = ToEnglish("耳")
	expected = "Ear"

	AssertTranslation(actual, expected, t)

	actual = ToEnglish("変数")
	expected = "Variable"

	AssertTranslation(actual, expected, t)
}

func TestToJapanese(t *testing.T) {
	actual := ToJapanese("Banana")
	expected := "バナナ"

	AssertTranslation(actual, expected, t)

	actual = ToJapanese("ear")
	expected = "耳"

	AssertTranslation(actual, expected, t)

	actual = ToJapanese("VARIABLE")
	expected = "変数"

	AssertTranslation(actual, expected, t)

	actual = ToJapanese("CoFfEe")
	expected = "コーヒー"

	AssertTranslation(actual, expected, t)
}

func AssertTranslation(actual string, expected string, t *testing.T) {
	if actual == "" {
		t.Errorf("翻訳を実行できませんでした。")
	}

	if actual != expected {
		t.Errorf("翻訳結果が異なります %v, %v", actual, expected)
	}
}
