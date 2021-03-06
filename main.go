package main

import (
    "fmt"
    "time"

    "github.com/nsf/termbox-go"

    cctr "./chunkcontroller"
    "./graphic"
)

const CELL_STR = "██"

type RunState int
const (
    RUN = iota
    STOP
)

func main() {
    if err := termbox.Init(); err != nil {
        panic(err)
    }
    defer termbox.Close()

    w, h := termbox.Size()
    field := graphic.ScreenField{w, h, [2]int{0, 0}, [2]int{0, 0}}

    cc := cctr.NewChunkcontroller()
    field.W, field.H = termbox.Size()
    field.H -= 5
    graphic.DrawField(cc, field, CELL_STR)

    for i := 31; i < 64; i++ {
        cc.SetCell(0, 0, i, 63, 1, true)
        cc.SetCell(0, 0, 63, i, 1, true)
    }
    bin, _ := cc.GetNeighborhood(1, 1, 0, 0)
    fmt.Println(fmt.Sprintf("%08b\n", bin))


    // draw
    updateInterval := 500
    ticker := time.NewTicker(time.Millisecond * time.Duration(updateInterval))
    updaterFunc := func() {
        for range ticker.C {
            field.W, field.H = termbox.Size()
            field.H -= 5
            err := cc.UpdateField()
            if err != nil {
                panic(err)
            }
            graphic.DrawField(cc, field, CELL_STR)
        }
    }

    graphic.DrawField(cc, field, CELL_STR)
    termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

    runstate := STOP
MAINLOOP:
    for {
        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey:
            switch ev.Key {
            case termbox.KeyEsc:
                ticker.Stop()
                break MAINLOOP

            case termbox.KeySpace:
                if runstate == STOP {
                    ticker = time.NewTicker(time.Millisecond * time.Duration(updateInterval))
                    go updaterFunc()
                    runstate = RUN
                } else if runstate == RUN {
                    ticker.Stop()
                    runstate = STOP
                    graphic.DrawField(cc, field, CELL_STR)
                    // graphic.DrawChunk(&ch, [2]int{ox, oy}, CELL_STR)
                }

            case termbox.KeyArrowUp:
                if updateInterval-50 <= 0 {
                    continue
                }
                updateInterval -= 50
                if runstate == RUN {
                    ticker.Stop()
                    ticker = time.NewTicker(time.Millisecond * time.Duration(updateInterval))
                    go updaterFunc()
                }

            case termbox.KeyArrowDown:
                if updateInterval+50 <= 0 {
                    continue
                }
                updateInterval += 50
                if runstate == RUN {
                    ticker.Stop()
                    ticker = time.NewTicker(time.Millisecond * time.Duration(updateInterval))
                    go updaterFunc()
                }
            }
        case termbox.EventMouse:
            switch ev.Key {
            case termbox.MouseLeft:
                mx := ev.MouseX
                my := ev.MouseY
                cx := (mx/len([]rune(CELL_STR)))
                graphic.DrawBottomMessage(fmt.Sprintf("%d, %d     ", cx, my), 1, 0)
                graphic.DrawField(cc, field, CELL_STR)
            }
        }
    }

    fmt.Println("Stop")
}
