package meander

type Facade interface {
	Public() any
}

func Public(o any) any {
	//インターフェースを保持しているかチェック
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
