package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	// 创建邮件对象
	m := gomail.NewMessage()

	// 邮件发送人
	m.SetHeader("From", "utest@21kunpeng.com")

	// 邮件收件人列表
	to := []string{"1005271076@qq.com"}
	m.SetHeader("To", to...)

	// 邮件抄送人列表
	cc := []string{"1005271076@qq.com"}
	m.SetHeader("Cc", cc...)

	// 邮件主题和内容
	m.SetHeader("Subject", "【快手集团招聘通知】")

	//m.SetBody("text/plain", "吕宏林, 您好!\n"+
	//	"非常高兴通知您，您已通过快手技术面考试，接下来将安排hr终面，请您保持电话畅通，3个工作日内会跟你约定面试时间!\n" +
	//	"\n" +
	//	"\n" +
	//	"快手\n" +
	//	"人力资源部")
	m.SetBody("text/plain", "尊敬的[吕宏林]先生："+
		"衷心恭喜您！经过我们公司的严格筛选和面试流程，我很高兴地通知您，您已成功通过我们公司的面试，并被录用为[资深java开发工程师]。在这里，我代表"+
		"整个团队对您的加入表示热烈的欢迎，相信您将会成为我们团队中不可或缺的一员。\n"+
		"根据我们的评估和考量，我们愿意向您提供年薪65万元的优厚薪酬，并为您提供完善的福利待遇、良好的职业发展空间以及广阔的发展前景。我们期待与您共同开启美好的职业旅程，并相信您的专业素养和工作能力将为公司带来新的活力和动力。\n"+
		"为了确保您顺利入职，后续我们会与您进行详细的薪酬福利谈判和签约事宜，同时也将为您提供必要的入职指导和培训计划。如果您有任何疑问或需要进一步了解相关信息，欢迎随时与我们联系。\n"+
		"最后，再次祝贺您成功通过面试，期待与您共同开启美好的工作旅程！\n"+
		"诚挚的祝福，\n"+
		"快手招聘团队 敬上")

	// SMTP服务器配置
	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "utest@21kunpeng.com", "a123456E")

	// 发送邮件
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}

	println("Email sent successfully!")
}
