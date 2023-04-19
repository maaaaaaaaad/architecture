package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"math"
)

func main() {
	// 그래프의 범위 설정
	xMin, xMax := -math.Pi, math.Pi
	yMin, yMax := -1.2, 1.2

	// 삼각함수 그래프의 데이터 생성
	sinData := plotter.XYs{}
	cosData := plotter.XYs{}
	tanData := plotter.XYs{}

	for x := xMin; x <= xMax; x += 0.01 {
		sinData = append(sinData, plotter.XY{x, math.Sin(x)})
		cosData = append(cosData, plotter.XY{x, math.Cos(x)})
		tanData = append(tanData, plotter.XY{x, math.Tan(x)})
	}

	// 그래프 생성
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// 그래프 설정
	p.Title.Text = "삼각함수 그래프"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"
	p.X.Min = xMin
	p.X.Max = xMax
	p.Y.Min = yMin
	p.Y.Max = yMax

	// 그래프에 데이터 추가
	sinLine, err := plotter.NewLine(sinData)
	if err != nil {
		panic(err)
	}
	sinLine.Color = plot.ColorRed
	sinLine.Width = vg.Points(1)

	cosLine, err := plotter.NewLine(cosData)
	if err != nil {
		panic(err)
	}
	cosLine.Color = plot.ColorGreen
	cosLine.Width = vg.Points(1)

	tanLine, err := plotter.NewLine(tanData)
	if err != nil {
		panic(err)
	}
	tanLine.Color = plot.ColorBlue
	tanLine.Width = vg.Points(1)

	// 그래프에 선 추가
	p.Add(sinLine, cosLine, tanLine)

	// 그래프 저장
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "trig.png"); err != nil {
		panic(err)
	}
}
