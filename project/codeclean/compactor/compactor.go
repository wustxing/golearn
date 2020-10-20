package compactor

import (
	"fmt"
	"strings"
)

const(
	ELLIPSIS = "..."
	DELTA_END = "]"
	DELTA_START = "["
)

type Compactor struct{
	contextLength int
	expected      string
	actual        string
	prefixLength  int
	suffixLength  int
}

func NewCompactor(length int,expected string,actual string)*Compactor{
	return &Compactor{
		contextLength: length,
		expected:      expected,
		actual:        actual,
	}
}

func format(expected,actual string)string{
	return fmt.Sprintf("expected:<%s> actual:<%s>",expected,actual)
}

func(p *Compactor) formatCompactedComparison()string{
	if p.shouldNotCompact(){
		return format(p.expected,p.actual)
	}

	compactExpected,compactActual:=p.compactExpectedAndActual()
	return format(compactExpected, compactActual)
}

func(p *Compactor)compactExpectedAndActual()(string,string){
	p.findCommonPrefixAndSuffix()
	return p.compact(p.expected),p.compact(p.actual)
}

func(p *Compactor) compact(s string)string{
	builder:= strings.Builder{}
	builder.WriteString(p.startingEllipsis())
	builder.WriteString(p.startingContext())
	builder.WriteString(DELTA_START)
	builder.WriteString(p.delta(s))
	builder.WriteString(DELTA_END)
	builder.WriteString(p.endingContext())
	builder.WriteString(p.endingEllipsis())
	return builder.String()
}

func(p *Compactor)findCommonPrefixAndSuffix(){
	p.findCommonPrefix()
	p.suffixLength=0

	for ;!p.suffixOverlapsPrefix(p.suffixLength);p.suffixLength++{
		if p.charFromEnd(p.expected,p.suffixLength)!=p.charFromEnd(p.actual,p.suffixLength){
			break
		}
	}
}

func(p *Compactor)shouldNotCompact()bool{
	return p.expected==""||p.actual==""||p.areStringsEqual()
}


func (p *Compactor)charFromEnd(s string, i int)byte{
	return s[len(s)-i-1]
}

func(p *Compactor)suffixOverlapsPrefix(suffixLength int)bool{
	return len(p.actual)-suffixLength<=p.prefixLength ||len(p.expected)-suffixLength<=p.prefixLength
}


func(p *Compactor)findCommonPrefix(){
	p.prefixLength = 0

	end:=min(len(p.expected),len(p.actual))
	for ;p.prefixLength <end;p.prefixLength++{
		if p.expected[p.prefixLength]!=p.actual[p.prefixLength]{
			break
		}
	}
}




func(p *Compactor)areStringsEqual()bool{
	return p.expected==p.actual
}

func(p *Compactor)startingEllipsis()string{
	if p.prefixLength>p.contextLength{
		return ELLIPSIS
	}
	return ""
}

func(p *Compactor)startingContext()string{
	start:=max(0,p.prefixLength-p.contextLength)
	end:=p.prefixLength
	return p.expected[start:end]
}
func(p *Compactor)delta(s string)string{
	start:=p.prefixLength
	end:=len(s)-p.suffixLength
	return s[start:end]
}

func(p *Compactor)endingContext()string{
	start :=len(p.expected)-p.suffixLength
	end:=min(start+p.contextLength,len(p.expected))
	return p.expected[start:end]
}

func(p *Compactor)endingEllipsis()string{
	if p.suffixLength>p.contextLength{
		return ELLIPSIS
	}
	return ""
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}
