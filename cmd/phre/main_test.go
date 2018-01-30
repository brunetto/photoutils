package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sort"
)

func TestNewSet(t *testing.T) {
	s := NewStringSet("t1", "t2")
	assert.True(t, s.Contains("t1"))
	assert.True(t, s.Contains("t2"))
}

func TestSet_AddSingle(t *testing.T) {
	s := StringSet{}
	s.Add("t1")
	_, exists := s["t1"]
	assert.True(t, exists)
}

func TestSet_AddSingleTwice(t *testing.T) {
	s := StringSet{}
	s.Add("t1")
	s.Add("t2")
	assert.Equal(t, 2, len(s))
	_, exists := s["t1"]
	assert.True(t, exists)
	_, exists = s["t2"]
	assert.True(t, exists)
}

func TestSet_AddMultiple(t *testing.T) {
	s := StringSet{}
	s.Add("t1", "t2")
	_, exists := s["t1"]
	assert.True(t, exists)
	_, exists = s["t2"]
	assert.True(t, exists)
}

func TestSet_Contains(t *testing.T) {
	s := StringSet{}
	s.Add("t1")
	assert.True(t, s.Contains("t1"))
	assert.False(t, s.Contains("t2"))
}

func TestSet_RemoveSingle(t *testing.T) {
	s := StringSet{}
	s.Add("t1")
	s.Remove("t1")
	_, exists := s["t1"]
	assert.False(t, exists)
}

func TestSet_RemoveMultiple(t *testing.T) {
	s := StringSet{}
	s.Add("t1")
	s.Add("t2")
	s.Add("t3")
	s.Remove("t1", "t2")
	_, exists := s["t1"]
	assert.False(t, exists)
	_, exists = s["t2"]
	assert.False(t, exists)
	_, exists = s["t3"]
	assert.True(t, exists)
}

func TestSplitOnExtension(t *testing.T) {
	fbn, ext := SplitOnExtension("fbn.ext")
	assert.Equal(t, "fbn", fbn)
	assert.Equal(t, ".ext", ext)
}

func TestNewFnDataSet(t *testing.T) {}

func TestFnDataSet_Add(t *testing.T) {
	s := NewFnDataSet()
	s.Add("t1", ".e1")
	s.Add("t1", ".e1")
	s.Add("t1", ".e2")
	s.Add("t2", ".e1")
	assert.Equal(t, 2, len(s))
	t1, exists := s["t1"]
	assert.True(t, exists)
	assert.Equal(t, 2, len(t1.Ext))
	t2, exists := s["t2"]
	assert.True(t, exists)
	assert.Equal(t, 1, len(t2.Ext))
}
func TestFnDataSet_Contains(t *testing.T) {
	s := NewFnDataSet()
	s.Add("t1", ".e1")
	assert.True(t, s.Contains("t1"))
	assert.False(t, s.Contains("t2"))
}
func TestFnDataSet_Remove(t *testing.T) {
	s := NewFnDataSet()
	s.Add("t1", ".e1")
	s.Remove("t1")
	assert.False(t, s.Contains("t1"))
}
func TestFnDataSet_ToSlice_and_Sort(t *testing.T) {
	s := NewFnDataSet()
	s.Add("t1", ".e1")
	s.Add("t2", ".e2")

	ss := s.ToSlice()
	sort.Sort(ss)

	assert.Equal(t, 2, len(ss))
	assert.EqualValues(t, &FnData{BaseName:"t1", Ext:StringSet{".e1": struct{}{}}}, ss[0])
	assert.EqualValues(t, &FnData{BaseName:"t2", Ext:StringSet{".e2": struct{}{}}}, ss[1])

}