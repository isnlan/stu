package generate_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	mok "stu/generate/mocks"
	"testing"
)

func TestBlo(t *testing.T) {
	var fake = &mok.FakeBlockPuller{}
	fake.PullBlock(12)
	fmt.Println(fake.PullBlockCallCount())

	fake.PullBlockReturns(errors.New("no number"))
	err := fake.PullBlock(15)
	fmt.Println(err)

	call := fake.PullBlockArgsForCall(1)
	fmt.Println("--", call)

	fake.PullBlockCalls(func(u uint64) error {
		fmt.Println("-->", u)
		return errors.New(fmt.Sprintf("%d", u))
	})
	err = fake.PullBlock(67)
	fmt.Println(err)
}

func TestBlo2(t *testing.T) {
	fake := &mok.Configurator{}
	fake.On("Configure", mock.AnythingOfType("string"), mock.AnythingOfType("int"))
	fake.Configure("fsf", 12)

}
