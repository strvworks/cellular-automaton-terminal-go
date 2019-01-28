package chunkcontroller

import (
    "errors"

    "../chunk"
)

type Chunkcontroller struct {
    Chunkset map[[2]int]*chunk.Chunk
}

func NewChunkcontroller() (*Chunkcontroller) {
    var cc Chunkcontroller
    cc.Chunkset = map[[2]int]*chunk.Chunk{}
    return &cc
}

func (cc *Chunkcontroller) GetChunk(x, y int) (*chunk.Chunk, error) {
    if val, ok := cc.Chunkset[[2]int{x, y}]; ok {
        return val, nil
    } else {
        return nil, errors.New("That chunk is not found")
    }
}

func (cc *Chunkcontroller) GetCell(cx, cy, x, y int) (int, error) {
    ch, err := cc.GetChunk(cx, cy)
    if err != nil {
        return nil, err
    }
    cell, err := ch.GetCell(x, y)
    if err != nil {
        return nil, err
    }
}

func (cc *Chunkcontroller) SetChunk(c *chunk.Chunk, x, y int) {
    cc.Chunkset[[2]int{x, y}] = c
}

func (cc *Chunkcontroller) NewChunk(x, y int) {
    var c chunk.Chunk
    cc.Chunkset[[2]int{x, y}] = &c
}
