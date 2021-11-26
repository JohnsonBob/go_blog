package pkg

import (
	"github.com/robfig/cron"
	"go_blog/models"
	"go_blog/pkg/util"
)

func StartClearDataBaseCron() {
	c := cron.New()
	if err := c.AddFunc("0 * * * * *", func() {
		util.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	}); err != nil {
		util.Println(err)
	}
	if err := c.AddFunc("0 * * * * *", func() {
		util.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	}); err != nil {
		util.Println(err)
	}

	c.Start()
	//t1 := time.NewTimer(time.Second * 10)
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 10)
	//	}
	//}
}
