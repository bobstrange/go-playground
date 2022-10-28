package main

type Consumer struct {
	// 定義はいろいろされているとする
	ActiveFlg bool
}

type Consumers []Consumer

// DB などの問い合わせ結果をレシーバーにしてメソッドを呼び出せる様に
// 更に戻り値をレシーバーと同じにしておくと、メソッドチェーンもできる
func (c Consumers) ActiveConsumer() Consumers {
	resp := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.ActiveFlg {
			resp = append(resp, v)
		}
	}
	return resp
}
