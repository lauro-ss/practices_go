package main

type Food string

type Animal interface {
	Walk() string
	Feed(Food) string
	String() string
}
