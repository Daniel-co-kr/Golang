package main

import "fmt"

func main() {

    ctx := context.Background() //최상위 context	

    ctx, cancel := context.WithCancel(ctx) //취소 가능한 context
    defer cancel() // main 함수 실행 후 호출
    
    // goroutine
    go doSomething(ctx)

	//set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//set deadline
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	//return value also
	ctx := context.WithValue(context.Background(), "key", "value")
	value := ctx.Value("key").(string) // "value"

	//parent 기반 context 생성
	rootCtx := context.Background()  // 최상위 부모 context 생성
	parentCtx, parentCancel := context.WithTimeout(rootCtx, 10*time.Second)
	defer parentCancel()  // 부모가 취소되면 자식도 취소됨

	childCtx, childCancel := context.WithTimeout(parentCtx, 5*time.Second)
	defer childCancel()  // 부모와 자식이 모두 취소되면 자식도 취소됨

}

func doSomething(ctx context.Context) error { //context를 첫 번째 매개변수로 설정하는 것이 관례
    select {
    // 취소 신호 핸들링, 상위 컨텍스트에서 cancel() 함수가 실행되어 취소 신호가 전달 된 경우 이 케이스에 해당함
    case <-ctx.Done():
        return ctx.Err()
    default:
        // 실제 작업 코드들
    }
    return nil
}