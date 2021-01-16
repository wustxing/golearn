package compactor

import "testing"

func assertEqual(t *testing.T,a,b string){
	if a!=b{
		t.Logf("%s,%s",a,b)
		t.Fail()
	}
}

func TestMessage(t *testing.T){
	ret:=NewCompactor(0,"b","c").formatCompactedComparison()
	assertEqual(t,"expected:<[b]> actual:<[c]>", ret)
}

func TestStartSame(t *testing.T){
	ret:=NewCompactor(1,"ba","bc").formatCompactedComparison()
	assertEqual(t,"expected:<b[a]> actual:<b[c]>",ret)
}

func TestEndSame(t *testing.T){
	ret:=NewCompactor(1,"ab","cb").formatCompactedComparison()
	assertEqual(t,"expected:<[a]b> actual:<[c]b>",ret)
}

func TestSame(t *testing.T){
	ret:=NewCompactor(1,"ab","ab").formatCompactedComparison()
	assertEqual(t,"expected:<ab> actual:<ab>",ret)
}

func TestNoContextStartAndEndSame(t *testing.T){
	ret:=NewCompactor(0,"abc","adc").formatCompactedComparison()
	assertEqual(t,"expected:<...[b]...> actual:<...[d]...>",ret)
}

func TestStartAndEndContext(t *testing.T){
	ret:=NewCompactor(1,"abc","adc").formatCompactedComparison()
	assertEqual(t,"expected:<a[b]c> actual:<a[d]c>",ret)
}

func TestStartAndEndContextWithEllipses(t *testing.T){
	ret:=NewCompactor(1,"abcde","abfde").formatCompactedComparison()
	assertEqual(t,"expected:<...b[c]d...> actual:<...b[f]d...>",ret)
}

func TestComparisonErrorStartSameComplete(t *testing.T){
	ret:=NewCompactor(2,"ab","abc").formatCompactedComparison()
	assertEqual(t,"expected:<ab[]> actual:<ab[c]>",ret)
}

func TestComparisonErrorStartSameComplete1(t *testing.T){
	ret:=NewCompactor(0,"abfecd","abdcd").formatCompactedComparison()
	assertEqual(t,"expected:<...[fe]...> actual:<...[d]...>",ret)
}





