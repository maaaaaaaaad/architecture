package main

import (
    "math"
    "os"

    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
)

func main() {
    xMin, xMax := -math.Pi, math.Pi
    yMin, yMax := -1.2, 1.2

    sinData := plotter.XYs{}
    cosData := plotter.XYs{}
    tanData := plotter.XYs{}

    for x := xMin; x <= xMax; x += 0.01 {
        sinData = append(sinData, plotter.XY{x, math.Sin(x)})
        cosData = append(cosData, plotter.XY{x, math.Cos(x)})
        tanData = append(tanData, plotter.XY{x, math.Tan(x)})
    }

    p, err := plot.New()
    if err != nil {
        panic(err)
    }

    p.Title.Text = "삼각함수 그래프"
    p.X.Label.Text = "x"
    p.Y.Label.Text = "y"
    p.X.Min = xMin
    p.X.Max = xMax
    p.Y.Min = yMin
    p.Y.Max = yMax

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

    p.Add(sinLine, cosLine, tanLine)

    if err := p.Save(10*vg.Inch, 10*vg.Inch, "trig.png"); err != nil {
        panic(err)
    }

    f, err := os.Open("trig.png")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    buf := make([]byte, 1<<16)
    for {
        n, err := f.Read(buf)
        if err != nil && n == 0
        _, err = os.Stdout.Write(buf[:n])
        if err != nil {
            panic(err)
        }
        if n == 0 {
            break
        }
    }
}
