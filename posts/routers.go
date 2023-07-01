package posts

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) {
	posts := router.Group("api/posts")
	{
		posts.GET("/questions", getAllQuestionsDetailHandler)
		posts.GET("/question/:qid", getQuestionHandler)
		// posts.GET("/question/paginated", getLimitedQuestionsHandler)
		// posts.GET("/question/tags", getQuestionsByTagsHandler)

		posts.POST("/question", createNewQuestionHandler)
		// posts.PUT("/question/:qid", updateQuestionHandler)
		// posts.PUT("/question/:qid/tags",updateQuestionTagsHandler)
		// posts.DELETE("/question/:qid", deleteQuestionHandler)
		posts.POST("/question/:qid/answer", createAnswerHandler)
		// posts.PUT("/answer/:aid", updateAnswerHandler)
		// posts.DELETE("/answer/:aid", deleteAnswerHandler)

		// posts.POST("/answer/:aid/comment", createCommentHandler)
		// posts.PUT("/comment/:cid", updateCommentHandler)
		// posts.DELETE("/comment/:cid", deleteCommentHandler)
		posts.GET("/follow", getAllStarredQuestions)
		posts.GET("/follow/:qid", getQuestionFollowingStatus)
		posts.PUT("/follow/:qid", updateQuestionFollowingStatus)
	}
}
