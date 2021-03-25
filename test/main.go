package main

import (
	"fmt"
	"gatewayMock/test/retriever"
)

type Retriever interface {
	Get(string) string
}

const url = "www.baidu.com"

func main(){
	var r Retriever

	fakeRetriever := retriever.Retriever{
		Contexts: "wanglong",
	}

	r = &fakeRetriever

	if _, ok := r.(*retriever.Retriever); ok {
		fmt.Println(r.Get(url))
	}
}