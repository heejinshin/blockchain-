package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var b *blockchain 

func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}

//싱글톤
// 최초 한번만 메모리 할당, 메모리 낭비 방지
// blockchain에 다른 패키지에서 접근하고 공유하는걸 쉽게함
// 전역으로 선언해서 해당 인스턴스가 절대적으로 한개만 존재함을 보충