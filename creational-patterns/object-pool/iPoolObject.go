package main

type iPoolObject interface {
	getID() string // this is any id which can ve used to compare two different pool objects
}
