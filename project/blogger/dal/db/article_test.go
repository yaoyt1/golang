package db

import (
	"fmt"
	"testing"
	"yytGithub/project/blogger/model"
)

func init() {
	dns := "root:123456@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInstallArticle(t *testing.T) {
	articledatail := &model.ArticleDetailModel{}
	articledatail.UserName = "用户名"
	articledatail.Summary = "golang我们使用go get命令的时候会出现有些包拉取失败， 这是因为我们国内特殊网络的原因。"
	articledatail.Title = "解决golang go get被墙的问题"
	articledatail.ArticleInfoModel.CategoryId = 1
	articledatail.Content = "解决方案1：设置GOPROXY 代理\n我们先运用go env 命令： 先查看自己的GO111MODULE , GOPROXY\n运行export GO111MODULE=on 来开启go module 。\n设置代理： export GOPROXY=https://goproxy.io,direct\n设置完后，相关配置 如下图"

	articleId, err := InstallArticle(articledatail)
	if err != nil {
		fmt.Println("插入失败")
	}
	fmt.Println("插入成功id=", articleId)
	t.Logf("插入成功，文章Id=%d", articleId)
}

func TestSelectAboutArticleList(t *testing.T) {
	items, err := SelectAboutArticleList(1, 10)
	if err != nil {
		t.Errorf("查询相关文章失败：%s", err)
		return
	}
	t.Logf("查询相关文章成功，%#v", items)
}

func TestSelectArticleDetailById(t *testing.T) {
	items, err := SelectArticleDetailById(1)
	if err != nil {
		t.Errorf("查询文章详情失败：%s", err)
		return
	}
	t.Logf("查询文章详情成功，%#v", items)
}

func TestSelectUpDownArticle(t *testing.T) {
	upInfo, downInfo, err := SelectUpDownArticle(1)
	if err != nil {
		t.Errorf("查询上下文章失败：%s", err)
		return
	}
	t.Logf("查询上下文章成功，upInfo:%#v\ndownInfo:%#v", upInfo, downInfo)
}
