package main

import "fmt"

type Floenk int

type Wallet struct{
    balance Floenk
}

func (f Floenk) String() string {
    return fmt.Sprintf("%d FLK", f)
}
