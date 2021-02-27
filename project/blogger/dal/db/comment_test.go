package db

import "testing"

func TestInserComment(t *testing.T) {
	commentId, err := InserComment("yyt", "2860889218@qq.com", "文章中有错误", 1)
	if err != nil {
		t.Errorf("评论文章失败：%s\n", err)
		return
	}
	t.Logf("文章评论成功：评论ID=%d\n", commentId)
}

func TestSelectCommentByArticleId(t *testing.T) {
	itmes, err := SelectCommentByArticleId(1)
	if err != nil {
		t.Errorf("查询评论失败：%s\n", err)
		return
	}
	t.Logf("查询文章评论成功：%#v\n", itmes)
}
