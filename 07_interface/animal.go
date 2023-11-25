package main

type Food string

type Animal interface {
	Feed(Food) string
	String() string
}
