package main

import (
	"fmt"
	"time"
)

type Portion int

const (
	Regular Portion = iota
	Small
	Large
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

// よく利用されるバリエーションを関数として提供する

func NewKakeUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

// 構造体を利用したオプション

type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdonWithOption(opt Option) *Udon {
	// ゼロ値のデフォルト処理
	if opt.ebiten == 0 && time.Now().Hour() < 10 {
		opt.ebiten = 1
	}
	return &Udon{
		men:      opt.men,
		aburaage: opt.aburaage,
		ebiten:   opt.ebiten,
	}
}

// ビルダー利用 ↑ の Option を流用

func NewOption(p Portion) *Option {
	return &Option{men: p}
}
func (o *Option) Aburaage() *Option {
	o.aburaage = true
	return o
}

func (o *Option) Ebiten(n uint) *Option {
	o.ebiten = n
	return o
}

func (o *Option) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

func main() {
	tempraUdon := NewUdon(Large, false, 2)
	fmt.Println("tempraUdon: ", tempraUdon)

	optionalUdon := NewUdonWithOption(Option{men: Large})
	fmt.Println("optionalUdon: ", optionalUdon)

	builderUdon := NewOption(Large).Aburaage().Order()
	fmt.Println("builderUdon: ", builderUdon)
}
