package example

import (
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

/*
 * @author taohu
 * @date 2019/11/8
 * @time 下午7:12
 * @desc please describe function
**/

func TestReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserRepository(ctrl)
	// 期望FindOne(1)返回张三用户
	repo.EXPECT().FindOne(1).Return(&User{Name: "张三"}, nil)
	// 期望FindOne(2)返回李四用户
	repo.EXPECT().FindOne(2).Return(&User{Name: "李四"}, nil)
	// 期望FindOne(3)返回找不到用户的错误
	//repo.EXPECT().FindOne(4).Return(nil, errors.New("user not found"))

	// 验证一下结果
	log.Println(repo.FindOne(1))
	log.Println(repo.FindOne(2))
	log.Println(repo.FindOne(3))
	//log.Println(repo.FindOne(4))
}
