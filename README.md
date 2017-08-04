# golymer
Create HTML [custom elements](https://www.w3.org/TR/custom-elements/#custom-element) with [go](https://golang.org) ([gopherjs](https://github.com/gopherjs/gopherjs))

under construction

contribution of all kind is welcome

![Caution image](caution.png)

With golymer you can create your own HTML custom elements, just by registering an go struct. The innerHTML of the shadowDOM has automatic data bindings to the struct fields (and fields of the struct fields ...).

```go
package main

import "github.com/microo8/golymer"

const myAwesomeTemplate = `
<style>
	:host {
		background-color: blue;
		width: 500px;
		height: [[FooAttr]]px;
	}
</style>
<p>[[privateProperty]]</p>
<input type="text" value="{{BarAttr}}"/>
`

type MyAwesomeElement struct {
	golymer.Element
	FooAttr int
	BarAttr string
	privateProperty float64
}

func NewMyAwesomeElement() *MyAwesomeElement {
	e := &MyAwesomeElement{
		FooAttr: 800,
	}
	e.Template = myAwesomeTemplate
	return e
}

func main() {
	//pass the element constructor to the Define function
	err := golymer.Define(NewMyAwesomeElement)
	if err != nil {
		panic(err)
	}
}
```

Then just run `$ gopherjs build`, import the generated script to your html `<script src="my_awesome_element.js"></script>` and you can use your new element

```html
<my-awesome-element foo-attr="1" bar-attr="hello"></my-awesome-element>
```
