package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var messages = []string{
	"早安，开启美好一天！",
	"天冷了，多穿衣服",
	"不顺心的话就找我叭",
	"多喝水哦~",
	"要天天开心吖~",
	"保持好心情",
	"照顾好自己",
	"别熬夜",
	"早点休息",
	"今天过的开心嘛",
	"今天也要加油哦！",
	"梦想成真",
	"保持微笑吖",
	"愿你所有烦恼都消失",
	"多吃水果",
	"要按时吃饭",
	"你值得被世界温柔以待",
	"好好爱自己",
	"顺顺利利",
	"愿你笑容常在。",

	"愿你被温柔以待，也被世界偏爱。",
	"阳光正好，微风不燥，一切都刚刚好。",
	"慢一点没关系，生活不需要太匆忙。",
	"请你相信，美好的事正在发生。",
	"别害怕改变，它正带你去更好的地方。",
	"你已经很棒啦，别总苛责自己。",
	"好好吃饭，好好睡觉，好好生活。",
	"心怀浪漫宇宙，也珍惜人间日常。",
	"今天的小美好，也值得被收藏。",
	"偶尔停下来歇歇脚，也是前进的一部分。",
	"天气晴朗，适合和自己握个手言和。",
	"你不需要成为任何人，你本身就很好。",
	"别忘了抬头看看天空，它一直都在。",
	"所有的坚持都值得，所有的等待都不算晚。",
	"别急，一切都会在该来的时候出现。",
	"你的温柔也值得被温柔以待。",
	"请对自己温柔一点，你已经走了很远。",
	"笑一笑，风都会变得柔软。",
	"万事尽意，目之所及皆是欢喜。",
	"愿此刻的你，内心安然、心底有光。",

	"Take a deep breath, you're doing fine.",
	"Small steps still move you forward.",
	"You are enough, just as you are.",
	"Be kind to yourself today.",
	"Let the light in, even if it's small.",
	"Keep a little sunshine in your pocket.",
	"Pause. Breathe. Continue.",
	"You make ordinary days special.",
	"Every sunrise brings new hope.",
	"Be gentle — with yourself and the world.",
}

func randomWarmColor() color.NRGBA {
	r := uint8(rand.Intn(56) + 200)  // 200~255
	g := uint8(rand.Intn(106) + 150) // 150~255
	b := uint8(rand.Intn(80) + 120)  // 120~200
	return color.NRGBA{r, g, b, 255}
}

func showPopup(a fyne.App) {
	msg := messages[rand.Intn(len(messages))]
	bgColor := randomWarmColor()

	w := a.NewWindow("比心")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(200, 60))

	label := canvas.NewText(msg, color.RGBA{51, 51, 51, 255})
	label.TextSize = 18
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: true}

	bg := canvas.NewRectangle(bgColor)
	content := container.NewMax(bg, container.NewCenter(label))
	w.SetContent(content)

	w.SetPadded(false)
	w.CenterOnScreen()

	// 模拟随机偏移
	const offsetX = 400
	const offsetY = 300
	dx := rand.Intn(2*offsetX+1) - offsetX
	dy := rand.Intn(2*offsetY+1) - offsetY

	// 注意：w.Move(...) 在多数平台可能无效，但尝试
	if mover, ok := w.(interface{ Move(pos fyne.Position) }); ok {
		mover.Move(fyne.NewPos(float32(dx), float32(dy)))
	}

	w.Show()

	// 显示时间，之后关闭
	time.AfterFunc(time.Second*2, func() {
		w.Close()
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a := app.NewWithID("com.example.warmlove")
	a.Settings().SetTheme(theme.LightTheme())

	const total = 50
	const interval = 100 * time.Millisecond

	for i := 0; i < total; i++ {
		go func(i int) {
			time.Sleep(time.Duration(i) * interval)
			showPopup(a)
		}(i)
	}

	a.Run()
}
