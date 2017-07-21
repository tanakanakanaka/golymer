package main

import (
	"time"

	"github.com/microo8/golymer"
)

const testElemTemplate = `
<style>
	:host {
		display: block;
		box-shadow: 0px 6px 10px #000;
		height: [[height]]px;
	}

	h1 {
		font-size: 100px;
	}
</style>

<h1 id="heading" height="{{Value}}" int="{{intValue2}}">
	<span id="meh" style="background-color: [[BackgroundColor]];">[[content]]</span>
</h1>
<test-elem-two id="two" display="[[Display]]" counter="{{intValue}}"></test-elem-two>

<form>
	<h2 id="formHeading">[[inputObject.Heading]]</h2>
	<input id="inputName" type="text" value="{{inputObject.Name}}">
	<input id="inputAge" type="number" value="{inputObject.Age}}">
	<input id="inputDate" type="date" value="{{inputObject.Date}}">
	<input id="inputActive" type="checkbox" checked="{{inputObject.Active}}">

	<div id="divName" value="{{divObject.Name}}">[[divObject.Name]]</div>
	<div id="divAge" value="{{divObject.Age}}">[[divObject.Age]]</div>
	<div id="divDate" value="{{divObject.Date}}">[[divObject.Date]]</div>
	<div id="divActive" checked="{{divObject.Active}}">[[divObject.Active]]</div>
</form>
`

//TestElem ...
type TestElem struct {
	golymer.Element
	content         string
	height          int
	Display         string
	BackgroundColor string
	Value           string
	intValue        int
	intValue2       int
	inputObject     *TestDataObject
	divObject       *TestDataObject
}

//NewTestElem ...
func NewTestElem() *TestElem {
	elem := &TestElem{
		content:         "Hello world!",
		height:          100,
		Display:         "block",
		BackgroundColor: "red",
		inputObject: &TestDataObject{
			Age:    28,
			Name:   "John",
			Date:   time.Now(),
			Active: true,
		},
		divObject: &TestDataObject{
			Age:    28,
			Name:   "John",
			Date:   time.Now(),
			Active: true,
		},
	}
	elem.Template = testElemTemplate
	return elem
}

//TestDataObject ...
type TestDataObject struct {
	Heading string
	Age     int
	Name    string
	Date    time.Time
	Active  bool
}

//TestElemTwo ...
type TestElemTwo struct {
	golymer.Element
	Display string
	Value   string
	Counter int
}

//NewTestElemTwo ...
func NewTestElemTwo() *TestElemTwo {
	elem := &TestElemTwo{
		Display: "none",
		Value:   "foobar",
	}
	elem.Template = `
	<style>
		:host {
			display: [[Display]];
			background-color: red;
			width: 10vw;
			height: 10vh;
		}
	</style>
	test-elem-two
	`
	return elem
}

func main() {
	err := golymer.Define(NewTestElem)
	if err != nil {
		panic(err)
	}
	err = golymer.Define(NewTestElemTwo)
	if err != nil {
		panic(err)
	}
}
