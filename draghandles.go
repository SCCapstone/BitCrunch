package main

import (
	"github.com/PuerkitoBio/goquery" // for css file parsing, may change if function found in gin
	// https://github.com/robertkrimen/otto --secondary parsing file if need be
	"github.com/gin-gonic/gin"
)

/*
want to do something like the following javascript functions/listeners

	func handleDragStart(e) {
		this.style.opacity = 0.4;
	  }

	  func handleDragEnd(e) {
		this.style.opacity = 1;
	  }

	  let items = document.querySelectorAll('.container .box');
	  items.forEach(function (item) {
		item.addEventListener('dragstart', handleDragStart);
		item.addEventListener('dragend', handleDragEnd);
	  });
*/
func handleDragPageStart(c *gin.Context) {
	_, err := goquery.NewDocument(".draggable")
	if err != nil {
		// error occurred, do something
	}
	/*
		for s:= range(doc.Nodes) {
			single := doc.Eq(s)

		}
	*/
}
