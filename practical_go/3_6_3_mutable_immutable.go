package main

import "math/big"

// Mutable

type Currency string

type MutableMoney struct {
	currency Currency
	amount   *big.Int
}

func (m MutableMoney) Currency() Currency {
	return m.currency
}

func (m *MutableMoney) SetCurrency(c Currency) {
	m.currency = c
}

// Immutable
// Entity の場合は Mutable にしたほうが良い場合が多いらしい
// ValueObject の場合に Immutable にするっぽい Time など
type ImmutableMoney struct {
	currency Currency
	amount   *big.Int
}

func (im ImmutableMoney) Currency() Currency {
	return im.currency
}

func (im ImmutableMoney) SetCurrency(c Currency) ImmutableMoney {
	return ImmutableMoney{
		currency: c,
		amount:   im.amount,
	}
}
