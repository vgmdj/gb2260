package main

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/vgmdj/gb2260/gbdata"
	"github.com/vgmdj/utils/area"
	"strconv"
	"testing"
)

func TestGenerate(t *testing.T) {
	ast := assert.New(t)

	years := []int{
		2018, 2017, 2016, 2015, 2014,
		2013, 2012, 2011, 2010, 2009,
		2008, 2007, 2006, 2005, 2004,
		2003, 2002, 2001, 2000, 1999,
		1998, 1997, 1996, 1995, 1994,
		1993, 1992, 1991, 1990, 1989,
		1988, 1987, 1986}

	list := new(area.Selector).Revisions()

	for k, v := range list {
		ast.Equal(v == strconv.Itoa(years[k]), true)
	}

}
